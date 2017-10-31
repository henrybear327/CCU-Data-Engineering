package main

/*
godoc -http=":6060"

Merge sort

Procedures:
	1. Split data into chunks, sort before write
	2. Winner tree
*/

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"sort"
	"strconv"
	"time"
)

type ProgramArgument struct {
	inputFilename     *string
	outputFilename    *string
	temporaryFilePath *string

	chunkSize   int
	totalChunks *int

	useChunkSize          *bool
	preserveInputFile     *bool
	preserveTemporaryFile *bool

	isDebug *bool

	useParallel *bool
	depth       *int

	cpuprofile *string
	memprofile *string
}

var config ProgramArgument

func parseCommandLineArgument() {
	fmt.Println("Parse command line argument")

	// Define flags
	config.inputFilename = flag.String("i", "in", "The file to be sorted")
	config.outputFilename = flag.String("o", "out", "The file to store sorted result")
	config.temporaryFilePath = flag.String("tmp", "/tmp", "The path for generated temporary files to be stored")

	config.totalChunks = flag.Int("chunks", 1024, "Minimal chunks to be created")

	config.preserveInputFile = flag.Bool("pi", true, "Set to true to preserve the input file")
	config.preserveTemporaryFile = flag.Bool("pt", false, "Set to true to preserve the temporary file")

	config.isDebug = flag.Bool("d", false, "Set true for debug mode")

	config.useParallel = flag.Bool("p", true, "Default to parallel mode")
	config.depth = flag.Int("depth", 4, "Depth")

	config.cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	config.memprofile = flag.String("memprofile", "", "write memory profile to `file`")

	// parse flags
	flag.Parse()

	// print parsed data
	fmt.Printf("\n\n=================================================\n")
	fmt.Println("input filename: " + *config.inputFilename)
	fmt.Println("output filename: " + *config.outputFilename)
	fmt.Println("temporary file path: " + *config.temporaryFilePath)

	fmt.Println("total chunks to be created: " + strconv.FormatInt(int64(*config.totalChunks), 10))

	fmt.Println("preserve input file: " + strconv.FormatBool(*config.preserveInputFile))
	fmt.Println("preserve temporary file: " + strconv.FormatBool(*config.preserveTemporaryFile))

	fmt.Println("preserve input file: " + strconv.FormatBool(*config.preserveInputFile))

	fmt.Println("debug mode: " + strconv.FormatBool(*config.isDebug))
	fmt.Printf("=================================================\n\n")
}

func openFile(filename string, errorString string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		panic(errorString + " (" + filename + ")")
	}

	if *config.isDebug {
		fmt.Println("File \"" + filename + "\" opened successfully")
	}

	return file
}

func adjustChunkFactors(fileSize int) {
	fmt.Printf("\n=================================================\n")
	if fileSize < *config.totalChunks {
		*config.totalChunks = fileSize
		fmt.Println("Total chunks to be created is updated to " + strconv.FormatInt(int64(*config.totalChunks), 10))
	}

	config.chunkSize = fileSize / *config.totalChunks
	if fileSize%*config.totalChunks != 0 {
		*config.totalChunks++
		fmt.Println("Total chunks to be created is updated to " + strconv.FormatInt(int64(*config.totalChunks), 10))
	}
	fmt.Printf("The expected chunk size is %v byte(s)\n", config.chunkSize)

	// config.pivots = make([]string, *config.totalChunks)
	fmt.Printf("=================================================\n\n")
}

func openTempFile(index int) *os.File {
	filename := *config.temporaryFilePath + "/tmp_" + strconv.FormatInt(int64(index), 10)
	return openFile(filename, "Opening tmpfile error")
}

func createTempFile(index int) *os.File {
	filename := *config.temporaryFilePath + "/tmp_" + strconv.FormatInt(int64(index), 10)
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	if *config.isDebug {
		fmt.Println("tmp file \"" + filename + "\" created successfully")
	}

	return file
}

func createResultFile() *os.File {
	filename := *config.outputFilename
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	if *config.isDebug {
		fmt.Println("Result file \"" + filename + "\" created successfully")
	}

	return file
}

func parallelSort(data []string) {
	ch := make(chan []string, 1)
	mergesort(data, ch, 0)
	data = <-ch
}

func mergesort(data []string, out chan []string, dep int) {
	if *config.depth >= 4 {
		sort.Slice(data, func(i, j int) bool {
			return data[i] > data[j]
		})
		out <- data
		return
	}

	N := len(data)
	res1 := make(chan []string, 1)
	res2 := make(chan []string, 1)
	go mergesort(data[:N/2], res1, dep+1)
	go mergesort(data[N/2:], res2, dep+1)

	l, r := <-res1, <-res2
	i, j := 0, 0
	for ix := range data {
		switch {
		case i == len(l):
			data[ix] = r[j]
			j++
		case j == len(r):
			data[ix] = l[i]
			i++
		case l[i] > r[j]:
			data[ix] = l[i]
			i++
		default:
			data[ix] = r[j]
			j++
		}
	}

	out <- data
}

func writeTempFile(buffer []string, chunkIndex int) {
	// sort
	if *config.useParallel {
		parallelSort(buffer)
	} else {
		sort.Strings(buffer)
	}

	// write
	tmpFileFd := createTempFile(chunkIndex)
	fd := bufio.NewWriter(tmpFileFd)
	for i := 0; i < len(buffer); i++ {
		fd.WriteString(buffer[i])
	}
	fd.Flush()
	tmpFileFd.Close()
}

func splitDataIntoChunks() {
	fmt.Println("Split data into chunks")
	start := time.Now()

	/*
		For merge sort

		- Open the input file
		- Split the input file into k chunk files
			- Sort before writing out
	*/

	inputFile := openFile(*config.inputFilename, "Input file is not found!")
	stat, err := inputFile.Stat()
	if err != nil {
		panic(err)
	}

	adjustChunkFactors(int(stat.Size()))

	accumulatedSize := 0
	chunkIndex := 0

	scanner := bufio.NewScanner(inputFile)

	buffer := make([]string, 0)
	for scanner.Scan() {
		str := scanner.Text()
		str += "\n"
		// fmt.Printf("%v", str) // Println will add back the final '\n'

		buffer = append(buffer, str)
		accumulatedSize += len(str)

		if accumulatedSize >= config.chunkSize {
			if *config.isDebug {
				fmt.Printf("Accumulated bytes is %v\n", accumulatedSize)
			}

			if chunkIndex%100 == 0 {
				fmt.Printf("Finishing slice %v\n", chunkIndex)
			}

			accumulatedSize = 0

			writeTempFile(buffer, chunkIndex)

			chunkIndex++
			buffer = make([]string, 0)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	if accumulatedSize > 0 {
		if *config.isDebug {
			fmt.Printf("Accumulated bytes is %v\n", accumulatedSize)
		}
		if chunkIndex%100 == 0 {
			fmt.Printf("Finishing slice %v\n", chunkIndex)
		}

		writeTempFile(buffer, chunkIndex)

		chunkIndex++ // if accumulatedSize is 0, means the chunkIndex isn't used yet
	}

	*config.totalChunks = chunkIndex

	inputFile.Close()

	fmt.Printf("Time elapsed %v\n", time.Since(start))
}

func mergeChunks() {
	start := time.Now()

	fmt.Println("Merge chunks")

	/*
		Merge all chunks into one

		- open all file descriptors at once
		- perform winner tree
	*/
	var winnerTreeData WinnerTreeData
	winnerTreeData.winnerTreeInit()

	resultFd := createResultFile()
	fd := bufio.NewWriter(resultFd)

	for winnerTreeData.winnerTreeIsEmpty() == false {
		// fmt.Println("Top = " + winnerTreeData.winnerTreeTop())
		fd.WriteString(winnerTreeData.winnerTreeTop() + "\n")

		winnerTreeData.winnerTreePop()
		// winnerTreeData.winnerTreePrint()
	}

	fd.Flush()
	resultFd.Close()

	elapsed := time.Since(start)
	fmt.Printf("Time elapsed %v\n\n", elapsed)
}

func cleanup() {
	fmt.Println("Cleanup")
	start := time.Now()

	if *config.preserveInputFile == false {
		fmt.Println("Input file is being removed...")
		err := os.Remove(*config.inputFilename)
		if err != nil {
			panic(err)
		}
		fmt.Println("Done")
	}

	if *config.preserveTemporaryFile == false {
		fmt.Println("Temporary files are being removed...")
		for i := 0; i < *config.totalChunks; i++ {
			filename := *config.temporaryFilePath + "/tmp_" + strconv.FormatInt(int64(i), 10)

			if *config.isDebug {
				fmt.Println("Removing " + filename)
			}
			err := os.Remove(filename)
			if err != nil {
				panic(err)
			}
		}
		fmt.Println("Done")
	}

	elapsed := time.Since(start)
	fmt.Printf("Time elapsed %v\n\n", elapsed)
}

func readTest() {
	inputFile := openFile(*config.inputFilename, "Input file is not found!")

	scanner := bufio.NewScanner(inputFile)

	buf := make([]byte, 0, 100000)
	scanner.Buffer(buf, 100000)

	for scanner.Scan() {
		str := scanner.Text()
		str += "\n"
		fmt.Printf("%v", str) // Println will add back the final '\n'

		// tmp := scanner.Bytes()
		// fmt.Println(tmp)
	}

	// reader := bufio.NewReader(inputFile)
	// for line, err := reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {
	// 	fmt.Printf("%q\n", line)
	// }
}

func main() {
	start := time.Now()
	parseCommandLineArgument()

	// cou profiling
	if *config.cpuprofile != "" {
		f, err := os.Create(*config.cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	splitDataIntoChunks()
	mergeChunks()
	cleanup()

	// mem profiling
	if *config.memprofile != "" {
		f, err := os.Create(*config.memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}

	elapsed := time.Since(start)
	fmt.Printf("Total runtime %v\n\n", elapsed)

	// readTest()
}

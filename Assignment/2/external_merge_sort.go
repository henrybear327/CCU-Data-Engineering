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
	"os"
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

func closeTempFile(file **os.File) {
	stat, err := (*file).Stat()
	if err != nil {
		panic(err)
	}

	err = (*file).Close()
	if err != nil {
		panic(err)
	}

	if *config.isDebug {
		fmt.Println("tmp file \"" + stat.Name() + "\" is closed successfully\n")
	}
	*file = nil
}

var tempFileSortingTime time.Duration
var writeTempFileTime time.Duration

func writeTempFile(buffer []string, chunkIndex int) {
	// sort
	start := time.Now()

	sort.Strings(buffer)

	end := time.Now()
	elapsed := end.Sub(start)
	tempFileSortingTime += elapsed

	// write
	start = time.Now()

	tmpFileFd := createTempFile(chunkIndex)
	fd := bufio.NewWriter(tmpFileFd)
	for i := 0; i < len(buffer); i++ {
		fd.WriteString(buffer[i])
	}
	fd.Flush()
	closeTempFile(&tmpFileFd)

	end = time.Now()
	elapsed = end.Sub(start)
	writeTempFileTime += elapsed
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

		writeTempFile(buffer, chunkIndex)

		chunkIndex++ // if accumulatedSize is 0, means the chunkIndex isn't used yet
	}

	*config.totalChunks = chunkIndex

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("Time elapsed %v\n", elapsed)
	fmt.Printf("Time elapsed for sorting tmp %v\n", tempFileSortingTime)
	fmt.Printf("Time elapsed for writing tmp %v\n\n", writeTempFileTime)
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

	end := time.Now()
	elapsed := end.Sub(start)
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

	end := time.Now()
	elapsed := end.Sub(start)
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
	splitDataIntoChunks()
	mergeChunks()
	cleanup()

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("Total runtime %v\n\n", elapsed)

	// readTest()
}

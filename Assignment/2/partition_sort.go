/*
Partition sort

Procedures:
	1. Get pivot
	2. Binary search
	3. Sort chunks (merge sort, parallel)
	4. Merge file

TODO / Restrictions:
	1. Chunk size must not exceed available memory
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
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

	// pivots []string
}

var config ProgramArgument

func parseCommandLineArgument() {
	fmt.Println("Parse command line argument")

	// Define flags
	config.inputFilename = flag.String("i", "in", "The file to be sorted")
	config.outputFilename = flag.String("o", "out", "The file to store sorted result")
	config.temporaryFilePath = flag.String("tmp", "/tmp", "The path for generated temporary files to be stored")

	config.totalChunks = flag.Int("chunks", 1024, "Minimal chunks to be created")

	config.preserveInputFile = flag.Bool("pi", false, "Set to true to preserve the input file")
	config.preserveTemporaryFile = flag.Bool("pt", false, "Set to true to preserve the temporary file")

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
	fmt.Printf("=================================================\n\n")
}

func openFile(filename string, errorString string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		panic(errorString + " (" + filename + ")")
	}

	fmt.Println("File \"" + filename + "\" opened successfully")

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
	fmt.Printf("The expected chunk size is %v kilobyte(s)\n", config.chunkSize)

	// config.pivots = make([]string, *config.totalChunks)
	fmt.Printf("=================================================\n\n")
}

func openTempFile(index int) *os.File {
	filename := "tmp_" + strconv.FormatInt(int64(index), 10)
	return openFile(*config.temporaryFilePath+"/"+filename, "Opening tmpfile error")
}

func createTempFile(index int) *os.File {
	filename := "tmp_" + strconv.FormatInt(int64(index), 10)
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("tmp file \"" + filename + "\" created successfully")

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

	fmt.Println("tmp file \"" + stat.Name() + "\" is closed successfully\n")
	*file = nil
}

func writeTempFile(buffer []string, chunkIndex int) {
	sort.Strings(buffer)

	tmpFileFd := createTempFile(chunkIndex)
	for i := 0; i < len(buffer); i++ {
		tmpFileFd.WriteString(buffer[i])
	}
	closeTempFile(&tmpFileFd)
}

func splitDataIntoChunks() {
	fmt.Println("Split data into chunks")

	/*
		For merge sort...

		- Open the input file
		- Split the input file into k chunk files
			- Record the first record of each chunk
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
		// if tmpFileFd == nil {
		// 	tmpFileFd = createTempFile(chunkIndex)
		// }

		str := scanner.Text()
		str += "\n"
		fmt.Printf("%v", str) // Println will add back the final '\n'

		// if accumulatedSize == 0 {
		// 	// this is the pivot
		// 	config.pivots[chunkIndex] = str
		// }

		buffer = append(buffer, str)
		// tmpFileFd.WriteString(str)

		accumulatedSize += len(str)

		if accumulatedSize >= config.chunkSize {
			fmt.Printf("Accumulated bytes is %v\n", accumulatedSize)
			accumulatedSize = 0

			sort.Strings(buffer)

			writeTempFile(buffer, chunkIndex)

			chunkIndex++
			buffer = make([]string, 0)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	if accumulatedSize > 0 {
		fmt.Printf("Accumulated bytes is %v\n", accumulatedSize)

		sort.Strings(buffer)

		writeTempFile(buffer, chunkIndex)

		chunkIndex++ // if accumulatedSize is 0, means the chunkIndex isn't used yet
	}

	*config.totalChunks = chunkIndex
	// fmt.Printf("\n===== pivot(s) ====\n")
	// for i := 0; i < *config.totalChunks; i++ {
	// 	fmt.Printf("%v", config.pivots[i])
	// }
	// fmt.Printf("================\n\n")
}

func sortChunks() {
	fmt.Println("Sort chunks")

	/*
		For every chunk file
		- Open file
		- Read file
		- Sort file
	*/

	// for i := 0; i < *config.totalChunks; i++ {
	// 	tmpFile := openTempFile(i)
	// }
}

func mergeChunks() {
	fmt.Println("Merge chunks")

	/*
		Merge all chunks into one

		Handle deleting original file, temporary files
	*/
}

func cleanup() {
	fmt.Println("Cleanup")
}

func main() {
	parseCommandLineArgument()
	splitDataIntoChunks()
	sortChunks()
	mergeChunks()
	cleanup()
}

/*
Partition sort

Procedures:
	1. Get pivot
	2. Binary search
	3. Sort chunks (merge sort, parallel)
	4. Merge file

Restrictions:
	1. Chunk size must not exceed available memory
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type ProgramArgument struct {
	inputFilename     *string
	outputFilename    *string
	temporaryFilePath *string

	chunkSize   int64
	totalChunks *int64

	useChunkSize          *bool
	preserveInputFile     *bool
	preserveTemporaryFile *bool
}

var programArgument ProgramArgument

func parseCommandLineArgument() {
	fmt.Println("Parse command line argument")

	// Define flags
	programArgument.inputFilename = flag.String("i", "in", "The file to be sorted")
	programArgument.outputFilename = flag.String("o", "out", "The file to store sorted result")
	programArgument.temporaryFilePath = flag.String("tmp", "/tmp", "The path for generated temporary files to be stored")

	programArgument.totalChunks = flag.Int64("chunks", 1024, "Minimal chunks to be created")

	programArgument.preserveInputFile = flag.Bool("pi", false, "Set to true to preserve the input file")
	programArgument.preserveTemporaryFile = flag.Bool("pt", false, "Set to true to preserve the temporary file")

	// parse flags
	flag.Parse()

	// print parsed data
	fmt.Printf("\n\n=================================================\n")
	fmt.Println("input filename: " + *programArgument.inputFilename)
	fmt.Println("output filename: " + *programArgument.outputFilename)
	fmt.Println("temporary file path: " + *programArgument.temporaryFilePath)
	fmt.Println("total chunks to be created: " + strconv.FormatInt(*programArgument.totalChunks, 10))
	fmt.Println("preserve input file: " + strconv.FormatBool(*programArgument.preserveInputFile))
	fmt.Println("preserve temporary file: " + strconv.FormatBool(*programArgument.preserveTemporaryFile))
	fmt.Printf("=================================================\n\n")
}

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		// fmt.Println("File \"" + filename + "\" isn't found")
		panic("File \"" + filename + "\" isn't found")
		// return nil
	}

	fmt.Println("File \"" + filename + "\" opened successfully")

	return file
}

func setChunkFactors(fileSize int64) {
	fmt.Printf("\n=================================================\n")
	if fileSize < *programArgument.totalChunks {
		*programArgument.totalChunks = fileSize
		fmt.Println("Total chunks to be created is updated to " + strconv.FormatInt(*programArgument.totalChunks, 10))
	}

	programArgument.chunkSize = fileSize / *programArgument.totalChunks
	if fileSize%*programArgument.totalChunks != 0 {
		*programArgument.totalChunks++
		fmt.Println("Total chunks to be created is updated to " + strconv.FormatInt(*programArgument.totalChunks, 10))
	}
	fmt.Printf("The expected chunk size is %v kilobyte(s)\n", programArgument.chunkSize)
	fmt.Printf("=================================================\n\n")
}

func splitDataIntoChunks() {
	fmt.Println("Split data into chunks")

	/*
		- Open the input file
		- Split the input file into k chunk files
			- Record the first record of each chunk
	*/

	inputFile := openFile(*programArgument.inputFilename)

	stat, err := inputFile.Stat()
	if err != nil {
		panic(err)
	}

	setChunkFactors(stat.Size())
}

func sortChunks() {
	fmt.Println("Sort chunks")

	/*
		For every chunk file
		- Open file
		- Read file
		- Sort file
	*/
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

/*
Partition sort

1. Get pivot
2. Binary search
3. Sort chunks (merge sort, parallel)
4. Merge file
*/

package main

import (
	"flag"
	"fmt"
	"strconv"
)

type ProgramArgument struct {
	inputFilename     *string
	outputFilename    *string
	temporaryFilePath *string

	chunkSize *int

	preserveInputFile     *bool
	preserveTemporaryFile *bool
}

var programArgument ProgramArgument

func parseCommandLineArgument() {
	fmt.Println("Parse command line argument")

	// Define flags
	programArgument.inputFilename = flag.String("input", "in", "The file to be sorted")
	programArgument.outputFilename = flag.String("output", "out", "The file to store sorted result")
	programArgument.temporaryFilePath = flag.String("tempPath", "/tmp", "The path for generated temporary files to be stored")

	programArgument.chunkSize = flag.Int("chunkSize", 1024, "Chunk size in KB")

	programArgument.preserveInputFile = flag.Bool("preserveInputFile", false, "Preserve the input file or not")
	programArgument.preserveTemporaryFile = flag.Bool("preserveTemporaryFile", false, "Preserve the temporary file or not")

	// parse flags
	flag.Parse()

	// print parsed data
	fmt.Printf("\n\n=================================================\n")
	fmt.Println("input filename: " + *programArgument.inputFilename)
	fmt.Println("output filename: " + *programArgument.outputFilename)
	fmt.Println("temporary file path: " + *programArgument.temporaryFilePath)
	fmt.Println("chunk size in kb: " + strconv.Itoa(*programArgument.chunkSize))
	fmt.Println("preserve input file: " + strconv.FormatBool(*programArgument.preserveInputFile))
	fmt.Println("preserve temporary file: " + strconv.FormatBool(*programArgument.preserveTemporaryFile))
	fmt.Printf("=================================================\n\n")
}

func splitDataIntoChunks() {
	fmt.Println("Split data into chunks")

	/*
		- Open the input file
		- Split the input into k chunks
			- Record the first record of each chunk
	*/

}

func sortChunks() {
	fmt.Println("Sort chunks")

	/*
		Sort each chunk
	*/
}

func mergeChunks() {
	fmt.Println("Merge chunks")

	/*
		Merge all chunks into one

		Handle deleting original file, temporary files
	*/
}

func main() {
	parseCommandLineArgument()
	splitDataIntoChunks()
	sortChunks()
	mergeChunks()
}

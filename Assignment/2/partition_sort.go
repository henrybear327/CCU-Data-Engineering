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

	chunkSize   *int
	totalChunks *int

	useChunkSize          *bool
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
	programArgument.totalChunks = flag.Int("totalChunks", 1024, "Total chunks to be created")

	programArgument.useChunkSize = flag.Bool("useChunkSize", true, "Use chunk size to split the input file or not")
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
	fmt.Println("total chunks to be created: " + strconv.Itoa(*programArgument.totalChunks))
	fmt.Println("Use chunk size for splitting: " + strconv.FormatBool(*programArgument.useChunkSize))
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
	defer file.Close()

	// // get the file size
	// stat, err := file.Stat()
	// if err != nil {
	// 	return
	// }
	// fmt.Println(stat.Name())
	// fmt.Println(stat.Size())

	// // read the file
	// bs := make([]byte, stat.Size())
	// _, err = file.Read(bs)
	// if err != nil {
	// 	return
	// }
	// str := string(bs)
	// fmt.Println(str)

	return file
}

func splitDataIntoChunks() {
	fmt.Println("Split data into chunks")

	/*
		- Open the input file
		- Split the input file into k chunk files
			- Record the first record of each chunk
	*/

	inputFd := openFile(*programArgument.inputFilename)

	if inputFd != nil {
		fmt.Println("Ok")
	} else {
		fmt.Println("Error")
	}
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

func main() {
	parseCommandLineArgument()
	splitDataIntoChunks()
	sortChunks()
	mergeChunks()
}

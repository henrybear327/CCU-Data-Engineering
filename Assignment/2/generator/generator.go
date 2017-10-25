package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func handleFlag() (numbers int, outputFilename string, genRange int64) {
	_numbers := flag.Int("n", 1000, "How many numbers to generate")
	_outputFilename := flag.String("o", "random.out", "The name of the output file")
	_range := flag.Int64("r", 2000000000, "Range of the numbers to be generated")

	flag.Parse()

	numbers = *_numbers
	outputFilename = *_outputFilename
	genRange = *_range
	return numbers, outputFilename, genRange
}

func main() {
	rand.Seed(42)
	n, outputFilename, genRange := handleFlag()
	fmt.Printf("%v %v\n", n, outputFilename)

	file, err := os.Create(outputFilename)
	if err != nil {
		panic(err)
	}
	fd := bufio.NewWriter(file)

	for i := 0; i < n; i++ {
		// fmt.Println(rand.Intn(2000000000))
		fd.WriteString(strconv.FormatInt(rand.Int63n(genRange), 10) + "\n")
	}

	fd.Flush()
	file.Close()
}

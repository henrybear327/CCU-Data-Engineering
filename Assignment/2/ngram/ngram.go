package main

import (
	"flag"
	"fmt"
)

func handleFlag() (int, string) {
	ngram := flag.Int("n", 4, "ngram")
	inputFile := flag.String("i", "in", "Input file")
	flag.Parse()

	return *ngram, *inputFile
}

func main() {
	ngram, inputFile := handleFlag()
	fmt.Printf("ngram = %v, input file = %v\n", ngram, inputFile)

}

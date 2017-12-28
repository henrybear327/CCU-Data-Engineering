package main

import (
	"flag"
	"fmt"
)

func parseArgument() (string, string) {
	str := flag.String("str", "", "string to match")
	regex := flag.String("pat", "", "pattern to match")

	flag.Parse()

	if len(*str) == 0 {
		panic("No string")
	}

	if len(*regex) == 0 {
		panic("No pattern")
	}

	return *str, *regex
}

func regex2postfix(regex string) {
	fmt.Println("regex", regex, "to postfix")
}

func main() {
	fmt.Printf("Regex by Henry\n")

	str, regex := parseArgument()
	fmt.Printf("String is: %v\n", str)
	fmt.Printf("Pattern is: %v\n\n", regex)

	regex2postfix(regex)
}

package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, err := regexp.MatchString("a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	fmt.Println(matched, err)
	// matched, err = regexp.MatchString("bar.*", "seafood")
	// fmt.Println(matched, err)
	// matched, err = regexp.MatchString("a(b", "seafood")
	// fmt.Println(matched, err)
}

package main

import (
	"fmt"
	"regexp"
)

func main() {
	repeat := 30
	str := ""
	regex := ""

	// for i := 0; i < repeat*2; i++ { // exact
	// 	str += "a"
	// }

	for i := 0; i < repeat; i++ { // partial
		str += "a"
	}

	for i := 0; i < repeat; i++ {
		regex += "a?"
	}

	for i := 0; i < repeat; i++ {
		regex += "a"
	}

	fmt.Println(str)
	fmt.Println(regex)

	matched, err := regexp.MatchString(regex, str)
	fmt.Println(matched, err)
	// matched, err = regexp.MatchString("bar.*", "seafood")
	// fmt.Println(matched, err)
	// matched, err = regexp.MatchString("a(b", "seafood")
	// fmt.Println(matched, err)
}

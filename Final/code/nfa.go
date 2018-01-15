package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang-collections/collections/stack"
)

/* Helper functions */

func debugPrintf(format string, str ...interface{}) {
	fmt.Fprintf(os.Stderr, format, str...)
}

func debugPrintln(str ...interface{}) {
	fmt.Fprintln(os.Stderr, str...)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

/* Helper functions */

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

// StackData is the data structure that is stored in the stack
type StackData struct {
	components   int
	alternations int
}

func regex2postfix(regex string) string {
	/*
	   1. Scan the infix expression from left to right.
	   2. If the scanned character is an operand, output it.
	   3. Else,
	   ... 3.1 If the precedence of the scanned operator is greater than the precedence of the operator in the stack(or the stack is empty), push it.
	   ... 3.2 Else, Pop the operator from the stack until the precedence of the scanned operator is less-equal to the precedence of the operator residing on the top of the stack. Push the scanned operator to the stack.
	   4. If the scanned character is an ‘(‘, push it to the stack.
	   5. If the scanned character is an ‘)’, pop and output from the stack until an ‘(‘ is encountered.
	   6. Repeat steps 2-6 until infix expression is scanned.
	   7. Pop and output from the stack until it is not empty.

	   a+b*(c^d-e)^(f+g*h)-i
	   abcd^e-fgh*+^*+i-
	*/

	// Error checking
	// change concat sign to a special char for the future escaping to work

	res := ""
	s := stack.New()
	components := 0
	alternations := 0
	for _, c := range regex {
		// debugPrintf("%c\n", c)
		switch c {
		case '(':
			if components > 1 {
				res += "."
				components--
			}

			s.Push(StackData{components, alternations})
			components = 0
			alternations = 0
		case ')':
			// do cleanup
			// let the components before it concat
			components--
			for components > 0 {
				res += "."
				components--
			}

			// show all alternation
			for alternations > 0 {
				res += "|"
				alternations--
			}

			// restore the state
			top := s.Pop().(StackData)
			alternations = top.alternations
			components = top.components
			components++ // count yourself!
		case '|':
			// let the components before it concat
			components--
			for components > 0 {
				res += "."
				components--
			}

			alternations++
		case '*', '+', '?':
			res += string(c)
		default:
			if components > 1 {
				res += "."
				components--
			}
			components++

			res += string(c)
		}
	}

	components--
	for components > 0 { // add . to concat all components together
		res += "."
		components--
	}

	for alternations > 0 { // print all remaining alternations
		res += "|"
		alternations--
	}

	fmt.Println("infix", regex, "to postfix", res)
	return res
}

func postfix2nfa(postfix string) {
	// debugPrintln("postfix", postfix, "to nfa")
}

func matching(str string, regex string) {
	// debugPrintln("matching", str, "against", regex)
}

func main() {
	str, regex := parseArgument()
	// debugPrintf("String is: %v\n", str)
	// debugPrintf("Pattern is: %v\n\n", regex)

	postfix := regex2postfix(regex)

	postfix2nfa(postfix)

	matching(str, regex)
}

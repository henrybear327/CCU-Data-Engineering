package main

import (
	"flag"
	"fmt"
	"os"
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

	debugPrintln("regex", regex, "to postfix")

	return ""
}

func postfix2nfa(postfix string) {
	debugPrintln("postfix", postfix, "to nfa")
}

func matching(str string, regex string) {
	debugPrintln("matching", str, "against", regex)
}

func main() {
	str, regex := parseArgument()
	debugPrintf("String is: %v\n", str)
	debugPrintf("Pattern is: %v\n\n", regex)

	postfix := regex2postfix(regex)

	postfix2nfa(postfix)

	matching(str, regex)
}

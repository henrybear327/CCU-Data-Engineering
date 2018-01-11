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
func regex2postfix(regex string) string {
	fmt.Println("regex", regex, "to postfix")

	return ""
}

func postfix2nfa(postfix string) {
	fmt.Println("postfix", postfix, "to nfa")
}

func matching(str string, regex string) {
	fmt.Println("matching", str, "against", regex)
}

func main() {
	str, regex := parseArgument()
	fmt.Printf("String is: %v\n", str)
	fmt.Printf("Pattern is: %v\n\n", regex)

	postfix := regex2postfix(regex)

	postfix2nfa(postfix)

	matching(str, regex)
}

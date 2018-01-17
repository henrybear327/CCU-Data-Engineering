package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang-collections/collections/stack"
)

var (
	debugNumberingNFA map[*NFAState]int
	counter           int
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

// Character is the data structure for holding the character, concat signal, etc.
type Character struct {
	control int // 0 for nil, 1 for regular rune, 2 ...
	c       rune
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

// NFAState is a single state in NFA
type NFAState struct {
	control int  // 0 nil, 1 for regular rune
	c       rune // int32, represents a Unicode code point

	out1 *NFAState
	out2 *NFAState

	lastlist int // WTF?
}

// NFAFragment is a couple of connected NFAState during the "build" process of NFA
type NFAFragment struct {
	startingNFAState *NFAState

	outPtr []**NFAState // a list of outgoing ptr that will need to be connected to the next state
}

// go ptr still needs to be dereferenced
// just no need to worry about -> or .
func connectNFAFragments(first, second *NFAFragment) *NFAFragment {
	newFragment := NFAFragment{}

	for i := 0; i < len(first.outPtr); i++ {
		*first.outPtr[i] = second.startingNFAState
	}

	newFragment.startingNFAState = first.startingNFAState
	newFragment.outPtr = make([]**NFAState, len(second.outPtr))
	copy(newFragment.outPtr, second.outPtr)

	return &newFragment
}

func postfix2nfa(postfix string) *NFAState {
	// debugPrintln("postfix", postfix, "to nfa")

	s := stack.New()

	for _, c := range postfix {
		debugPrintln("\nWorking on", string(c))

		switch c {
		case '.':
			second := s.Pop().(*NFAFragment)
			first := s.Pop().(*NFAFragment)

			connected := connectNFAFragments(first, second)
			debugPrintNFA(connected.startingNFAState)
			// debugPrintOutptr(newFragment.outPtr)

			s.Push(connected)
		case '|':
			second := s.Pop().(*NFAFragment)
			first := s.Pop().(*NFAFragment)

			newState := NFAState{0, c, first.startingNFAState, second.startingNFAState, 0}
			debugPrintNFA(&newState)

			newFragment := NFAFragment{&newState, make([]**NFAState, 0)}
			newFragment.outPtr = append(newFragment.outPtr, first.outPtr...)
			newFragment.outPtr = append(newFragment.outPtr, second.outPtr...)
			debugPrintNFA(newFragment.startingNFAState)
			// debugPrintOutptr(newFragment.outPtr)

			s.Push(&newFragment)
		case '?':
		case '+':
		case '*':
		default:
			newState := NFAState{1, c, nil, nil, 0}
			debugPrintNFA(&newState)

			newFragment := NFAFragment{&newState, make([]**NFAState, 0)}
			newFragment.outPtr = append(newFragment.outPtr, &newState.out1)
			debugPrintOutptr(newFragment.outPtr)

			s.Push(&newFragment)
		}
	}

	debugPrintln("\n\n\n")
	return s.Pop().(*NFAFragment).startingNFAState
}

func dfsNFA(cur *NFAState, seen map[*NFAState]bool) {
	if cur == nil {
		return
	}

	if val := seen[cur]; val == true {
		return
	}

	debugPrintln("NFA state", getNumbering(cur), ":", string(cur.c), getNumbering(cur.out1), getNumbering(cur.out2))

	seen[cur] = true

	dfsNFA(cur.out1, seen)
	dfsNFA(cur.out2, seen)
}

func getNumbering(cur *NFAState) int {
	if cur == nil {
		return -1
	}

	if val := debugNumberingNFA[cur]; val == 0 {
		debugNumberingNFA[cur] = counter
		counter++
	}

	val := debugNumberingNFA[cur]
	return val
}

func debugPrintOutptr(outPtr []**NFAState) {
	for i := 0; i < len(outPtr); i++ {
		debugPrintf("%d (%p)-> ", getNumbering(*outPtr[i]), *outPtr[i])
	}
	debugPrintln(" nil")
}

func debugPrintNFA(cur *NFAState) {
	seen := make(map[*NFAState]bool)

	debugPrintln("Starting NFA state", getNumbering(cur))

	dfsNFA(cur, seen)
}

func matching(str string, regex string) {
	// debugPrintln("matching", str, "against", regex)
}

func main() {
	str, regex := parseArgument()
	// debugPrintf("String is: %v\n", str)
	// debugPrintf("Pattern is: %v\n\n", regex)

	postfix := regex2postfix(regex)

	counter = 1
	debugNumberingNFA = make(map[*NFAState]int)

	startingNFAState := postfix2nfa(postfix)
	debugPrintNFA(startingNFAState)

	matching(str, regex)
}

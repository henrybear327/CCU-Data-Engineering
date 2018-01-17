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

	// if len(*str) == 0 { // allow empty string
	// 	panic("No string")
	// }

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
	control int // -1 for match, 0 for epsilon / nil, 1 for regular rune, 2 ...
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
	control int  // -1 for match, 0 epsilon / nil, 1 for regular rune
	c       rune // int32, represents a Unicode code point

	out1 *NFAState
	out2 *NFAState

	timestamp int // when pushed into the list, update this
}

// NFAFragment is a couple of connected NFAState during the "build" process of NFA
type NFAFragment struct {
	startingNFAState *NFAState

	outPtr []**NFAState // a list of outgoing ptr that will need to be connected to the next state
}

// go ptr still needs to be dereferenced
// just no need to worry about -> or .
func connectNFAFragments(first *NFAFragment, second *NFAState) {
	for i := 0; i < len(first.outPtr); i++ {
		*(first.outPtr[i]) = second
	}
}

func postfix2nfa(postfix string) (*NFAState, int) {
	// debugPrintln("postfix", postfix, "to nfa")

	s := stack.New()
	totalStates := 0

	for _, c := range postfix {
		debugPrintln("\nWorking on", string(c))

		switch c {
		case '.':
			second := s.Pop().(*NFAFragment)
			first := s.Pop().(*NFAFragment)

			connectNFAFragments(first, second.startingNFAState)
			debugPrintNFA(first.startingNFAState)

			newFragment := NFAFragment{first.startingNFAState, second.outPtr}
			s.Push(&newFragment)
		case '|':
			second := s.Pop().(*NFAFragment)
			first := s.Pop().(*NFAFragment)

			newState := NFAState{0, c, first.startingNFAState, second.startingNFAState, 0}
			totalStates++
			debugPrintNFA(&newState)

			newFragment := NFAFragment{&newState, make([]**NFAState, 0)}
			newFragment.outPtr = append(newFragment.outPtr, first.outPtr...)
			newFragment.outPtr = append(newFragment.outPtr, second.outPtr...)
			debugPrintNFA(newFragment.startingNFAState)

			s.Push(&newFragment)
		case '?':
			first := s.Pop().(*NFAFragment)

			newState := NFAState{0, c, first.startingNFAState, nil, 0}
			totalStates++
			debugPrintNFA(&newState)

			newFragment := NFAFragment{&newState, make([]**NFAState, 0)}
			newFragment.outPtr = append(newFragment.outPtr, first.outPtr...)
			newFragment.outPtr = append(newFragment.outPtr, &newState.out2)
			debugPrintOutptr(newFragment.outPtr)

			s.Push(&newFragment)
		case '+':
			first := s.Pop().(*NFAFragment)

			newState := NFAState{0, c, first.startingNFAState, nil, 0}
			totalStates++
			debugPrintNFA(&newState)

			connectNFAFragments(first, &newState)

			newFragment := NFAFragment{first.startingNFAState, make([]**NFAState, 0)}
			newFragment.outPtr = append(newFragment.outPtr, &newState.out2)
			debugPrintOutptr(newFragment.outPtr)

			s.Push(&newFragment)
		case '*':
			first := s.Pop().(*NFAFragment)

			newState := NFAState{0, c, first.startingNFAState, nil, 0}
			totalStates++
			debugPrintNFA(&newState)

			connectNFAFragments(first, &newState)

			newFragment := NFAFragment{&newState, make([]**NFAState, 0)}
			newFragment.outPtr = append(newFragment.outPtr, &newState.out2)
			debugPrintOutptr(newFragment.outPtr)

			s.Push(&newFragment)
		default:
			newState := NFAState{1, c, nil, nil, 0}
			totalStates++
			debugPrintNFA(&newState)

			newFragment := NFAFragment{&newState, make([]**NFAState, 0)}
			newFragment.outPtr = append(newFragment.outPtr, &newState.out1)
			debugPrintOutptr(newFragment.outPtr)

			s.Push(&newFragment)
		}
	}

	debugPrintln("\n\n\n")

	matchState := NFAState{-1, ' ', nil, nil, 0}
	result := s.Pop().(*NFAFragment)
	connectNFAFragments(result, &matchState)

	return result.startingNFAState, totalStates
}

func dfsNFA(cur *NFAState, seen map[*NFAState]bool) {
	if cur == nil {
		return
	}

	if val := seen[cur]; val == true {
		return
	}

	// debugPrintln("NFA state", getNumbering(cur), ":", string(cur.c), getNumbering(cur.out1), getNumbering(cur.out2))

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

func addState(cur *NFAState, nextStates *[]*NFAState, timer int) {
	if cur == nil || cur.timestamp == timer { // nil, or already in list, return
		return
	}

	cur.timestamp = timer
	if cur.control == 0 { // epsilon edge, go!
		addState(cur.out1, nextStates, timer)
		addState(cur.out2, nextStates, timer)
		return
	}

	*nextStates = append(*nextStates, cur)
}

func isMatched(currentStates []*NFAState) bool {
	for _, c := range currentStates {
		// fmt.Println("checking", c)
		if c.control == -1 {
			return true // exact match
		}
	}
	return false
}

// returns -1 on mismatch, or the one-past-last index of the partial matched string, starting from theh beginning
func match(startingNFAState *NFAState, totalStates int, str string) int {
	// debugPrintln("matching", str)

	partialMatch := -1

	currentStates := make([]*NFAState, 0)
	addState(startingNFAState, &currentStates, 1) // timer starts at 1
	// fmt.Println("starting len", len(currentStates))
	nextStates := make([]*NFAState, 0)
	for t, c := range str {
		// go
		// fmt.Println("loop, matching", string(c))

		for _, j := range currentStates {
			if j.control == 1 && j.c == c {
				// fmt.Println("Edge matched! Advancing", getNumbering(j))
				addState(j.out1, &nextStates, t+2) // timers starts at 1, add the starting round => +2
			}
		}

		// fmt.Println("len", len(nextStates))
		currentStates = nextStates
		nextStates = make([]*NFAState, 0)

		if isMatched(currentStates) {
			partialMatch = t + 1
		}
	}

	if isMatched(currentStates) && partialMatch == -1 { // empty string case
		partialMatch = 0
	}

	return partialMatch // -1 no match, else: longest match position
}

func main() {
	str, regex := parseArgument()
	// debugPrintf("String is: %v\n", str)
	// debugPrintf("Pattern is: %v\n\n", regex)

	postfix := regex2postfix(regex)

	counter = 1
	debugNumberingNFA = make(map[*NFAState]int)

	startingNFAState, totalStates := postfix2nfa(postfix)
	debugPrintNFA(startingNFAState)
	if totalStates > len(postfix) {
		panic("NFA state overflow!")
	}

	res := match(startingNFAState, totalStates, str)
	fmt.Println("res", res)
	if res == len(str) {
		fmt.Println("Exact match:", str)
	} else if 0 <= res && res < len(str) {
		fmt.Println("Partial match:", str[0:res])
	} else if res == -1 {
		fmt.Println("Mismatch")
	} else {
		panic("WTF is going on with match() return value")
	}
}

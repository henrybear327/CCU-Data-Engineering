/*
https://swtch.com/~rsc/regexp/regexp1.html

Awesome tutorial!

Studied the sample and expanded a bit. More to be done...
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/golang-collections/collections/stack"
)

var (
	debugNumberingNFA map[*NFAState]int
	counter           int
)

const (
	// ANSIColorRed is terminal color
	ANSIColorRed = "\x1b[31m"
	// ANSIColorGreen is terminal color
	ANSIColorGreen = "\x1b[32m"
	// ANSIColorReset is terminal color
	ANSIColorReset = "\x1b[0m"
	// ANSIColorBlue is terminal color
	ANSIColorBlue = "\x1b[33m"
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

/* NFA debugging code */
func dfsNFA(cur *NFAState, seen map[*NFAState]bool) {
	if cur == nil {
		return
	}

	if val := seen[cur]; val == true {
		return
	}

	debugPrintln("NFA state", getNumbering(cur), ":", cur.c.control, string(cur.c.c), getNumbering(cur.out1), getNumbering(cur.out2))

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

/* Core starts here */

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
	control int // -1 for match, 0 for epsilon / nil, 1 for regular rune, 2 for control, 3 for . (match any)
	c       rune
}

func addMatchAnyCharacter(components *int, res *[]Character) {
	if *components > 1 {
		*res = append(*res, Character{2, '.'}) // concat
		*components--
	}
	*components++

	*res = append(*res, Character{3, '.'}) // match any
}

func addValidCharacter(components *int, res *[]Character, c rune) {
	if *components > 1 {
		*res = append(*res, Character{2, '.'})
		*components--
	}
	*components++

	*res = append(*res, Character{1, c})
}

func isOperator(c rune) bool {
	switch c {
	case '+', '?', '*', '|', '(', ')', '.', '\\':
		return true
	default:
		return false
	}
}

func regex2postfix(regex string) []Character {
	res := make([]Character, 0)
	s := stack.New()
	components := 0
	alternations := 0
	prev := rune(' ')
	for _, c := range regex {
		// debugPrintf("%c\n", c)
		if c == '\\' && prev != '\\' {
			// pass this round, first \ is matched
			debugPrintln("Pass, first ", string(c))
			prev = c
		} else if prev == '\\' { // has \ preceding this current one
			if isOperator(c) { // escape!
				// treat this as normal character
				debugPrintln("Escaped, ", string(c))
				addValidCharacter(&components, &res, c)
			} else {
				// no need to escape, e.g. \\a?\\\
				debugPrintln("Makeup, ", string(prev), string(c))
				addValidCharacter(&components, &res, prev)
				addValidCharacter(&components, &res, c)
			}
			prev = ' '
		} else { // normal case
			debugPrintln("Switch", string(c))
			switch c {
			case '(':
				if components > 1 {
					res = append(res, Character{2, '.'})
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
					res = append(res, Character{2, '.'})
					components--
				}

				// show all alternation
				for alternations > 0 {
					res = append(res, Character{2, '|'})
					alternations--
				}

				// restore the state
				if s.Len() == 0 {
					panic("Empty stack! WTF")
				}

				top := s.Pop().(StackData)
				alternations = top.alternations
				components = top.components
				components++ // count yourself!
			case '|':
				// let the components before it concat
				components--
				for components > 0 {
					res = append(res, Character{2, '.'})
					components--
				}

				alternations++
			case '*', '+', '?':
				res = append(res, Character{2, c})
			case '.':
				addMatchAnyCharacter(&components, &res)
			default:
				addValidCharacter(&components, &res, c)
			}

			prev = c
		}
	}

	if prev == '\\' {
		addValidCharacter(&components, &res, prev) // add the last \ if required
	}

	components--
	for components > 0 { // add . to concat all components together
		res = append(res, Character{2, '.'})
		components--
	}

	for alternations > 0 { // print all remaining alternations
		res = append(res, Character{2, '|'})
		alternations--
	}

	// debugPrintln("infix", regex, "to postfix", res)
	// debugPrintln("infix", regex, "to postfix")
	fmt.Printf("regex %v to postfix: ", regex)
	for _, c := range res {
		if c.control == 1 {
			// debugPrintf("%v%v%v", ANSIColorGreen, string(c.c), ANSIColorReset)
			fmt.Printf("%v%v%v", ANSIColorGreen, string(c.c), ANSIColorReset)
		} else if c.control == 2 {
			// debugPrintf("%v%v%v", ANSIColorRed, string(c.c), ANSIColorReset)
			fmt.Printf("%v%v%v", ANSIColorRed, string(c.c), ANSIColorReset)
		} else if c.control == 3 {
			// debugPrintf("%v%v%v", ANSIColorRed, string(c.c), ANSIColorReset)
			fmt.Printf("%v%v%v", ANSIColorBlue, string(c.c), ANSIColorReset)
		} else {
			panic("WTF is being filled to the postfix Character struct?")
		}
	}
	// debugPrintln("")
	fmt.Println("")

	return res
}

// NFAState is a single state in NFA
type NFAState struct {
	c Character

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

func postfix2nfa(postfix []Character) (*NFAState, int) {
	// debugPrintln("postfix", postfix, "to nfa")

	s := stack.New()
	totalStates := 0

	for _, c := range postfix {
		debugPrintln("\nWorking on", c.control, string(c.c))

		if c.control == 2 { // for operators
			switch c.c {
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

				newState := NFAState{Character{0, c.c}, first.startingNFAState, second.startingNFAState, 0}
				totalStates++
				debugPrintNFA(&newState)

				newFragment := NFAFragment{&newState, make([]**NFAState, 0)}
				newFragment.outPtr = append(newFragment.outPtr, first.outPtr...)
				newFragment.outPtr = append(newFragment.outPtr, second.outPtr...)
				debugPrintNFA(newFragment.startingNFAState)

				s.Push(&newFragment)
			case '?':
				first := s.Pop().(*NFAFragment)

				newState := NFAState{Character{0, c.c}, first.startingNFAState, nil, 0}
				totalStates++
				debugPrintNFA(&newState)

				newFragment := NFAFragment{&newState, make([]**NFAState, 0)}
				newFragment.outPtr = append(newFragment.outPtr, first.outPtr...)
				newFragment.outPtr = append(newFragment.outPtr, &newState.out2)
				debugPrintOutptr(newFragment.outPtr)

				s.Push(&newFragment)
			case '+':
				first := s.Pop().(*NFAFragment)

				newState := NFAState{Character{0, c.c}, first.startingNFAState, nil, 0}
				totalStates++
				debugPrintNFA(&newState)

				connectNFAFragments(first, &newState)

				newFragment := NFAFragment{first.startingNFAState, make([]**NFAState, 0)}
				newFragment.outPtr = append(newFragment.outPtr, &newState.out2)
				debugPrintOutptr(newFragment.outPtr)

				s.Push(&newFragment)
			case '*':
				first := s.Pop().(*NFAFragment)

				newState := NFAState{Character{0, c.c}, first.startingNFAState, nil, 0}
				totalStates++
				debugPrintNFA(&newState)

				connectNFAFragments(first, &newState)

				newFragment := NFAFragment{&newState, make([]**NFAState, 0)}
				newFragment.outPtr = append(newFragment.outPtr, &newState.out2)
				debugPrintOutptr(newFragment.outPtr)

				s.Push(&newFragment)
			default:
				panic("Unrecognized operator WTF")
			}
		} else if c.control == 1 || c.control == 3 { // for valid characters, and match any
			newState := NFAState{Character{c.control, c.c}, nil, nil, 0}
			totalStates++
			debugPrintNFA(&newState)

			newFragment := NFAFragment{&newState, make([]**NFAState, 0)}
			newFragment.outPtr = append(newFragment.outPtr, &newState.out1)
			debugPrintOutptr(newFragment.outPtr)

			s.Push(&newFragment)
		} else {
			panic("WTF is wrong with control value is NFA construction")
		}
	}

	debugPrintln("\n\n\n")

	matchState := NFAState{Character{-1, ' '}, nil, nil, 0}
	result := s.Pop().(*NFAFragment)
	connectNFAFragments(result, &matchState)

	return result.startingNFAState, totalStates
}

func addState(cur *NFAState, nextStates *[]*NFAState, timer int) {
	if cur == nil || cur.timestamp == timer { // nil, or already in list, return
		return
	}

	cur.timestamp = timer
	if cur.c.control == 0 { // epsilon edge, go!
		addState(cur.out1, nextStates, timer)
		addState(cur.out2, nextStates, timer)
		return
	}

	*nextStates = append(*nextStates, cur)
}

func isMatched(currentStates []*NFAState) bool {
	for _, c := range currentStates {
		// fmt.Println("checking", c)
		if c.c.control == -1 {
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
			if (j.c.control == 1 && j.c.c == c) || j.c.control == 3 { // match only valid characters
				// fmt.Println("Edge matched! Advancing", getNumbering(j))
				addState(j.out1, &nextStates, t+2) // timers starts at 1, add the starting round => +2
			}
		}

		// fmt.Println("len", len(nextStates))
		currentStates = nextStates
		nextStates = make([]*NFAState, 0)

		if isMatched(currentStates) {
			partialMatch = t + utf8.RuneLen(c)
		}
	}

	if isMatched(currentStates) && partialMatch == -1 { // empty string case
		partialMatch = 0
	}

	return partialMatch // -1 no match, else: longest match position
}

func main() {
	str, regex := parseArgument()
	fmt.Printf("Regex is: %v\n", regex)
	fmt.Printf("String is: %v\n\n", str)

	postfix := regex2postfix(regex)
	debugPrintln(postfix)

	counter = 1
	debugNumberingNFA = make(map[*NFAState]int)

	startingNFAState, totalStates := postfix2nfa(postfix)
	debugPrintNFA(startingNFAState)
	if totalStates > len(postfix) {
		panic("NFA state overflow!")
	}

	res := match(startingNFAState, totalStates, str)
	debugPrintln("res", res)
	if res == len(str) {
		fmt.Printf("Exact match: %v%v%v\n", ANSIColorGreen, str, ANSIColorReset)
	} else if 0 <= res && res < len(str) {
		fmt.Printf("Partial match: %v%v%v%v\n", ANSIColorGreen, str[0:res], ANSIColorReset, str[res:])
	} else if res == -1 {
		fmt.Println("Mismatch")
	} else {
		panic("WTF is going on with match() return value")
	}
}

// Author: Kevin Delassus - G00270791
// Project Discription: A GoLang application which can build a non-deterministic
// finite automaton (NFA) from a regular expression.

package main

import (
	"fmt"
)

//Struts which represent the NFA as a linked collection.
type state struct {
	symbol rune
	edge1 *state
	edge2 *state
}
type nfa struct {
	initial *state
	accept *state
}

func intoport(infix string) string{

	specials := map[rune]int{'*': 10, '.':9,'|':8}

	pofix, s := []rune{}, []rune{}

	for _, r := range infix{
		switch{

		case r == '(':
			s = append(s,r)
		case r == ')':
			for s[len(s)-1] != '('{
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			s = s[:len(s)-1]
		case specials[r] > 0:
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]]{
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			s = append(s,r)
		default:
			pofix = append(pofix, r)
		}
	}

	for len(s) > 0{
		pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
	}

	return string(pofix)
}

func poregtonfa(pofix string) *nfa{

	nfastack := []*nfa{}

	for _, r := range pofix {
		switch r {
		case '.': // Catenation

			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			frag1.accept.edge1 = frag2.initial

			nfastack = append(nfastack, &nfa{initial:frag1.initial, accept:frag2.accept})
		case '|': //Alternation

			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept


			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		case '*': //Zero or More

			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2:&accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		case '+': //One or More

			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			accept := state{}
			//initial := state{edge1: frag.initial, edge2:&accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfastack = append(nfastack,&nfa{initial: frag.initial, accept: &accept})

		default: // Literal Characters
			accept := state{}
			initial := state{symbol:r, edge1: &accept }
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}

	}
	// Returning nfastack at first index
	return nfastack[0]
}

func match(po string , s string) bool{

	ismatch := false

	ponfa := poregtonfa(po)

	current := []*state{}
	next := []*state{}

	current = addState(current[:], ponfa.initial, ponfa.accept)

	for _, r := range s{
		for _, c := range current{
			if c.symbol == r {
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		current, next = next, []*state{}
	}

	for _, c := range current{
		if c == ponfa.accept {
			ismatch = true
			break
		}
	}

	return  ismatch
}
func addState(l []*state, s *state, a *state) []*state {

	l = append(l,s)

	if s != a && s.symbol == 0{
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}
	return l
}
//A function for getting user input from the console for selecting option
func getOptionInput() int{

	var input int
	fmt.Scanln(&input)

	return input
}
// A function for getting a string from the user through the console
func getInput()string{

	var input string
	fmt.Scanln(&input)

	return input

}
// A function for selecting the available options i.e. pofix or infix notation
func option() {

	// Displaying the options to the user
	fmt.Println("================================================================")
	fmt.Println("Select 1 for Pofix\nSelect 2 for Infix\nSelect 0 to Exit")
	fmt.Println("================================================================")
	//Getting user input
	opt := getOptionInput()

	// Keep looping unless the user enters 0
	for opt != 0 {

		switch opt {
		case 1: // Pofix
			// The predefined Pofix Expression
			pofixExp := "ab.c*|"
			//Displaying the predefined pofix expression to the user
			fmt.Println("Polix Expression is ", pofixExp)
			//Asking user for a string
			fmt.Println("Enter a String: ")
			// Running the match function and displaying the result to the user
			fmt.Println("================================================================\n" +
				"Result: ",match(pofixExp, getInput()))

		case 2: // Infix

			// The predefined Infix Expression
			infixExp := "a.b.c"
			// Displaying the predefined infix expression to the user
			fmt.Println("Infix Expression is", infixExp)
			// Asking the user for a string
			fmt.Println("Enter a String: ")
			// Running the match and intoport function and displaying the result to the user
			fmt.Println("================================================================\n" +
				"Result: ",match(intoport(infixExp), getInput()))

		default: // Invalid Character

			// Displaying the options if the user enters a wrong character
			fmt.Println("================================================================")
			fmt.Println("Select 1 for Pofix \nSelect 2 for Infix\nSelect 0 to Exit")
			fmt.Println("================================================================")
		}

		// Once the user has been exited from the switch statement the user is shown the options again.
		fmt.Println("================================================================")
		fmt.Println("Select 1 for Pofix\nSelect 2 for Infix\nSelect 0 to Exit")
		fmt.Println("================================================================")
		// Getting user input
		opt = getOptionInput()
	}
}
// Main function
func main() {
	//Displaying the options to the user
	option()
}
// Author: Kevin Delassus - G00270791
// Project Discription: A GoLang application which can build a non-deterministic
// finite automaton (NFA) from a regular expression.

// Code adapted from : https://swtch.com/~rsc/regexp/regexp1.html

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
// Shunting Yard Algorithm
// Adapted from : https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e
func intoport(infix string) string{

	// Creating a map which maps special characters into integers
	specials := map[rune]int{'*': 10, '.':9,'|':8}
	// Pofix and s is a empty array of runes
	pofix, s := []rune{}, []rune{}

	// Loop over infix string
	for _, r := range infix{

		switch{
		case r == '(': // if the the character is open bracket
			// Adding s to the stack
			s = append(s,r)
		case r == ')': // if the character is closed bracket
			// Popping everything off the stack until we get a closing bracket
			for s[len(s)-1] != '('{
				// Add whatever is before the closing bracket
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			// Discard bracket
			s = s[:len(s)-1]
		case specials[r] > 0: // if the character integer is greater then 0
			//
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]]{
				//get rid of last element on stack
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			s = append(s,r)
		default: // Anything else
			// Taking the character and sticking it on the end of polix array
			pofix = append(pofix, r)
		}
	}

	for len(s) > 0{
		pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
	}

	// returning result which is being cast to a string
	return string(pofix)
}
// Thompson Algorithm
func poregtonfa(pofix string) *nfa{

	nfastack := []*nfa{}

	for _, r := range pofix {
		switch r {
		case '.': // Catenation
			// Popping a character off the nfa stack
			frag2 := nfastack[len(nfastack)-1]
			// Removing that last item from the stack
			nfastack = nfastack[:len(nfastack)-1]
			// Popping another character from the nfa stack
			frag1 := nfastack[len(nfastack)-1]
			// Removing that last item from the stack
			nfastack = nfastack[:len(nfastack)-1]
			// Concatenate the two frags
			frag1.accept.edge1 = frag2.initial
			//Pushes a new fragment to the NFA stack which contains frag1 initial state and frag2 accept state
			nfastack = append(nfastack, &nfa{initial:frag1.initial, accept:frag2.accept})
		case '+': //One or More
			// Popping a character off the nfa stack
			frag := nfastack[len(nfastack)-1]
			// Removing that last item from the stack
			nfastack = nfastack[:len(nfastack)-1]
			//Creating a new accept state
			accept := state{}
			//initial := state{edge1: frag.initial, edge2:&accept}
			// Joining the states
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept
			// Popping the new fragment on the NFA stack
			// The new fragment is the old fragment with two new extra states
			nfastack = append(nfastack,&nfa{initial: frag.initial, accept: &accept})

		case '|': //Alternation
			// Popping a character off the nfa stack
			frag2 := nfastack[len(nfastack)-1]
			// Removing that last item from the stack
			nfastack = nfastack[:len(nfastack)-1]
			// Popping another character from the nfa stack
			frag1 := nfastack[len(nfastack)-1]
			// Removing that last item from the stack
			nfastack = nfastack[:len(nfastack)-1]
			// Creating a new initial state with two edges which point at the initial states of the fragments
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			//Creating a new state
			accept := state{}
			// Joining the states
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept
			// Popping the new fragment on the NFA stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		case '*': //Zero or More
			// Popping a character off the nfa stack
			frag := nfastack[len(nfastack)-1]
			// Removing that last item from the stack
			nfastack = nfastack[:len(nfastack)-1]
			//Creating a new state
			accept := state{}
			// Creating a new initial state with one edges which point at the accept states and frag which points at the initial state of the fragments
			initial := state{edge1: frag.initial, edge2:&accept}
			// Joining the states
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept
			// Popping the new fragment on the NFA stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		case '?': //Zero or one:

			// Popping a character off the nfa stack
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			// Creating a new initial state with one edges which point at the accept states and frag which points at the initial state of
			initial := state{edge1: frag.initial, edge2: frag.accept}
			// Append fragment to the stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: frag.accept})

		default: // Literal Characters
			// Creating a new accept state
			accept := state{}
			// Creating a new initial state  which set the symbol to r, and edge1 points to accept state
			initial := state{symbol:r, edge1: &accept }
			//Popping the new fragment onto the NFA stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}
	}
	// Returning nfastack at first index
	return nfastack[0]
}
// A function which uses the poregtonfa function to determine whether a regular expression
// matches a string.
func match(po string , s string) bool{
	// Setting result variable ismatch to false
	ismatch := false
	// Calling the poregtonfa function and assigning ponfa with the return value
	ponfa := poregtonfa(po)
	// Creating an array of pointers to state
	current := []*state{}
	next := []*state{}

	// finding out the initial state
	current = addState(current[:], ponfa.initial, ponfa.accept)

	// Looping through s a character at a time
	for _, r := range s{
		// Looping through current state
		for _, c := range current{
			// If c.symbol is the same current symbol that im currently reading
			if c.symbol == r {
				// Using add state function and storing the result in next
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		// Swapping current with next
		current, next = next, []*state{}
	}
	// Loop through the current array
	for _, c := range current{
		// if the state that im running through in the currant array is equaled to the accept state of ponfa.
		if c == ponfa.accept {
			// Set ismatch to true
			ismatch = true
			//break out of for loop
			break
		}
	}
	// Returning the result
	return  ismatch
}
// Takes the current slice and adds state s and goes to s checking
// if its one of the states with e arrows coming from it
func addState(l []*state, s *state, a *state) []*state {
	// Append the state that has been passed in
	l = append(l,s)
	if s != a && s.symbol == 0{
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}
	// returning the list of pointers
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
	fmt.Println("Select 1 for Postfix\nSelect 2 for Infix\nSelect 0 to Exit")
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
			fmt.Println("Postfix Expression is ", pofixExp)
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
			fmt.Println("Select 1 for Postfix \nSelect 2 for Infix\nSelect 0 to Exit")
			fmt.Println("================================================================")
		}

		// Once the user has been exited from the switch statement the user is shown the options again.
		fmt.Println("================================================================")
		fmt.Println("Select 1 for Postfix\nSelect 2 for Infix\nSelect 0 to Exit")
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
package main

import (
	"fmt"
)

//import "fmt"

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
		case '.':
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			frag1.accept.edge1 = frag2.initial

			nfastack = append(nfastack, &nfa{initial:frag1.initial, accept:frag2.accept})
		case '|':
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept


			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		case '*':

			/*e = pop();
			s = state(Split, e.start, NULL);
			patch(e.out, s);
			push(frag(s, list1(&s->out1)));*/

			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2:&accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		case '+':

			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			accept := state{}
			//initial := state{edge1: frag.initial, edge2:&accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfastack = append(nfastack,&nfa{initial: frag.initial, accept: &accept})

		default:
			accept := state{}
			initial := state{symbol:r, edge1: &accept }
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}

	}

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
//A function for getting user input from the console
func getOptionInput() int{

	var input int
	fmt.Scanln(&input)

	return input
}
func getInput()string{

	var input string
	fmt.Scanln(&input)

	return input

}
func option() {

	fmt.Println("================================================================")
	fmt.Println("Select 1 for Pofix\nSelect 2 for Infix\nSelect 0 to Exit")
	fmt.Println("================================================================")
	opt := getOptionInput()

	for opt != 0 {

		switch opt {
		case 1:
			pofixExp := "ab.c*|"
			fmt.Println("Polix Expression is ", pofixExp)
			fmt.Println("Enter a String: ")
			fmt.Println("================================================================\n" +
				"Result: ",match(pofixExp, getInput()))
			break
		case 2:
			infixExp := "a.b.c"
			fmt.Println("Infix Expression is", infixExp)
			fmt.Println("Enter a String: ")
			fmt.Println("================================================================\n" +
				"Result: ",match(intoport(infixExp), getInput()))
			break
		default:
			fmt.Println("================================================================")
			fmt.Println("Select 1 for Pofix \nSelect 2 for Infix\nSelect 0 to Exit")
			fmt.Println("================================================================")
		}
		fmt.Println("================================================================")
		fmt.Println("Select 1 for Pofix\nSelect 2 for Infix\nSelect 0 to Exit")
		fmt.Println("================================================================")
		opt = getOptionInput()
	}
}
func main() {
	option()
}
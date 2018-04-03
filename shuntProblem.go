package main

import (
	"fmt"
	"bufio"
	"os"
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
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2:&accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		default:
			accept := state{}
			initial := state{symbol:r, edge1: &accept }
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}

	}

	return nfastack[0]
}

func pomatch(po string , s string) bool{

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
func getInput() string{

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please Enter Input: ")
	input, _ := reader.ReadString('\n')

	return input
}
func main(){

	//ab.c*|
	//def

	fmt.Println(pomatch(getInput(), getInput()))
}
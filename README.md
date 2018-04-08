# READ me

A GoLang application which can build a non-deterministic finite automaton (NFA) from a regular expression. 

## Problem Statement Given

You must write a program in the Go programming language that can build a non-deterministic finite automaton (NFA) from a regular expression, and can use the NFA to check if the regular expression matches any given string of text. You must write the program from scratch and cannot use the regexp package from the Go standard library nor any other external library.
A regular expression is a string containing a series of characters, some of which may have a special meaning. For example, the three characters “.”, “|”, and “” have the special meanings “concatenate”, “or”, and “Kleene star” respectively. So, 0.1 means a 0 followed by a 1, 0|1 means a 0 or a 1, and 1 means any number of 1’s. These special characters must be used in your submission.
Other special characters you might consider allowing as input are brackets “()” which can be used for grouping, “+” which means “at least one of”, and “?” which means “zero or one of”. You might also decide to remove the concatenation character, so that 1.0 becomes 10, with the concatenation implicit.
You may initially restrict the non-special characters your program works with to 0 and 1, if you wish. However, you should at least attempt to expand these to all of the digits, and the characters a to z, and A to Z.
You are expected to be able to break this project into a number of smaller tasks that are easier to solve, and to plug these together after they have been completed. You might do that for this project as follows:

1. Parse the regular expression from infix to postfix notation.
2. Build a series of small NFA’s for parts of the regular expression.
3. Use the smaller NFA’s to create the overall NFA.
4. Implement the matching algorithm using the NFA.

## Running Non Deterministic Finite Automaton Application

To run the problem sheet you first need to make sure that you have GO on your PC.If you dont go to the link below and download and install GO.

https://golang.org/

To complete the next step GIT is also required on you PC. If you don't have GIT installed go to the link below and download and install GIT

https://git-scm.com/

Once both has been installed you are ready to run the go program Open your Console again and navigate to where you would like to clone the problem sheet repository. To clone enter:

    $ git clone https://github.com/kevind992/Non-Deterministic-Finite-Automaton-Application.git

Once the repository has been cloned, navigate into the problem sheet folder. To build the GO file enter:

    $ go build shuntProblem.go

To run the program, enter:

    $ shuntProblem.exe

The Non Deterministic Finite Automaton Application should now be running on the console.
For both Infix and Pofix, expressions are already predetermined. This can be changed in the option function of the code.
To enter a Pofix notation select '1' , for a infix notation select '2' or to exit select '0'.
A string needs to be entered to see if it matches the expression. Once the string is entered a result of true or false will be displayed.

## Example of Running Code

- Example 1

![newsnapshot_1](https://user-images.githubusercontent.com/31921534/38460544-00394922-3ab4-11e8-93a3-384130b9707d.JPG)

- Example 2

![newsnapshot_2](https://user-images.githubusercontent.com/31921534/38460549-1bfa6efc-3ab4-11e8-9eff-9ace94db36fb.JPG)

## Research and Coding

Before I attempted coding out the problem I did some research on on Regex and the Shunting Yard and Thompson's algoritms

- https://swtch.com/~rsc/regexp/regexp1.html
- https://en.wikipedia.org/wiki/Regular_expression
- https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_Expressions
- https://en.wikipedia.org/wiki/Shunting-yard_algorithm
- https://en.wikipedia.org/wiki/Thompson%27s_construction

After I compleated researching I started following weekly tutorial videos which our Graph Theory lecture provided us with.
I coded along side his videos taking notes along the way.

## Documented Extras

- Added an option to pick between infix and postfix. Once the option was picked you would be prompted to enter a string to see if it   matches with the expression given.
- Added the '+' opperator into Thompsons algorithm

```
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
```			
- Added '?' opperator into Thompsons algorithm
- I included the Shunting Yard Algorithm into my project by including the following code

```
match(intoport(infixExp), getInput()))
```

## Technology User

The Go Programming Language
- https://golang.org/

I coded using JetBrains GoLand
- https://www.jetbrains.com/go/

Git Distributed Version Control System
- https://git-scm.com/

## Recommendations for Future Development

- I would include different go files instead of having all the code in one go file. The main and option function could have been on its own and the the Shunting Yard and Thompsons Algorithms could have had another go file. It would make the code alot less cluttered.

## References
- https://swtch.com/~rsc/regexp/regexp1.html
- https://en.wikipedia.org/wiki/Regular_expression
- https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_Expressions
- https://en.wikipedia.org/wiki/Shunting-yard_algorithm
- https://en.wikipedia.org/wiki/Thompson%27s_construction

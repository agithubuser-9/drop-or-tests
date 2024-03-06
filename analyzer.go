package main

import "fmt"

func gotReduced(stack *[]string, checkpointBackwardsIndex int) bool {

	if DigitWasReduced(stack, checkpointBackwardsIndex) {
		return true
	}

	subStackString := SliceOfCharactersToString((*stack)[checkpointBackwardsIndex:])

	grammar := map[string]string{
		"ExprOpExpr": "Expr",
		"Expr-Expr":  "Expr",
		"(Expr)":     "Expr",
		"Expr-":      "Expr",
		"+":          "Op",
		"-":          "Op",
		"*":          "Op",
	}

	respectiveValue, belongsGrammar := grammar[subStackString]

	if belongsGrammar {
		PopStringsFromSlice(stack, checkpointBackwardsIndex)
		PushCharactersStringIntoSlice(stack, respectiveValue)
	}

	return belongsGrammar

}

func goBackwardsToReduce(stack *[]string) {

	topStackIndex := len(*stack) - 1

	for checkpointBackwards := topStackIndex; checkpointBackwards >= 0; checkpointBackwards-- {

		/*
			it tries to reduce and if it got reduced it
			goes back to top and tries to reduce again
		*/
		if gotReduced(stack, checkpointBackwards) {

			// stack size may have changed
			topStackIndex = len(*stack) - 1
			checkpointBackwards = topStackIndex

			Test("reduced\n", nil)
		}

	}

}

func CheckSyntax(characters []string) {

	var stack []string
	fmt.Println()

	// if it were string instead of slice it would return Runes
	for _, char := range characters {

		Test("shift", nil)
		stack = append(stack, char)

		goBackwardsToReduce(&stack)

	}

	resultingStack := SliceOfCharactersToString(stack)

	/*
		fmt.Println( len(resultingStack) )
		fmt.Println( len("Expr") )
		fmt.Println(resultingStack)

		return resultingStack == "Expr"


		"Expr" has length 4 and
		resultingStack length is 6
	*/

	fmt.Println(resultingStack)
	fmt.Println()

}

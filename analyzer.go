package main

func gotReduced(stack *[]string, checkpointBackwardsIndex int) bool {

	if DigitWasReduced(stack, checkpointBackwardsIndex) {
		return true
	}

	subStackString := SliceOfCharactersToString((*stack)[checkpointBackwardsIndex:])

	grammar := map[string]string{
		"ExprOpExpr": "Expr",
		"(Expr)":     "Expr",
		"-Expr":      "Expr",
		"num":        "Expr",
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

	topStack := len(*stack) - 1
	for checkpointBackwards := topStack; checkpointBackwards >= 0; checkpointBackwards-- {
		/*
			it tries to reduce and if it got reduced it
			goes back to top and tries to reduce again
		*/
		if gotReduced(stack, checkpointBackwards) {
			str := string(checkpointBackwards)
			Test("reduced", &str)
			checkpointBackwards = topStack
		}
	}

}

func CheckResult(characters []string) bool {

	var stack []string

	// if it were string instead of slice it would return Runes
	for _, char := range characters {

		Test("shift", nil)
		stack = append(stack, char)

		goBackwardsToReduce(&stack)

	}

	resultingStack := SliceOfCharactersToString(stack)

	return resultingStack == "Expr"

}

package main

import "unicode"

var gotDigit = false

func getTopByteStack(stack *[]string) byte {

	topStack := PeekStringFromSlice(stack)
	topIsSingleElement := len(topStack) != 1

	// it's checked again for reusability sake of this function
	if topIsSingleElement {
		/*
			It throws an error and stops the execution.
			For simplicity's sake Panic function was used
			even though it's not recommended over Error Value.
		*/
		panic("The top of the stack is not a single character.")
	}

	return topStack[0]

}

/*
  For the sake of maintainability, I had established the
  convention of only handling Strings and Bytes. A character
  is Rune, hence there is not a "runeIsDigit" function.
*/
func byteIsDigit(byteCharacter byte) bool {
	return unicode.IsDigit(rune(byteCharacter))
}

func DigitWasReduced(slice *[]string, checkpointBackwardsIndex int) bool {

	topStackIndex := len(*slice) - 1

	if checkpointBackwardsIndex < topStackIndex {
		/*
			if there are digits before top, they had
			already been replaced, they were on top
		*/
		return false
	}

	topByteSlice := getTopByteStack(slice)
	topByteSliceIsDigit := byteIsDigit(topByteSlice)

	// number followed by another
	if topByteSliceIsDigit && gotDigit {
		// just "Expr" for the whole sequence
		PopStringsFromSlice(slice, topStackIndex)
		return true
	}

	// if it got there, it's the first digit for now
	if topByteSliceIsDigit {

		gotDigit = true

		PopStringsFromSlice(slice, topStackIndex)
		PushCharactersStringIntoSlice(slice, "Expr")

		return gotDigit
	}

	// if it got here, it's not a number
	gotDigit = false
	return gotDigit

}

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

	topElementIndex := len(*slice) - 1
	isSingleElement := checkpointBackwardsIndex == topElementIndex

	/*
	  based on the class' subject logic, if
	  backward's stack range is larger than
	  1, it has already been checked and it's only
	  need to do it oncesince there is no combination
	  between a number and other grammar categories
	   at least before the number had been replaced.
	*/
	if isSingleElement && gotDigit {
		return false
	}

	topByteSlice := getTopByteStack(slice)
	topByteSliceIsDigit := byteIsDigit(topByteSlice)

	// if it got there, it's another detected number
	if topByteSliceIsDigit && gotDigit {
		/*
			the respective grammar value is already
			in place, the number just get "poped"
		*/
		PopStringsFromSlice(slice, topElementIndex)
		return true
	}

	// if it got there, it's the first digit for now
	if topByteSliceIsDigit {

		gotDigit = true

		PopStringsFromSlice(slice, topElementIndex)
		PushCharactersStringIntoSlice(slice, "num")

		return true
	}

	// if it got there, any number was detected
	if gotDigit {
		gotDigit = false
	}

	return true

}

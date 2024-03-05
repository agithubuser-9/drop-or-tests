package main

func PeekStringFromSlice(slice *[]string) string {
	topIndex := len(*slice) - 1
	return (*slice)[topIndex]
}

func PeekStringsFromSlice(slice *[]string, checkpointBackwards int) []string {
	// everything before the checkpointBackwards
	// [ everything : checkpoint ]
	topStringsSlice := (*slice)[:checkpointBackwards]
	return topStringsSlice
}

func PopStringFromSlice(slice *[]string) {
	topIndex := len(*slice) - 1
	// [ everything : topIndex ]
	*slice = (*slice)[:topIndex]
}

func PopStringsFromSlice(slice *[]string, checkpointBackwards int) {
	// [ everything : checkpoint ]
	*slice = (*slice)[:checkpointBackwards]
}

func PushCharactersStringIntoSlice(slice *[]string, str string) {

	// a string is not made of characteres, but bytes that represent characteres
	// a slice is a mutable-sized array, preferable for further use
	for index := 0; index < len(str); index++ {
		*slice = append(*slice, string(str[index]))
	}

}

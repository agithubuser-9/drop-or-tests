package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// a very simple way to check what's happening
func Test(message string, optionalValue *string) {

	hasOptionalValue := optionalValue != nil

	if hasOptionalValue {
		fmt.Println(message, " at ", *optionalValue, "° time")
	} else {
		fmt.Println(message)
	}

}

func GetUserSingleCommand() string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("give me a command, mr. user:")

	line, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	line = strings.TrimSuffix(line, "\r\n")

	return line

}

func StringToSliceOfCharacters(str string) []string {

	var sliceOfCharacters []string

	for i := 0; i < len(str); i++ {

		characterString := string(str[i])
		sliceOfCharacters = append(sliceOfCharacters, characterString)

	}

	return sliceOfCharacters

}

// the function converts slice to string
func SliceOfCharactersToString(slice []string) string {
	str := strings.Join(slice, "")
	return str
}

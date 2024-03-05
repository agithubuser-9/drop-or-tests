package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserSingleCommand() string {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("give me a command, mr. user:")

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

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

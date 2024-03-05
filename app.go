package main

import "fmt"

func main() {

	command := GetUserSingleCommand()
	commandCharacters := StringToSliceOfCharacters(command)

	goodResult := CheckResult(commandCharacters)

	if goodResult {
		fmt.Println("Success!")
	} else {
		fmt.Println("Fail!")
	}

}

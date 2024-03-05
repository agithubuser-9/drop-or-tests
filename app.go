package main

import "fmt"

func main() {

	command := GetUserSingleCommand()
	commandCharacters := StringToSliceOfCharacters(command)

	goodResult := CheckResult(commandCharacters)

	if goodResult {
		fmt.Println("Accepted!")
	} else {
		fmt.Println("Rejected!")
	}

}

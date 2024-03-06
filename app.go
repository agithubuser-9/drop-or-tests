package main

// import "fmt"

func main() {

	command := GetUserSingleCommand()
	commandCharacters := StringToSliceOfCharacters(command)

	CheckSyntax(commandCharacters)
	

	/*
		correctSyntax := CheckSyntax(commandCharacters)

		if correctSyntax {
			// I got a warning writing fmt.Println("\nAccepted!\n")
			fmt.Println("\nAccepted!")
			fmt.Println()
		} else {
			fmt.Println("\nRejected!")
			fmt.Println()
		}
	*/

}

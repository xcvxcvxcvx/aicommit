package main

import (
	"fmt"
)

func main() {
	prompt, err := GetPrompt()
	diff := GenerateGit()

	if err != nil {
		fmt.Println(err)
	} else {
		message := prompt + "\nHere is the diff:\n" + diff

		response, error := GenerateCommitMessage(message)

		if error != nil {
			fmt.Println(error)
		} else {
			fmt.Println(response)
		}
	}
}

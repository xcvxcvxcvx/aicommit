package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: aicommit <commit-msg-file>")
		os.Exit(1)
	}

	source := ""

	if len(args) >= 3 {
		source = args[2]
	}

	if source != "" {
		os.Exit(0)
	}

	commitMsgFile := args[1]

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
			prependCommentToCommitMsg(commitMsgFile, response)
		}
	}
}

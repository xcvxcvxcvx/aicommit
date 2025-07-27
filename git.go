package main

import (
	"fmt"
	"os"
	"os/exec"
)

func GenerateGit() string {
	cmd := exec.Command("git", "diff", "--cached")
	out, err := cmd.Output()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when trying git diff: %v\n", err)
		os.Exit(1)
	}

	action := os.Getenv("GIT_REFLOG_ACTION")
	if action != "" && action != "commit" {
		os.Exit(0)
	}

	return string(out)
}

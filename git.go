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
		fmt.Fprintf(os.Stderr, "Ошибка при выполнении git diff: %v\n", err)
		os.Exit(1)
	}
	return string(out)
}

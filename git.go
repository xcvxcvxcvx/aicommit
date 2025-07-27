package main

import (
	"bytes"
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

	lines := bytes.Split(out, []byte{'\n'})
	var hasContent bool
	for _, line := range lines {
		trimmed := bytes.TrimSpace(line)
		if len(trimmed) > 0 && !bytes.HasPrefix(trimmed, []byte("#")) {
			hasContent = true
			break
		}
	}

	if !hasContent {
		os.Exit(0)
	}

	return string(out)
}

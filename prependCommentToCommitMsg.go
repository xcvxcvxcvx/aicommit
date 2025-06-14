package main

import (
	"os"
	"strings"
)

func prependCommentToCommitMsg(filePath, comment string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	lines := strings.Split(comment, "\n")
	for i, line := range lines {
		if line != "" {
			lines[i] = "# " + line
		}
	}
	commentWithHashes := strings.Join(lines, "\n") + "\n\n"

	newContent := []byte(commentWithHashes + string(data))
	return os.WriteFile(filePath, newContent, 0644)
}

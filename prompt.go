package main

import (
	"fmt"
	"os"
)

func GetPrompt() (string, error) {
	data, err := os.ReadFile("./prompt.txt")
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	return string(data), err
}

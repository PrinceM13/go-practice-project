package main

import (
	"errors"
	"fmt"
)

func main() {
	title, content, err := getNoteData()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Your note = ", title, content)
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note Title:")

	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	content, err := getUserInput("Note Content:")

	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	return title, content, nil
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var value string
	fmt.Scan(&value)

	if value == "" {
		return "", errors.New("input cannot be empty")
	}

	return value, nil
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/PrinceM13/go-practice-project/single-note/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title:")
	content := getUserInput("Note Content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n") // Remove the newline character
	text = strings.TrimSuffix(text, "\r") // Remove the carriage return character if present

	return text
}

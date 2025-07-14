package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/PrinceM13/go-practice-project/singlenote/customcalculator"
	"github.com/PrinceM13/go-practice-project/singlenote/note"
	"github.com/PrinceM13/go-practice-project/singlenote/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	intRes := customcalculator.Add(1, 2) + 1
	floatRes := customcalculator.Add(1.1, 2.2) + 0.5
	stringRes := customcalculator.Add("Hello, ", "World!") + "!"
	fmt.Println("Result of integer addition:", intRes)
	fmt.Println("Result of float addition:", floatRes)
	fmt.Println("Result of string concatenation:", stringRes)

	printSomething(1)
	printSomething(2.5)
	printSomething("Hello, World!")

	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	todoText := getUserInput("Todo Text:")
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	printSomething(todo)

	err = outputData(userNote)
	if err != nil {
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}
}

// any type is an alias for interface{}
func printSomething(value interface{}) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer value:", intVal)
		return
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float value:", floatVal)
		return
	}

	strVal, ok := value.(string)
	if ok {
		fmt.Println("String value:", strVal)
		return
	}

	fmt.Println("Unknown type:", value)

	// switch v := value.(type) {
	// case int:
	// 	fmt.Println("Integer value:", v)
	// case float64:
	// 	fmt.Println("Float value:", v)
	// case string:
	// 	fmt.Println("String value:", v)
	// default:
	// 	fmt.Println("Unknown type:", v)
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Error saving note:", err)
		return err
	}

	fmt.Println("Note saved successfully!")

	return nil
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

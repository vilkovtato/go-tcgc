package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vilkovtato/note/note"
	"github.com/vilkovtato/note/todo"
)

type Saver interface {
	Save() error
}

type outputable interface {
	Saver
	Display()
}

func main() {
	title, content, err := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		fmt.Println("Saving the todo failed.")
		return
	}

    outputData(userNote)
}

func printSometheing(value any) {

	switch value.(type) {
	case int:
		fmt.Println("It's an int:", value)
	case string:
		fmt.Println("It's a string", value)
	
	}
	
	fmt.Println(value)
}


func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}

func saveData(data Saver) error {
	return data.Save()
}

func getNoteData() (string, string, error) {

	title := getUserInput("Enter title: ")
	content := getUserInput("Enter content: ")

	return title, content, nil
}

func getUserInput(prompt string) string {

	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return ""
	}

	text = strings.TrimSuffix(text, "\n")

	return text
}

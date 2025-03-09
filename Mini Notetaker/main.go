package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text")

	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	todo, err1 := todo.New(todoText)
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	err = displayData(todo)
	if err != nil {
		return
	}

	displayData(note) // no need for err as program ends here anyways
}

func displayData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the data failed!")
		return err
	}
	fmt.Println("Saving the data Succeeded! ")
	return nil
}

func getNoteData() (string, string) {
	noteTitle := getUserInput("Note Title")
	noteContent := getUserInput("Note Content")

	return noteTitle, noteContent
}

func getUserInput(prompt string) (string) {
	fmt.Print(prompt + ": ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

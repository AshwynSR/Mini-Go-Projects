package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()
	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	note.Display()
	err = note.Save()

	if err != nil {
		fmt.Println("Saving the Note failed!")
		return
	}
	fmt.Println("Saving the Note Succeeded! ")
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

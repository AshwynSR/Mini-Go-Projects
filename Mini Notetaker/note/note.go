package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	CreatedTime time.Time `json:"created_time"`
}


func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

func (note *Note) Display() {
	fmt.Println("Title: ", note.Title)
	fmt.Println("Content: ", note.Content)
	fmt.Println("Created At: ", note.CreatedTime)
}

func New(noteTitle, noteContent string) (*Note, error ){
	if noteTitle == "" || noteContent == "" {
		return new(Note), errors.New("missing mandatory fields for note")
	}
	return &Note{Title: noteTitle, Content: noteContent, CreatedTime: time.Now()}, nil
}

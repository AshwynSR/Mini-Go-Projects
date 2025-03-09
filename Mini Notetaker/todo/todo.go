package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text	string `json:"text"`
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

func (todo *Todo) Display() {
	fmt.Println("Todo Content: ", todo.Text)
}

func New(todoContent string) (*Todo, error ){
	if todoContent == "" {
		return new(Todo), errors.New("missing mandatory fields for note")
	}
	return &Todo{Text: todoContent}, nil
}

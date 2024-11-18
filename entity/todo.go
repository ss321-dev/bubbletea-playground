package entity

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const todoFile = ".todo.json"

type Todo struct {
	Task      string `json:"task"`
	CreatedAt string `json:"created_at"`
}

func (t Todo) String() string {
	return t.Task
}

type Todos []Todo

func NewTodo(task string) Todo {
	return Todo{Task: task, CreatedAt: time.Now().Format(time.DateOnly)}
}

func SaveTodo(todos Todos) error {
	filePath, err := filepath.Abs(fmt.Sprintf("./%v", todoFile))
	if err != nil {
		return err
	}

	fileData, err := json.Marshal(todos)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, fileData, 0644)
}

func LoadTodo() (Todos, error) {
	filePath, err := filepath.Abs(fmt.Sprintf("./%v", todoFile))
	if err != nil {
		return nil, err
	}

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return Todos{}, nil
		}
		return nil, err
	}

	var todos Todos
	if err := json.Unmarshal(fileData, &todos); err != nil {
		return nil, err
	}
	return todos, nil
}

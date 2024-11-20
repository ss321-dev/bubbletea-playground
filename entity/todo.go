package entity

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

const todoFileName = ".todo.json"

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

func SaveTodo(path string, todos Todos) error {
	fileData, err := json.Marshal(todos)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(path, todoFileName), fileData, 0644)
}

func LoadTodo(path string) (Todos, error) {
	fileData, err := os.ReadFile(filepath.Join(path, todoFileName))
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

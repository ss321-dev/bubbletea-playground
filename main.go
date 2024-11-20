package main

import (
	"log"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/ss321-dev/bubbletea-playground/entity"
	"github.com/ss321-dev/bubbletea-playground/model"
	"github.com/ss321-dev/bubbletea-playground/view"
)

func main() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	path := filepath.Dir(filepath.Clean(exePath))

	todos, err := entity.LoadTodo(path)
	if err != nil {
		log.Fatal(err)
	}

	viewState := view.Menu
	if len(os.Args) > 1 {
		viewState = view.NewState(os.Args[1])
	}

	state := model.NewState(todos, viewState)
	defer func() {
		if err := entity.SaveTodo(path, state.Todos); err != nil {
			log.Fatal(err)
		}
	}()

	teaProgram := tea.NewProgram(state)
	state.SetExitFunc(func() {
		teaProgram.Send(tea.KeyMsg{Type: model.ExitMessage})
	})
	if _, err := teaProgram.Run(); err != nil {
		log.Fatal(err)
	}
}

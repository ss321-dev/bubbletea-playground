package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/ss321-dev/bubbletea-playground/entity"
	"github.com/ss321-dev/bubbletea-playground/model"
	"github.com/ss321-dev/bubbletea-playground/view"
)

func main() {
	todos, err := entity.LoadTodo()
	if err != nil {
		log.Fatal(err)
	}

	viewState := view.Menu
	if len(os.Args) > 1 {
		viewState = view.NewState(os.Args[1])
	}

	state := model.NewState(todos, viewState)
	defer func() {
		if err := entity.SaveTodo(state.Todos); err != nil {
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

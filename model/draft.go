package model

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/ss321-dev/bubbletea-playground/entity"
	"github.com/ss321-dev/bubbletea-playground/funcs"
)

type draft struct {
	index *int
	input textinput.Model
}

func (d *draft) Init(index *int, todos entity.Todos) {
	textInputModel := textinput.New()
	if index != nil {
		textInputModel.SetValue(todos[*index].Task)
	}
	textInputModel.Focus()
	textInputModel.Placeholder = "please input"
	d.input = textInputModel
	d.index = index
}

func (d *draft) SetMsg(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	d.input, cmd = d.input.Update(msg)
	return cmd
}

func (d *draft) View() string {
	return d.input.View()
}

func (d *draft) ToTodo() entity.Todo {
	return entity.NewTodo(d.input.Value())
}

func (d *draft) Index() (int, bool) {
	return funcs.ToValue(d.index), d.index != nil
}

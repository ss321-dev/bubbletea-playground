package model

import (
	"slices"
	"sync"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/ss321-dev/bubbletea-playground/entity"
	"github.com/ss321-dev/bubbletea-playground/funcs"
	"github.com/ss321-dev/bubbletea-playground/option"
	"github.com/ss321-dev/bubbletea-playground/view"
)

var exitFunc sync.Once

const ExitMessage tea.KeyType = 9999

type State struct {
	Todos     entity.Todos
	ViewState view.State
	Draft     draft
	Cursor    int
	ExitFunc  func()
	Err       error
}

func NewState(todo entity.Todos, viewState view.State) *State {
	return &State{Todos: todo, ViewState: viewState, Cursor: 0}
}

func (m *State) SetExitFunc(fn func()) tea.Cmd {
	m.ExitFunc = fn
	return nil
}

func (m *State) Exit() {
	exitFunc.Do(m.ExitFunc)
}

func (m *State) Init() tea.Cmd {
	m.Draft.Init(nil, nil)
	return nil
}

func (m *State) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	// 出力して終了するコマンドはここに記載する
	if keyMsg.Type == ExitMessage {
		switch m.ViewState {
		case view.List:
			return m, tea.Sequence(tea.Printf(view.ListView(m.Todos)), tea.Quit)
		case view.Count:
			return m, tea.Sequence(tea.Printf(view.CountView(len(m.Todos))), tea.Quit)
		case view.Done:
			return m, tea.Sequence(tea.Printf(view.DoneView()), tea.Quit)
		}
	}

	switch m.ViewState {
	case view.Menu:
		return m.handleMenu(keyMsg.Type)
	case view.Input:
		return m.handleInput(msg, keyMsg.Type)
	case view.Update:
		return m.handleUpdate(keyMsg.Type)
	case view.Delete:
		return m.handleDelete(keyMsg.Type)
	}

	if slices.Contains([]tea.KeyType{tea.KeyCtrlC, tea.KeyEsc}, keyMsg.Type) {
		return m, tea.Quit
	}
	return m, nil
}

func (m *State) View() string {
	switch m.ViewState {
	case view.List, view.Count, view.Done:
		go m.Exit() // 出力して終了するコマンドはUpdate()に記載する
	case view.Input:
		return view.InputView(m.Draft.View())
	case view.Update:
		return view.UpdateView(m.Cursor, m.Todos)
	case view.Delete:
		return view.DeleteView(m.Cursor, m.Todos)
	case view.Menu:
		return view.MenuView(m.Cursor)
	}
	return ""
}

func (m *State) handleMenu(key tea.KeyType) (tea.Model, tea.Cmd) {
	switch key {
	case tea.KeyCtrlC, tea.KeyEsc:
		return m, tea.Quit
	case tea.KeyEnter:
		m.ViewState = view.NewState(option.ActionOptions[m.Cursor])
	case tea.KeyDown:
		if m.Cursor+1 <= len(option.ActionOptions)-1 {
			m.Cursor++
		}
	case tea.KeyUp:
		if m.Cursor-1 >= 0 {
			m.Cursor--
		}
	}
	return m, nil
}

func (m *State) handleDelete(key tea.KeyType) (tea.Model, tea.Cmd) {
	if len(m.Todos) == 0 {
		return m, tea.Sequence(tea.Println("todo empty."), tea.Quit)
	}

	switch key {
	case tea.KeyCtrlC, tea.KeyEsc:
		return m, tea.Quit
	case tea.KeyEnter:
		m.ViewState = view.Done
		m.Todos = funcs.Remove(m.Todos, m.Cursor)
		return m, nil
	case tea.KeyDown:
		if m.Cursor+1 <= len(m.Todos)-1 {
			m.Cursor++
		}
	case tea.KeyUp:
		if m.Cursor-1 >= 0 {
			m.Cursor--
		}
	}
	return m, nil
}

func (m *State) handleUpdate(key tea.KeyType) (tea.Model, tea.Cmd) {
	if len(m.Todos) == 0 {
		return m, tea.Sequence(tea.Println("todo empty."), tea.Quit)
	}

	switch key {
	case tea.KeyCtrlC, tea.KeyEsc:
		return m, tea.Quit
	case tea.KeyEnter:
		m.Draft.Init(funcs.ToPtr(m.Cursor), m.Todos)
		m.ViewState = view.Input
	case tea.KeyDown:
		if m.Cursor+1 <= len(m.Todos)-1 {
			m.Cursor++
		}
	case tea.KeyUp:
		if m.Cursor-1 >= 0 {
			m.Cursor--
		}
	}
	return m, nil
}

func (m *State) handleInput(msg tea.Msg, key tea.KeyType) (tea.Model, tea.Cmd) {
	switch key {
	case tea.KeyCtrlC, tea.KeyEsc:
		return m, tea.Quit
	case tea.KeyEnter:
		m.ViewState = view.Done
		index, ok := m.Draft.Index()
		if ok {
			m.Todos[index] = m.Draft.ToTodo()
		} else {
			m.Todos = append(m.Todos, m.Draft.ToTodo())
		}
		return m, nil
	}
	return m, m.Draft.SetMsg(msg)
}

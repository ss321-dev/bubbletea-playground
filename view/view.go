package view

import (
	"fmt"
	"strconv"
	"strings"
	
	"github.com/ss321-dev/bubbletea-playground/entity"
	"github.com/ss321-dev/bubbletea-playground/option"
)

const pleaseExitText = "(ctrl+c to quit)"

type State string

const (
	List   State = "LIST_VIEW"
	Input  State = "INPUT_VIEW"
	Update State = "UPDATE_VIEW"
	Delete State = "DELETE_VIEW"
	Count  State = "COUNT_VIEW"
	Menu   State = "MENU_VIEW"
	Done   State = "DONE_VIEW"
)

func NewState(arg string) State {
	switch arg {
	case option.List:
		return List
	case option.Add:
		return Input
	case option.Update:
		return Update
	case option.Delete:
		return Delete
	case option.Count:
		return Count
	default:
		return Menu
	}
}

func ListView(todos entity.Todos) string {
	builder := strings.Builder{}
	for i, todo := range todos {
		builder.WriteString(fmt.Sprintf("#%d: %s", i+1, fmt.Sprint(todo)))
		if len(todos) != (i + 1) {
			builder.WriteString("\n")
		}
	}
	return builder.String()
}

func InputView(textView string) string {
	return fmt.Sprintf("%s\n\n%s\n", textView, pleaseExitText)
}

func UpdateView(current int, todos entity.Todos) string {
	return viewList(current, todos)
}

func DeleteView(current int, todos entity.Todos) string {
	return viewList(current, todos)
}

func CountView(count int) string {
	return strconv.Itoa(count)
}

func MenuView(current int) string {
	return viewList(current, option.ActionOptions)
}

func DoneView() string {
	return "completed."
}

func viewList[T any](current int, collection []T) string {
	builder := strings.Builder{}
	for i, v := range collection {
		cursor := " "
		if current == i {
			cursor = ">"
		}
		builder.WriteString(fmt.Sprintf("%s #%d: %s\n", cursor, i+1, fmt.Sprint(v)))
	}
	builder.WriteString(fmt.Sprintf("\n%s", pleaseExitText))
	return builder.String()
}

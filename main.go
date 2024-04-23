package main

import (
	go_edit "sebmentation-fault/go-edit/pkg"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	gfile go_edit.GoEditFile
	fmode go_edit.GoEditMode
}

func initialModel() model {
	return nil
}

func (m model) Init() tea.Cmd {
	return
}

func (m model) View() string {
	return ""
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return
}


func main() {
	program = tea.NewProgram(initialModel())
}

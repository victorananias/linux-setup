package main

import tea "github.com/charmbracelet/bubbletea"

type PasswordModel struct{}

func (m *PasswordModel) Init() tea.Cmd {
	return nil
}

func (m *PasswordModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *PasswordModel) View() string {
	return "Enter your password \n"
}

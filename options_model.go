package main

import tea "github.com/charmbracelet/bubbletea"

type OptionsModel struct{}

func (m *OptionsModel) Init() tea.Cmd {
	return nil
}

func (m *OptionsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *OptionsModel) View() string {
	return ""
}

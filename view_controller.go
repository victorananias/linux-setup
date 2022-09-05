package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/victorananias/setup-linux/views"
)

type model struct {
	view    int
	aptView *views.AptView
}

func InitialModel() model {
	return model{
		view:    1,
		aptView: views.NewAptView(),
	}
}

func (m model) Init() tea.Cmd {
	switch m.view {
	case 1:
		return m.aptView.Init()
	}
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.view {
	case 1:
		return m.aptView.Update(m)
	}
	return nil, nil
}

func (m model) View() string {
	switch m.view {
	case 1:
		return m.aptView.View()
	}
	return ""
}

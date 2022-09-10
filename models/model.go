package models

import (
	"fmt"
	"sort"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/victorananias/setup-linux/packages"
)

var mainModel *model

type model struct {
	choices []string
	cursor  int
}

var CurrentModel string = APT_MODEL

func StartModels() *model {
	choices := packages.GetAptPackages()
	sort.Strings(choices)
	selectedAptPackages = func() map[int]struct{} {
		m := make(map[int]struct{}, 0)
		for i := 0; i < len(choices); i++ {
			m[i] = struct{}{}
		}
		return m
	}()
	return &model{
		choices: choices,
	}
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyCtrlC.String(), "q":
			return m, tea.Quit
		case tea.KeyUp.String(), "k":
			m.CursorUp()
		case tea.KeyDown.String(), "j":
			m.CursorDown()
		case tea.KeySpace.String():
			_, ok := selectedAptPackages[m.cursor]
			if ok {
				delete(selectedAptPackages, m.cursor)
			} else {
				selectedAptPackages[m.cursor] = struct{}{}
			}
		case tea.KeyEnter.String():
			if m.cursor == len(m.choices) {

			} else {
				m.cursor = len(m.choices)
			}
		}
	}

	return m, nil
}

func (m *model) View() string {
	var cursorStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#b728b5"))
	s := "What apt packages should we install?\n\n"

	for i, choice := range m.choices {
		checked := "[ ]"
		if _, ok := selectedAptPackages[i]; ok {
			checked = "[x]"
		}
		if m.cursor == i {
			choice = cursorStyle.Render(choice)
			checked = cursorStyle.Render(checked)
		}
		s += fmt.Sprintf("  %s %s (%d)\n", checked, choice, m.cursor)
	}

	if m.cursor == len(m.choices) {
		s += cursorStyle.Render("\n  Install selected\n")
	} else {
		s += "\n  Install selected\n"
	}

	s += "\nPress q to quit\nPress space to select \nPress enter to install\n"

	return s
}

func (m *model) CursorDown() {
	if m.cursor < len(m.choices) {
		m.cursor++
	} else if m.cursor == len(m.choices) {
		m.cursor = 0
	}
}

func (m *model) CursorUp() {
	if m.cursor > 0 {
		m.cursor--
	} else if m.cursor == 0 {
		m.cursor = len(m.choices)
	}
}

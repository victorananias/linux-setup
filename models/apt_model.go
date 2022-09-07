package models

import (
	"fmt"
	"sort"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/victorananias/setup-linux/packages"
)

var selectedAptPackages map[int]struct{}

var APT_MODEL = "AptModel"

type AptModel struct {
	choices []string
	cursor  int
}

func NewAptModel() *AptModel {
	choices := packages.GetAptPackages()
	sort.Strings(choices)
	selectedAptPackages = func() map[int]struct{} {
		m := make(map[int]struct{}, 0)
		for i := 0; i < len(choices); i++ {
			m[i] = struct{}{}
		}
		return m
	}()
	return &AptModel{
		choices: choices,
	}
}

func (m *AptModel) ChangeCurrentModel() {
	CurrentModel = INSTALL_APT_PACKAGES_MODEL
}

func (m *AptModel) Init() tea.Cmd {
	return nil
}

func (m *AptModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			} else if m.cursor == 0 {
				m.cursor = len(m.choices)
			}
		case "down", "j":
			if m.cursor < len(m.choices) {
				m.cursor++
			} else if m.cursor == len(m.choices) {
				m.cursor = 0
			}
		case " ":
			_, ok := selectedAptPackages[m.cursor]
			if ok {
				delete(selectedAptPackages, m.cursor)
			} else {
				selectedAptPackages[m.cursor] = struct{}{}
			}
		case "enter":
			if m.cursor == len(m.choices) {
				tea.Model = NewInstallAptPackagesModel()
			} else {
				m.cursor = len(m.choices)
			}
		}
	}

	return m, nil
}

func (m *AptModel) View() string {
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

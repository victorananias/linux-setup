package views

import (
	"fmt"
	"sort"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/victorananias/setup-linux/packages"
)

type AptView struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func NewAptView() *AptView {
	choices := packages.GetAptPackages()
	sort.Strings(choices)
	selected := func() map[int]struct{} {
		m := make(map[int]struct{}, 0)
		for i := 0; i < len(choices); i++ {
			m[i] = struct{}{}
		}
		return m
	}()
	return &AptView{
		choices:  choices,
		selected: selected,
	}
}

func (view *AptView) Init() tea.Cmd {
	return nil
}

func (view *AptView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return view, tea.Quit
		case "up", "k":
			if view.cursor > 0 {
				view.cursor--
			} else if view.cursor == 0 {
				view.cursor = len(view.choices)
			}
		case "down", "j":
			if view.cursor < len(view.choices) {
				view.cursor++
			} else if view.cursor == len(view.choices) {
				view.cursor = 0
			}
		case " ":
			_, ok := view.selected[view.cursor]
			if ok {
				delete(view.selected, view.cursor)
			} else {
				view.selected[view.cursor] = struct{}{}
			}
		case "enter":
			if view.cursor == len(view.choices) {

			} else {
				view.cursor = len(view.choices)
			}
		}
	}

	return view, nil
}

func (m *AptView) View() string {
	var cursorStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#b728b5"))
	s := "What apt packages should we install?\n\n"

	for i, choice := range m.choices {
		checked := "[ ]"
		if _, ok := m.selected[i]; ok {
			checked = "[x]"
		}
		if m.cursor == i {
			choice = cursorStyle.Render(choice)
			checked = cursorStyle.Render(checked)
		}
		s += fmt.Sprintf("  %s %s\n", checked, choice)
	}

	if m.cursor == len(m.choices) {
		s += cursorStyle.Render("\n  Install selected\n")
	} else {
		s += "\n  Install selected\n"
	}

	s += "\nPress q to quit\nPress space to select \nPress enter to install\n"

	return s
}

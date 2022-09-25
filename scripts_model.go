package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ScriptsModel struct {
	choices       []*Script
	cursor        int
	selected      map[int]struct{}
	dir           string
	installButton int
}

type RunSelected struct {
	Selected []*Script
}

func NewScriptsModel(dir string) *ScriptsModel {
	scripts := GetScriptsFrom(dir)
	scripts = sortAlphabetically(scripts)
	return &ScriptsModel{
		choices:       scripts,
		selected:      make(map[int]struct{}),
		dir:           dir,
		installButton: len(scripts),
	}
}

func GetScriptsFrom(dir string) []*Script {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	scripts := []*Script{}

	for _, file := range files {
		scripts = append(scripts, NewScript(file.Name(), dir))
	}

	return scripts
}

func (m *ScriptsModel) Init() tea.Cmd {
	return nil
}

func (m *ScriptsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.cursor >= len(m.choices) {
				break
			}
			isSelected := m.isSelected(m.cursor)
			if isSelected {
				m.deselectPackage(m.cursor)
			} else {
				m.selectPackage(m.cursor)
			}
		case tea.KeyEnter.String():
			if m.cursor == m.installButton {
				return m, func() tea.Msg {
					return RunSelected(RunSelected{Selected: m.SelectedScripts()})
				}
			} else {
				m.cursor = m.installButton
			}
		}
	}

	return m, nil
}

func (m *ScriptsModel) View() string {
	var cursorStyle = m.GetCursorStyle()
	v := "\n\nWhat packages should be installed?\n\n"

	for i, choice := range m.choices {
		title := choice.name
		checked := "[ ]"
		if m.isSelected(i) {
			checked = "[x]"
		}
		if m.cursor == i {
			title = cursorStyle.Render(title)
			checked = cursorStyle.Render(checked)
		}
		v += fmt.Sprintf("  %s %s\n", checked, title)
	}

	if m.cursor == len(m.choices) {
		v += cursorStyle.Render("\n  Install selected\n")
	} else {
		v += "\n  Install selected\n"
	}

	v += "\nPress q to quit\nPress space to select \nPress enter to install\n"

	return v
}

func (m *ScriptsModel) GetCursorStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#b728b5"))
}

func (m *ScriptsModel) isSelected(element int) bool {
	if _, ok := m.selected[element]; ok {
		return true
	}
	return false
}

func (m *ScriptsModel) deselectPackage(element int) {
	delete(m.selected, element)
}

func (m *ScriptsModel) selectPackage(element int) {
	m.selected[m.cursor] = struct{}{}
}

func (m *ScriptsModel) SelectedScripts() []*Script {
	scripts := make([]*Script, 0, len(m.selected))
	for k := range m.selected {
		scripts = append(scripts, m.choices[k])
	}
	return scripts
}

func (m *ScriptsModel) CursorDown() {
	if m.cursor < len(m.choices) {
		m.cursor++
	} else if m.cursor == len(m.choices) {
		m.cursor = 0
	}
}

func (m *ScriptsModel) CursorUp() {
	if m.cursor > 0 {
		m.cursor--
	} else if m.cursor == 0 {
		m.cursor = len(m.choices)
	}
}

func sortAlphabetically(list []*Script) []*Script {
	sort.Slice(list, func(i, j int) bool {
		return list[i].name < list[j].name
	})
	return list
}

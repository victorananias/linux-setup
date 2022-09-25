package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type sessionState int

const (
	scriptsView sessionState = iota
	runScriptsView
)

type MainModel struct {
	state      sessionState
	scripts    tea.Model
	runScripts tea.Model
	selected   []*Script
	dir        string
}

var optionDir = "/Install"

func NewModel() *MainModel {
	return &MainModel{
		state:   scriptsView,
		scripts: NewScriptsModel(dirs.OPTIONS + optionDir),
		dir:     dirs.OPTIONS + optionDir,
	}
}

func (m *MainModel) Init() tea.Cmd {
	return nil
}

func (m *MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case RunSelected:
		m.selected = msg.Selected
		m.state = runScriptsView
		action := m.selected[0].action
		m.runScripts = NewRunningScriptsModel(action, m.selected)
		cmd := m.runScripts.Init()
		cmds = append(cmds, cmd)
	}

	switch m.state {
	case scriptsView:
		newScripts, newCmd := m.scripts.Update(msg)
		scriptsModel, ok := newScripts.(*ScriptsModel)
		if !ok {
			panic("could not perform assertion on scripts model")
		}
		m.scripts = scriptsModel
		cmd = newCmd
	case runScriptsView:
		newRunScripts, newCmd := m.runScripts.Update(msg)
		runScriptsModel, ok := newRunScripts.(RunningScriptsModel)
		if !ok {
			panic("could not perform assertion on scripts model")
		}
		m.runScripts = runScriptsModel
		cmd = newCmd
	}
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m *MainModel) View() string {
	switch m.state {
	case scriptsView:
		return m.scripts.View()
	case runScriptsView:
		return m.runScripts.View()
	}
	return ""
}

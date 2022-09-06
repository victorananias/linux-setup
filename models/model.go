package views

import (
	tea "github.com/charmbracelet/bubbletea"
)

var mainModel *model

type model struct {
	CurrentModel string
}

var models = map[string]tea.Model{
	APT_MODEL:                  NewAptModel(),
	INSTALL_APT_PACKAGES_MODEL: NewInstallAptPackagesModel(),
}

func StartModels() model {
	mainModel = &model{
		CurrentModel: APT_MODEL,
	}
	return *mainModel
}

func (m model) Init() tea.Cmd {
	return models[m.CurrentModel].Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return models[m.CurrentModel].Update(m)
}

func (m model) View() string {
	return models[m.CurrentModel].View()
}

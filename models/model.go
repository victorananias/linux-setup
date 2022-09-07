package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

var mainModel *model

type model struct {
}

var CurrentModel string = APT_MODEL

var models *map[string]tea.Model = &map[string]tea.Model{
	APT_MODEL:                  NewAptModel(),
	INSTALL_APT_PACKAGES_MODEL: NewInstallAptPackagesModel(),
}

func StartModels() *model {
	mainModel = &model{}
	return mainModel
}

func (m *model) Init() tea.Cmd {
	return (*models)[CurrentModel].Init()
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return (*models)[CurrentModel].Update(m)
}

func (m *model) View() string {
	return (*models)[CurrentModel].View()
}

package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

var INSTALL_APT_PACKAGES_MODEL = "InstallAptPackagesModel"

func NewInstallAptPackagesModel() *InstallAptPackagesModel {
	return &InstallAptPackagesModel{}
}

type InstallAptPackagesModel struct{}

func (view *InstallAptPackagesModel) Init() tea.Cmd {
	return nil
}

func (view *InstallAptPackagesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return view, nil
}

func (m *InstallAptPackagesModel) View() string {
	s := "Installing apt packages\n\n"
	for _, pkg := range selectedAptPackages {
		s += fmt.Sprintf("%s\n", pkg)
	}
	return s
}

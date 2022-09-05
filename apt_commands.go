package main

import (
	"fmt"
)

func InstallAptPackages(packages []string) {
	for _, apt_package := range packages {
		ExecAptInstall(apt_package)
	}
}

func ExecAptInstall(command string) {
	ExecApt(fmt.Sprintf("install %s", command))
}

func UpdateSystem() {
	ExecApt("update -y")
}

func UpgradeAptPackages() {
	ExecApt("upgrade -y")
}

func ExecApt(command string) {
	ExecSudo(fmt.Sprintf("apt %s", command))
}

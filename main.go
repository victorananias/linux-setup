package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	OS_READ  = 04
	OS_WRITE = 02
	OS_EX    = 01
)

func main() {
	prepareScripts()
	start()
}

func start() {
	p := tea.NewProgram(NewModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func prepareScripts() {
	createTempDirectory()
	createTempOptionsDirectory()
}

func createTempOptionsDirectory() {
	directories, err := os.ReadDir(dirs.OPTIONS)
	if err != nil {
		panic("failed to read " + dirs.OPTIONS)
	}
	for _, directory := range directories {
		if !directory.IsDir() {
			continue
		}
		directoryPath := path.Join(dirs.OPTIONS, directory.Name())

		files, err := os.ReadDir(directoryPath)
		if err != nil {
			panic(fmt.Sprintf("failed to read %s/ files", dirs.OPTIONS))
		}
		for _, dirEntry := range files {
			filePath := path.Join(directoryPath, dirEntry.Name())
			contentAsBytes, err := os.ReadFile(filePath)
			if err != nil {
				panic(fmt.Sprintf("failed to read %s", filePath))
			}
			contentAsText := string(contentAsBytes)
			contentAsText = strings.ReplaceAll(contentAsText, "sudo", `echo "$SUDO_PASSWORD" | sudo -S`)
			contentAsBytes = []byte(contentAsText)

			err = ioutil.WriteFile(filePath, contentAsBytes, 0)
			if err != nil {
				panic(fmt.Sprintf("failed to write %s", filePath))
			}
		}
	}
}

func createTempDirectory() {
	_, err := os.ReadDir(dirs.TMP)
	if err == nil {
		err = os.RemoveAll(dirs.TMP)
		if err != nil {
			panic(err.Error())
		}
	}
	err = CopyDirectory(dirs.ORINAL_OPTIONS, dirs.OPTIONS)
	if err != nil {
		panic(fmt.Sprintf("failed to create %s", dirs.TMP))
	}
}

package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/victorananias/setup-linux/models"
)

func main() {
	p := tea.NewProgram(models.StartModels())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	// models.CurrentModel = models.APT_MODEL
	// print(models.CurrentModel + "\n")
	// aptModel := models.NewAptModel()
	// aptModel.ChangeCurrentModel()
	// print(models.CurrentModel + "\n")
}

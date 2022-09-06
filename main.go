package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	views "github.com/victorananias/setup-linux/models"
)

func main() {
	p := tea.NewProgram(views.StartModels())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

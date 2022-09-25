package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type RunningScriptsModel struct {
	scripts  []*Script
	index    int
	width    int
	height   int
	spinner  spinner.Model
	progress progress.Model
	done     bool
	action   string
}

type scriptExecuted *Script

var (
	currentScriptNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	subtleStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color("239"))
	doneStyle              = lipgloss.NewStyle().Margin(1, 2)
	checkMark              = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
)

func NewRunningScriptsModel(action string, scripts []*Script) RunningScriptsModel {
	p := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(40),
		progress.WithoutPercentage(),
	)
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	return RunningScriptsModel{
		scripts:  scripts,
		spinner:  s,
		progress: p,
		action:   action,
	}
}

func (m RunningScriptsModel) Init() tea.Cmd {
	return tea.Batch(m.executeScript(m.scripts[m.index]), m.spinner.Tick, m.successMessage())
}

func (m RunningScriptsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		}
	case scriptExecuted:
		if m.index >= len(m.scripts)-1 {
			// Everything's been installed. We're done!
			m.done = true
			return m, tea.Quit
		}

		// Update progress bar
		progressCmd := m.progress.SetPercent(float64(m.index) / float64(len(m.scripts)-1))

		m.index++
		return m, tea.Batch(
			progressCmd,
			m.successMessage(),                  // print success message above our program
			m.executeScript(m.scripts[m.index]), // download the next package
		)
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case progress.FrameMsg:
		newModel, cmd := m.progress.Update(msg)
		if newModel, ok := newModel.(progress.Model); ok {
			m.progress = newModel
		}
		return m, cmd
	}
	return m, nil
}

func (m RunningScriptsModel) View() string {
	n := len(m.scripts)
	w := lipgloss.Width(fmt.Sprintf("%d", n))

	if m.done {
		return doneStyle.Render(fmt.Sprintf("Done! "+m.action+"ed %d.\n", n))
	}

	pkgCount := fmt.Sprintf(" %*d/%*d", w, m.index, w, n-1)

	spin := m.spinner.View() + " "
	prog := m.progress.View()
	cellsAvail := max(0, m.width-lipgloss.Width(spin+prog+pkgCount))

	action := m.scripts[m.index].action
	scriptName := currentScriptNameStyle.Render(m.scripts[m.index].name)
	info := lipgloss.NewStyle().MaxWidth(cellsAvail).Render(action + "ing " + scriptName)

	cellsRemaining := max(0, m.width-lipgloss.Width(spin+info+prog+pkgCount))
	gap := strings.Repeat(" ", cellsRemaining)

	return spin + info + gap + prog + pkgCount
}

func (m *RunningScriptsModel) successMessage() tea.Cmd {
	return tea.Printf("%s %s", checkMark, m.scripts[m.index].name)
}

func (m *RunningScriptsModel) executeScript(script *Script) tea.Cmd {
	return func() tea.Msg {
		commands.Exec(script.path)
		return scriptExecuted(script)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

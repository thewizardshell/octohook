package styles

import "github.com/charmbracelet/lipgloss"

var (
	Title = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205"))

	Muted = lipgloss.NewStyle().
		Foreground(lipgloss.Color("241"))
)

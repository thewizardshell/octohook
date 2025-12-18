package styles

import "github.com/charmbracelet/lipgloss"

var (
	Purple = lipgloss.Color("99")
	Orange = lipgloss.Color("208")
	Gray   = lipgloss.Color("241")
	Green  = lipgloss.Color("82")
	Red    = lipgloss.Color("196")

	Title = lipgloss.NewStyle().
		Bold(true).
		Foreground(Purple)

	Box = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(Purple).
		Padding(1, 2)

	Ok = lipgloss.NewStyle().
		Foreground(Green).
		Bold(true)

	Fail = lipgloss.NewStyle().
		Foreground(Red).
		Bold(true)

	Running = lipgloss.NewStyle().
		Foreground(Orange)

	Muted = lipgloss.NewStyle().
		Foreground(Gray)
)

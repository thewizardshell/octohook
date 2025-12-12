package update

import (
	"octohook/internal/model"

	tea "github.com/charmbracelet/bubbletea"
)

func Update(msg tea.Msg, m *model.Model) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return tea.Quit
		}
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
	}
	return nil
}

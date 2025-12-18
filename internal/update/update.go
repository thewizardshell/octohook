package update

import (
	"octohook/internal/model"

	tea "github.com/charmbracelet/bubbletea"
)

// Init initializes the model and returns initial commands
func Init(m *model.Model) tea.Cmd {
	if m.HookCmd != nil {
		return tea.Batch(
			m.Spinner.Tick,
			m.HookCmd,
		)
	}
	return m.Spinner.Tick
}

// Update handles all message updates for the model
func Update(msg tea.Msg, m *model.Model) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return tea.Quit
		}

	case model.TestUpdate:
		found := false
		for i, t := range m.Tests {
			if t.Name == msg.Name {
				m.Tests[i].Status = msg.Status
				m.Tests[i].Output = msg.Output
				found = true
				break
			}
		}
		if !found {
			m.Tests = append(m.Tests, model.TestState{
				Name:   msg.Name,
				Status: msg.Status,
				Output: msg.Output,
			})
		}
		if msg.Status == model.StatusFail {
			m.Failed = true
		}
		return model.WaitForUpdate(m.Updates)

	case model.HookFinished:
		m.Done = true
		m.Running = false
		return tea.Quit

	default:
		if !m.Done {
			var cmd tea.Cmd
			m.Spinner, cmd = m.Spinner.Update(msg)
			return cmd
		}
	}
	return nil
}

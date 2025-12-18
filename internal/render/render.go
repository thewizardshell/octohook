package render

import (
	"octohook/internal/model"
	"octohook/internal/update"
	"octohook/internal/view"

	tea "github.com/charmbracelet/bubbletea"
)

// App wraps a Model and implements tea.Model interface
type App struct {
	Model *model.Model
}

// Init implements tea.Model
func (a *App) Init() tea.Cmd {
	return update.Init(a.Model)
}

// Update implements tea.Model
func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := update.Update(msg, a.Model)
	return a, cmd
}

// View implements tea.Model
func (a *App) View() string {
	return view.Render(a.Model)
}

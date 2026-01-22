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

// ResumeApp is a separate Bubble Tea app for showing the Resume
type ResumeApp struct {
	Model *model.Model
}

func NewResumeApp(m *model.Model) *ResumeApp {
	return &ResumeApp{Model: m}
}

// Init implements tea.Model
func (a *ResumeApp) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model
func (a *ResumeApp) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" || msg.String() == "enter" {
			return a, tea.Quit
		}
	}
	return a, nil
}

// View implements tea.Model
func (a *ResumeApp) View() string {
	return view.RenderResume(a.Model)
}

package main

import (
	"fmt"
	"octohook/internal/model"
	"octohook/internal/update"
	"octohook/internal/view"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type OctohookApp struct {
	model *model.Model
}

func (o OctohookApp) Init() tea.Cmd {
	return nil
}

func (o OctohookApp) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := update.Update(msg, o.model)
	return o, cmd
}

func (o OctohookApp) View() string {
	return view.Render(o.model)
}
func main() {
	p := tea.NewProgram(OctohookApp{model: model.NewModel()})
	if err := p.Start(); err != nil {
		fmt.Printf("Error starting program: %s\n", err)
		os.Exit(1)
	}

}

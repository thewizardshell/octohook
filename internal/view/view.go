package view

import (
	"octohook/internal/model"
	"octohook/internal/styles"
)

func Render(m *model.Model) string {
	s := styles.Title.Render("Octohook 🐙\n")
	s += styles.Muted.Render("A terminal app for Git Hooks\n")
	return s
}

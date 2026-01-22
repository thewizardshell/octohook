package view

import (
	"fmt"
	"octohook/internal/model"
	"octohook/internal/styles"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	badgeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("0")).
			Background(styles.Orange).
			Padding(0, 1).
			Bold(true)

	dividerStyle = lipgloss.NewStyle().
			Foreground(styles.Gray)

	hookNameStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("255"))

	footerStyle = lipgloss.NewStyle().
			Foreground(styles.Gray).
			Italic(true)
)

func Render(m *model.Model) string {
	var sections []string

	logo := styles.Title.Render("◉ octohook")
	badge := badgeStyle.Render(m.HookName)

	header := lipgloss.JoinHorizontal(
		lipgloss.Center,
		logo,
		"  ",
		badge,
	)
	sections = append(sections, header)

	divider := dividerStyle.Render(strings.Repeat("─", 44))
	sections = append(sections, divider)

	if len(m.Tests) == 0 {
		empty := styles.Muted.Italic(true).Render("  No hooks running...")
		sections = append(sections, empty)
	} else {
		stats := renderStats(m.Tests)
		sections = append(sections, stats)
		sections = append(sections, "")

		// Hook list
		for i, hook := range m.Tests {
			isLast := i == len(m.Tests)-1
			sections = append(sections, renderHook(m, hook, isLast))
		}
	}

	sections = append(sections, "")
	sections = append(sections)

	content := lipgloss.JoinVertical(lipgloss.Left, sections...)
	return "\n" + styles.Box.Render(content) + "\n"
}

func RenderResume(m *model.Model) string {
	var sections []string

	// Header del Resume
	divider := dividerStyle.Render(strings.Repeat("─", 44))
	sections = append(sections, divider)
	title := styles.Fail.Render("✗ Resume")
	sections = append(sections, title)
	sections = append(sections, divider)
	sections = append(sections, "")

	if len(m.Tests) > 0 {
		// Primero mostrar los que pasaron
		for _, test := range m.Tests {
			if test.Status == model.StatusOk {
				icon := styles.Ok.Render("✓")
				name := hookNameStyle.Render(test.Name)
				sections = append(sections, fmt.Sprintf("%s %s", icon, name))
			}
		}

		// Luego mostrar los que fallaron con detalles
		for _, test := range m.Tests {
			if test.Status == model.StatusFail {
				icon := styles.Fail.Render("✗")
				name := hookNameStyle.Render(test.Name)
				sections = append(sections, fmt.Sprintf("%s %s", icon, name))

				prefix := dividerStyle.Render("  └─")
				sections = append(sections, prefix)

				lines := strings.Split(strings.TrimSpace(test.Output), "\n")
				for _, line := range lines {
					sections = append(sections, fmt.Sprintf("     %s", line))
				}
				sections = append(sections, "")
			}
		}
	}

	content := lipgloss.JoinVertical(lipgloss.Left, sections...)

	hint := styles.Muted.Render("Press q, enter or ctrl+c to exit")

	return "\n" + content + "\n\n" + hint + "\n"
}

func renderHook(m *model.Model, test model.TestState, isLast bool) string {
	var icon string
	var statusStyle lipgloss.Style

	prefix := "├─"
	if isLast {
		prefix = "└─"
	}
	prefixStyled := dividerStyle.Render("  " + prefix)

	switch test.Status {
	case model.StatusPending:
		icon = "○"
		statusStyle = styles.Muted
	case model.StatusRunning:
		icon = m.Spinner.View()
		statusStyle = styles.Running
	case model.StatusOk:
		icon = "✓"
		statusStyle = styles.Ok
	case model.StatusFail:
		icon = "✗"
		statusStyle = styles.Fail
	}

	iconStyled := statusStyle.Render(icon)
	name := hookNameStyle.Render(test.Name)

	return fmt.Sprintf("%s %s %s", prefixStyled, iconStyled, name)
}

func renderStats(tests []model.TestState) string {
	var passed, failed, running, pending int

	for _, t := range tests {
		switch t.Status {
		case model.StatusOk:
			passed++
		case model.StatusFail:
			failed++
		case model.StatusRunning:
			running++
		case model.StatusPending:
			pending++
		}
	}

	var parts []string
	if passed > 0 {
		parts = append(parts, styles.Ok.Render(fmt.Sprintf("✓ %d", passed)))
	}
	if failed > 0 {
		parts = append(parts, styles.Fail.Render(fmt.Sprintf("✗ %d", failed)))
	}
	if running > 0 {
		parts = append(parts, styles.Running.Render(fmt.Sprintf("● %d", running)))
	}
	if pending > 0 {
		parts = append(parts, styles.Muted.Render(fmt.Sprintf("○ %d", pending)))
	}

	return "  " + strings.Join(parts, "  ")
}

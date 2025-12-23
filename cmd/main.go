package main

import (
	"fmt"
	"octohook/internal/config"
	"octohook/internal/hooks"
	"octohook/internal/model"
	"octohook/internal/render"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: octohook <command>\n")
		fmt.Fprintf(os.Stderr, "Commands:\n")
		fmt.Fprintf(os.Stderr, "  install           Install hooks to .git/hooks/\n")
		fmt.Fprintf(os.Stderr, "  uninstall         Uninstall hooks from .git/hooks/\n")
		fmt.Fprintf(os.Stderr, "  uninstall-hook    Uninstall a specific hook from .git/hooks/\n")
		fmt.Fprintf(os.Stderr, "  pre-commit        Run pre-commit hook\n")
		fmt.Fprintf(os.Stderr, "  post-commit       Run post-commit hook\n")
		fmt.Fprintf(os.Stderr, "  pre-push          Run pre-push hook\n")
		fmt.Fprintf(os.Stderr, "  post-push         Run post-push hook\n")
		os.Exit(1)
	}

	command := os.Args[1]

	cfg, err := config.LoadConfig("octohook.yml")
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	if command == "install" {
		if err := hooks.InstallAll(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if command == "uninstall" {
		if err := hooks.UninstallAll(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if command == "uninstall-hook" {
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "Usage: octohook uninstall-hook <hook-name>\n")
			os.Exit(1)
		}
		hookName := os.Args[2]
		switch hookName {
		case "pre-commit", "post-commit", "pre-push", "post-push":
			if err := hooks.Uninstall(hookName); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Uninstalled %s hook\n", hookName)
			os.Exit(0)
		default:
			fmt.Fprintf(os.Stderr, "Unknown hook name: %s\n", hookName)
			os.Exit(1)
		}
	}

	var hook *config.Hook
	switch command {
	case "pre-commit":
		hook = cfg.PreCommit
	case "post-commit":
		hook = cfg.PostCommit
	case "pre-push":
		hook = cfg.PrePush
	case "post-push":
		hook = cfg.PostPush
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	if hook == nil {
		fmt.Printf("No %s hook configured\n", command)
		os.Exit(0)
	}

	updates, cmd := hooks.StartHook(hook, cfg.InvalidateCacheOn)
	m := model.NewHookModel(command, cmd)
	m.Updates = updates
	app := &render.App{Model: m}
	p := tea.NewProgram(app)

	finalModel, err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	finalApp := finalModel.(*render.App)
	if finalApp.Model.Failed {
		fmt.Fprintln(os.Stderr)

		divider := "────────────────────────────────────────────"
		fmt.Fprintln(os.Stderr, "\033[38;5;241m"+divider+"\033[0m")
		fmt.Fprintln(os.Stderr, "\033[1;38;5;196m✗ Resume\033[0m")
		fmt.Fprintln(os.Stderr, "\033[38;5;241m"+divider+"\033[0m")
		fmt.Fprintln(os.Stderr)

		if len(finalApp.Model.Tests) > 0 {
			for _, test := range finalApp.Model.Tests {
				if test.Status == model.StatusFail {
					fmt.Fprintf(os.Stderr, "\033[1;38;5;196m✗ %s\033[0m\n", test.Name)
					fmt.Fprintln(os.Stderr, "\033[38;5;241m  └─\033[0m")

					lines := strings.Split(strings.TrimSpace(test.Output), "\n")
					for _, line := range lines {
						fmt.Fprintf(os.Stderr, "     %s\n", line)
					}
					fmt.Fprintln(os.Stderr)
				}
			}
		}
		os.Exit(1)
	}

	os.Exit(0)
}

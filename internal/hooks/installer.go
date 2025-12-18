package hooks

import (
	"fmt"
	"octohook/internal/config"
	"os"
	"path/filepath"
	"strings"
)

// InstallAll installs all hooks defined in the config
func InstallAll(cfg *config.Config) error {
	hooks := []struct {
		name string
		hook *config.Hook
	}{
		{"pre-commit", cfg.PreCommit},
		{"post-commit", cfg.PostCommit},
		{"pre-push", cfg.PrePush},
		{"post-push", cfg.PostPush},
	}

	installed := 0
	for _, h := range hooks {
		if h.hook != nil {
			if err := Install(h.name); err != nil {
				return fmt.Errorf("failed to install %s: %w", h.name, err)
			}
			installed++
		}
	}

	if installed == 0 {
		return fmt.Errorf("no hooks configured in octohook.yml")
	}

	fmt.Printf("Installed %d hook(s)\n", installed)
	return nil
}

// Install creates a git hook script
func Install(hookType string) error {
	hookPath := filepath.Join(".git", "hooks", hookType)

	script := `#!/bin/sh
octohook ` + hookType + `
`

	return os.WriteFile(hookPath, []byte(script), 0755)
}

func ListHooks() ([]string, error) {
	var hooks []string
	pathFiles, err := os.ReadDir(".git/hooks")
	if err != nil {
		return nil, err
	}
	for _, file := range pathFiles {
		read_files, err := os.ReadFile(filepath.Join(".git", "hooks", file.Name()))
		if err != nil {
			return nil, err
		}
		match := strings.Contains(string(read_files), "octohook")
		if match {
			hooks = append(hooks, file.Name())
		}
	}
	return hooks, nil
}

func UninstallAll() error {
	list, err := ListHooks()
	if err != nil {
		return fmt.Errorf("failed to list installed hooks: %w", err)
	}
	for _, hook := range list {
		if err := Uninstall(hook); err != nil {
			return fmt.Errorf("failed to uninstall %s: %w", hook, err)
		}
	}
	return nil
}

// Uninstall removes a git hook
func Uninstall(hookType string) error {
	hookPath := filepath.Join(".git", "hooks", hookType)
	return os.Remove(hookPath)
}

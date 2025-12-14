package git

import (
	"os/exec"
	"strings"
)

func GetStagedFiles() ([]string, error) {
	c := exec.Command("git", "diff", "--cached", "--name-only")
	output, err := c.Output()
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(string(output))
	if trimmed == "" {
		return []string{}, nil
	}

	return strings.Split(trimmed, "\n"), nil
}

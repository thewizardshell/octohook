package git

import (
	"os/exec"
	"strings"
)

func GetStatus() ([]string, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(string(output))
	if trimmed == "" {
		return []string{}, nil
	}

	lines := strings.Split(trimmed, "\n")
	files := make([]string, 0, len(lines))

	for _, line := range lines {
		if len(line) < 4 {
			continue
		}
		// Skip the status flags (first 3 chars) and get the filename
		file := strings.TrimSpace(line[3:])
		files = append(files, file)
	}

	return files, nil
}

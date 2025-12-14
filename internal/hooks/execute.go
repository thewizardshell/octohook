package hooks

import (
	"octohook/internal/model"
	"os/exec"
)

func execute(cmd model.Command) (bool, string) {
	c := exec.Command(cmd.Run, cmd.Args...)
	output, err := c.CombinedOutput()
	if err != nil {
		return false, string(output)
	}
	return err == nil, string(output)
}

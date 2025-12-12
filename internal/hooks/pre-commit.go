package hooks

import (
	"octohook/internal/model"
)

func PreCommitHook(m *model.Model) (bool, string) {
	return RunHook(m.Hooks.PreCommit)
}

package hooks

import (
	"octohook/internal/model"
)


func RunHook(hook model.Hook) (bool,string){
	if !hook.Enabled {
		return true, "skipped"
	}

	for _, cmd := range hook.Command {
		passed, output := execute(cmd)
		if !passed {
			return false, output
		}
	return true, "success"
}

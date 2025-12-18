package hooks

import (
	"octohook/internal/cache"
	"octohook/internal/config"
	"octohook/internal/git"
	"octohook/internal/model"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
	tea "github.com/charmbracelet/bubbletea"
)

// getSuffix extracts the suffix of a glob pattern after the last '*'.
// It is used to infer file relationships (e.g. service → test file mapping).
func getSuffix(pattern string) string {
	idx := strings.LastIndex(pattern, "*")
	if idx == -1 {
		return ""
	}
	return pattern[idx+1:]
}

// RunHook executes a hook based on the current git staged files.
//
// The process:
//   - Retrieves staged files
//   - Builds a dependency graph of the project
//   - Determines affected files and related tests
//   - Applies test filters based on hook configuration
//   - Uses cache to skip unchanged tests
//   - Executes tests and streams updates through the updates channel
//
// The updates channel is always closed when execution finishes.

func RunHook(hook *config.Hook, updates chan<- model.TestUpdate) {
	if hook == nil {
		close(updates)
		return
	}

	files, err := git.GetStagedFiles()
	if err != nil {
		close(updates)
		return
	}

	if len(files) == 0 {
		close(updates)
		return
	}

	graph := BuildGraph(".")
	if graph == nil {
		close(updates)
		return
	}

	affected := []string{}
	for _, f := range files {
		normalizedPath := filepath.ToSlash(f)
		result := graph.FindAffected(normalizedPath)
		affected = append(affected, result...)
	}

	testOnly, err := FilterTestOnly(affected, hook.Path.Test)
	if err != nil {
		close(updates)
		return
	}

	for _, f := range files {
		testIndir, err := FindTestInDir(f, hook.Path.Test)
		if err != nil {
			continue
		}
		testOnly = append(testOnly, testIndir...)
	}

	for _, to := range affected {
		testIndir, err := FindTestInDir(to, hook.Path.Test)
		if err != nil {
			continue
		}
		testOnly = append(testOnly, testIndir...)
	}

	testToServices := make(map[string][]string)
	for _, file := range files {
		for _, pattern := range hook.Path.Services {
			match, err := doublestar.PathMatch(pattern, file)
			if err != nil {
				continue
			}
			if match {
				serviceSuffix := getSuffix(pattern)
				testSuffix := getSuffix(hook.Path.Test[0])
				testFile := strings.TrimSuffix(file, serviceSuffix) + testSuffix
				testToServices[testFile] = append(testToServices[testFile], file)
				break
			}
		}
	}

	for _, test := range testOnly {
		updates <- model.TestUpdate{
			Name:   test,
			Status: model.StatusRunning,
		}

		filesToHash := []string{test}
		if relatedServices, ok := testToServices[test]; ok {
			filesToHash = append(filesToHash, relatedServices...)
		}

		currentHash, _ := cache.HashFiles(filesToHash)
		cached, _ := cache.Load(test)

		if cached != nil && cached.Hash == currentHash {
			updates <- model.TestUpdate{
				Name:   test,
				Status: model.StatusOk,
				Output: "cached",
			}
			continue
		}

		target := test
		if hook.UseDirectory {
			target = "./" + filepath.Dir(test)
		}

		args := append(hook.Arg, target)
		cmd := exec.Command(hook.Command, args...)
		output, err := cmd.CombinedOutput()
		passed := err == nil

		cache.Save(test, cache.CacheEntry{
			Hash:   currentHash,
			Passed: passed,
			Output: string(output),
		})

		if passed {
			updates <- model.TestUpdate{
				Name:   test,
				Status: model.StatusOk,
			}
		} else {
			updates <- model.TestUpdate{
				Name:   test,
				Status: model.StatusFail,
				Output: string(output),
			}
		}
	}

	close(updates)
}

// StartHook initializes and runs a hook asynchronously.
//
// It returns:
//   - A channel that streams TestUpdate events
//   - A Bubble Tea command that waits for hook execution to finish
//
// This function is the entry point used by the TUI layer
func StartHook(hook *config.Hook) (chan model.TestUpdate, tea.Cmd) {
	updates := make(chan model.TestUpdate)
	go RunHook(hook, updates)
	return updates, model.WaitForUpdate(updates)
}

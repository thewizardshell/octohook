package model

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type HookStatus int

const (
	StatusPending HookStatus = iota
	StatusRunning
	StatusOk
	StatusFail
)

// TestState represents the current state of a test
type TestState struct {
	Name   string
	Status HookStatus
	Output string
}

// TestUpdate is a message sent when a test changes state
type TestUpdate struct {
	Name   string
	Status HookStatus
	Output string
}

// HookFinished is sent when all tests are done
type HookFinished struct {
	Failed bool
}

type Model struct {
	HookName string
	Tests    []TestState
	Spinner  spinner.Model
	Running  bool
	HookCmd  tea.Cmd
	Done     bool
	Failed   bool
	Updates  chan TestUpdate
}

func NewModel() *Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	return &Model{
		Spinner: s,
		Running: false,
		Tests:   []TestState{},
		Done:    false,
		Failed:  false,
	}
}

func NewHookModel(hookName string, hookCmd tea.Cmd) *Model {
	m := NewModel()
	m.HookName = hookName
	m.Running = true
	m.HookCmd = hookCmd
	return m
}

func WaitForUpdate(updates <-chan TestUpdate) tea.Cmd {
	return func() tea.Msg {
		msg, ok := <-updates
		if !ok {
			return HookFinished{Failed: false}
		}
		return msg
	}
}

package model

type Hook struct {
	Enabled bool
	Command []Command
}

type Command struct {
	Name  string
	Run   string
	Files []string
	Args  []string
}

type GitHooks struct {
	PreCommit        Hook
	PostCommit       Hook
	PrePush          Hook
	PostPush         Hook
	PreRebase        Hook
	PostRebase       Hook
	PreMerge         Hook
	PostMerge        Hook
	CommitMsg        Hook
	PrepareCommitMsg Hook
}

type Model struct {
	Cursor int
	Width  int
	Height int
	Hooks  GitHooks
}

func NewModel() *Model {
	return &Model{
		Cursor: 0,
		Width:  80,
		Height: 24,
	}
}

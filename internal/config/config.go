package config

// Path defines glob patterns for matching service files and their corresponding tests.
// Used for intelligent test execution - only runs tests for changed services.
type Path struct {
	Services []string `yaml:"services"`
	Test     []string `yaml:"test"`
}

// Hook represents a git hook configuration.
// Cache defaults to true if omitted in yml.
// UseDirectory tells the runner to pass the directory instead of the file (needed for Go tests).
type Hook struct {
	Command      string   `yaml:"command"`
	Arg          []string `yaml:"arg,omitempty"`
	Path         Path     `yaml:"path,omitempty"`
	Cache        bool     `yaml:"cache,omitempty"`
	UseDirectory bool     `yaml:"use_directory,omitempty"`
}

// Config represents the complete octohook.yml structure.
// Each hook type is optional - only configure what you need.
type Config struct {
	PreCommit  *Hook `yaml:"pre-commit,omitempty"`
	PostCommit *Hook `yaml:"post-commit,omitempty"`
	PrePush    *Hook `yaml:"pre-push,omitempty"`
	PostPush   *Hook `yaml:"post-push,omitempty"`
}

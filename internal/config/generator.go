package config

import "gopkg.in/yaml.v3"

func Init(cfg *Config) (string, error) {
	generateHook(cfg, "PreCommit")
	generateHook(cfg, "PostCommit")
	generateHook(cfg, "PrePush")
	generateHook(cfg, "PostPush")

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func generateHook(cfg *Config, hookType string) {
	hook := &Hook{
		Command: "",
		Arg:     []string{""},
		Path: Path{
			Services: []string{""},
			Test:     []string{""},
		},
		Cache: true,
	}

	switch hookType {
	case "PreCommit":
		cfg.PreCommit = hook
	case "PostCommit":
		cfg.PostCommit = hook
	case "PrePush":
		cfg.PrePush = hook
	case "PostPush":
		cfg.PostPush = hook
	}
}

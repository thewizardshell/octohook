package cache

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type CacheEntry struct {
	Hash   string `json:"hash"`
	Passed bool   `json:"passed"`
	Output string `json:"output"`
}

func InitCache() error {
	return os.MkdirAll(".octohook/cache", 0755)
}

func Save(service string, entry CacheEntry) error {
	err := InitCache()
	if err != nil {
		return err
	}

	cachePath := filepath.Join(".octohook", "cache", service+".json")

	// Create subdirectories if needed
	dir := filepath.Dir(cachePath)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	return os.WriteFile(cachePath, data, 0644)

}

func Load(service string) (*CacheEntry, error) {
	cachePath := filepath.Join(".octohook", "cache", service+".json")
	data, err := os.ReadFile(cachePath)
	if err != nil {
		return nil, err
	}
	var entry CacheEntry
	err = json.Unmarshal(data, &entry)
	return &entry, err
}

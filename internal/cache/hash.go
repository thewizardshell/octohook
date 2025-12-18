package cache

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"sort"
)

func HashFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hasher := sha256.New()
	_, err = io.Copy(hasher, f)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func HashFiles(paths []string) (string, error) {
	hasher := sha256.New()

	// Sort paths to ensure consistent hash regardless of order
	sorted := make([]string, len(paths))
	copy(sorted, paths)
	sort.Strings(sorted)

	for _, path := range sorted {
		f, err := os.Open(path)
		if err != nil {
			return "", err
		}

		_, err = io.Copy(hasher, f)
		f.Close()
		if err != nil {
			return "", err
		}
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

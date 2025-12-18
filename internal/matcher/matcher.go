package matcher

import (
	"fmt"
	"octohook/internal/config"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
)

// getSuffix extracts the file suffix from a glob pattern.
// It finds the last "*" and returns everything after it.
//
// Examples:
//   - "**/*.go" → ".go"
//   - "**/*_test.go" → "_test.go"
//   - "**/*.test.ts" → ".test.ts"
func getSuffix(pattern string) string {
	idx := strings.LastIndex(pattern, "*")
	if idx == -1 {
		return ""
	}
	return pattern[idx+1:]
}

// Match takes staged files and returns their corresponding test files.
// It matches each file against the service patterns and transforms
// matching files to their test equivalents using the naming convention.
//
// Example:
//
//	files = ["src/auth.ts", "src/user.ts"]
//	path.Services = ["**/*.ts"]
//	path.Test = ["**/*.test.ts"]
//
//	Returns: ["src/auth.test.ts", "src/user.test.ts"]
//
// The transformation works by:
//  1. Checking if file matches any service pattern
//  2. Extracting suffix from service pattern (.ts)
//  3. Extracting suffix from test pattern (.test.ts)
//  4. Replacing service suffix with test suffix (auth.ts → auth.test.ts)
func Match(files []string, path config.Path) ([]string, error) {
	var tests []string

	for _, file := range files {
		matched := false
		for _, pattern := range path.Services {
			match, err := doublestar.PathMatch(pattern, file)
			if err != nil {
				return nil, err
			}
			if match {
				matched = true
				serviceSuffix := getSuffix(pattern)
				testSuffix := getSuffix(path.Test[0])

				testFile := strings.TrimSuffix(file, serviceSuffix) + testSuffix
				tests = append(tests, testFile)
				break
			}
		}
		if !matched {
			fmt.Printf("Warning: no match found for file: %s\n", file)
		}
	}
	return tests, nil
}

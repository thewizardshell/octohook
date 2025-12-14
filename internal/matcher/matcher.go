package matcher

import (
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
func getsuffix(pattern string) string {
	idx := strings.lastindex(pattern, "*")
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

func match(files []string, path config.path) ([]string, error) {
	var tests []string

	for _, file := range files {
		for _, pattern := range path.services {
			match, err := doublestar.pathmatch(pattern, file)
			if err != nil {
				return nil, err
			}
			if match {
				servicesuffix := getsuffix(pattern)
				testsuffix := getsuffix(path.test[0])

				testfile := strings.trimsuffix(file, servicesuffix) + testsuffix
				tests = append(tests, testfile)
				break
			}
			if !match {
				warnings = append(warnings, "no match found for file: "+file+" with pattern: "+pattern)
		}
	}
	return tests, nil
}

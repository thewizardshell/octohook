package hooks

import (
	"os"
	"path/filepath"

	"github.com/bmatcuk/doublestar/v4"
)

// Graph represents a reverse dependency graph where each file
// maps to the files that depend on it.
type Graph struct {
	reverse map[string][]string
}

// FindAffected returns all files that are transitively affected
// by a given file using a breadth-first traversal of the graph.

func (g *Graph) FindAffected(file string) []string {
	queue := []string{file}
	visited := make(map[string]bool)
	affected := []string{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}
		visited[current] = true

		dependents := g.reverse[current]
		for _, dep := range dependents {

			affected = append(affected, dep)
			queue = append(queue, dep)
		}

	}
	return affected

}

// FilterTestOnly filters a list of files and returns only those
// that match at least one test glob pattern.
func FilterTestOnly(files []string, patternTest []string) ([]string, error) {
	var tests []string

	for _, test := range files {
		for _, pattern := range patternTest {
			match, err := doublestar.PathMatch(pattern, test)
			if err != nil {
				continue
			}
			if match {
				tests = append(tests, test)
				break
			}
		}
	}
	return tests, nil
}

// BuildGraph scans the project directory and builds a reverse
// dependency graph based on import relationships.
//
// Supported languages:
//   - Go
//   - JavaScript
//   - TypeScript
//   - Python
//
// The graph is used to determine which files and tests are
// affected by a change.
func BuildGraph(projectPath string) *Graph {
	reverse := make(map[string][]string)
	packageFiles := make(map[string][]string)

	moduleName := getModuleName(projectPath)

	err := filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)
		validExt := map[string]bool{".go": true, ".js": true, ".ts": true, ".py": true}

		if !validExt[ext] {
			return nil
		}

		relPath := path
		if filepath.IsAbs(path) {
			relPath, _ = filepath.Rel(projectPath, path)
		}
		relPath = filepath.ToSlash(relPath)

		dir := filepath.Dir(relPath)
		packageFiles[dir] = append(packageFiles[dir], relPath)

		return nil
	})
	if err != nil {
		return nil
	}

	err = filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			dirName := info.Name()
			if dirName == ".git" || dirName == "node_modules" || dirName == ".cache" {
				return filepath.SkipDir
			}
			return nil
		}

		ext := filepath.Ext(path)
		validExt := map[string]bool{".go": true, ".js": true, ".ts": true, ".py": true}

		if !validExt[ext] {
			return nil
		}

		relPath := path
		if filepath.IsAbs(path) {
			relPath, _ = filepath.Rel(projectPath, path)
		}
		relPath = filepath.ToSlash(relPath)

		imports := extractImports(path)

		for _, imported := range imports {
			importPath := importToFilePath(imported, moduleName, projectPath, relPath, ext)
			if importPath != "" {
				reverse[importPath] = append(reverse[importPath], relPath)
			}
			reverse[imported] = append(reverse[imported], relPath)
		}

		if ext == ".go" {
			dir := filepath.Dir(relPath)
			for _, pkgFile := range packageFiles[dir] {
				if pkgFile != relPath {
					reverse[relPath] = append(reverse[relPath], pkgFile)
				}
			}
		}

		return nil
	})
	if err != nil {
		return nil
	}
	return &Graph{reverse: reverse}
}

// FindTestInDir searches for test files inside the directory
// of a given file, matching against the provided test patterns.
func FindTestInDir(file string, patternTest []string) ([]string, error) {
	dir := filepath.Dir(file)
	listFiles, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	list := []string{}
	for _, l := range listFiles {
		for _, pattern := range patternTest {
			match, err := doublestar.PathMatch(pattern, l.Name())
			if err != nil {
				continue
			}
			if match {
				list = append(list, filepath.Join(dir, l.Name()))
				break
			}
		}
	}
	return list, nil
}

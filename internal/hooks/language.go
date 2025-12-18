package hooks

import (
	"os"
	"path/filepath"
	"regexp"
)

func extractImports(filePath string) []string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil
	}
	text := string(content)
	ext := filepath.Ext(filePath)
	imports := []string{}

	switch ext {
	case ".go":
		// Go: import "path" or "path" in import block
		re := regexp.MustCompile(`(?m)(?:import\s+"([^"]+)"|^\s+"([^"]+)")`)
		matches := re.FindAllStringSubmatch(text, -1)
		for _, match := range matches {
			if match[1] != "" {
				imports = append(imports, match[1])
			} else if match[2] != "" {
				imports = append(imports, match[2])
			}
		}

	case ".ts", ".js":
		// TypeScript/JavaScript: import from "./path" or require("./path")
		// Captures: "./user.service" -> "./user.service"
		re := regexp.MustCompile(`(?:import\s+.*?from\s+["'](\.[^"']+)["']|require\(["'](\.[^"']+)["']\))`)
		matches := re.FindAllStringSubmatch(text, -1)
		for _, match := range matches {
			if match[1] != "" {
				imports = append(imports, match[1])
			} else if match[2] != "" {
				imports = append(imports, match[2])
			}
		}

	case ".py":
		re := regexp.MustCompile(`(?m)^(?:from\s+([^\s]+)\s+import|import\s+([^\s]+))`)
		matches := re.FindAllStringSubmatch(text, -1)
		for _, match := range matches {
			if match[1] != "" {
				imports = append(imports, match[1])
			} else if match[2] != "" {
				imports = append(imports, match[2])
			}
		}
	}

	return imports
}

func importToFilePath(importPath, moduleName, projectPath, currentFile, ext string) string {
	if ext == ".go" {
		if moduleName != "" && len(importPath) > len(moduleName) && importPath[:len(moduleName)] == moduleName {
			// Remove module prefix: "octohook/internal/cache" -> "internal/cache"
			relPath := importPath[len(moduleName):]
			if len(relPath) > 0 && relPath[0] == '/' {
				relPath = relPath[1:]
			}
			return relPath
		}
		return ""
	}

	if len(importPath) > 0 && (importPath[0] == '.' || importPath[0] == '/') {
		currentDir := filepath.Dir(currentFile)

		targetPath := filepath.Join(currentDir, importPath)
		targetPath = filepath.ToSlash(filepath.Clean(targetPath))

		// Add the appropriate extension
		if ext == ".ts" {
			return targetPath + ".ts"
		} else if ext == ".js" {
			return targetPath + ".js"
		} else if ext == ".py" {
			return targetPath + ".py"
		}

		return targetPath
	}

	return ""
}

func getModuleName(projectPath string) string {
	goModPath := filepath.Join(projectPath, "go.mod")
	content, err := os.ReadFile(goModPath)
	if err != nil {
		return ""
	}
	re := regexp.MustCompile(`(?m)^module\s+(.+)$`)
	match := re.FindStringSubmatch(string(content))
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

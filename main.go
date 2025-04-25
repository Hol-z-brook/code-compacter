package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Define flags
	includeConfig := flag.Bool("config", false, "Include configuration files (package.json, tsconfig.json, etc.)")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: go run main.go [options] /path/to/astro/project")
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	projectPath := flag.Args()[0]
	outputFile := "compacted.md"

	if err := processProject(projectPath, outputFile, *includeConfig); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Output written to %s\n", outputFile)
}

func processProject(projectPath string, outputFile string, includeConfig bool) error {
	// Define file extensions to process
	extensions := map[string]bool{
		".astro":  true,
		".ts":     true,
		".js":     true,
		".svelte": true,
		".css":    true,
		".scss":   true,
	}

	// Define directories to exclude
	excludeDirs := map[string]bool{
		"node_modules": true,
		".git":         true,
		"dist":         true,
		"examples":     true,
		"build":        true,
		".next":        true,
		".svelte-kit":  true,
		".astro":       true,
		"coverage":     true,
		".vscode":      true,
		"out":          true,
		"public":       true,
	}

	// Define security-sensitive files to exclude (always excluded)
	securityFiles := []string{
		".env",
		".env.*",
		".env.local",
		".env.development",
		".env.production",
		".env.test",
		"*.pem",
		"*.key",
		"*.cert",
		"*.crt",
		"id_rsa",
		"id_rsa.pub",
		"*.log",
		"npm-debug.log*",
		"yarn-debug.log*",
		"yarn-error.log*",
	}

	// Define optional configuration files to exclude (unless includeConfig is true)
	configFiles := []string{
		"package.json",
		"package-lock.json",
		"yarn.lock",
		"pnpm-lock.yaml",
		"tsconfig.json",
		"astro.config.*",
		"svelte.config.*",
		"vite.config.*",
		"*.config.*",
		"README.md",
		"LICENSE",
		"CHANGELOG.md",
	}

	// Create or truncate the output file
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer file.Close()

	// Write the header
	_, err = file.WriteString("# Project Code Overview\n\n")
	if err != nil {
		return err
	}

	// Walk through the project directory
	return filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip excluded directories
		if info.IsDir() {
			if excludeDirs[info.Name()] {
				return filepath.SkipDir
			}
			return nil
		}

		// Skip non-relevant files
		if !extensions[strings.ToLower(filepath.Ext(path))] {
			return nil
		}

		// Skip files in excluded directories
		for _, dir := range strings.Split(path, string(os.PathSeparator)) {
			if excludeDirs[dir] {
				return nil
			}
		}

		// Skip security-sensitive files (always excluded)
		fileName := filepath.Base(path)
		for _, pattern := range securityFiles {
			matched, err := filepath.Match(pattern, fileName)
			if err != nil {
				return err
			}
			if matched {
				return nil
			}
		}

		// Skip configuration files unless includeConfig is true
		if !includeConfig {
			for _, pattern := range configFiles {
				matched, err := filepath.Match(pattern, fileName)
				if err != nil {
					return err
				}
				if matched {
					return nil
				}
			}
		}

		// Read the file
		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error reading file %s: %w", path, err)
		}

		// Get relative path for output
		relPath, err := filepath.Rel(projectPath, path)
		if err != nil {
			return err
		}

		// Write file header and content with markdown formatting
		_, err = file.WriteString(fmt.Sprintf("## File: `%s`\n\n```%s\n%s\n```\n\n",
			relPath,
			getLanguageFromExtension(filepath.Ext(path)),
			string(content)))
		return err
	})
}

func getLanguageFromExtension(ext string) string {
	switch strings.ToLower(ext) {
	case ".astro":
		return "astro"
	case ".ts":
		return "typescript"
	case ".js":
		return "javascript"
	case ".svelte":
		return "svelte"
	case ".css":
		return "css"
	case ".scss":
		return "scss"
	default:
		return "text"
	}
}

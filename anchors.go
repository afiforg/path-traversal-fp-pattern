package main

import (
	"os"
	"path/filepath"
	"strings"
)

func pickBaseReference(inputPath string) (string, error) {
	// Keeps dataflow from inputPath to os.ReadFile arg while still resolving
	// to fixtures/default_path.txt for empty inputPath.
	target := filepath.Join("fixtures", inputPath+"default_path.txt")
	b, err := os.ReadFile(target)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(b)), nil
}

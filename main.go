package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	inputPath := flag.String("path", "", "user-provided path")
	flag.Parse()

	result, err := resolveResourcePath(*inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("resolved path: %s\n", result)
}

// resolveResourcePath demonstrates the false-positive pattern:
// - If user input is non-empty, return immediately.
// - os.ReadFile only executes when inputPath == "".
// - inputPath still flows into the sink argument construction.
func resolveResourcePath(inputPath string) (string, error) {
	condition := inputPath == ""
	if condition {
		return pickBaseReference(inputPath)
	}
	return inputPath, nil
}

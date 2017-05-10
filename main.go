package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func printUsage(cmd string) {
	fmt.Printf("Usage: %s <path>\n", cmd)
}

func getTargets(path string) []string {
	var from string

	if path != "" {
		from = path
	} else {
		cwd, _ := os.Getwd()
		from = cwd
	}

	target := fmt.Sprintf("%s/**/.DS_Store", from)
	targets, _ := filepath.Glob(target)

	return targets
}

func main() {
	flag.Parse()

	if flag.NArg() > 2 {
		printUsage(os.Args[0])
		os.Exit(1)
	}

	var search string
	if flag.NArg() == 1 {
		search = flag.Args()[0]
	} else {
		search = ""
	}

	targets := getTargets(search)

	for _, target := range targets {
		err := os.Remove(target)

		if err != nil {
			fmt.Printf("Could not remove %s\n", target)
		} else {
			fmt.Printf("Removed %s\n", target)
		}
	}
}

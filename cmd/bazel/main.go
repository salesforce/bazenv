package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"github.com/salesforce/bazenv/pkg/bazenv"
)

func main() {
	bazenv.EnsureBazenvDir()

	// Respond to bazenv sniff used by doctor command
	if len(os.Args) == 2 && os.Args[1] == "bazenv" {
		fmt.Println("yes")
		os.Exit(0)
	}

	bazelProfile, err := bazenv.ReadBazenvFile()
	if err != nil {
		fmt.Println("Error reading global or local bazenv_version file: " + err.Error())
		os.Exit(1)
	}

	bazelDir, err := bazenv.ResolveBazelDirectory(bazelProfile)
	if err != nil {
		fmt.Println("Error finding bazil: " + err.Error())
		os.Exit(1)
	}

	// Execute the selected real bazel entry point
	binary := filepath.Join(bazelDir, "bin", "bazel")
	args := os.Args
	env := os.Environ()

	err = syscall.Exec(binary, args, env)
	if err != nil {
		fmt.Println("Error executing bazil: " + err.Error())
		os.Exit(1)
	}
}

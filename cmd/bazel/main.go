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
	// binary := "/usr/local/Cellar/bazel/0.23.2/libexec/bin/bazel"
	binary := filepath.Join(bazelDir, "bin", "bazel")
	args := os.Args
	env := os.Environ()

	err = syscall.Exec(binary, args, env)
	if err != nil {
		fmt.Println("Error executing bazil: " + err.Error())
		os.Exit(1)
	}
}

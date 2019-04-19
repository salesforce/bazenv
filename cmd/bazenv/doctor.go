package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/subcommands"
	"github.com/mitchellh/go-homedir"
	"github.com/salesforce/bazenv/pkg/bazenv"
)

type doctorCmd struct{}

func (*doctorCmd) Name() string {
	return "doctor"
}

func (*doctorCmd) Synopsis() string {
	return "looks for problems with your bazenv environment"
}

func (g *doctorCmd) Usage() string {
	return "bazenv doctor\n" + g.Synopsis()
}

func (g *doctorCmd) SetFlags(f *flag.FlagSet) {}

func (g *doctorCmd) Execute(_ context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	// Sanity check input
	if len(f.Args()) > 0 {
		fmt.Println("Too many arguments.")
		return subcommands.ExitUsageError
	}

	// Check: Bazel executable is bazenv stub
	isBazenv, err := bazenv.SniffIsBazenvStub()
	if err != nil {
		failure("Error identifying bazel executable: " + err.Error())
	}
	if !isBazenv {
		failure("Your bazel executable is not the bazenv shim. Make sure the bazenv bazel shim is on your path with" +
			" higher priority than any other bazel executable.")
	} else {
		success("Found bazenv bazel shim.")
	}

	// Check: Bazenv directory exists in home directory and is writable
	homedir, err := homedir.Dir()
	if err != nil {
		failure("Error accessing home directory: " + err.Error())
	}

	exists, writable := true, false
	versionsDir := filepath.Join(homedir, bazenv.BazenvDir, bazenv.BazenvVersionsDir)

	_, err = os.Stat(versionsDir)
	if os.IsNotExist(err) {
		exists = false
	}

	tmp, err := os.Create(filepath.Join(versionsDir, "touch.tmp"))
	tmp.Close()
	err = os.Remove(filepath.Join(versionsDir, "touch.tmp"))

	if err == nil {
		writable = true
	}

	if exists && writable {
		success("Bazenv working directory exists and is writable.")
	} else {
		failure("~/.bazenv/bazel directory is missing or not writeable.")
	}

	// Check: At least one bazel version is installed
	versions, err := bazenv.ListBazelVersions()
	if err != nil {
		failure("Error checking for bazel versions: " + err.Error())
	}
	if len(versions) == 0 {
		failure("No bazel version installed. Use 'bazenv install' to download bazel, or 'bazenv add' to link to an" +
			" existing bazel directory.")
	} else {
		success("At least one bazel version installed.")
	}

	// Check: A global bazel version is selected
	globalVersion, err := bazenv.ReadGlobalBazenvFile()
	if err != nil {
		failure("Error reading global bazenv version: " + err.Error())
	}
	if globalVersion == nil {
		failure("No global bazel version selected. Use 'bazenv global' to select one.")
	} else {
		success("Global bazel version selected.")
	}

	// Check: Global bazel version is executable

	return subcommands.ExitSuccess
}

func success(message string) {
	fmt.Println("✅ " + message)
}

func failure(message string) {
	fmt.Println("❌ " + message)
}

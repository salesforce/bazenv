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

	allSucceeded := true
	homedir, err := homedir.Dir()
	if err != nil {
		failure("Error accessing home directory: " + err.Error())
		allSucceeded = false
	}

	allSucceeded = allSucceeded && checkBazenvStub()
	allSucceeded = allSucceeded && checkWritableBazenvHome(homedir)
	allSucceeded = allSucceeded && checkBazelInstalled(homedir)
	allSucceeded = allSucceeded && checkBazelSelected(homedir)
	allSucceeded = allSucceeded && checkBazelExecutable(homedir)

	if !allSucceeded {
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

// Check: Bazel executable is bazenv stub
func checkBazenvStub() bool {
	allSucceeded := true
	isBazenv, err := bazenv.SniffIsBazenvStub()
	if err != nil {
		failure("Error identifying bazel executable: " + err.Error())
		allSucceeded = false
	}
	if !isBazenv {
		failure("Your bazel executable is not the bazenv shim. Make sure the bazenv bazel shim is on your path with" +
			" higher priority than any other bazel executable.")
		allSucceeded = false
	} else {
		success("Found bazenv bazel shim.")
	}
	return allSucceeded
}

// Check: Bazenv directory exists in home directory and is writable
func checkWritableBazenvHome(homedir string) bool {
	allSucceeded := true
	exists, writable := true, false
	versionsDir := filepath.Join(homedir, bazenv.BazenvDir, bazenv.BazenvVersionsDir)

	_, err := os.Stat(versionsDir)
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
		allSucceeded = false
	}

	return allSucceeded
}

// Check: At least one bazel version is installed
func checkBazelInstalled(homedir string) bool {
	allSucceeded := true
	versions, err := bazenv.ListBazelVersions()
	if err != nil {
		failure("Error checking for bazel versions: " + err.Error())
		allSucceeded = false
	}
	if len(versions) == 0 {
		failure("No bazel version installed. Use 'bazenv install' to download bazel, or 'bazenv add' to link to an" +
			" existing bazel directory.")
		allSucceeded = false
	} else {
		success("At least one bazel version installed.")
	}
	return allSucceeded
}

// Check: A global bazel version is selected
func checkBazelSelected(homedir string) bool {
	allSucceeded := true
	globalVersion, err := bazenv.ReadGlobalBazenvFile()
	if err != nil {
		failure("Error reading global bazenv version: " + err.Error())
		allSucceeded = false
	}
	if globalVersion == nil {
		failure("No global bazel version selected. Use 'bazenv global' to select one.")
		allSucceeded = false
	} else {
		success("Global bazel version selected.")
	}
	return allSucceeded
}

// Check: Global bazel version is executable
func checkBazelExecutable(homedir string) bool {
	allSucceeded := true
	globalVersion, _ := bazenv.ReadGlobalBazenvFile()
	if globalVersion != nil {
		path, err := bazenv.ResolveBazelDirectory(*globalVersion)
		if err != nil {
			failure("Error resolving bazel directory: " + err.Error())
			allSucceeded = false
		}
		version, err := bazenv.SniffBazelVersion(path)
		if err != nil {
			failure("Could not verify global bazel version is a valid bazel install: " + err.Error())
			allSucceeded = false
		}
		success("Global bazel version is a valid bazel install (bazel " + version + ").")
	}
	return allSucceeded
}

func success(message string) {
	fmt.Println("✅ " + message)
}

func failure(message string) {
	fmt.Println("❌ " + message)
}

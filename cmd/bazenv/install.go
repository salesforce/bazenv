package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/google/subcommands"
	"github.com/mitchellh/go-homedir"
	"github.com/salesforce/bazenv/pkg/bazenv"
)

type installCmd struct{}

func (*installCmd) Name() string {
	return "install"
}

func (*installCmd) Synopsis() string {
	return "download and install bazel from github"
}

func (g *installCmd) Usage() string {
	return "bazenv install <version>\n" + g.Synopsis()
}

func (g *installCmd) SetFlags(f *flag.FlagSet) {}

func (g *installCmd) Execute(_ context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	// Sanity check input
	if len(f.Args()) != 1 && len(f.Args()) != 1 {
		fmt.Println("Must specify a bazel version to install.")
		return subcommands.ExitUsageError
	}

	// Download bazel
	version := f.Arg(0)
	filename, err := bazenv.DownloadBazelVersion(version)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("See https://github.com/bazelbuild/bazel/releases for available versions")
		return subcommands.ExitFailure
	}

	// Prepare for install
	homedir, err := homedir.Dir()
	if err != nil {
		fmt.Println(err.Error())
		return subcommands.ExitFailure
	}

	installDir := filepath.Join(homedir, bazenv.BazenvDir, bazenv.BazenvVersionsDir, version)
	os.MkdirAll(installDir, os.ModePerm)

	// Execute the installer
	fmt.Println("Installing " + filename)
	filepath := filepath.Join(homedir, bazenv.BazenvDir, filename)
	installArgs := []string{"--prefix=" + installDir, "--bin=" + installDir + "/bin", "--base=" + installDir + "/lib/bazel"}
	output, err := exec.Command(filepath, installArgs...).Output()
	if err != nil {
		fmt.Println("Error executing bazil installer: " + err.Error())
		return subcommands.ExitFailure
	}
	fmt.Println(string(output))

	// Delete the installer
	err = os.Remove(filepath)
	if err != nil {
		fmt.Println(err.Error())
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

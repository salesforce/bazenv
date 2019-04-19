package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"github.com/salesforce/bazenv/pkg/bazenv"
)

type addCmd struct{}

func (*addCmd) Name() string {
	return "add"
}

func (*addCmd) Synopsis() string {
	return "add an bazel install directory to bazenv"
}

func (g *addCmd) Usage() string {
	return "bazenv add <path> [name]\n" + g.Synopsis()
}

func (g *addCmd) SetFlags(f *flag.FlagSet) {}

func (g *addCmd) Execute(_ context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	// Sanity check input
	if len(f.Args()) != 1 && len(f.Args()) != 2 {
		fmt.Println("Must specify the path to a bazel install, and optionally a custom name.")
		return subcommands.ExitUsageError
	}

	path := f.Arg(0)
	version, err := bazenv.SniffBazelVersion(path)
	if err != nil {
		fmt.Println(err.Error())
		return subcommands.ExitFailure
	}

	// Override version name if provided
	if len(f.Args()) == 2 {
		version = f.Arg(1)
	}

	err = bazenv.AddBazelVersion(path, version)
	if err != nil {
		fmt.Println("Error adding bazel version to bazenv: " + err.Error())
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

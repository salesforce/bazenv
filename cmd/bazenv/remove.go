package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"github.com/ryanmichela/bazenv/pkg/bazenv"
)

type removeCmd struct{}

func (*removeCmd) Name() string {
	return "remove"
}

func (*removeCmd) Synopsis() string {
	return "remove a bazel version from bazenv"
}

func (g *removeCmd) Usage() string {
	return "remove version\n" + g.Synopsis()
}

func (g *removeCmd) SetFlags(f *flag.FlagSet) {}

func (g *removeCmd) Execute(_ context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	// Sanity check input
	if len(f.Args()) != 1 {
		fmt.Println("Must specify the version to remove.")
		return subcommands.ExitUsageError
	}

	version := f.Arg(0)
	err := bazenv.RemoveBazelVersion(version)
	if err != nil {
		fmt.Println("Error removing bazel version: " + err.Error())
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

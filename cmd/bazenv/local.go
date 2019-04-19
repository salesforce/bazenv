package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"github.com/salesforce/bazenv/pkg/bazenv"
)

type localCmd struct{}

func (*localCmd) Name() string {
	return "local"
}

func (*localCmd) Synopsis() string {
	return "set the local bazel version (this directory, and child directories)"
}

func (l *localCmd) Usage() string {
	return "bazenv local <version name>\n" + l.Synopsis()
}

func (l *localCmd) SetFlags(f *flag.FlagSet) {}

func (l *localCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	// Sanity check input
	if len(f.Args()) != 1 {
		fmt.Println("Must specify a single bazel version name")
		return subcommands.ExitUsageError
	}

	// Set the global bazel version
	bazenv.SetLocalBazelVersion(f.Arg(0))
	return subcommands.ExitSuccess
}

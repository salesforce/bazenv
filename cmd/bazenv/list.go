package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"github.com/salesforce/bazenv/pkg/bazenv"
)

type listCmd struct{}

func (*listCmd) Name() string {
	return "list"
}

func (*listCmd) Synopsis() string {
	return "lists the bazel versions known to bazenv"
}

func (l *listCmd) Usage() string {
	return "list\n" + l.Synopsis()
}

func (l *listCmd) SetFlags(f *flag.FlagSet) {}

func (l *listCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	versions, err := bazenv.ListBazelVersions()
	if err != nil {
		fmt.Println("Error listing bazel versions: " + err.Error())
		return subcommands.ExitFailure
	}

	global, err := bazenv.ReadGlobalBazenvFile()
	if err != nil {
		fmt.Println("Error reading global bazel version: " + err.Error())
		return subcommands.ExitFailure
	}

	local, err := bazenv.FindAndReadLocalBazenvFile()
	if err != nil {
		fmt.Println("Error reading local bazel version: " + err.Error())
		return subcommands.ExitFailure
	}

	for _, name := range versions {
		printname := name
		if (global != nil && name == *global) || (local != nil && name == *local) {
			printname = "* " + printname + " "
		}
		if global != nil && name == *global {
			printname += "(global)"
		}
		if local != nil && name == *local {
			printname += "(local)"
		}

		fmt.Println(printname)
	}

	return subcommands.ExitSuccess
}

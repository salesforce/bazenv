/*
 *  Copyright (c) 2019, Salesforce.com, Inc.
 *  All rights reserved.
 *  Licensed under the BSD 3-Clause license.
 *  For full license text, see LICENSE.txt file in the repo root  or https://opensource.org/licenses/BSD-3-Clause
 */

package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/salesforce/bazenv/pkg/bazenv"

	"github.com/google/subcommands"
)

type whichCmd struct{}

func (*whichCmd) Name() string {
	return "which"
}

func (*whichCmd) Synopsis() string {
	return "print bazel's install directory"
}

func (g *whichCmd) Usage() string {
	return "bazenv which [version]\n" + g.Synopsis()
}

func (g *whichCmd) SetFlags(f *flag.FlagSet) {}

func (g *whichCmd) Execute(_ context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	// Sanity check input
	if len(f.Args()) > 1 {
		fmt.Println("Too many arguments")
		return subcommands.ExitUsageError
	}

	// Look up version if not provided
	var version string
	if len(f.Args()) == 0 {
		var err error
		version, err = bazenv.ReadBazenvFile()
		if err != nil {
			fmt.Println(err.Error())
			return subcommands.ExitFailure
		}
	} else {
		version = f.Arg(0)
	}

	// Find version
	path, err := bazenv.ResolveBazelDirectory(version)
	if err != nil {
		fmt.Println(err.Error())
		return subcommands.ExitFailure
	}

	fmt.Println(path)

	return subcommands.ExitSuccess
}

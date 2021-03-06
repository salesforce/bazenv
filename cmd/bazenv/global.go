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

	"github.com/google/subcommands"
	"github.com/salesforce/bazenv/pkg/bazenv"
)

type globalCmd struct{}

func (*globalCmd) Name() string {
	return "global"
}

func (*globalCmd) Synopsis() string {
	return "set the global bazel version"
}

func (g *globalCmd) Usage() string {
	return "bazenv global <version name>\n" + g.Synopsis()
}

func (g *globalCmd) SetFlags(f *flag.FlagSet) {}

func (g *globalCmd) Execute(_ context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	// Sanity check input
	if len(f.Args()) != 1 {
		fmt.Println("Must specify a single bazel version name")
		return subcommands.ExitUsageError
	}

	// Set the global bazel version
	bazenv.SetGlobalBazelVersion(f.Arg(0))
	return subcommands.ExitSuccess
}

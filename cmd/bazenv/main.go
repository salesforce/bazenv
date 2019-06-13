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
	"os"

	"github.com/google/subcommands"
	"github.com/salesforce/bazenv/pkg/bazenv"
)

func main() {
	bazenv.EnsureBazenvDir()

	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(&globalCmd{}, "")
	subcommands.Register(&localCmd{}, "")
	subcommands.Register(&listCmd{}, "")
	subcommands.Register(&addCmd{}, "")
	subcommands.Register(&removeCmd{}, "")
	subcommands.Register(&whichCmd{}, "")
	subcommands.Register(&installCmd{}, "")
	subcommands.Register(&doctorCmd{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}

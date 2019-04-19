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

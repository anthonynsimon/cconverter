package main

import (
	"context"
	"flag"
	"os"

	"github.com/anthonynsimon/xeclient/commands"
	"github.com/google/subcommands"
)

var (
	apiHost = flag.String("apiHost", "http://localhost:8080", "usage...")
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&commands.ConvertCmd{}, "")
	subcommands.Register(&commands.RatesCmd{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}

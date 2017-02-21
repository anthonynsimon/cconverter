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
	subcommands.Register(&commands.NormalizeCmd{}, "")
	subcommands.Register(&commands.ConvertCmd{}, "")
	subcommands.Register(&commands.RatesCmd{}, "")
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")

	flag.Parse()
	ctx := context.WithValue(context.Background(), "apiHost", *apiHost)
	os.Exit(int(subcommands.Execute(ctx)))
}

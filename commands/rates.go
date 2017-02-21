package commands

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

type RatesCmd struct {
	currency string
}

func (*RatesCmd) Name() string {
	return "rates"
}

func (*RatesCmd) Synopsis() string {
	return "Gets the latest rates from a provided currency."
}

func (*RatesCmd) Usage() string {
	return `rates [-currency]:
	For ex.: -currency=EUR
`
}

func (cmd *RatesCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.currency, "currency", "", "the currency code to get the rates from")
}

func (cmd *RatesCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	for _, arg := range f.Args() {
		// TODO: implement
		fmt.Printf("%s ", arg)
	}
	fmt.Println()
	return subcommands.ExitSuccess
}

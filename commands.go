package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

type convertCmd struct {
	from   string
	to     string
	amount string
}

func (*convertCmd) Name() string { return "convert" }
func (*convertCmd) Synopsis() string {
	return "Converts an 'amount' from the provided 'from' currency to the 'to' currency."
}
func (*convertCmd) Usage() string {
	return `convert [-from] [-to] [-amount]:
	For ex.: -from=EUR -to=USD -amount=100
`
}

func (cmd *convertCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.from, "from", "", "the base currency")
	f.StringVar(&cmd.to, "to", "", "the target currency after conversion")
	f.StringVar(&cmd.amount, "amount", "", "the amount to be converted, in the base currency")
}

func (cmd *convertCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	for _, arg := range f.Args() {
		// TODO: implement
		fmt.Printf("%s ", arg)
	}
	fmt.Println()
	return subcommands.ExitSuccess
}

type ratesCmd struct {
	currency string
}

func (*ratesCmd) Name() string     { return "rates" }
func (*ratesCmd) Synopsis() string { return "Gets the latest rates from a provided currency." }
func (*ratesCmd) Usage() string {
	return `rates [-currency]:
	For ex.: -currency=EUR
`
}

func (cmd *ratesCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.currency, "currency", "", "the currency code to get the rates from")
}

func (cmd *ratesCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	for _, arg := range f.Args() {
		// TODO: implement
		fmt.Printf("%s ", arg)
	}
	fmt.Println()
	return subcommands.ExitSuccess
}

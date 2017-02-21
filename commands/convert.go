package commands

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

type ConvertCmd struct {
	from   string
	to     string
	amount string
}

func (*ConvertCmd) Name() string {
	return "convert"
}

func (*ConvertCmd) Synopsis() string {
	return "Converts an 'amount' from the provided 'from' currency to the 'to' currency."
}

func (*ConvertCmd) Usage() string {
	return `convert [-from] [-to] [-amount]:
	For ex.: -from=EUR -to=USD -amount=100
`
}

func (cmd *ConvertCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.from, "from", "", "the base currency")
	f.StringVar(&cmd.to, "to", "", "the target currency after conversion")
	f.StringVar(&cmd.amount, "amount", "", "the amount to be converted, in the base currency")
}

func (cmd *ConvertCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	for _, arg := range f.Args() {
		// TODO: implement
		fmt.Printf("%s ", arg)
	}
	fmt.Println()
	return subcommands.ExitSuccess
}

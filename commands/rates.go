package commands

import (
	"context"
	"flag"
	"fmt"

	"github.com/anthonynsimon/xeclient/client"
	"github.com/anthonynsimon/xeclient/currency"
	"github.com/google/subcommands"
)

type RatesCmd struct {
	currencyCode string
}

func (*RatesCmd) Name() string {
	return "rates"
}

func (*RatesCmd) Synopsis() string {
	return "Gets the latest rates from a provided currency."
}

func (*RatesCmd) Usage() string {
	return `rates [-currency]:
	For ex.: -currencycode=EUR
`
}

func (cmd *RatesCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.currencyCode, "currencycode", "", "the currency code to get the rates from")
}

func (cmd *RatesCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	// TODO: validate inputs
	// currency := cmd.currency

	currencyCode, err := currency.Parse(cmd.currencyCode)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	apiHost := extractApiHost(ctx)
	apiClient := client.NewClient(apiHost)

	rates, err := apiClient.GetRates(currencyCode)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}
	fmt.Println(rates)

	return subcommands.ExitSuccess
}

package commands

import (
	"context"
	"flag"
	"fmt"

	"github.com/anthonynsimon/cconverter/client"
	"github.com/anthonynsimon/cconverter/currency"
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

	fmt.Println("--------------------------")
	fmt.Printf("Base Currency:\t%s\n\n", rates.Base)
	for code, rate := range rates.Rates {
		fmt.Printf("%s:\t\t%f\n", code, rate)
	}
	fmt.Println("--------------------------")

	return subcommands.ExitSuccess
}

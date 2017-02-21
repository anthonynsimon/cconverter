package commands

import (
	"context"
	"flag"
	"fmt"
	"strconv"

	"github.com/anthonynsimon/xeclient/client"
	"github.com/anthonynsimon/xeclient/currency"
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

func (cmd *ConvertCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	fromCurrency, err := currency.Parse(cmd.from)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}
	toCurrency, err := currency.Parse(cmd.to)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	amount, err := strconv.ParseFloat(cmd.amount, 64)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	apiHost := extractApiHost(ctx)
	apiClient := client.NewClient(apiHost)

	quote, err := apiClient.Convert(fromCurrency, toCurrency, amount)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	fmt.Println("--------------------------")
	fmt.Printf("From Currency:\t%s\n", quote.FromCurrency)
	fmt.Printf("To Currency:\t%s\n", quote.ToCurrency)
	fmt.Printf("Amount:\t\t%f\n", quote.AmountToConvert)
	fmt.Printf("Exchange Rate:\t%f\n", quote.ExchangeRate)
	fmt.Printf("Result:\t\t%f\n", quote.ConversionResult)
	fmt.Println("--------------------------")

	return subcommands.ExitSuccess
}

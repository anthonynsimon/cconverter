package commands

import (
	"context"
	"flag"
	"fmt"

	"github.com/anthonynsimon/cconverter/client"
	"github.com/anthonynsimon/cconverter/currency"
	"github.com/google/subcommands"
	"github.com/shopspring/decimal"
)

// ConvertCmd holds related data and methods for the Conversion operation.
type ConvertCmd struct {
	from   string
	to     string
	amount string
}

func (*ConvertCmd) Name() string {
	return "convert"
}

func (*ConvertCmd) Synopsis() string {
	return "converts an amount from one currency to another"
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
	// TODO: tidy this up
	// Validate inputs
	if (cmd.to == "") || (cmd.from == "") || (cmd.amount == "") {
		fmt.Printf("please specify the 'from' and 'to' currencies, as well as the amount to be converted\n\n")
		f.Usage()
		return subcommands.ExitFailure
	}

	// Parse provided inputs and validate them
	fromCurrency, err := currency.Parse(cmd.from)
	if err != nil {
		fmt.Printf(err.Error() + "\n\n")
		f.Usage()
		return subcommands.ExitFailure
	}
	toCurrency, err := currency.Parse(cmd.to)
	if err != nil {
		fmt.Printf(err.Error() + "\n\n")
		f.Usage()
		return subcommands.ExitFailure
	}
	amount, err := decimal.NewFromString(cmd.amount)
	if err != nil {
		fmt.Printf(err.Error() + "\n\n")
		f.Usage()
		return subcommands.ExitFailure
	}

	// Initialize client
	apiHost := extractAPIHost(ctx)
	apiClient := client.NewClient(apiHost)

	// Convert to target currency
	quote, err := apiClient.Convert(fromCurrency, toCurrency, amount)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	fmt.Println("--------------------------")
	fmt.Printf("From Currency:\t%s\n", quote.FromCurrency)
	fmt.Printf("To Currency:\t%s\n", quote.ToCurrency)
	fmt.Printf("Amount:\t\t%s\n", quote.AmountToConvert.String())
	fmt.Printf("Exchange Rate:\t%s\n", quote.ExchangeRate.String())
	fmt.Printf("Result:\t\t%s\n", quote.ConversionResult.String())
	fmt.Println("--------------------------")

	return subcommands.ExitSuccess
}

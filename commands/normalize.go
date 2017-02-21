package commands

import (
	"context"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/anthonynsimon/cconverter/client"
	"github.com/anthonynsimon/cconverter/currency"
	"github.com/google/subcommands"
)

var (
	ErrInvalidCSVFormat = errors.New("invalid csv record format. expected '[id], [amount], [currency code]'")
)

type NormalizeCmd struct {
	toCurrency string
	csvFile    string
	out        string
}

func parseCsvRecord(record []string) (*csvRecord, error) {
	if len(record) != 3 {
		return nil, ErrInvalidCSVFormat
	}
	result := &csvRecord{}
	result.ID = record[0]

	amount, err := strconv.ParseFloat(record[1], 64)
	if err != nil {
		return nil, ErrInvalidCSVFormat
	}

	result.Amount = amount

	currencyCode, err := currency.Parse(record[2])
	if err != nil {
		return nil, ErrInvalidCSVFormat
	}

	result.CurrencyCode = currencyCode

	return result, nil
}

type csvRecord struct {
	ID           string
	Amount       float64
	CurrencyCode currency.Currency
}

func (*NormalizeCmd) Name() string {
	return "normalize"
}

func (*NormalizeCmd) Synopsis() string {
	return "normalizes a CSV file so that all values are in the same currency"
}

func (*NormalizeCmd) Usage() string {
	return `normalize [-csvfile] [-to] [-out]:
	For ex.: -csvfile=./myfile.csv -to=USD -out=./normalized.csv
`
}

func (cmd *NormalizeCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.toCurrency, "to", "", "the target currency for the normalization")
	f.StringVar(&cmd.csvFile, "csvfile", "", "the path to the csv file to be normalized")
	f.StringVar(&cmd.out, "out", "./normalized.csv", "the file to which the normalized output will be saved to")
}

func (cmd *NormalizeCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if cmd.toCurrency == "" {
		fmt.Printf("please specify a target currency\n\n")
		f.Usage()
		return subcommands.ExitFailure
	}
	toCurrency, err := currency.Parse(cmd.toCurrency)
	if err != nil {
		fmt.Printf("bad target currency: %s\n\n", err)
		f.Usage()
		return subcommands.ExitFailure
	}
	if cmd.csvFile == "" {
		fmt.Printf("csv file path cannot be empty\n\n")
		f.Usage()
		return subcommands.ExitFailure
	}

	fileIn, err := os.Open(cmd.csvFile)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}
	defer fileIn.Close()

	fileOut, err := os.Create(cmd.out)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}
	defer fileOut.Close()

	csvWriter := csv.NewWriter(fileOut)
	csvReader := csv.NewReader(fileIn)

	apiHost := extractApiHost(ctx)
	apiClient := client.NewClient(apiHost)

	for {
		rawRecord, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return subcommands.ExitFailure
		}

		record, err := parseCsvRecord(rawRecord)
		if err != nil {
			fmt.Println(err)
			return subcommands.ExitFailure
		}

		fmt.Printf("Converting %s %s to %s\n", strconv.FormatFloat(record.Amount, 'f', -1, 64), record.CurrencyCode, toCurrency)

		quote, err := apiClient.Convert(record.CurrencyCode, toCurrency, record.Amount)
		if err != nil {
			fmt.Println(err)
			return subcommands.ExitFailure
		}

		conversionResult := strconv.FormatFloat(quote.ConversionResult, 'f', 4, 64)

		err = csvWriter.Write([]string{record.ID, conversionResult, string(toCurrency)})
		if err != nil {
			fmt.Println(err)
			return subcommands.ExitFailure
		}

		csvWriter.Flush()

	}
	return subcommands.ExitSuccess
}

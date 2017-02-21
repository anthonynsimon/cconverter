# cconverter
[![Build Status](https://travis-ci.org/anthonynsimon/cconverter.svg?branch=master)](https://travis-ci.org/anthonynsimon/cconverter) 

CLI for currency exchange rate querying and conversion.

## Setup

Simply use `go get` on the repo to download and build from sources for your platform:

```
$ go get github.com/anthonynsimon/cconverter
$ cconverter
Usage: cconverter <flags> <subcommand> <subcommand args>

Subcommands:
        normalize        normalizes a CSV file so that all values are in the same currency
        convert          converts an amount from one currency to another
        rates            fetches the latest rates for the provided currency
        commands         list all command names
        flags            describe all known top-level flags
        help             describe subcommands and their syntax
```

Tested on Go 1.7.5

## Usage

Make sure that the [API server](https://github.com/anthonynsimon/cconverter-api) is already up and running at this point.

For convinience, the CLI assumes that the companion API server is running on `localhost:8080`. But you can config that, simply use the `-apiHost`
flag, like so: 
```
cconverter -apiHost=myhost:myport <subcommand> <subcommand args>
```

### Normalize CSV files

If someone happens to want to normalize the prices of a CSV file, and the layout looks
something like this:

```
item,price,currency
```

You can simply use the following command:

```
$ cconverter normalize -csvfile=FILE_PATH -to=CURRENCY -out=OPTIONAL_DESTINATION_PATH

# Help for normalize command
$ cconverter help normalize
normalize [-csvfile] [-to] [-out]:
        For ex.: -csvfile=./myfile.csv -to=USD -out=./normalized.csv
  -csvfile string
        the path to the csv file to be normalized
  -out string
        the file to which the normalized output will be saved to (default "./normalized.csv")
  -to string
        the target currency for the normalization
```

Example:

```
$ cat myfile.csv
A,500,EUR
B,750.5576,GBP
C,10.0,JPY
D,568.621,EUR
E,1,USD

$ cconverter normalize -csvfile=./myfile.csv -to=USD -out=./normalized.csv
Converting 500.0 EUR to USD
Converting 750.55760 GBP to USD
Converting 10.00 JPY to USD
Converting 568.621 EUR to USD
Converting 1.0 USD to USD

$ cat normalized.csv
A,526.8500,USD
B,931.9674,USD
C,0.0880,USD
D,599.1559,USD
E,1.0000,USD
```

### Convert between currencies

To convert from a base currency to any other supported currency:

```
$ cconverter convert -from=BASE_CURRENCY -to=TARGET_CURRENCY -amount=PRICE
```
Example:

```
$ cconverter convert -from=usd -to=eur -amount=100
--------------------------
From Currency:  USD
To Currency:    EUR
Amount:         100
Exchange Rate:  0.94904
Result:         94.904
--------------------------
```

### Query Exchange Rates

To query rates for any supported base currency:

```
$ cconverter rates -currencycode=CURRENCY
```
Example:

```
$ cconverter rates -currencycode=usd
--------------------------
Base Currency:  USD

BRL:            3.096
BGN:            1.8561
CZK:            25.644
RON:            4.285
CAD:            1.3135
GBP:            0.80535
ILS:            3.7023
JPY:            113.67
NZD:            1.4005
HRK:            7.0722
KRW:            1147.1
MXN:            20.431
PLN:            4.0869
PHP:            50.289
SEK:            8.9864
TRY:            3.6266
AUD:            1.3061
CNY:            6.8848
HKD:            7.7615
MYR:            4.4579
NOK:            8.3646
DKK:            7.0545
RUB:            57.738
SGD:            1.4225
THB:            35.035
ZAR:            13.16
CHF:            1.0097
HUF:            291.26
IDR:            13373
INR:            66.936
EUR:            0.94904
--------------------------
```

It currently supports the following currency codes (case insensitive):

`AUD`, `BGN`, `BRL`, `CAD`, `CHF`, `CNY`, `CZK`, `DKK`, `GBP`, `HKD`, `HRK`, `HUF`,
`IDR`, `ILS`, `INR`, `JPY`, `KRW`, `MXN`, `MYR`, `NOK`, `NZD`, `USD`, `PHP`, `PLN`,
`RON`, `RUB`, `SEK`, `SGD`, `THB`, `TRY`, `ZAR`, `EUR`.
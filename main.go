package main

import (
	"flag"
	"os"

	gbill "github.com/go-expenses/pkg"
)

type filterArgs []string

var (
	csvFile    string
	filterargs filterArgs
)

func (i *filterArgs) String() string {
	return ""
}

func (i *filterArgs) Set(arg string) error {
	*i = append(*i, arg)
	return nil
}

func init() {
	flag.StringVar(&csvFile, "csv", "", "CSV file to load into gBill")
	flag.Var(&filterargs, "filter", "Filters to be applied")
}

func main() {
	flag.Parse()

	wallet := &gbill.Wallet{}

	wallet.Go(csvFile, filterargs, os.Stdout)
}

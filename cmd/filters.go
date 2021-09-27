package cmd

import (
	"flag"
)

type flagArray []string

var (
	CsvFile    string
	FilterArgs flagArray
)

func (i *flagArray) String() string {
	return ""
}

func (i *flagArray) Set(arg string) error {
	*i = append(*i, arg)
	return nil
}

func init() {
	flag.StringVar(&CsvFile, "csv", "", "CSV file to load into gBill")
	flag.Var(&FilterArgs, "filter", "Filters to be applied")

	flag.Parse()
}

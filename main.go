package main

import (
	"flag"
	"os"

	gbill "github.com/go-expenses/pkg"
)

type flagArray []string

var (
	csvFile    string
	filterArgs flagArray
)

func (i *flagArray) String() string {
	return ""
}

func (i *flagArray) Set(arg string) error {
	*i = append(*i, arg)
	return nil
}

func init() {
	flag.StringVar(&csvFile, "csv", "", "CSV file to load into gBill")
	flag.Var(&filterArgs, "filter", "Filters to be applied")
}

func main() {
	flag.Parse()

	bills, err := gbill.LoadBillsFromCSV(csvFile)
	if err != nil {
		panic(err)
	}

	filters := NewBillFilters([]string{})

	wallet := &gbill.Wallet{
		Bills:   bills,
		Filters: filters,
		Output:  os.Stdout,
	}

	wallet.ShowMeTheMoney()
}

func NewBillFilters(options []string) []gbill.BillFilter {
	var billFilters []gbill.BillFilter

	for _, f := range filterArgs {
		switch name := f; name {
		case "top_by_category":
			billFilters = append(billFilters, gbill.NewTopByCategory(options))
		case "highest_paid":
			billFilters = append(billFilters, gbill.NewHighestPaid(options))
		case "lowest_paid":
			billFilters = append(billFilters, gbill.NewLowestPaid(options))
		}
	}

	return billFilters
}

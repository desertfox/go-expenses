package main

import (
	"os"

	"github.com/go-expenses/cmd"
	gbill "github.com/go-expenses/pkg"
)

func main() {

	bills, err := gbill.LoadBillsFromCSV(cmd.CsvFile)
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

	for _, f := range cmd.FilterArgs {
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

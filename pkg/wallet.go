package gbill

import (
	"io"
)

type Wallet struct {
	Bills   []*Bill
	Filters []BillFilter
	Output  io.Writer
}

type BillFilter interface {
	CalculateBill([]*Bill) []string
}

func (w *Wallet) Go(csvFile string, filterArgs []string, output io.Writer) {
	w.loadBills(csvFile)

	w.loadFilters(filterArgs)

	w.loadOutput(output)

	w.flashCash()
}

func (w *Wallet) loadBills(csvFile string) {
	bills, err := loadBillsFromCSV(csvFile)
	if err != nil {
		panic(err)
	}
	w.Bills = bills
}

func (w *Wallet) loadFilters(filterArgs []string) {
	filters := NewFiltersFromArgs(filterArgs)

	w.Filters = NewBillFilters(filters)
}

func NewBillFilters(filters []Filter) []BillFilter {
	var billFilters []BillFilter

	for _, f := range filters {
		switch name := f.Name; name {
		case "top_by_category":
			billFilters = append(billFilters, NewTopByCategory(f))
		case "highest_paid":
			billFilters = append(billFilters, NewHighestPaid(f))
		case "lowest_paid":
			billFilters = append(billFilters, NewLowestPaid(f))
		}
	}

	return billFilters
}

func (w *Wallet) loadOutput(o io.Writer) {
	w.Output = o
}

func (w *Wallet) flashCash() {
	lines := []string{}
	for _, bf := range w.Filters {
		newLines := bf.CalculateBill(w.Bills)
		lines = append(lines, newLines...)
	}

	for _, line := range lines {
		_, _ = w.Output.Write([]byte(line))
	}
}

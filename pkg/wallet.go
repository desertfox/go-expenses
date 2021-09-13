package gbill

import "fmt"

type Wallet struct {
	Bills   []*Bill
	Filters []BillFilter
	Output  Printer
}

type BillFilter interface {
	CalculateBill([]*Bill) []string
}

type Printer interface {
	PrintF(string, string)
}

func (w *Wallet) Go(csvFile string, filterArgs []string) {
	w.LoadBills(csvFile)

	w.LoadFilters(filterArgs)

	w.LoadOutput()

	w.FlashCash()
}

func (w *Wallet) LoadBills(csvFile string) {
	bills, err := loadBillsFromCSV(csvFile)
	if err != nil {
		panic(err)
	}
	w.Bills = bills
}

func (w *Wallet) LoadFilters(filterArgs []string) {
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

func (w *Wallet) LoadOutput() {
	//w.Output = o
}

func (w *Wallet) FlashCash() {
	lines := []string{}
	for _, bf := range w.Filters {
		newLines := bf.CalculateBill(w.Bills)
		lines = append(lines, newLines...)
	}

	for _, line := range lines {
		//w.Output.PrintF("%v", line)
		fmt.Printf("%v\n", line)
	}
}

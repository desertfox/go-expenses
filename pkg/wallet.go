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

func (w *Wallet) ShowMeTheMoney() {
	lines := []string{}
	for _, bf := range w.Filters {
		newLines := bf.CalculateBill(w.Bills)
		lines = append(lines, newLines...)
	}

	for _, line := range lines {
		_, _ = w.Output.Write([]byte("\n"))
		_, _ = w.Output.Write([]byte(line))
	}
}

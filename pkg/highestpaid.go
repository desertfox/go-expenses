package gbill

import (
	"bytes"
	"text/template"
)

type HighestPaid struct {
	Filter
}

func NewHighestPaid(f Filter) BillFilter {
	return &HighestPaid{f}
}

func (hp *HighestPaid) CalculateBill(bills []*Bill) []string {
	highestBill := *bills[0]
	for _, bill := range bills {
		if bill.Amount > highestBill.Amount {
			highestBill = *bill
		}
	}

	const format = `{{ .Name }}:{{ .PaidOn }}:{{ .Category }}:{{ .Amount }}`
	t := template.Must(template.New("line").Parse(format))

	buf := bytes.NewBufferString("")
	err := t.Execute(buf, highestBill)
	if err != nil {
		panic(err)
	}

	return []string{buf.String()}
}

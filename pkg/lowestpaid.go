package gbill

import (
	"bytes"
	"text/template"
)

type LowestPaid struct {
	Filter
}

func NewLowestPaid(f Filter) BillFilter {
	return &LowestPaid{f}
}

func (lp *LowestPaid) CalculateBill(bills []*Bill) []string {
	lowestBill := *bills[0]
	for _, bill := range bills {
		if bill.Amount < lowestBill.Amount {
			lowestBill = *bill
		}
	}

	const format = `{{ .Name }}:{{ .PaidOn }}:{{ .Category }}:{{ .Amount }}`
	t := template.Must(template.New("line").Parse(format))

	buf := bytes.NewBufferString("")
	err := t.Execute(buf, lowestBill)
	if err != nil {
		panic(err)
	}

	return []string{buf.String()}
}

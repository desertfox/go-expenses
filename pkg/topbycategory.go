package gbill

import (
	"bytes"
	"text/template"
)

type CategoryTotal struct {
	Total    int
	Count    int
	Category string
}

func newCategoryTotal(c string) *CategoryTotal {
	return &CategoryTotal{0, 0, c}
}

func (ct *CategoryTotal) add(a int) {
	ct.Total = ct.Total + a
}

func (ct *CategoryTotal) addCount() {
	ct.Count++
}

type TopByCategory struct {
	Filter
}

func NewTopByCategory(f Filter) BillFilter {
	return &TopByCategory{f}
}

func (tbc *TopByCategory) CalculateBill(bills []*Bill) []string {
	cts := make(map[string]*CategoryTotal)

	for _, bill := range bills {
		if cts[bill.Category] == nil {
			cts[bill.Category] = newCategoryTotal(bill.Category)
		}

		(*cts[bill.Category]).add(bill.Amount)
		(*cts[bill.Category]).addCount()
	}

	const format = `{{ .Category }}:{{ .Count }}:{{ .Total }}`
	t := template.Must(template.New("line").Parse(format))
	//Sort by options
	lines := make([]string, len(cts))
	for _, billTotal := range cts {
		buf := bytes.NewBufferString("")
		err := t.Execute(buf, billTotal)
		if err != nil {
			panic(err)
		}
		lines = append(lines, buf.String())
	}

	return lines
}

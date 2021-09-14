package gbill

type Filter struct {
	Name    string
	Options []string
}

func NewFilter(n string, o []string) Filter {
	return Filter{n, o}
}

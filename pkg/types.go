package gbill

type Bill struct {
	name      string
	category  string
	frequency string
	amount    int64
	dueOn     time.Time
	paidOn    time.Time
}

type Bills struct {
	Bills []Bill
}

package gbill

func NewWallet() *Wallet {
	return &Wallet{
		Bills: []*Bill{},
	}
}

func (w *Wallet) AddBuild(b *Bill) {
	w.Bills = append(w.Bills, b)
}

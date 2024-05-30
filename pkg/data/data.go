package data

type Data struct {
	BillsDue           int
	Date               string
	Selected           bool
	ConfirmationNumber string
}

func NewData() *Data {
	return &Data{}
}

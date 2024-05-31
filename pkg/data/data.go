package data

type Data struct {
	BillsDue           int
	Date               string
	Selected           bool
	ConfirmationNumber int
	Message            string
}

func NewData() *Data {
	return &Data{}
}

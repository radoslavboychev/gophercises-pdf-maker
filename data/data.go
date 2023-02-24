package data

type Data struct {
	UnitName       string
	PricePerUnit   float64
	UnitsPurchased float64
}

type CompanyInfo struct {
	PhoneNum       string
	Email          string
	Website        string
	CompanyAddress Address
}

type Address struct {
	Street    string
	CityState string
	PostCode  int
}

type ClientData struct {
	Name          string
	Address       string
	CityState     string
	PostCode      string
	InvoiceNumber string
	DateOfIssue   string
}

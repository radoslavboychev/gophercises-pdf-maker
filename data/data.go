package data

type Data struct {
	UnitName       string
	PricePerUnit   int
	UnitsPurchased int
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

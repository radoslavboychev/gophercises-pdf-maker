package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"github.com/radoslavboychev/gophercises-pdf-maker/data"
	"github.com/radoslavboychev/gophercises-pdf-maker/render"
)

var myData = []data.Data{
	{
		UnitName:       "2x6 Lumber - 8'",
		PricePerUnit:   3.75,
		UnitsPurchased: 220,
	}, {
		UnitName:       "Drywall Sheet",
		PricePerUnit:   8.22,
		UnitsPurchased: 50,
	}, {
		UnitName:       "Paint",
		PricePerUnit:   14.55,
		UnitsPurchased: 3,
	},
	{
		UnitName:       "Aluminum Coating",
		PricePerUnit:   150,
		UnitsPurchased: 2,
	},
}

var myCompany = data.CompanyInfo{
	PhoneNum: "(814) 977-7556",
	Email:    "jon@calhoun.io",
	Website:  "Gophercises.com",
	CompanyAddress: data.Address{
		Street:    "123 Fake St",
		CityState: "Some Town, PA",
		PostCode:  12345,
	},
}

var myClient = data.ClientData{
	Name:          "Client Name",
	Address:       "1 Client Address",
	CityState:     "City,State,Country",
	PostCode:      "POSTAL CODE",
	InvoiceNumber: "00000000123",
	DateOfIssue:   "05/29/2018",
}

func main() {

	dataAddress := fmt.Sprintf(myCompany.CompanyAddress.Street + "\n" + myCompany.CompanyAddress.CityState + "\n" + fmt.Sprint(myCompany.CompanyAddress.PostCode))
	dataPhoneNo := fmt.Sprintf("%v\n%v\n%v", myCompany.PhoneNum, myCompany.Email, myCompany.Website)

	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := pdf.GetPageSize()
	fmt.Printf("Generated PDF!")
	pdf.AddPage()

	// render the top banner
	render.DrawBanner(pdf, w, dataPhoneNo, dataAddress)

	// render the labels for the top half
	render.DrawTop(pdf, w, h, myClient, myData)

	// render miscellaneous assets
	render.DrawAssets(pdf)

	// render all products from order line by line
	render.RenderLine(pdf, myData)

	err := pdf.OutputFileAndClose("p1.pdf")
	if err != nil {
		panic(err)
	}
}

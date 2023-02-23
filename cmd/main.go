package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"github.com/radoslavboychev/gophercises-pdf-maker/data"
)

var myData = []data.Data{
	{
		UnitName:       "2x6 Lumber - 8'",
		PricePerUnit:   375,
		UnitsPurchased: 220,
	}, {
		UnitName:       "Drywall Sheet",
		PricePerUnit:   822,
		UnitsPurchased: 50,
	}, {
		UnitName:       "Paint",
		PricePerUnit:   1455,
		UnitsPurchased: 3,
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

func main() {

	dataAddress := fmt.Sprintf(myCompany.CompanyAddress.Street + "\n" + myCompany.CompanyAddress.CityState + "\n" + fmt.Sprint(myCompany.CompanyAddress.PostCode))
	dataPhoneNo := fmt.Sprintf(myCompany.PhoneNum + "\n" + myCompany.Email + "\n" + myCompany.Website)

	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := pdf.GetPageSize()
	fmt.Printf("width=%v, height=%v\n", w, h)
	pdf.AddPage()

	// COMPANY INFO 2 -  ADDRESS BAR

	renderBanner(pdf, w, dataPhoneNo, dataAddress)

	// pdf.CellFormat(0, 18, myCompany.CompanyAddress.Street+"\n"+myCompany.CompanyAddress.CityState+"\n"+fmt.Sprint(myCompany.CompanyAddress.PostCode), "", "", false)
	err := pdf.OutputFileAndClose("p1.pdf")
	if err != nil {
		panic(err)
	}
}

// RENDER THE BANNER ON TOP
func renderBanner(pdf *gofpdf.Fpdf, w float64, dataPhoneNo string, dataAddress string) {

	// PURPLE RECTANGLE
	pdf.SetFillColor(102, 61, 79)
	pdf.SetDrawColor(102, 61, 79)
	pdf.Rect(0, 0, w, 100, "FD")
	pdf.SetFillColor(100, 200, 200)

	// "INVOICE" text
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("arial", "B", 40)
	pdf.Text(40, 68, "INVOICE")

	// COMPANY INFO 1 - PHONE
	pdf.SetFont("arial", "B", 16)
	pdf.SetXY(340, 30)
	pdf.MultiCell(0, 17, dataPhoneNo, "", "", false)

	// COMPANY INFO 2 -  ADDRESS BAR
	pdf.SetFont("arial", "B", 13)
	pdf.MoveTo(485, 30)
	pdf.MultiCell(0, 14, dataAddress, "", gofpdf.AlignRight, false)

	// IMAGE MIDDLE
	pdf.MoveTo(340, 30)
	pdf.ImageOptions(".././images/gopher.png", 240, 12, 53, 0, false, gofpdf.ImageOptions{
		ReadDpi: true,
	}, 0, "")
}

func drawGrid(pdf *gofpdf.Fpdf) {
	w, h := pdf.GetPageSize()
	pdf.SetFont("courier", "", 12)
	pdf.SetTextColor(80, 80, 80)
	pdf.SetDrawColor(200, 200, 200)
	for x := 0.0; x < w; x = x + (w / 20.0) {
		pdf.Line(x, 0, x, h)
		_, lineHt := pdf.GetFontSize()
		pdf.Text(x, lineHt, fmt.Sprintf("%d", int(x)))
	}
	for y := 0.0; y < h; y = y + (w / 20.0) {
		pdf.Line(0, y, w, y)
		pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}
}

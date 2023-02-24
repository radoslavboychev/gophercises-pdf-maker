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
	fmt.Printf("width=%v, height=%v\n", w, h)
	pdf.AddPage()

	// render the top banner
	renderBanner(pdf, w, dataPhoneNo, dataAddress)

	// render the labels for the top half
	renderTop(pdf)

	// render miscellaneous assets
	renderAssets(pdf)

	renderProducts(pdf)

	err := pdf.OutputFileAndClose("p1.pdf")
	if err != nil {
		panic(err)
	}
}

func renderProducts(pdf *gofpdf.Fpdf) {
	pdf.SetTextColor(64, 64, 64)
	pdf.SetFont("times", "", 16)

	pdf.Text(40, 326+40, "2x6 Lumber - 8'")
	pdf.Text(374, 326+40, "$3.75")
	pdf.Text(469, 326+40, "220")
	pdf.Text(530, 326+40, "$825.00")
	drawLine(pdf, 340+40)

	pdf.Text(40, 326+80, "2x6 Lumber - 10'")
	pdf.Text(374, 326+80, "$5.55")
	pdf.Text(469, 326+80, "18")
	pdf.Text(530, 326+80, "$99.90")
	drawLine(pdf, 340+80)

	pdf.Text(40, 326+120, "2x4 Lumber - 8'")
	pdf.Text(374, 326+120, "$2.99")
	pdf.Text(469, 326+120, "80")
	pdf.Text(530, 326+120, "$239.20")
	drawLine(pdf, 340+120)

	pdf.Text(40, 326+160, "Drywall Sheet")
	pdf.Text(374, 326+160, "$8.22")
	pdf.Text(469, 326+160, "50")
	pdf.Text(530, 326+160, "$411.00")
	drawLine(pdf, 340+160)

	pdf.Text(40, 326+200, "Paint")
	pdf.Text(374, 326+200, "$14.55")
	pdf.Text(469, 326+200, "3")
	pdf.Text(530, 326+200, "$43.65")
	drawLine(pdf, 340+200)

	pdf.Text(40, 326+240, "Paint")
	pdf.Text(374, 326+240, "$14.55")
	pdf.Text(469, 326+240, "3")
	pdf.Text(530, 326+240, "$43.65")
	drawLine(pdf, 340+240)
}

func drawLine(pdf *gofpdf.Fpdf, y float64) {
	pdf.SetFillColor(64, 64, 64)
	pdf.SetDrawColor(64, 64, 64)
	pdf.Rect(17, y, 570, 1, "FD")
}

// RENDER MISCELLANEOUS ASSETS
func renderAssets(pdf *gofpdf.Fpdf) {
	pdf.SetFillColor(102, 61, 79)
	pdf.SetDrawColor(102, 61, 79)
	pdf.Rect(17, 270, 570, 3, "FD")
	pdf.SetFillColor(100, 200, 200)
}

// RENDER THE LABELS
func renderTop(pdf *gofpdf.Fpdf) {

	// 'description' label
	pdf.SetTextColor(198, 198, 198)
	pdf.SetFont("times", "", 16)
	pdf.Text(40, 305, "Description")

	// 'Price Per Unit' label
	pdf.SetTextColor(198, 198, 198)
	pdf.SetFont("times", "", 16)
	pdf.Text(320, 305, "Price Per Unit")

	// 'Subtotal' label
	pdf.SetTextColor(198, 198, 198)
	pdf.SetFont("times", "", 16)
	pdf.Text(340, 630, "Subtotal")

	// 'Quantity' label
	pdf.SetTextColor(198, 198, 198)
	pdf.SetFont("times", "", 16)
	pdf.Text(440, 305, "Quantity")

	// 'Amount' label
	pdf.SetTextColor(64, 64, 64)
	pdf.SetFont("times", "", 16)
	pdf.Text(526, 305, "Amount")

	pdf.Text(523, 630, "$1838.53")

	// BILLED TO

	// 'billed to' label
	pdf.SetTextColor(198, 198, 198)
	pdf.SetFont("times", "", 16)
	pdf.Text(40, 150, "Billed To")

	// 'billed to' text
	pdf.SetTextColor(64, 64, 64)
	pdf.SetFont("times", "", 14)
	pdf.SetXY(38, 158)
	pdf.MultiCell(0, 19, myClient.Name+"\n"+myClient.Address+"\n"+myClient.CityState+"\n"+myClient.PostCode, "", "", false)

	// INVOICE NUMBER

	// 'invoice number' label
	pdf.SetTextColor(198, 198, 198)
	pdf.SetFont("times", "", 16)
	pdf.Text(240, 150, "Invoice Number")

	// 'invoice number' text
	pdf.SetTextColor(64, 64, 64)
	pdf.SetFont("arial", "", 14)
	pdf.Text(240, 170, myClient.InvoiceNumber)

	// DATE OF ISSUE

	// 'date of issue' label
	pdf.SetTextColor(198, 198, 198)
	pdf.SetFont("times", "", 16)
	pdf.Text(240, 210, "Date of Issue")

	// 'date of issue' text
	pdf.SetTextColor(64, 64, 64)
	pdf.SetFont("arial", "", 14)
	pdf.Text(240, 230, myClient.DateOfIssue)

	// INVOICE TOTAL

	// 'invoice total' label
	pdf.SetTextColor(198, 198, 198)
	pdf.SetFont("times", "", 16)
	pdf.Text(480, 150, "Invoice Total")

	// 'invoice total' text
	pdf.SetTextColor(102, 61, 79)
	pdf.SetFont("times", "", 50)
	pdf.Text(385, 200, "$1838.53")
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
	pdf.SetFont("arial", "", 16)
	pdf.SetXY(340, 30)
	pdf.MultiCell(0, 18, dataPhoneNo, "", "", false)

	// COMPANY INFO 2 -  ADDRESS BAR
	pdf.SetFont("arial", "", 13)
	pdf.MoveTo(485, 30)
	pdf.MultiCell(0, 14, dataAddress, gofpdf.BorderNone, gofpdf.AlignRight, false)

	// IMAGE MIDDLE
	pdf.ImageOptions(".././images/gopher.png", 240, 14, 53, 0, false, gofpdf.ImageOptions{
		ReadDpi: true,
	}, 0, "")
}

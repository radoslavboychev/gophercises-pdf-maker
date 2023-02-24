package render

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"github.com/radoslavboychev/gophercises-pdf-maker/data"
	"github.com/radoslavboychev/gophercises-pdf-maker/util"
)

const (
	xIndent = 40.0
)

// draws the divider lines
func DrawLine(pdf *gofpdf.Fpdf, y float64) {
	pdf.SetFillColor(64, 64, 64)
	pdf.SetDrawColor(64, 64, 64)
	pdf.Rect(17, y, 570, 1, "FD")
}

// RENDER MISCELLANEOUS ASSETS
func DrawAssets(pdf *gofpdf.Fpdf) {
	pdf.SetFillColor(102, 61, 79)
	pdf.SetDrawColor(102, 61, 79)
	pdf.Rect(17, 270, 570, 3, "FD")
	pdf.SetFillColor(100, 200, 200)
}

// RENDER THE LABELS
func DrawTop(pdf *gofpdf.Fpdf, w, h float64, myClient data.ClientData, myData []data.Data) {
	// _, lineHt := pdf.GetFontSize()
	// 'description' label
	pdf.SetTextColor(198, 198, 198)
	pdf.SetFont("times", "", 16)
	pdf.CellFormat(w-(xIndent*2), 430, "Description", gofpdf.BorderNone, 0, gofpdf.AlignLeft, false, 0, "")

	// 'Price Per Unit' label
	pdf.SetTextColor(198, 198, 198)
	pdf.SetFont("times", "", 16)
	// pdf.CellFormat(w-(xIndent), 430, "Price Per Unit", gofpdf.BorderNone, 0, "", false, 0, "")
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

	pdf.Text(523, 630, fmt.Sprintf("$%.2f", calculateSubtotal(myData)))

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
	pdf.Text(385, 200, fmt.Sprintf("$%.2f", calculateSubtotal(myData)))
}

// RENDER THE BANNER ON TOP
func DrawBanner(pdf *gofpdf.Fpdf, w float64, dataPhoneNo string, dataAddress string) {

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

func RenderLine(pdf *gofpdf.Fpdf, myData []data.Data) {
	pdf.SetTextColor(64, 64, 64)
	pdf.SetFont("times", "", 16)
	var offset float64 = 40.00
	var priceX float64 = 374
	var unitsPurchasedX float64 = 469
	var amountX float64 = 530

	for _, line := range myData {
		pdf.Text(40, 326+offset, line.UnitName)
		pdf.Text(priceX, 326+offset, util.Currencyfy(line.PricePerUnit))
		pdf.Text((unitsPurchasedX), 326+offset, fmt.Sprint(line.UnitsPurchased))
		pdf.Text((amountX), 326+offset, util.Currencyfy(line.UnitsPurchased*line.PricePerUnit))
		DrawLine(pdf, 340+offset)
		offset += 40

	}
}

func calculateSubtotal(myData []data.Data) float64 {
	var sum float64
	for _, d := range myData {
		sum += float64(d.PricePerUnit * d.UnitsPurchased)
	}
	return sum
}

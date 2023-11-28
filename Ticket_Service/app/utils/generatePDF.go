package utils

import (
	"fmt"
	"os"

	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/common/structs"
	"github.com/jung-kurt/gofpdf"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func GeneratePDF(success bool, userId string, bookingId string, data interface{}) (string, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 32)
	pdf.Write(16, "Tessera\n")

	if success {
		p := message.NewPrinter(language.Indonesian)

		successData := data.(commonStructs.SuccessPDFPayload)
		pdf.SetFont("Arial", "", 18)
		pdf.Write(16, "Payment Receipt\n")
		pdf.SetFont("Arial", "", 12)
		pdf.Cell(30, 10, "Status")
		pdf.Write(10, ": Success\n")
		pdf.Cell(30, 10, "Booking ID")
		pdf.Write(10, fmt.Sprintf(": %s\n", bookingId))
		pdf.Cell(30, 10, "Seat")
		pdf.Write(10, fmt.Sprintf(": %s\n", successData.Seat))
		pdf.Cell(30, 10, "Total Price")
		pdf.Write(10, p.Sprintf(": Rp%d,00\n", successData.Price))

		// Generate QR
		GenerateQR(bookingId)

		pdf.ImageOptions(fmt.Sprintf("./out/%s.png", bookingId), 10, 90, 50, 0, false, gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true}, 0, "")
	} else {
		failedData := data.(commonStructs.FailedPDFPayload)
		pdf.SetFont("Arial", "", 18)
		pdf.Write(16, "Failure Occurs\n")
		pdf.SetFont("Arial", "B", 12)
		pdf.Cell(20, 10, "Error")
		pdf.SetFont("Arial", "", 12)
		pdf.Cell(0, 10, fmt.Sprintf(": %s\n", failedData.ErrorMessage))
	}

	pdfPath := "./public"
	pdfName := fmt.Sprintf("/reports/%s_%s.pdf", userId, bookingId)

	if err := pdf.OutputFileAndClose(fmt.Sprintf("%s%s", pdfPath, pdfName)); err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	if success {
		// delete qr image
		if err := os.Remove(fmt.Sprintf("./out/%s.png", bookingId)); err != nil {
			fmt.Println(err)
			return "", err
		}
	}

	return pdfName, nil
}

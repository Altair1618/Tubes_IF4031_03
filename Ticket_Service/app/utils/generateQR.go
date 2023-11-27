package utils

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

func GenerateQR(data string) error {
	qrCode, _ := qrcode.New(data, qrcode.Medium)
	fileName := fmt.Sprintf("./out/%s.png", data)

	if err := qrCode.WriteFile(256, fileName); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

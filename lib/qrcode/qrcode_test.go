package main

import (
	"fmt"
	"testing"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

// go get -u github.com/yeqown/go-qrcode/v2
func TestQrCodeWithLogo(t *testing.T) {
	// defer func() {
	// 	if i := recover(); i != nil {
	// 		t.Fatal(i)
	// 	}
	// }()
	qrc, err := qrcode.New("https://guangchang.tech")
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
		return
	}
	w0, err := standard.New("result.png",
		standard.WithHalftone("icon.png"),
		standard.WithQRWidth(23),
	)
	if err != nil {
		fmt.Println("could not new")
	}

	if err := qrc.Save(w0); err != nil {
		fmt.Printf("could not save QRCode: %v", err)
		return
	}
}

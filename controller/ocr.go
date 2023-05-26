package controller

import (
	"net/http"
	"os/exec"
)

func Ocr(w http.ResponseWriter, r *http.Request) {
	// client := gosseract.NewClient()
	// defer client.Close()
	// client.SetImage("./Surat Dokter.jpg")
	// text, _ := client.Text()
	// fmt.Println(text)
	exec.Command(`C:\Program Files (x86)\Tesseract-OCR\tesseract.exe`, "./images/photo_ktp.jpg", "output")
}

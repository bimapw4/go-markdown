package main

import (
	"fmt"
	"go-markdown/router"
	"go-markdown/serivce/response"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	port := ":9000"

	cmd := exec.Command(`C:\Program Files (x86)\Tesseract-OCR\tesseract.exe`, "./images/photo_ktp.jpg", "output")
	fmt.Println(cmd)

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	}
	fmt.Println("Output: ", string(out))

	fmt.Println("server is running on port", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		response.NewResponse().WithCode(http.StatusOK).WithStatus("success").WithData(map[string]interface{}{
			"message": "hello world",
		}).ParseResponse(w, r)
	})

	router.Router()

	err = http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal(err)
	}
}

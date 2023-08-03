package main

import (
	"fmt"
	"go-markdown/router"
	"go-markdown/serivce/response"
	"log"
	"net/http"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func main() {
	port := ":9000"

	// cmd := exec.Command(`C:\Users\Lenovo\go\bin\pdfcpu.exe`, "stamp", "add", "-mode", "text", "--", "This is a stamp", "sc:1", "in.pdf", "out.pdf")
	// fmt.Println(cmd)

	// out, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("could not run command: ", err)
	// }
	// fmt.Println("Output: ", string(out))

	// Add a "Demo" stamp to all pages of in.pdf along the diagonal running from lower left to upper right.
	onTop := false
	update := false

	wm, _ := api.TextWatermark("Footer stamp", "c:.5 1 1, pos:bc", onTop, update, types.POINTS)
	api.AddWatermarksFile("in.pdf", "out.pdf", nil, wm, nil)

	fmt.Println("server is running on port", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		response.NewResponse().WithCode(http.StatusOK).WithStatus("success").WithData(map[string]interface{}{
			"message": "hello world",
		}).ParseResponse(w, r)
	})

	router.Router()

	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal(err)
	}
}

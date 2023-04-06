package main

import (
	"fmt"
	"go-markdown/router"
	"go-markdown/serivce/response"
	"log"
	"net/http"
)

func main() {
	port := ":9000"

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

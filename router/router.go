package router

import (
	"go-markdown/controller"
	"net/http"
)

func Router() {

	http.HandleFunc("/html-parse", controller.HtmlParseHandler)
	http.HandleFunc("/markdown-parse", controller.MarkdownParseHandler)
	http.HandleFunc("/generate", controller.GenerateKey)
	http.HandleFunc("/encrypt", controller.Encrypt)
	http.HandleFunc("/decrypt", controller.Decrypt)
	http.HandleFunc("/decrypt-aes", controller.DecryptAES)

	http.HandleFunc("/lalal", controller.TestBack)

}

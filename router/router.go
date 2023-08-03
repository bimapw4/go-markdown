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

	http.HandleFunc("/ocr", controller.Ocr)
	http.HandleFunc("/face", controller.Face)
	http.HandleFunc("/pigo", controller.Pigo)

	http.HandleFunc("/lalal", controller.TestBack)

	http.HandleFunc("/test-encrypt-aes", controller.TestEncryptAes)
	http.HandleFunc("/test-decrypt-aes", controller.TestingDecryptAes)

	http.HandleFunc("/localization", controller.Localization)

	http.HandleFunc("/mustache", controller.TestMustache)

}

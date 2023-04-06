package controller

import (
	"encoding/json"
	"go-markdown/serivce/response"
	"net/http"

	"github.com/gomarkdown/markdown"
)

type payloadMarkdown struct {
	Markdown string `json:"markdown"`
}

func MarkdownParseHandler(w http.ResponseWriter, r *http.Request) {
	payload := payloadMarkdown{}
	json.NewDecoder(r.Body).Decode(&payload)

	md := []byte(payload.Markdown)
	html := markdown.ToHTML(md, nil, nil)
	response.NewResponse().WithCode(http.StatusAccepted).WithData(map[string]interface{}{
		"html": string(html),
	}).ParseResponse(w, r)
}

package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Code   int
	Status string                 `json:"status,omitempty"`
	Data   map[string]interface{} `json:"data,omitempty"`
}

func (res *Response) WithCode(code int) *Response {
	res.Code = code

	return res
}

func (res *Response) WithStatus(status string) *Response {
	res.Status = status
	return res
}

func (res *Response) WithData(data map[string]interface{}) *Response {
	res.Data = data
	return res
}

func (res *Response) ParseResponse(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(res.Data)

	if err != nil {
		fmt.Errorf(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(json)
}

func NewResponse() *Response {

	return &Response{}
}

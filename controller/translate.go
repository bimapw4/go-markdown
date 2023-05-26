package controller

import (
	"fmt"
	"go-markdown/serivce/response"
	"go-markdown/serivce/translate"
	"net/http"
)

func Localization(w http.ResponseWriter, r *http.Request) {
	translate, err := translate.LocalizationTemplate(translate.Params{
		Tag:  "login",
		Lang: "id",
	})

	if err != nil {
		fmt.Println(err)
	}

	response.NewResponse().WithCode(http.StatusAccepted).WithData(map[string]interface{}{
		"translate": translate,
	}).ParseResponse(w, r)
}

package controller

import (
	"encoding/json"
	"fmt"
	"go-markdown/serivce/response"
	"log"
	"math/rand"
	"net/http"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/avast/retry-go"
)

type payloadHtml struct {
	Html string `json:"html"`
}

func HtmlParseHandler(w http.ResponseWriter, r *http.Request) {

	payload := payloadHtml{}
	json.NewDecoder(r.Body).Decode(&payload)

	opt := &md.Options{
		EmDelimiter: "*", // default: **
	}

	converter := md.NewConverter("", true, opt)

	markdown, err := converter.ConvertString(payload.Html)

	if err != nil {
		log.Println("error disini", err)
	}

	response.NewResponse().WithCode(http.StatusAccepted).WithData(map[string]interface{}{
		"markdown": markdown,
	}).ParseResponse(w, r)

}

func TestBack(w http.ResponseWriter, r *http.Request) {
	// client := &http.Client{
	// 	Timeout: time.Second * 5,
	// }
	// Membuat fungsi request HTTP yang akan digunakan

	request := func() error {
		// resp, err := client.Get("https://be-kopimage-v2.dcidev.id/jajal")
		// if err != nil {
		// 	return err
		// }
		// defer resp.Body.Close()

		// if resp.StatusCode != http.StatusOK {
		// 	fmt.Printf("response status body %d\n", resp.Body)
		// 	return fmt.Errorf("response status body %d", resp.Body)
		// }

		// fmt.Println("Response OK")

		min := 0
		max := 3
		i := rand.Intn(max - min)
		if i != 1 {
			fmt.Printf("response status body %d\n", i)
			// 	return fmt.Errorf("response status body %d", resp.Body)
			return fmt.Errorf("response status body %d", i)
		}
		return nil
	}

	err := retry.Do(request, retry.Delay(time.Second), retry.Attempts(4))
	if err != nil {
		fmt.Println("Request failed:", err)
	}
	fmt.Println("Request succeeded")
}

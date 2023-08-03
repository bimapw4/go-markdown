package controller

import (
	"encoding/json"
	"fmt"
	"go-markdown/serivce/response"
	"net/http"
	"regexp"
	"strings"

	"github.com/cbroglie/mustache"
)

var dummy = `{
	"logo":"",
	"subject":"full boss",
	"title":{
		"title_id":"ini itu full",
		"title_en":"this is full"
	},
	"button":{
		"button_active":true,
		"button_link":"https://encrypted-tbn0.gstatic.com/",
		"button_text":{
			"button_text_id":"sasa",
			"button_text_en":"sasa"
		}
	},
	"image":{
		"image_active":false,
		"image_url":"https://chat.openai.com/"
	},
	"paragraph":[
		{
			"id":"<b style='color:#0000'>sasasasassa paragraph</b>",
			"en":"<b style='color:red'>sasasasassa paragraph</b>"
		},
		{
			"id":"",
			"en":""
		}
	],
	"footer":[
		{
			"id":"<b style='color:#0000'>sasasasassa footer</b>",
			"en":"<b style='color:red'>sasasasassa footer</b>"
		},
		{
			"id":"",
			"en":""
		}
	]
}`

type (
	// AdditionalDataFormEmailContent struct {
	// 	ID            string         `json:"id"`
	// 	ApplicationID string         `json:"application_id"`
	// 	EmailType     string         `json:"email_type"`
	// 	EmailContent  types.JSONText `json:"email_content"`
	// 	UploadBy      string         `json:"upload_by"`
	// 	IsActive      bool           `json:"is_active"`
	// 	CreatedAt     time.Time      `json:"created_at"`
	// 	UpdatedAt     time.Time      `json:"updated_at"`
	// }

	// PayloadAdditionalDataFormEmailContent struct {
	// 	ID            string          `json:"id"`
	// 	PrivyID       string          `json:"privy_id"`
	// 	ApplicationID string          `json:"application_id"`
	// 	EmailType     string          `json:"email_type"`
	// 	EmailContent  SubEmailContent `json:"email_content"`
	// 	UploadBy      string          `json:"upload_by"`
	// 	IsActive      bool            `json:"is_active"`
	// 	CreatedAt     time.Time       `json:"created_at"`
	// 	UpdatedAt     time.Time       `json:"updated_at"`
	// }

	SubEmailContent struct {
		Logo      string                   `json:"logo"`
		Subject   string                   `json:"subject"`
		Title     TitleMultiLanguage       `json:"title"`
		Button    SubButton                `json:"button"`
		Image     SubImage                 `json:"image"`
		Paragraph []ParagraphMultiLanguage `json:"paragraph"`
		Footer    []FooterMultiLanguage    `json:"footer"`
	}

	SubEmail struct {
		ApplicationName string                   `json:"application_name"`
		Name            string                   `json:"name"`
		MerchantName    string                   `json:"merchant_name"`
		Logo            string                   `json:"logo"`
		Subject         string                   `json:"subject"`
		Title           TitleMultiLanguage       `json:"title"`
		Button          SubButton                `json:"button"`
		Image           SubImage                 `json:"image"`
		Paragraph       []ParagraphMultiLanguage `json:"paragraph"`
		Footer          []FooterMultiLanguage    `json:"footer"`
		DataFormFilled  bool                     `json:"data_form_filled"`
	}

	TitleMultiLanguage struct {
		TitleEn string `json:"title_en"`
		TitleId string `json:"title_id"`
	}

	SubButton struct {
		ButtonActive bool                   `json:"button_active"`
		ButtonLink   string                 `json:"button_link"`
		ButtonText   SubButtonMultilanguage `json:"button_text"`
	}

	SubButtonMultilanguage struct {
		ButtonTextID string `json:"button_text_id"`
		ButtonTextEN string `json:"button_text_en"`
	}

	SubImage struct {
		ImageActive bool   `json:"image_active"`
		ImageUrl    string `json:"image_url"`
	}

	ParagraphMultiLanguage struct {
		Id string `json:"id"`
		En string `json:"en"`
	}

	FooterMultiLanguage struct {
		Id string `json:"id"`
		En string `json:"en"`
	}
)

func convertCustomTag(str string) string {
	customtag := map[string]string{
		"<*":  "<b>",
		"<_":  "<i>",
		"<=":  "<u>",
		"<~":  "<del>",
		"<@":  "<p>",
		"</*": "</b>",
		"</_": "</i>",
		"</=": "</u>",
		"</~": "</del>",
		"</@": "</p>",
	}

	for i, v := range customtag {
		str = strings.Replace(str, i, v, -1)
	}

	return str
}

func convertCustomStylingOld(template string, tag ...string) string {

	for _, v := range tag {

		regexPattern := map[string]string{
			"regexstyle1":       `{{color:(\w+)}}`,
			"regexstyle2":       `{{color:(#\w+)}}`,
			"regextag":          `<(\w+)>`,
			"regexcontent":      `\[([^]]+)\]`,
			"regexreplaceinput": `(<[^>]+>.*?<\/[^>]+>)`,
		}

		resultRgx := map[string]string{}

		for i, vp := range regexPattern {
			switch i {
			case "regexstyle1", "regexstyle2":

				regexstyle := regexPattern["regexstyle1"]
				if strings.Contains(v, "{{color:#") {
					regexstyle = regexPattern["regexstyle2"]
				}

				repattern := regexp.MustCompile(regexstyle)
				matchpattern := repattern.FindStringSubmatch(v)
				resultRgx["regexstyle"] = ""
				if len(matchpattern) > 0 {
					resultRgx["regexstyle"] = matchpattern[1]
				}

			default:
				repattern := regexp.MustCompile(vp)
				matchpattern := repattern.FindStringSubmatch(v)
				resultRgx[i] = ""
				if len(matchpattern) > 0 {
					resultRgx[i] = matchpattern[len(matchpattern)-1]
				}
			}
		}

		var tagOpen, tagClose string
		if resultRgx["regextag"] != "" {
			tagOpen = fmt.Sprintf("<%s>", resultRgx["regextag"])
			tagClose = fmt.Sprintf("</%s>", resultRgx["regextag"])
		}

		if resultRgx["regexstyle"] != "" && resultRgx["regextag"] != "" {
			tagOpen = fmt.Sprintf(`<%s style="color:%s">`, resultRgx["regextag"], resultRgx["regexstyle"])
		}

		content := fmt.Sprintf(`%s%s%s`, tagOpen, resultRgx["regexcontent"], tagClose)
		if resultRgx["regexreplaceinput"] != "" {
			content = strings.Replace(v, resultRgx["regexreplaceinput"], fmt.Sprintf(`%s%s%s`, tagOpen, resultRgx["regexcontent"], tagClose), -1)
		}

		template = strings.Replace(template, v, content, -1)
	}

	return template
}

func convertCustomStyling(str string) string {

	pattern := `\{\{.*?\}\}`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(str, -1)

	for i := 0; i < len(matches); i++ {

		regexPattern := map[string]string{
			"regexstyle1":       `{{color:(\w+)}}`,
			"regexstyle2":       `{{color:(#\w+)}}`,
			"regextag":          `<(\w+)>`,
			"regexcontent":      `\[([^]]+)\]`,
			"regexreplaceinput": `(<[^>]+>((({{.*?}})\[(.*?)\])|(\[(.*?)\]))<\/[^>]+>)`,
		}

		resultRgx := map[string]string{}

		for i, vp := range regexPattern {
			switch i {
			case "regexstyle1", "regexstyle2":

				regexstyle := regexPattern["regexstyle1"]
				if strings.Contains(str, "{{color:#") {
					regexstyle = regexPattern["regexstyle2"]
				}

				repattern := regexp.MustCompile(regexstyle)
				matchpattern := repattern.FindStringSubmatch(str)
				resultRgx["regexstyle"] = ""
				if len(matchpattern) > 0 {
					resultRgx["regexstyle"] = matchpattern[1]
				}

			default:
				repattern := regexp.MustCompile(vp)
				matchpattern := repattern.FindStringSubmatch(str)
				resultRgx[i] = ""
				if len(matchpattern) > 0 {
					resultRgx[i] = matchpattern[1]
				}
			}
		}

		var tagOpen, tagClose string
		if resultRgx["regextag"] != "" {
			tagOpen = fmt.Sprintf("<%s>", resultRgx["regextag"])
			tagClose = fmt.Sprintf("</%s>", resultRgx["regextag"])
		}

		if resultRgx["regexstyle"] != "" && resultRgx["regextag"] != "" {
			tagOpen = fmt.Sprintf(`<%s style="color:%s">`, resultRgx["regextag"], resultRgx["regexstyle"])
		}

		content := fmt.Sprintf(`%s%s%s`, tagOpen, resultRgx["regexcontent"], tagClose)
		if resultRgx["regexreplaceinput"] != "" {
			content = strings.Replace(str, resultRgx["regexreplaceinput"], fmt.Sprintf(`%s%s%s`, tagOpen, resultRgx["regexcontent"], tagClose), -1)
		}

		str = content
	}

	return str
}

func TestMustache(w http.ResponseWriter, r *http.Request) {

	var data SubEmail
	json.Unmarshal([]byte(dummy), &data)

	fmt.Printf("Sasasa ==== %#v \n", data)

	data.ApplicationName = "bima pw test mustachae"
	data.Name = "bima pw lalalal"

	for i, v := range data.Paragraph {
		data.Paragraph[i].Id = convertCustomStyling(convertCustomTag(v.Id))
		data.Paragraph[i].En = convertCustomStyling(convertCustomTag(v.En))
	}

	str, err := mustache.RenderFile("./email.mustache", data)
	if err != nil {
		fmt.Println("sa == ", err)
	}

	response.NewResponse().WithCode(http.StatusAccepted).WithData(map[string]interface{}{
		"output": str,
	}).ParseResponse(w, r)

	fmt.Println("sasasa === ", str)
	// for _, v := range data.Paragraph {

	// 	rawcont := []string{v.Id, v.En}
	// 	str = convertCustomStyling(str, rawcont...)

	// }

	// fmt.Println("sasaa", str)
}

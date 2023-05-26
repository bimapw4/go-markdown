package translate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Params struct {
	Tag  string
	Lang string
}

type Result struct {
	Activity     string `json:"activity"`
	ActivityText string `json:"activity_text"`
}

type localization struct {
	Activity     lang `json:"activity"`
	ActivityText lang `json:"activity_text"`
}

type lang struct {
	En string `json:"en"`
	Id string `json:"id"`
}

func LocalizationTemplate(param Params, additional ...map[string]interface{}) (*Result, error) {
	jsonFile, err := os.Open("central_log.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	var tag map[string]interface{}
	json.Unmarshal(byteValue, &tag)

	fmt.Println("sasasa =========== ", tag)

	jsonLocalization := tag[param.Tag]

	var localization localization

	byteLocalization, err := json.Marshal(jsonLocalization)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(byteLocalization, &localization)

	var result Result
	switch param.Lang {
	case "id":
		result = Result{
			Activity:     localization.Activity.Id,
			ActivityText: localization.ActivityText.Id,
		}
	case "en":
		result = Result{
			Activity:     localization.Activity.Id,
			ActivityText: localization.ActivityText.Id,
		}
	default:
		return nil, errors.New("language must be set")
	}

	for _, data := range additional {
		for i, v := range data {
			pattern := fmt.Sprintf("{{%s}}", i)
			if strings.Contains(result.ActivityText, pattern) {
				result.ActivityText = strings.ReplaceAll(result.ActivityText, pattern, fmt.Sprint(v))
			}
		}
	}

	return &result, nil
}

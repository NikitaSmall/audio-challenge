/*
 * This package holds the main logic. Here tasks are parsed, created, processed.
 * This file holds API-specific (Yandex-API currently) functions.
 * Also this file contains main parse functions.
 */
package task

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/elgs/gojq"
	"github.com/nikitasmall/audio-challenge/util"
)

// messageRequest sends a request to parse captured voice to text
// and returns raw response body.
func messageRequest(message io.Reader) []byte {
	uuid := util.RandStringRunes(32)

	url := fmt.Sprintf("%s/asr_xml?uuid=%s&key=%s&topic=notes&lang=ru-RU",
		os.Getenv("YANDEX_SPEECH_RECOGNITION_URL"),
		uuid,
		os.Getenv("YANDEX_SPEECH_API_KEY"),
	)

	resp, err := http.Post(url, "audio/x-wav", message)
	if err != nil {
		log.Panic(err)
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	return res
}

// messageRequest sends a request to break text to logical blocks,
// such as date, names, addresses. Returns raw response body.
func paramsRequest(text string) []byte {
	url := fmt.Sprintf("%s/?key=%s&text=%s&layers=Fio,GeoAddr,Date",
		os.Getenv("YANDEX_MARKUP_URL"),
		os.Getenv("YANDEX_SPEECH_API_KEY"),
		text,
	)

	resp, err := http.Get(url)

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	return res
}

// newTask tries to parse provided API result and
// returns a new task or message about failed attempt.
func newTask(rawTask []byte) (*BaseTask, error) {
	parsedResult := util.ParseXML(rawTask)

	if len(parsedResult.Variants) == 0 {
		log.Print("Unsuccessful recognition")
		return nil, errors.New("Unsuccessful recognition")
	}

	return &BaseTask{
		RawQuery: parsedResult.Variants[0],
		Status:   false,
	}, nil
}

// getQueryParams gets the request, parse it and tries
// to get special information such as user name, his addr or order date if possible.
func (task *BaseTask) getQueryParams() (string, string, time.Time) {
	jsonParams := paramsRequest(task.RawQuery)
	params := util.ParseJSON(jsonParams)

	return parseName(params), parseAddr(params), parseTime(params)
}

// parseName tries to get name information from parsed json.
func parseName(params *gojq.JQ) string {
	var name string

	// check special field in response json to get a name if possible
	if nameArray, err := params.QueryToArray("Fio"); err == nil {
		for _, n := range nameArray {
			nameField := n.(map[string]interface{})
			name = nameField["FirstName"].(string)
		}
	}

	return name
}

// parseAddr tries to get address information from parsed json.
func parseAddr(params *gojq.JQ) string {
	var addr string

	// check special field in response json to get an address if possible
	if addrArray, err := params.QueryToArray("GeoAddr"); err == nil {
		for _, a := range addrArray {
			addrField := a.(map[string]interface{})

			// append few lines of address information.
			// every line holds single element such as street name or building number.
			for _, field := range addrField["Fields"].([]interface{}) {
				if len(addr) == 0 {
					addr = fmt.Sprintf("%s%s", addr, field.(map[string]interface{})["Name"].(string))
				} else {
					addr = fmt.Sprintf("%s, %s", addr, field.(map[string]interface{})["Name"].(string))
				}
			}
		}
	}

	return addr
}

// this terrible func parses incoming message
// to get as much as possible information about date
func parseTime(params *gojq.JQ) time.Time {
	date := time.Now()

	// check special field in response json to get date info if possible
	if dateArray, err := params.QueryToArray("Date"); err == nil {
		for _, d := range dateArray {
			dateField := d.(map[string]interface{})

			// check that date field is not relative
			absoluteDay, okD := dateField["Day"]
			absoluteMonth, okM := dateField["Month"]
			_, okR := dateField["RelativeDay"]
			_, okDur := dateField["Duration"]

			// setting absolute date if actual
			if okD && okM && !(okR || okDur) {
				date = time.Date(date.Year(), time.Month(absoluteMonth.(float64)), int(absoluteDay.(float64)), date.Hour(), date.Minute(), date.Second(), 0, date.Location())
			}

			// setting up relative date if possible
			if _, ok := dateField["RelativeDay"]; ok {
				var relativeM, relativeD int

				// checking month shift
				if dateM, ok := dateField["Month"]; ok {
					relativeM = int(dateM.(float64))
				}

				// checking day shift
				if dateD, ok := dateField["Day"]; ok {
					relativeD = int(dateD.(float64))
				}

				date = date.AddDate(0, relativeM, relativeD)
			}

			// setting up relative time if possible
			if duration, ok := dateField["Duration"]; ok {
				dur := duration.(map[string]interface{})
				var durString string

				// checking hour shift
				if h, ok := dur["Hour"]; ok {
					durString = fmt.Sprintf("%fh", h.(float64))
				}

				// checking minutes shift
				if m, ok := dur["Min"]; ok {
					durString = fmt.Sprintf("%s%fm", durString, m.(float64))
				}

				parsedDuration, err := time.ParseDuration(durString)
				if err != nil {
					log.Println("bad parse")
					break
				}

				date = date.Add(parsedDuration)
			}
		}
	}

	return date
}

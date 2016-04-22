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

// ProcessMessage sends the message file to the Yandex API and returns parsed task
func ProcessMessage(message io.Reader) (Tasker, error) {
	parsedResult := messageRequest(message)

	t, err := newTask(parsedResult)
	if err != nil {
		return nil, err
	}

	task, err := t.defineTask()
	if err != nil {
		return nil, err
	}

	return task, nil
}

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

func newTask(result []byte) (*BaseTask, error) {
	parsedResult := util.ParseXML(result)

	if len(parsedResult.Variants) == 0 {
		log.Print("Unsuccessful recognition")
		return nil, errors.New("Unsuccessful recognition")
	}

	return &BaseTask{
		RawQuery: parsedResult.Variants[0],
		Status:   false,
	}, nil
}

// this terrible func parses incoming message to get as much as possible information abpit data
func parseTime(params *gojq.JQ) time.Time {
	date := time.Now()

	if dateArray, err := params.QueryToArray("Date"); err == nil {
		for _, d := range dateArray {
			dateField := d.(map[string]interface{})
			log.Println(d)

			// check that date field is not relative
			absoluteDay, okD := dateField["Day"]
			absoluteMonth, okM := dateField["Month"]
			_, okR := dateField["RelativeDay"]
			_, okDur := dateField["Duration"]
			if okD && okM && !(okR || okDur) {
				date = time.Date(date.Year(), time.Month(absoluteMonth.(float64)), int(absoluteDay.(float64)), date.Hour(), date.Minute(), date.Second(), 0, date.Location())
			}

			if _, ok := dateField["RelativeDay"]; ok {
				var relativeM, relativeD int

				if dateM, ok := dateField["Month"]; ok {
					relativeM = int(dateM.(float64))
				}

				if dateD, ok := dateField["Day"]; ok {
					relativeD = int(dateD.(float64))
				}

				date = date.AddDate(0, relativeM, relativeD)
			}

			if duration, ok := dateField["Duration"]; ok {
				dur := duration.(map[string]interface{})
				var durString string

				if h, ok := dur["Hour"]; ok {
					durString = fmt.Sprintf("%fh", h.(float64))
				}

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

package task

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/nikitasmall/audio-challenge/util"
)

// ParseMessage sends the message file to the Yandex API and returns basic task
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

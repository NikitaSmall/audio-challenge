package task

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/nikitasmall/audio-challenge/config"
	"github.com/nikitasmall/audio-challenge/util"
)

// ParseMessage sends the message file to the Yandex API and returns basic task
func ParseMessage() (*BaseTask, error) {
	parsedResult := messageRequest(os.Getenv("MESSAGE_FILE"))

	return newTask(parsedResult)
}

func messageRequest(messagePath string) []byte {
	message := messageBody(messagePath)
	uuid := config.RandStringRunes(32)

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

func messageBody(messagePath string) io.Reader {
	file, err := os.Open(messagePath)
	if err != nil {
		log.Panic(err)
	}

	return file
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

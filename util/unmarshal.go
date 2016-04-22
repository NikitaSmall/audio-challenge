package util

import (
	"encoding/xml"
	"log"

	"github.com/elgs/gojq"
)

// XMLMessage is a struct of Yandex API message
type XMLMessage struct {
	XMLName  xml.Name `xml:"recognitionResults"`
	Variants []string `xml:"variant"`
}

// ParseXML parses xml data into XMLMessage object
func ParseXML(data []byte) XMLMessage {
	log.Println(string(data))

	m := XMLMessage{}
	err := xml.Unmarshal(data, &m)
	if err != nil {
		log.Panic(err)
	}

	return m
}

// ParseJSON parses json data into
func ParseJSON(data []byte) *gojq.JQ {
	log.Println(string(data))

	parser, err := gojq.NewStringQuery(string(data))
	if err != nil {
		log.Panic(err)
	}

	return parser
}

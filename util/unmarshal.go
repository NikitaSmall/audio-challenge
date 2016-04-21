package util

import (
	"encoding/xml"
	"log"
)

// XMLMessage is a struct of Yandex API message
type XMLMessage struct {
	XMLName  xml.Name `xml:"recognitionResults"`
	Variants []string `xml:"variant"`
}

// ParseXML parses xml data into map[string]interface{} object
func ParseXML(data []byte) XMLMessage {
	log.Println(string(data))

	m := XMLMessage{}
	err := xml.Unmarshal(data, &m)
	if err != nil {
		log.Panic(err)
	}

	return m
}

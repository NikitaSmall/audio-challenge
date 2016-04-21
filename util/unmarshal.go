package util

import (
	"encoding/xml"
	"log"
)

// Message is a struct of Yandex API message
type Message struct {
	XMLName  xml.Name `xml:"recognitionResults"`
	Variants []string `xml:"variant"`
}

// ParseXML parses xml data into map[string]interface{} object
func ParseXML(data []byte) Message {
	log.Println(string(data))

	m := Message{}
	err := xml.Unmarshal(data, &m)
	if err != nil {
		log.Panic(err)
	}

	return m
}

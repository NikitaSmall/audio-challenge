/*
 * This package holds utility functions.
 * This file holds functions to unmarshal xml or json input.
 * The result stored in XMLMessage struct for XML and
 * in gojq.JQ (which is []interface{} for real) for json.
 */
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

// ParseJSON parses json data into gojq.JQ. There is no good simple struct
// for json input because of great variability for API answers in json.
func ParseJSON(data []byte) *gojq.JQ {
	log.Println(string(data))

	parser, err := gojq.NewStringQuery(string(data))
	if err != nil {
		log.Panic(err)
	}

	return parser
}

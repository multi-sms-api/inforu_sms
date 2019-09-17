package inforusms

import "encoding/xml"

// Content holds message information
type Content struct {
	XMLName xml.Name `xml:"Content"`
	Type    string   `xml:"Type,attr"`
	Message string   `xml:"Message"`
}

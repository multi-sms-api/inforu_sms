package inforusms

import "encoding/xml"

// Recipients holds contact information about phone numbers, seperated by
// semi-colon
type Recipients struct {
	XMLName     xml.Name `xml:"Recipients"`
	PhoneNumber string   `xml:"PhoneNumber"`
	GroupNumber string   `xml:"GroupNumber,omitempty"`
}

package inforusms

import "encoding/xml"

// DeliveryInfo holds information regarding delivery status arrived back as
// callback
type DeliveryInfo struct {
	XMLName           xml.Name       `xml:"IncomingData"`
	PhoneNumber       string         `xml:"PhoneNumber"`
	Network           string         `xml:"Network"`
	Status            DeliveryStatus `xml:"Status"`
	StatusDescription string         `xml:"StatusDescription"`
	CustomerMessageID int64          `xml:"CustomerMessageId"`
	SegmentsNumber    int            `xml:"SegmentsNumber"`
}

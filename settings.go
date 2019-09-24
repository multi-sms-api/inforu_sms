package inforusms

import "encoding/xml"

// Settings holds information regarding the sender and the sending of a message
type Settings struct {
	XMLName                 xml.Name `xml:"Settings"`
	Sender                  string   `xml:"Sender"`
	CustomerMessageID       string   `xml:"CustomerMessageID,omitempty"`
	CustomerParameter       string   `xml:"CustomerParameter,omitempty"`
	MessageInterval         int      `xml:"MessageInterval,omitempty"`
	TimeToSend              string   `xml:"TimeToSend,omitempty"`
	MaxSegments             int      `xml:"MaxSegments,omitempty"`
	DelayInSeconds          int64    `xml:"DelayInSeconds,omitempty"`
	DeliveryNotificationURL string   `xml:"DeliveryNotificationUrl,omitempty"`
	Priority                int      `xml:"priority,omitempty"`
}

package inforusms

import (
	"encoding/xml"
	"net/url"
	"time"
)

// DeliveryInfo holds information regarding delivery status arrived back as
// callback
type DeliveryInfo struct {
	XMLName           xml.Name       `xml:"IncomingData"`
	ActionType        string         `xml:"ActionType"`
	OriginalMessage   string         `xml:"OriginalMessage"`
	Price             string         `xml:"Price"`
	CustomerParam     string         `xml:"CustomerParam"`
	SenderNumber      string         `xml:"SenderNumber"`
	PhoneNumber       string         `xml:"PhoneNumber"`
	Network           string         `xml:"Network"`
	Status            DeliveryStatus `xml:"Status"`
	StatusDescription string         `xml:"StatusDescription"`
	CustomerMessageID int64          `xml:"CustomerMessageId"`
	SegmentsNumber    int            `xml:"SegmentsNumber"`
	RetriesNumber     uint           `xml:"RetriesNumber"`
	ID                string         `xml:"id"`
	NotificationDate  time.Time      `xml:"NotificationDate"`
	ProjectID         string         `xml:"ProjectId"`
	BillingCodeID     string         `xml:"BillingCodeId"`
}

// FormToDeliveryInfo takes form fields and place converts them into DeliveryInfo
func FormToDeliveryInfo(form url.Values) *DeliveryInfo {
	dateTime, err := time.Parse(TimeFormat, form.Get("NotificationDate"))
	if err != nil {
		dateTime = time.Now()
	}

	result := DeliveryInfo{
		ActionType:        form.Get("ActionType"),
		OriginalMessage:   form.Get("OriginalMessage"),
		Price:             form.Get("Price"),
		CustomerParam:     form.Get("CustomerParam"),
		SenderNumber:      form.Get("SenderNumber"),
		PhoneNumber:       form.Get("PhoneNumber"),
		Network:           form.Get("Network"),
		Status:            DeliveryStatusFromString(form.Get("Status")),
		StatusDescription: form.Get("StatusDescription"),
		CustomerMessageID: strToInt(form.Get("CustomerToMessageId"), 64),
		SegmentsNumber:    int(strToInt(form.Get("SegmentsNumber"), 16)),
		RetriesNumber:     uint(strToUint(form.Get("RetriesNumber"), 16)),
		ID:                form.Get("id"),
		NotificationDate:  dateTime,
		ProjectID:         form.Get("ProjectId"),
		BillingCodeID:     form.Get("BillingCodeId"),
	}

	return &result
}

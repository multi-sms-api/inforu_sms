package inforusms

import (
	"net/http"
	"os"
	"testing"
	"time"
)

var (
	userName, password   string
	fromNumber, toNumber string
	client               http.Client
	handler              HTTPHandler
)

func init() {
	userName = os.Getenv("INFORU_USERNAME")
	password = os.Getenv("INFORU_PASSWORD")
	fromNumber = os.Getenv("INFORU_FROM_NUMBER")
	toNumber = os.Getenv("INFORU_TO_NUMBER")
	client = http.Client{
		Timeout: time.Second * 15,
	}
	handler = HTTPHandler{
		Client: &client,
	}
}

func TestSendSMSHTTP(t *testing.T) {
	if testing.Short() {
		return
	}
	xmlToSend := InforuXML{
		Auth: UserAuth{
			UserName: userName,
			Password: password,
		},
		Content: Content{
			Type:    "sms",
			Message: "Test1",
		},
		Recipients: Recipients{
			PhoneNumber: toNumber,
		},
		Settings: Settings{
			Sender: fromNumber,
		},
	}

	_, err := xmlToSend.SendSMS(handler)
	if err != nil {
		t.Errorf("Error sending SMS: %s", err)
	}
}

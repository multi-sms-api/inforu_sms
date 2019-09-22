package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	inforusms "github.com/ik5/inforu_sms"
)

var (
	userName, password   string
	fromNumber, toNumber string
	callbackPort         string
	client               http.Client
	handler              inforusms.HTTPHandler
)

var done chan bool

func init() {
	userName = os.Getenv("INFORU_USERNAME")
	password = os.Getenv("INFORU_PASSWORD")
	fromNumber = os.Getenv("INFORU_FROM_NUMBER")
	toNumber = os.Getenv("INFORU_TO_NUMBER")
	callbackPort = os.Getenv("INFORU_PORT")
	client = http.Client{
		Timeout: time.Second * 15,
	}
	handler = inforusms.HTTPHandler{
		Client: &client,
	}
}

func sendSMS() error {
	xmlToSend := inforusms.InforuXML{
		Auth: inforusms.UserAuth{
			UserName: userName,
			Password: password,
		},
		Content: inforusms.Content{
			Type:    "sms",
			Message: "Test1",
		},
		Recipients: inforusms.Recipients{
			PhoneNumber: toNumber,
		},
		Settings: inforusms.Settings{
			Sender:                  fromNumber,
			DeliveryNotificationURL: "http://62.219.162.140/get_sms",
		},
	}

	return xmlToSend.SendSMS(handler)
}

func serveHTTP() {
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%s", callbackPort), nil))
}

func main() {

	done = make(chan bool)

	inforusms.Callback(handler, "/get_sms", nil, nil)

	go serveHTTP()
	fmt.Println(sendSMS())
	<-done
}

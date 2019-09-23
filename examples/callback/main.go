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

func sendSMS() (*http.Response, error) {
	xmlToSend := inforusms.InforuXML{
		Auth: inforusms.UserAuth{
			UserName: userName,
			Password: password,
		},
		Content: inforusms.Content{
			Type:    "sms",
			Message: "Test1 עם עברית",
		},
		Recipients: inforusms.Recipients{
			PhoneNumber: toNumber,
		},
		Settings: inforusms.Settings{
			Sender:                  fromNumber,
			DeliveryNotificationURL: "http://127.0.0.1:80/get_sms",
			CustomerMessageID:       "1",
			CustomerParameter:       fmt.Sprintf("test1-%s", time.Now()),
		},
	}

	return xmlToSend.SendSMS(handler)
}

func serveHTTP() {
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%s", callbackPort), nil))
}

func callback(w http.ResponseWriter, r *http.Request, info *inforusms.DeliveryInfo, err error) {
	if err != nil {
		panic(err)
	}
	if info != nil {
		fmt.Printf("Full struct: %+v\n", *info)
	}
	done <- true
}

func main() {
	done = make(chan bool)
	inforusms.Callback(handler, "/get_sms", nil, callback)
	go serveHTTP()
	fmt.Printf("Going to send sms: ")
	_, err := sendSMS()
	if err != nil {
		panic(err)
	}
	fmt.Println("Done, waiting for delivery report.")
	<-done
}

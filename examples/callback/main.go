package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	inforusms "github.com/multi-sms-api/inforu_sms"
)

var (
	userName, password   string
	fromNumber, toNumber string
	callbackPort         string
	callbackAddress      string
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
	callbackAddress = os.Getenv("INFORU_ADDRESS")
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
			Message: `Test1 לשליחה מאוד ארוכה של סמס, עם הרבה מאוד מלל ולא רק כמה תווים. מה עוד אפשר לבדוק האם זה עובד או לא. האם ככול שהטקסט יותר ארוך ככה יהיו בעיות? Maybe some English text as well to make it work harder and making sure that we have a very long text message, and see if there will be any issues with it, so we knows that too many fragments are bad for us. What do you think?`,
		},
		Recipients: inforusms.Recipients{
			PhoneNumber: toNumber,
		},
		Settings: inforusms.Settings{
			Sender:                  fromNumber,
			DeliveryNotificationURL: callbackAddress,
			CustomerMessageID:       "1",
			CustomerParameter:       fmt.Sprintf("test1-%s", time.Now()),
		},
	}

	return xmlToSend.SendSMS(handler)
}

func serveHTTP() {
	listenTo := fmt.Sprintf(":%s", callbackPort)
	fmt.Println("Going to listen on: ", listenTo)
	fmt.Println(http.ListenAndServe(listenTo, nil))
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

package inforusms

import (
	"net/http"
	"net/url"

	"github.com/multi-sms-api/smshandler"
)

// HTTPHandler perform HTTP actions, and implement
type HTTPHandler struct {
	Client *http.Client
}

// DoHTTP sends an HTTP Request for sending an SMS
func (h HTTPHandler) DoHTTP(
	method, contentType, address string, fields url.Values, body []byte) (resp *http.Response, err error) {
	response := &XMLResponse{}
	return smshandler.DoHTTP(h.Client, method, contentType, address, fields, body, response)
}

// OnGettingSMS is an HTTP server handler when incoming SMS arrives.
// If mux exists, it will use it for a server, otherwise it will
// use http.HandleFunc.
func (h HTTPHandler) OnGettingSMS(path string, mux *http.ServeMux, httpHandler http.HandlerFunc) {
	smshandler.OnGettingSMS(path, mux, httpHandler)
}

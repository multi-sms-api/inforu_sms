package inforusms

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

// HTTPHandler perform HTTP actions, and implement
type HTTPHandler struct {
	Client *http.Client
}

// DoHTTP sends an HTTP Request for sending an SMS
func (h HTTPHandler) DoHTTP(
	method, contentType, address string, fields url.Values, body []byte) (resp *http.Response, err error) {

	var request *http.Request

	fullAddress := fmt.Sprintf("%s?%s", address, fields.Encode())

	switch method {
	case http.MethodGet:
		request, err = http.NewRequest(http.MethodGet, fullAddress, nil)
	case http.MethodPost:
		request, err = http.NewRequest(http.MethodPost, fullAddress, nil)
	}

	if err != nil {
		return nil, err
	}

	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}
	request.Close = true

	ctx, cancel := context.WithTimeout(request.Context(), h.Client.Timeout)
	defer cancel()
	defer h.Client.CloseIdleConnections()

	resp, err = h.Client.Do(request.WithContext(ctx))

	if strings.Contains(os.Getenv("SMSHTTPDEBUG"), "dump=true") {
		dump, err := httputil.DumpRequestOut(request, true)
		fmt.Printf(">>>> dump request: %s \nerr: %s\n", dump, err)

		dump, err = httputil.DumpResponse(resp, true)
		fmt.Printf(">>>> dump response: %s \nerr: %s\n", dump, err)
	}

	if resp.Body == nil {
		return
	}

	var respBody []byte
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode == http.StatusOK {
		var status XMLResponse
		status, err = FromXMLResponse(respBody)
		if err != nil {
			return
		}
		if status.Status != StatusOK {
			err = ToError(status)
		}
	}

	return
}

// OnGettingSMS is an HTTP server handler when incoming SMS arrives.
// If mux exists, it will use it for a server, otherwise it will
// use http.HandleFunc.
func (h HTTPHandler) OnGettingSMS(path string, mux *http.ServeMux, httpHandler http.HandlerFunc) {
	if mux != nil {
		mux.HandleFunc(path, httpHandler)
		return
	}

	http.HandleFunc(path, httpHandler)
}

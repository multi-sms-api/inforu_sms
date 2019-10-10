package inforusms

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ik5/smshandler"
)

// Callback for having callback function
func Callback(h smshandler.HTTPHandler, path string, mux *http.ServeMux,
	onCallback func(http.ResponseWriter, *http.Request, *DeliveryInfo, error)) {

	handleCallback := func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		if !(contentType == "application/x-www-form-urlencoded" && r.ContentLength > 0) {
			// Not a valid request, but do what ever you want from it
			if onCallback != nil {
				onCallback(w, r, nil, fmt.Errorf(
					"Content-Type (%s) or ContentLength (%d) are invalid",
					contentType, r.ContentLength,
				))
			}
			return
		}
		_ = r.ParseForm()
		if strings.Contains(os.Getenv("SMSHTTPDEBUG"), "dump=true") {
			fmt.Printf(">>>> Request dump: %+v\n", r)
			for k, v := range r.Form {
				fmt.Printf("\t %s=%v\n", k, v)
			}
		}

		deliveryInfo := FormToDeliveryInfo(r.Form)
		if onCallback != nil {
			onCallback(w, r, deliveryInfo, nil)
		}
	}

	h.OnGettingSMS(path, mux, handleCallback)
}

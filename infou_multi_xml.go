package inforusms

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// InforuMultiXML generate multiple messages in one XML message
type InforuMultiXML struct {
	XMLName   xml.Name    `xml:"InforuRoot"`
	InforuXML []InforuXML `xml:"Inforu"`
}

// SendSMS sends the given SMS to InforU based on HTTP client
func (x InforuMultiXML) SendSMS(h HTTPHandler) (*http.Response, error) {
	field := url.Values{}
	buf, err := xml.Marshal(x)
	if err != nil {
		return nil, err
	}
	field.Set(HTTPArg, string(buf))
	if strings.Contains(os.Getenv("SMSHTTPDEBUG"), "dump=true") {
		fmt.Printf(">>>> dump InforuMultiXML: %s\n", buf)
	}
	resp, err := h.DoHTTP(HTTPMethod, HTTPContentType, HTTPSAPIAddress, nil, []byte(field.Encode()))
	return resp, err
}

package inforusms

import (
	"encoding/xml"
	"strings"
)

// XMLResponse holds the content for the answer arrived back for the
// request
//
//  <Result>
//    <Status></Status>
//    <Description></Description>
//    <NumberOfRecipients></NumberOfRecipients>
//  </Result>
type XMLResponse struct {
	XMLName            xml.Name       `xml:"Result"`
	Status             ResponseStatus `xml:"Status"`
	Description        string         `xml:"Description"`
	NumberOfRecipients int64          `xml:"NumberOfRecipients"`
}

// FromXMLResponse turns the body to XMLResponse, or error if something bad
// happened
func FromXMLResponse(buf []byte) (XMLResponse, error) {
	var response XMLResponse

	err := xml.Unmarshal(buf, &response)
	return response, err
}

// ToError converts XMLResponse to SMSError. If everything is ok, it will return
// nil
func (r XMLResponse) ToError() *SMSError {
	if r.Status == StatusOK {
		return nil
	}
	result := SMSError{
		Status:      r.Status,
		Description: r.Description,
		Effected:    r.NumberOfRecipients,
	}

	if strings.HasPrefix(strings.ToLower(result.Description), "error: ") {
		result.Description = result.Description[7:]
	}

	return &result
}

// IsOK Implements the interface to know if response holds success or not
func (r XMLResponse) IsOK() bool {
	return r.Status == StatusOK
}

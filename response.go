package inforusms

import "encoding/xml"

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

package inforusms

import "encoding/xml"

// UserAuth holds fields for User Authentication
type UserAuth struct {
	XMLName  xml.Name `xml:"User"`
	UserName string   `xml:"Username"`
	Password string   `xml:"Password"`
}

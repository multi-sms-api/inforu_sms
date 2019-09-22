package inforusms

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// InforuXML holds XML information for simple message to be sent, can support
// multiple phone numbers separated by semi-colon.
//
// Simple XML:
//
//  <Inforu>
//     <User>
//       <Username>MyUsername</Username>
//       <Password>MyPassword</Password>
//     </User>
//     <Content Type="sms">
//       <Message>This is a test SMS Message</Message>
//     </Content>
//     <Recipients>
//       <PhoneNumber>0501111111;0502222222</PhoneNumber>
//     </Recipients>
//     <Settings>
//       <Sender>0501111111</Sender>
//     </Settings>
//   </Inforu>
//
// Advanced XML
//
//   <Inforu>
//     <User>
//       <Username>MyUserName</Username>
//       <Password>MyPassword</Password>
//     </User>
//     <Content Type="sms">
//       <Message> This is a test SMS Message </Message>
//     </Content>
//     <Recipients>
//       <PhoneNumber>0501111111;0502222222</PhoneNumber>
//       <GroupNumber>5</GroupNumber>
//     </Recipients>
//     <Settings>
//       <Sender>Inforu</Sender>
//       <CustomerMessageID>112233</CustomerMessageID>
//       <CustomerParameter>AffId4</CustomerParameter>
//       <MessageInterval>0</MessageInterval>
//       <TimeToSend>12/05/2013 12:23</TimeToSend>
//       <DelayInSeconds>60</DelayInSeconds>
//       <DeliveryNotificationUrl>http://mysite.co.il/Notif.aspx</DeliveryNotificationUrl>
//       <MaxSegments>0</MaxSegments>
//     </Settings>
//   </Inforu>
type InforuXML struct {
	XMLName    xml.Name   `xml:"Inforu"`
	Auth       UserAuth   `xml:"User"`
	Content    Content    `xml:"Content"`
	Recipients Recipients `xml:"Recipients"`
	Settings   Settings   `xml:"Settings"`
}

// SendSMS sends the given SMS to InforU based on HTTP client
func (x InforuXML) SendSMS(h HTTPHandler) (*http.Response, error) {
	field := url.Values{}
	buf, err := xml.Marshal(x)
	if err != nil {
		return nil, err
	}
	field.Set(HTTPArg, string(buf))
	if strings.Contains(os.Getenv("SMSHTTPDEBUG"), "dump=true") {
		fmt.Printf(">>>> dump XML: %s\n", buf)
	}
	resp, err := h.DoHTTP(HTTPMethod, HTTPContentType, HTTPSAPIAddress, field, []byte(field.Encode()))
	return resp, err
}

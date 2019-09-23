package inforusms

import (
	"database/sql/driver"
	"encoding/xml"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// HTTP(s) address to send the API
const (
	HTTPSAPIAddress = `https://api.inforu.co.il/SendMessageXml.ashx`
	HTTPArg         = `InforuXML`
	HTTPMethod      = `POST`
	HTTPContentType = `application/x-www-form-urlencoded`
	TimeFormat      = `01/02/2006 15:04`
)

// ResponseStatus holds information regarding the given response back from the
// server
type ResponseStatus int

// DeliveryStatus holds callback delivery status that returned by the server
type DeliveryStatus int

// Statuses
const (
	StatusOK                          ResponseStatus = 1
	StatusFailed                      ResponseStatus = -1
	StatusBadUserNameOrPassword       ResponseStatus = -2
	StatusUserNameNotExist            ResponseStatus = -3
	StatusPasswordNotExists           ResponseStatus = -4
	StatusRecipientsDataNotExists     ResponseStatus = -6
	StatusMessageTextNotExists        ResponseStatus = -9
	StatusIllegalXML                  ResponseStatus = -11
	StatusUserQuotaExceeded           ResponseStatus = -13
	StatusProjectQuotaExceeded        ResponseStatus = -14
	StatusCustomerQuotaExceeded       ResponseStatus = -15
	StatusWrongDateTime               ResponseStatus = -16
	StatusNoValidRecipients           ResponseStatus = -18
	StatusInvalidSenderNumber         ResponseStatus = -20
	StatusInvalidSenderName           ResponseStatus = -21
	StatusUserBlocked                 ResponseStatus = -22
	StatusUserAuthenticationError     ResponseStatus = -26
	StatusNetworkTypeNotSupported     ResponseStatus = -28
	StatusNotAllNetworkTypesSupported ResponseStatus = -29
	StatusSenderIdentification        ResponseStatus = -90
)

// DeliveryStatus statuses
const (
	DeliveryStatusDelivered             DeliveryStatus = 2
	DeliveryStatusNotDelivered          DeliveryStatus = -2
	DeliveryStatusBlockedByInforuMobile DeliveryStatus = -4
)

// responseStatusMap holds information regarding status numbers and it's
// corresponding text
var responseStatusMap = map[ResponseStatus]string{
	StatusOK:                          "OK",
	StatusFailed:                      "Failed",
	StatusBadUserNameOrPassword:       "Bad user name or password",
	StatusUserNameNotExist:            "User name does not exists",
	StatusPasswordNotExists:           "Password does not exists",
	StatusRecipientsDataNotExists:     "Recipients data does not exists",
	StatusMessageTextNotExists:        "Message text does not exists",
	StatusIllegalXML:                  "Illegal XML",
	StatusUserQuotaExceeded:           "User Quota Exceeded",
	StatusProjectQuotaExceeded:        "Project Quota Exceeded",
	StatusCustomerQuotaExceeded:       "Customer Quota Exceeded",
	StatusWrongDateTime:               "Wrong Date Time",
	StatusNoValidRecipients:           "No valid recipients",
	StatusInvalidSenderNumber:         "Invalid sender number",
	StatusInvalidSenderName:           "Invalid sender name",
	StatusUserBlocked:                 "User is blocked",
	StatusUserAuthenticationError:     "User Authentication Error",
	StatusNetworkTypeNotSupported:     "Network type is not supported",
	StatusNotAllNetworkTypesSupported: "Not all network types are supported",
	StatusSenderIdentification:        "Invalid sender identification",
}

var deliveryStatusMap = map[DeliveryStatus]string{
	DeliveryStatusDelivered:             "Delivered",
	DeliveryStatusNotDelivered:          "Not delivered",
	DeliveryStatusBlockedByInforuMobile: "Blocked by InforuMobile",
}

func (s ResponseStatus) String() string {
	str, found := responseStatusMap[s]
	if !found {
		return strconv.Itoa(int(s))
	}

	return str
}

func (ds DeliveryStatus) String() string {
	str, found := deliveryStatusMap[ds]
	if !found {
		return strconv.Itoa(int(ds))
	}
	return str
}

// IsOK check to see if given ResponseStatus does not contain any error
func (s ResponseStatus) IsOK() bool {
	return s == StatusOK
}

// IsDelivered check to see if given DeliveryStatus delivered
func (ds DeliveryStatus) IsDelivered() bool {
	return ds == DeliveryStatusDelivered
}

// MarshalXML implements XML marshaling for ResponseStatus
func (s ResponseStatus) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(int(s), start)
}

// MarshalXML implements XML marshaling for DeliveryStatus
func (ds DeliveryStatus) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(int(ds), start)
}

// UnmarshalXML implement unmarshaling for ResponseStatus
func (s *ResponseStatus) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var i int
	err := d.DecodeElement(&i, &start)
	if err != nil {
		return err
	}

	*s = ResponseStatus(i)
	return nil
}

//UnmarshalXML implement unmarshaling for DeliveryStatus
func (ds *DeliveryStatus) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var i int
	err := d.DecodeElement(&i, &start)
	if err != nil {
		return err
	}

	*ds = DeliveryStatus(i)
	return nil
}

// Value implements the database interface of Value
func (s ResponseStatus) Value() (driver.Value, error) {
	result := int(s)
	return result, nil
}

// Value implements the database interface of Value
func (ds DeliveryStatus) Value() (driver.Value, error) {
	result := int(ds)
	return result, nil
}

// Scan implements the database interface for Scan
func (s *ResponseStatus) Scan(src interface{}) error {
	if src == nil {
		return errors.New("src cannot be nil")
	}

	switch src.(type) {
	case int, int8, int16, int32, int64:
		*s = ResponseStatus(reflect.ValueOf(src).Int())
	case float32, float64:
		*s = ResponseStatus(int(reflect.ValueOf(src).Float()))
	default:
		return fmt.Errorf("Invalid type of src: %T", src)
	}

	return nil
}

// Scan implements the database interface for Scan
func (ds *DeliveryStatus) Scan(src interface{}) error {
	if src == nil {
		return errors.New("src cannot be nil")
	}

	switch src.(type) {
	case int, int8, int16, int32, int64:
		*ds = DeliveryStatus(reflect.ValueOf(src).Int())
	case float32, float64:
		*ds = DeliveryStatus(int(reflect.ValueOf(src).Float()))
	default:
		return fmt.Errorf("Invalid type of src: %T", src)
	}

	return nil
}

// DeliveryStatusFromString initialize DeliveryStatus based on string parsing.
// If failed it returns 0
func DeliveryStatusFromString(str string) DeliveryStatus {
	i := strToInt(str, 16)
	return DeliveryStatus(i)
}

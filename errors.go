package inforusms

import "fmt"

// SMSError holds information about sending information
type SMSError struct {
	Status      ResponseStatus
	Type        string
	Description string
}

func (e SMSError) Error() string {
	return fmt.Sprintf("%d: %s", e.Status, e.Description)
}

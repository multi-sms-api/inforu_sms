package inforusms

import (
	"fmt"
)

// SMSError holds information about sending information
type SMSError struct {
	Status      ResponseStatus
	Description string
	Effected    int64
}

func (e SMSError) Error() string {
	return fmt.Sprintf("%d: %s", e.Status, e.Description)
}

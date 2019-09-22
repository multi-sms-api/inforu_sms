package inforusms

import (
	"fmt"
	"strings"
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

// ToError converts XMLResponse to SMSError. If everything is ok, it will return
// nil
func ToError(returnStatus XMLResponse) *SMSError {
	if returnStatus.Status == StatusOK {
		return nil
	}
	result := SMSError{
		Status:      returnStatus.Status,
		Description: returnStatus.Description,
		Effected:    returnStatus.NumberOfRecipients,
	}

	if strings.HasPrefix(strings.ToLower(result.Description), "error: ") {
		result.Description = result.Description[7:]
	}

	return &result
}

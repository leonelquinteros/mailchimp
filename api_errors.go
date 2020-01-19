package mailchimp

import (
	"fmt"
)

// ErrorResponse represents the error response format from Mailchimp's API
type ErrorResponse struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

// Error implementation
func (er ErrorResponse) Error() string {
	return fmt.Sprintf("Status %d: %s - %s", er.Status, er.Title, er.Detail)
}

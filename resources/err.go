package resources

import (
	"encoding/json"
	"fmt"
)

type APIResponse struct {
	Success bool              `json:"success,omitempty"`
	Data    json.RawMessage   `json:"data,omitempty"`
	Error2  *APIResponseError `json:"error,omitempty"`

	Debug *APIResponseDebug `json:"debug,omitempty"`
}

// type APIResponseData json.RawMessage //interface{}
type APIResponseError struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	DevMessage string `json:"dev_message,omitempty"`

	SID    *string `json:"sid,omitempty"`
	Action *action `json:"action,omitempty"`

	RequestId string `json:"requestId,omitempty"`
}
type APIResponseDebug interface{}

type action struct {
	Type    string            `json:"type"`
	SubType string            `json:"sub_type"`
	Title   string            `json:"title"`
	Text    string            `json:"text"`
	Options map[string]string `json:"options,omitempty"`
}

func (e *APIResponseError) Error() string {
	if e != nil {
		return fmt.Sprintf("[%d] %s (%s)", e.Code, e.Message, e.DevMessage)
	} else {
		return fmt.Sprintf("200")
	}
}

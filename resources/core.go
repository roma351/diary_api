package resources

import "encoding/json"

type Core struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data,omitempty"`
	Error   *CoreError      `json:"error,omitempty"` // Error2

	Debug interface{} `json:"debug,omitempty"`
}

type CoreError struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	DevMessage string `json:"dev_message,omitempty"`

	SID    *string          `json:"sid,omitempty"`
	Action *CoreErrorAction `json:"action,omitempty"`
}

type CoreErrorAction struct {
	Type    string            `json:"type"`
	SubType string            `json:"sub_type"`
	Title   string            `json:"title"`
	Text    string            `json:"text"`
	Options map[string]string `json:"options,omitempty"`
}

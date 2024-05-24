package diary_api

import (
	"fmt"
	"github.com/roma351/diary_api/resources"
)

type ApiError struct {
	Code       int
	Message    string
	DevMessage string

	SID    *string
	Action *resources.CoreErrorAction

	RequestId string `json:"requestId"`
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("[%d] %s (%s)", e.Code, e.Message, e.DevMessage)
}

func newError(code int, err error, requestId string) *ApiError {
	var devMsg string
	if err != nil {
		devMsg = err.Error()
	}
	return &ApiError{
		Code:       code,
		Message:    "Школьный Дневник не отвечает. Повторите позже",
		DevMessage: devMsg,

		RequestId: requestId,
	}
}

func newErrorApi(coreErr *resources.CoreError, requestId string) *ApiError {
	return &ApiError{
		Code:       coreErr.Code,
		Message:    coreErr.Message,
		DevMessage: coreErr.DevMessage,

		SID:    coreErr.SID,
		Action: coreErr.Action,

		RequestId: requestId,
	}
}

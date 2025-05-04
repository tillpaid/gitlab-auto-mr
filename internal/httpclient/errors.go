package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type APIError struct {
	StatusCode int
	Message    string
	Raw        []byte
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("API error (status %d): %s", e.StatusCode, e.Message)
	}
	return fmt.Sprintf("API error (status %d)", e.StatusCode)
}

func ParseAPIError(statusCode int, raw []byte) *APIError {
	var msg string

	var out bytes.Buffer
	if err := json.Indent(&out, raw, "", "  "); err == nil {
		msg = out.String()
	} else {
		msg = string(raw)
	}

	return &APIError{
		StatusCode: statusCode,
		Message:    msg,
		Raw:        raw,
	}
}

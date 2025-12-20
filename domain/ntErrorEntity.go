package domain

import "errors"

type NTErrorEntity struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	Object    string `json:"object"` // error
	RequestId string `json:"request_id"`
	Status    int    `json:"status"`
}

var ErrNotionErrorResponse = errors.New("notion error response")

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"github.com/codatio/client-sdk-go/sync-for-payables/v3/pkg/models/shared"
)

// ErrorMessage - The request made is not valid.
type ErrorMessage struct {
	// `True` if the error occurred transiently and can be retried.
	CanBeRetried *string `json:"canBeRetried,omitempty"`
	// Unique identifier used to propagate to all downstream services and determine the source of the error.
	CorrelationID *string `json:"correlationId,omitempty"`
	// Machine readable error code used to automate processes based on the code returned.
	DetailedErrorCode *int64 `json:"detailedErrorCode,omitempty"`
	// A brief description of the error.
	Error_ *string `json:"error,omitempty"`
	// Codat's service the returned the error.
	Service *string `json:"service,omitempty"`
	// The HTTP status code returned by the error.
	StatusCode *int64 `json:"statusCode,omitempty"`
	// A human-readable object describing validation decisions Codat has made. If an operation has failed because of validation errors, they will be detailed here.
	Validation *shared.ErrorValidation `json:"validation,omitempty"`
}

var _ error = &ErrorMessage{}

func (e *ErrorMessage) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

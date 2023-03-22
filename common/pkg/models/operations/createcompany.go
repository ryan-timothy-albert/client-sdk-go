// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type CreateCompanyRequestBody struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name"`
}

// CreateCompany401ApplicationJSON - Your API request was not properly authorized.
type CreateCompany401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

// CreateCompany200ApplicationJSON - OK
type CreateCompany200ApplicationJSON struct {
	Created           *string                  `json:"created,omitempty"`
	CreatedByUserName *string                  `json:"createdByUserName,omitempty"`
	DataConnections   []map[string]interface{} `json:"dataConnections,omitempty"`
	Description       *string                  `json:"description,omitempty"`
	ID                *string                  `json:"id,omitempty"`
	Name              *string                  `json:"name,omitempty"`
	Platform          *string                  `json:"platform,omitempty"`
	Redirect          *string                  `json:"redirect,omitempty"`
}

type CreateCompanyResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	CreateCompany200ApplicationJSONObject *CreateCompany200ApplicationJSON
	// Your API request was not properly authorized.
	CreateCompany401ApplicationJSONObject *CreateCompany401ApplicationJSON
}

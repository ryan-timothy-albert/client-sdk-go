// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteCompanyConnectionRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

// DeleteCompanyConnection404ApplicationJSON - One or more of the resources you referenced could not be found.
// This might be because your company or data connection id is wrong, or was already deleted.
type DeleteCompanyConnection404ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

// DeleteCompanyConnection401ApplicationJSON - Your API request was not properly authorized.
type DeleteCompanyConnection401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type DeleteCompanyConnectionResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Your API request was not properly authorized.
	DeleteCompanyConnection401ApplicationJSONObject *DeleteCompanyConnection401ApplicationJSON
	// One or more of the resources you referenced could not be found.
	// This might be because your company or data connection id is wrong, or was already deleted.
	DeleteCompanyConnection404ApplicationJSONObject *DeleteCompanyConnection404ApplicationJSON
}

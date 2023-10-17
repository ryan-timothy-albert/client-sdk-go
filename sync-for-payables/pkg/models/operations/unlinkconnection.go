// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
	"net/http"
)

type UnlinkConnectionUpdateConnection struct {
	// The current authorization status of the data connection.
	Status *shared.DataConnectionStatus `json:"status,omitempty"`
}

func (o *UnlinkConnectionUpdateConnection) GetStatus() *shared.DataConnectionStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type UnlinkConnectionRequest struct {
	RequestBody *UnlinkConnectionUpdateConnection `request:"mediaType=application/json"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

func (o *UnlinkConnectionRequest) GetRequestBody() *UnlinkConnectionUpdateConnection {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *UnlinkConnectionRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *UnlinkConnectionRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type UnlinkConnectionResponse struct {
	// OK
	Connection *shared.Connection
	// HTTP response content type for this operation
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UnlinkConnectionResponse) GetConnection() *shared.Connection {
	if o == nil {
		return nil
	}
	return o.Connection
}

func (o *UnlinkConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UnlinkConnectionResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *UnlinkConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UnlinkConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

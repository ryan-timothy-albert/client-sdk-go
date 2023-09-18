// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v4/pkg/models/shared"
	"net/http"
)

type GetCommerceDisputeRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a dispute.
	DisputeID string `pathParam:"style=simple,explode=false,name=disputeId"`
}

func (o *GetCommerceDisputeRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetCommerceDisputeRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetCommerceDisputeRequest) GetDisputeID() string {
	if o == nil {
		return ""
	}
	return o.DisputeID
}

type GetCommerceDisputeResponse struct {
	// OK
	CommerceDispute *shared.CommerceDispute
	ContentType     string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetCommerceDisputeResponse) GetCommerceDispute() *shared.CommerceDispute {
	if o == nil {
		return nil
	}
	return o.CommerceDispute
}

func (o *GetCommerceDisputeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCommerceDisputeResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetCommerceDisputeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCommerceDisputeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/shared"
	"net/http"
)

type GetPushOperationRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Push operation key.
	PushOperationKey string `pathParam:"style=simple,explode=false,name=pushOperationKey"`
}

func (o *GetPushOperationRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetPushOperationRequest) GetPushOperationKey() string {
	if o == nil {
		return ""
	}
	return o.PushOperationKey
}

type GetPushOperationResponse struct {
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	// OK
	PushOperation *shared.PushOperation
	StatusCode    int
	RawResponse   *http.Response
}

func (o *GetPushOperationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPushOperationResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetPushOperationResponse) GetPushOperation() *shared.PushOperation {
	if o == nil {
		return nil
	}
	return o.PushOperation
}

func (o *GetPushOperationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPushOperationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

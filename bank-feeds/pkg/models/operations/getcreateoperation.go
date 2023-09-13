// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/bank-feeds/v2/pkg/models/shared"
	"net/http"
)

type GetCreateOperationRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Push operation key.
	PushOperationKey string `pathParam:"style=simple,explode=false,name=pushOperationKey"`
}

func (o *GetCreateOperationRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetCreateOperationRequest) GetPushOperationKey() string {
	if o == nil {
		return ""
	}
	return o.PushOperationKey
}

type GetCreateOperationResponse struct {
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	// OK
	PushOperation *shared.PushOperation
	StatusCode    int
	RawResponse   *http.Response
}

func (o *GetCreateOperationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCreateOperationResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetCreateOperationResponse) GetPushOperation() *shared.PushOperation {
	if o == nil {
		return nil
	}
	return o.PushOperation
}

func (o *GetCreateOperationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCreateOperationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-expenses/pkg/models/shared"
	"net/http"
)

type CreateExpenseTransactionRequest struct {
	CreateExpenseRequest *shared.CreateExpenseRequest `request:"mediaType=application/json"`
	CompanyID            string                       `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *CreateExpenseTransactionRequest) GetCreateExpenseRequest() *shared.CreateExpenseRequest {
	if o == nil {
		return nil
	}
	return o.CreateExpenseRequest
}

func (o *CreateExpenseTransactionRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type CreateExpenseTransactionResponse struct {
	ContentType string
	// OK
	CreateExpenseResponse *shared.CreateExpenseResponse
	// The request made is not valid.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *CreateExpenseTransactionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateExpenseTransactionResponse) GetCreateExpenseResponse() *shared.CreateExpenseResponse {
	if o == nil {
		return nil
	}
	return o.CreateExpenseResponse
}

func (o *CreateExpenseTransactionResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *CreateExpenseTransactionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateExpenseTransactionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
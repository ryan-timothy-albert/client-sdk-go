// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-expenses/v4/pkg/models/shared"
	"net/http"
)

type CreateReimbursableExpenseTransactionRequest struct {
	CreateReimbursableExpenseRequest *shared.CreateReimbursableExpenseRequest `request:"mediaType=application/json"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *CreateReimbursableExpenseTransactionRequest) GetCreateReimbursableExpenseRequest() *shared.CreateReimbursableExpenseRequest {
	if o == nil {
		return nil
	}
	return o.CreateReimbursableExpenseRequest
}

func (o *CreateReimbursableExpenseTransactionRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type CreateReimbursableExpenseTransactionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	CreateReimbursableExpenseResponse *shared.CreateReimbursableExpenseResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateReimbursableExpenseTransactionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateReimbursableExpenseTransactionResponse) GetCreateReimbursableExpenseResponse() *shared.CreateReimbursableExpenseResponse {
	if o == nil {
		return nil
	}
	return o.CreateReimbursableExpenseResponse
}

func (o *CreateReimbursableExpenseTransactionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateReimbursableExpenseTransactionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-expenses/v4/pkg/models/shared"
	"net/http"
)

type CreateTransferTransactionRequest struct {
	CreateTransferRequest *shared.CreateTransferRequest `request:"mediaType=application/json"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// The unique identifier for your SMB's transaction.
	TransactionID string `pathParam:"style=simple,explode=false,name=transactionId"`
}

func (o *CreateTransferTransactionRequest) GetCreateTransferRequest() *shared.CreateTransferRequest {
	if o == nil {
		return nil
	}
	return o.CreateTransferRequest
}

func (o *CreateTransferTransactionRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateTransferTransactionRequest) GetTransactionID() string {
	if o == nil {
		return ""
	}
	return o.TransactionID
}

type CreateTransferTransactionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	CreateTransferResponse *shared.CreateTransferResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateTransferTransactionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTransferTransactionResponse) GetCreateTransferResponse() *shared.CreateTransferResponse {
	if o == nil {
		return nil
	}
	return o.CreateTransferResponse
}

func (o *CreateTransferTransactionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTransferTransactionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
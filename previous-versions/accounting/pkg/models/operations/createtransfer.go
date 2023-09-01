// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"net/http"
)

type CreateTransferRequest struct {
	Transfer     *shared.Transfer `request:"mediaType=application/json"`
	CompanyID    string           `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string           `pathParam:"style=simple,explode=false,name=connectionId"`
}

func (o *CreateTransferRequest) GetTransfer() *shared.Transfer {
	if o == nil {
		return nil
	}
	return o.Transfer
}

func (o *CreateTransferRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateTransferRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type CreateTransferResponse struct {
	ContentType string
	// Success
	CreateTransferResponse *shared.CreateTransferResponse
	// The request made is not valid.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *CreateTransferResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTransferResponse) GetCreateTransferResponse() *shared.CreateTransferResponse {
	if o == nil {
		return nil
	}
	return o.CreateTransferResponse
}

func (o *CreateTransferResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *CreateTransferResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTransferResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

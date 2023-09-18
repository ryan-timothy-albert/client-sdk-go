// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v4/pkg/models/shared"
	"net/http"
)

type GetAccountingTransferRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	TransferID   string `pathParam:"style=simple,explode=false,name=transferId"`
}

func (o *GetAccountingTransferRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetAccountingTransferRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetAccountingTransferRequest) GetTransferID() string {
	if o == nil {
		return ""
	}
	return o.TransferID
}

type GetAccountingTransferResponse struct {
	// Success
	AccountingTransfer *shared.AccountingTransfer
	ContentType        string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetAccountingTransferResponse) GetAccountingTransfer() *shared.AccountingTransfer {
	if o == nil {
		return nil
	}
	return o.AccountingTransfer
}

func (o *GetAccountingTransferResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountingTransferResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetAccountingTransferResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountingTransferResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

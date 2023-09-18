// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v4/pkg/models/shared"
	"net/http"
)

type GetAccountingAccountTransactionRequest struct {
	AccountTransactionID string `pathParam:"style=simple,explode=false,name=accountTransactionId"`
	CompanyID            string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID         string `pathParam:"style=simple,explode=false,name=connectionId"`
}

func (o *GetAccountingAccountTransactionRequest) GetAccountTransactionID() string {
	if o == nil {
		return ""
	}
	return o.AccountTransactionID
}

func (o *GetAccountingAccountTransactionRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetAccountingAccountTransactionRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type GetAccountingAccountTransactionResponse struct {
	// Success
	AccountingAccountTransaction *shared.AccountingAccountTransaction
	ContentType                  string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetAccountingAccountTransactionResponse) GetAccountingAccountTransaction() *shared.AccountingAccountTransaction {
	if o == nil {
		return nil
	}
	return o.AccountingAccountTransaction
}

func (o *GetAccountingAccountTransactionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountingAccountTransactionResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetAccountingAccountTransactionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountingAccountTransactionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

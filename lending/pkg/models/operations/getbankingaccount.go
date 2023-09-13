// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v2/pkg/models/shared"
	"net/http"
)

type GetBankingAccountRequest struct {
	// Unique identifier for an account
	AccountID    string `pathParam:"style=simple,explode=false,name=accountId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

func (o *GetBankingAccountRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *GetBankingAccountRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetBankingAccountRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type GetBankingAccountResponse struct {
	// Success
	BankingAccount *shared.BankingAccount
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
}

func (o *GetBankingAccountResponse) GetBankingAccount() *shared.BankingAccount {
	if o == nil {
		return nil
	}
	return o.BankingAccount
}

func (o *GetBankingAccountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetBankingAccountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetBankingAccountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

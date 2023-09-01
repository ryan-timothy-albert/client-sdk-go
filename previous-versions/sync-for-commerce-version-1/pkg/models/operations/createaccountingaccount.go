// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"net/http"
)

type CreateAccountingAccountRequest struct {
	AccountingAccount *shared.AccountingAccount `request:"mediaType=application/json"`
	CompanyID         string                    `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID      string                    `pathParam:"style=simple,explode=false,name=connectionId"`
	TimeoutInMinutes  *int                      `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

func (o *CreateAccountingAccountRequest) GetAccountingAccount() *shared.AccountingAccount {
	if o == nil {
		return nil
	}
	return o.AccountingAccount
}

func (o *CreateAccountingAccountRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateAccountingAccountRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateAccountingAccountRequest) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

type CreateAccountingAccountResponse struct {
	// Success
	AccountingCreateAccountResponse *shared.AccountingCreateAccountResponse
	ContentType                     string
	// The request made is not valid.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *CreateAccountingAccountResponse) GetAccountingCreateAccountResponse() *shared.AccountingCreateAccountResponse {
	if o == nil {
		return nil
	}
	return o.AccountingCreateAccountResponse
}

func (o *CreateAccountingAccountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAccountingAccountResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *CreateAccountingAccountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAccountingAccountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

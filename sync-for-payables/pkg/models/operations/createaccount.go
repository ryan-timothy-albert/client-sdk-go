// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/shared"
	"net/http"
)

type CreateAccountRequest struct {
	Account          *shared.Account `request:"mediaType=application/json"`
	CompanyID        string          `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID     string          `pathParam:"style=simple,explode=false,name=connectionId"`
	TimeoutInMinutes *int            `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

func (o *CreateAccountRequest) GetAccount() *shared.Account {
	if o == nil {
		return nil
	}
	return o.Account
}

func (o *CreateAccountRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateAccountRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateAccountRequest) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

type CreateAccountResponse struct {
	ContentType string
	// Success
	CreateAccountResponse *shared.CreateAccountResponse
	// The request made is not valid.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *CreateAccountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAccountResponse) GetCreateAccountResponse() *shared.CreateAccountResponse {
	if o == nil {
		return nil
	}
	return o.CreateAccountResponse
}

func (o *CreateAccountResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *CreateAccountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAccountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

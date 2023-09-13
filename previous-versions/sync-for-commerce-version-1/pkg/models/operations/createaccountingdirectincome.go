// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"net/http"
)

type CreateAccountingDirectIncomeRequest struct {
	AccountingDirectIncome *shared.AccountingDirectIncome `request:"mediaType=application/json"`
	CompanyID              string                         `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID           string                         `pathParam:"style=simple,explode=false,name=connectionId"`
	TimeoutInMinutes       *int                           `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

func (o *CreateAccountingDirectIncomeRequest) GetAccountingDirectIncome() *shared.AccountingDirectIncome {
	if o == nil {
		return nil
	}
	return o.AccountingDirectIncome
}

func (o *CreateAccountingDirectIncomeRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateAccountingDirectIncomeRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateAccountingDirectIncomeRequest) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

type CreateAccountingDirectIncomeResponse struct {
	// Success
	AccountingCreateDirectIncomeResponse *shared.AccountingCreateDirectIncomeResponse
	ContentType                          string
	// The request made is not valid.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *CreateAccountingDirectIncomeResponse) GetAccountingCreateDirectIncomeResponse() *shared.AccountingCreateDirectIncomeResponse {
	if o == nil {
		return nil
	}
	return o.AccountingCreateDirectIncomeResponse
}

func (o *CreateAccountingDirectIncomeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAccountingDirectIncomeResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *CreateAccountingDirectIncomeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAccountingDirectIncomeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
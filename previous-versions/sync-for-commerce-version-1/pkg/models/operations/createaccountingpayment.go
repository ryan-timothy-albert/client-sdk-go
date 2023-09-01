// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"net/http"
)

type CreateAccountingPaymentRequest struct {
	AccountingPayment *shared.AccountingPayment `request:"mediaType=application/json"`
	CompanyID         string                    `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID      string                    `pathParam:"style=simple,explode=false,name=connectionId"`
	TimeoutInMinutes  *int                      `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

func (o *CreateAccountingPaymentRequest) GetAccountingPayment() *shared.AccountingPayment {
	if o == nil {
		return nil
	}
	return o.AccountingPayment
}

func (o *CreateAccountingPaymentRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateAccountingPaymentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateAccountingPaymentRequest) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

type CreateAccountingPaymentResponse struct {
	// Success
	AccountingCreatePaymentResponse *shared.AccountingCreatePaymentResponse
	ContentType                     string
	// The request made is not valid.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *CreateAccountingPaymentResponse) GetAccountingCreatePaymentResponse() *shared.AccountingCreatePaymentResponse {
	if o == nil {
		return nil
	}
	return o.AccountingCreatePaymentResponse
}

func (o *CreateAccountingPaymentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAccountingPaymentResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *CreateAccountingPaymentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAccountingPaymentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

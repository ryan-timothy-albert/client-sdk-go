// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/pkg/models/shared"
	"net/http"
)

type GetAccountingCustomerRequest struct {
	CompanyID  string `pathParam:"style=simple,explode=false,name=companyId"`
	CustomerID string `pathParam:"style=simple,explode=false,name=customerId"`
}

func (o *GetAccountingCustomerRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetAccountingCustomerRequest) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

type GetAccountingCustomerResponse struct {
	// Success
	AccountingCustomer *shared.AccountingCustomer
	ContentType        string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetAccountingCustomerResponse) GetAccountingCustomer() *shared.AccountingCustomer {
	if o == nil {
		return nil
	}
	return o.AccountingCustomer
}

func (o *GetAccountingCustomerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountingCustomerResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetAccountingCustomerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountingCustomerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
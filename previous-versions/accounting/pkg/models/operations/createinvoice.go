// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"net/http"
)

type CreateInvoiceRequest struct {
	Invoice          *shared.Invoice `request:"mediaType=application/json"`
	CompanyID        string          `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID     string          `pathParam:"style=simple,explode=false,name=connectionId"`
	TimeoutInMinutes *int            `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

func (o *CreateInvoiceRequest) GetInvoice() *shared.Invoice {
	if o == nil {
		return nil
	}
	return o.Invoice
}

func (o *CreateInvoiceRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateInvoiceRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateInvoiceRequest) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

type CreateInvoiceResponse struct {
	ContentType string
	// Success
	CreateInvoiceResponse *shared.CreateInvoiceResponse
	// The request made is not valid.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *CreateInvoiceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateInvoiceResponse) GetCreateInvoiceResponse() *shared.CreateInvoiceResponse {
	if o == nil {
		return nil
	}
	return o.CreateInvoiceResponse
}

func (o *CreateInvoiceResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *CreateInvoiceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateInvoiceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

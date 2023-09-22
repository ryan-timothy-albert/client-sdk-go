// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
	"net/http"
)

type CreateBillPaymentRequest struct {
	BillPayment      *shared.BillPayment `request:"mediaType=application/json"`
	CompanyID        string              `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID     string              `pathParam:"style=simple,explode=false,name=connectionId"`
	TimeoutInMinutes *int                `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

func (o *CreateBillPaymentRequest) GetBillPayment() *shared.BillPayment {
	if o == nil {
		return nil
	}
	return o.BillPayment
}

func (o *CreateBillPaymentRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateBillPaymentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateBillPaymentRequest) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

type CreateBillPaymentResponse struct {
	ContentType string
	// Success
	CreateBillPaymentResponse *shared.CreateBillPaymentResponse
	// The request made is not valid.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *CreateBillPaymentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateBillPaymentResponse) GetCreateBillPaymentResponse() *shared.CreateBillPaymentResponse {
	if o == nil {
		return nil
	}
	return o.CreateBillPaymentResponse
}

func (o *CreateBillPaymentResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *CreateBillPaymentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateBillPaymentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"net/http"
)

type CreateBillRequest struct {
	Bill             *shared.Bill `request:"mediaType=application/json"`
	CompanyID        string       `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID     string       `pathParam:"style=simple,explode=false,name=connectionId"`
	TimeoutInMinutes *int         `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

func (o *CreateBillRequest) GetBill() *shared.Bill {
	if o == nil {
		return nil
	}
	return o.Bill
}

func (o *CreateBillRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateBillRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateBillRequest) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

type CreateBillResponse struct {
	ContentType string
	// Success
	CreateBillResponse *shared.CreateBillResponse
	// The request made is not valid.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *CreateBillResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateBillResponse) GetCreateBillResponse() *shared.CreateBillResponse {
	if o == nil {
		return nil
	}
	return o.CreateBillResponse
}

func (o *CreateBillResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *CreateBillResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateBillResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/shared"
	"net/http"
)

type UpdateBillRequest struct {
	Bill *shared.Bill `request:"mediaType=application/json"`
	// Unique identifier for a bill
	BillID       string `pathParam:"style=simple,explode=false,name=billId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// When updating data in the destination platform Codat checks the `sourceModifiedDate` against the `lastupdated` date from the accounting platform, if they're different Codat will return an error suggesting you should initiate another pull of the data. If this is set to `true` then the update will override this check.
	ForceUpdate      *bool `queryParam:"style=form,explode=true,name=forceUpdate"`
	TimeoutInMinutes *int  `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

func (o *UpdateBillRequest) GetBill() *shared.Bill {
	if o == nil {
		return nil
	}
	return o.Bill
}

func (o *UpdateBillRequest) GetBillID() string {
	if o == nil {
		return ""
	}
	return o.BillID
}

func (o *UpdateBillRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *UpdateBillRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateBillRequest) GetForceUpdate() *bool {
	if o == nil {
		return nil
	}
	return o.ForceUpdate
}

func (o *UpdateBillRequest) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

type UpdateBillResponse struct {
	ContentType string
	// The request made is not valid.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
	// Success
	UpdateBillResponse *shared.UpdateBillResponse
}

func (o *UpdateBillResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateBillResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *UpdateBillResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateBillResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateBillResponse) GetUpdateBillResponse() *shared.UpdateBillResponse {
	if o == nil {
		return nil
	}
	return o.UpdateBillResponse
}

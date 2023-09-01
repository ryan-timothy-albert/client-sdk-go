// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/shared"
	"net/http"
)

type CreateSupplierRequest struct {
	Supplier         *shared.Supplier `request:"mediaType=application/json"`
	CompanyID        string           `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID     string           `pathParam:"style=simple,explode=false,name=connectionId"`
	TimeoutInMinutes *int             `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

func (o *CreateSupplierRequest) GetSupplier() *shared.Supplier {
	if o == nil {
		return nil
	}
	return o.Supplier
}

func (o *CreateSupplierRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateSupplierRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateSupplierRequest) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

type CreateSupplierResponse struct {
	ContentType string
	// Success
	CreateSupplierResponse *shared.CreateSupplierResponse
	// The request made is not valid.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *CreateSupplierResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateSupplierResponse) GetCreateSupplierResponse() *shared.CreateSupplierResponse {
	if o == nil {
		return nil
	}
	return o.CreateSupplierResponse
}

func (o *CreateSupplierResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *CreateSupplierResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSupplierResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

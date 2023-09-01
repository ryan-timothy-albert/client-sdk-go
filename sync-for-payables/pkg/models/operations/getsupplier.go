// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/shared"
	"net/http"
)

type GetSupplierRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a supplier
	SupplierID string `pathParam:"style=simple,explode=false,name=supplierId"`
}

func (o *GetSupplierRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetSupplierRequest) GetSupplierID() string {
	if o == nil {
		return ""
	}
	return o.SupplierID
}

type GetSupplierResponse struct {
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
	// Success
	Supplier *shared.Supplier
}

func (o *GetSupplierResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSupplierResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetSupplierResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSupplierResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSupplierResponse) GetSupplier() *shared.Supplier {
	if o == nil {
		return nil
	}
	return o.Supplier
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"net/http"
)

type ListInvoiceAttachmentsRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for an invoice.
	InvoiceID string `pathParam:"style=simple,explode=false,name=invoiceId"`
}

func (o *ListInvoiceAttachmentsRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *ListInvoiceAttachmentsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListInvoiceAttachmentsRequest) GetInvoiceID() string {
	if o == nil {
		return ""
	}
	return o.InvoiceID
}

type ListInvoiceAttachmentsResponse struct {
	// Success
	AttachmentsDataset *shared.AttachmentsDataset
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListInvoiceAttachmentsResponse) GetAttachmentsDataset() *shared.AttachmentsDataset {
	if o == nil {
		return nil
	}
	return o.AttachmentsDataset
}

func (o *ListInvoiceAttachmentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListInvoiceAttachmentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListInvoiceAttachmentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

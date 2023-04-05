// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type GetInvoiceRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for an invoice
	InvoiceID string `pathParam:"style=simple,explode=false,name=invoiceId"`
}

type GetInvoiceResponse struct {
	ContentType string
	// Success
	Invoice     *shared.Invoice
	StatusCode  int
	RawResponse *http.Response
}

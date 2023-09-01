// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/pkg/models/shared"
	"net/http"
)

type GetAccountingBillCreditNoteRequest struct {
	BillCreditNoteID string `pathParam:"style=simple,explode=false,name=billCreditNoteId"`
	CompanyID        string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *GetAccountingBillCreditNoteRequest) GetBillCreditNoteID() string {
	if o == nil {
		return ""
	}
	return o.BillCreditNoteID
}

func (o *GetAccountingBillCreditNoteRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type GetAccountingBillCreditNoteResponse struct {
	// Success
	AccountingBillCreditNote *shared.AccountingBillCreditNote
	ContentType              string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetAccountingBillCreditNoteResponse) GetAccountingBillCreditNote() *shared.AccountingBillCreditNote {
	if o == nil {
		return nil
	}
	return o.AccountingBillCreditNote
}

func (o *GetAccountingBillCreditNoteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountingBillCreditNoteResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetAccountingBillCreditNoteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountingBillCreditNoteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
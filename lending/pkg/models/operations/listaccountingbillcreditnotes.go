// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/pkg/models/shared"
	"net/http"
)

type ListAccountingBillCreditNotesRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Field to order results by. [Read more](https://docs.codat.io/using-the-api/ordering-results).
	OrderBy *string `queryParam:"style=form,explode=true,name=orderBy"`
	// Page number. [Read more](https://docs.codat.io/using-the-api/paging).
	Page *int `queryParam:"style=form,explode=true,name=page"`
	// Number of records to return in a page. [Read more](https://docs.codat.io/using-the-api/paging).
	PageSize *int `queryParam:"style=form,explode=true,name=pageSize"`
	// Codat query string. [Read more](https://docs.codat.io/using-the-api/querying).
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

func (o *ListAccountingBillCreditNotesRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *ListAccountingBillCreditNotesRequest) GetOrderBy() *string {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *ListAccountingBillCreditNotesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListAccountingBillCreditNotesRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListAccountingBillCreditNotesRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

type ListAccountingBillCreditNotesResponse struct {
	// Success
	AccountingBillCreditNotes *shared.AccountingBillCreditNotes
	ContentType               string
	// Your `query` parameter was not correctly formed
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *ListAccountingBillCreditNotesResponse) GetAccountingBillCreditNotes() *shared.AccountingBillCreditNotes {
	if o == nil {
		return nil
	}
	return o.AccountingBillCreditNotes
}

func (o *ListAccountingBillCreditNotesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAccountingBillCreditNotesResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *ListAccountingBillCreditNotesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAccountingBillCreditNotesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

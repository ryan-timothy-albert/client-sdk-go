// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
	"net/http"
)

type GetEnhancedCashFlowTransactionsRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Page number. [Read more](https://docs.codat.io/using-the-api/paging).
	Page *int `queryParam:"style=form,explode=true,name=page"`
	// Number of records to return in a page. [Read more](https://docs.codat.io/using-the-api/paging).
	PageSize *int `queryParam:"style=form,explode=true,name=pageSize"`
	// Codat query string. [Read more](https://docs.codat.io/using-the-api/querying).
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

type GetEnhancedCashFlowTransactionsResponse struct {
	ContentType string
	// OK
	EnhancedCashFlowTransactions *shared.EnhancedCashFlowTransactions
	StatusCode                   int
	RawResponse                  *http.Response
	// Your API request was not properly authorized.
	Schema *shared.Schema
}
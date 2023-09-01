// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type GetCreditNoteRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	CreditNoteID string `pathParam:"style=simple,explode=false,name=creditNoteId"`
}

// GetCreditNote409ApplicationJSON - The data type's dataset has not been requested or is still syncing.
type GetCreditNote409ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetCreditNoteResponse struct {
	ContentType string
	// Success
	CreditNote  *shared.CreditNote
	StatusCode  int
	RawResponse *http.Response
	// The data type's dataset has not been requested or is still syncing.
	GetCreditNote409ApplicationJSONObject *GetCreditNote409ApplicationJSON
	// Your API request was not properly authorized.
	Schema *shared.Schema
}
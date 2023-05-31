// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type GetAccountTransactionRequest struct {
	AccountTransactionID string `pathParam:"style=simple,explode=false,name=accountTransactionId"`
	CompanyID            string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID         string `pathParam:"style=simple,explode=false,name=connectionId"`
}

// GetAccountTransaction409ApplicationJSON - The data type's dataset has not been requested or is still syncing.
type GetAccountTransaction409ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetAccountTransactionResponse struct {
	// Success
	AccountTransaction *shared.AccountTransaction
	ContentType        string
	StatusCode         int
	RawResponse        *http.Response
	// The data type's dataset has not been requested or is still syncing.
	GetAccountTransaction409ApplicationJSONObject *GetAccountTransaction409ApplicationJSON
	// Your API request was not properly authorized.
	Schema *shared.Schema
}

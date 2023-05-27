// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
	"net/http"
)

type GetBankFeedsRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetBankFeedsResponse struct {
	// Success
	BankFeedAccounts []shared.BankFeedAccount
	ContentType      string
	StatusCode       int
	RawResponse      *http.Response
	// Your API request was not properly authorized.
	Schema *shared.Schema
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type CreateBankTransactionsRequest struct {
	BankTransactions *shared.BankTransactions `request:"mediaType=application/json"`
	// Unique identifier for an account
	AccountID               string `pathParam:"style=simple,explode=false,name=accountId"`
	AllowSyncOnPushComplete *bool  `queryParam:"style=form,explode=true,name=allowSyncOnPushComplete"`
	CompanyID               string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID            string `pathParam:"style=simple,explode=false,name=connectionId"`
	TimeoutInMinutes        *int   `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type CreateBankTransactionsResponse struct {
	ContentType string
	// Success
	CreateBankTransactionsResponse *shared.CreateBankTransactionsResponse
	StatusCode                     int
	RawResponse                    *http.Response
	// Your API request was not properly authorized.
	Schema *shared.Schema
}

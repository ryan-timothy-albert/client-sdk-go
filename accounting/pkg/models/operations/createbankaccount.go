// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type CreateBankAccountRequest struct {
	BankAccount             *shared.BankAccount `request:"mediaType=application/json"`
	AllowSyncOnPushComplete *bool               `queryParam:"style=form,explode=true,name=allowSyncOnPushComplete"`
	CompanyID               string              `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID            string              `pathParam:"style=simple,explode=false,name=connectionId"`
	TimeoutInMinutes        *int                `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type CreateBankAccountResponse struct {
	ContentType string
	// Success
	CreateBankAccountResponse *shared.CreateBankAccountResponse
	StatusCode                int
	RawResponse               *http.Response
}

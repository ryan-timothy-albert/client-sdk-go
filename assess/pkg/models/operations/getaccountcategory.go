// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
	"net/http"
)

type GetAccountCategoryRequest struct {
	// Nominal account id
	AccountID    string `pathParam:"style=simple,explode=false,name=accountId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetAccountCategoryResponse struct {
	// OK
	CategorisedAccount *shared.CategorisedAccount
	ContentType        string
	StatusCode         int
	RawResponse        *http.Response
}

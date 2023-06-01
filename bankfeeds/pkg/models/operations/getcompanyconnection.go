// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
	"net/http"
)

type GetCompanyConnectionRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetCompanyConnectionResponse struct {
	// OK
	Connection  *shared.Connection
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Your API request was not properly authorized.
	Schema *shared.Schema
}

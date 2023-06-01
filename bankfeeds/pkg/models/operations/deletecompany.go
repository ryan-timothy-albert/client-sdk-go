// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
	"net/http"
)

type DeleteCompanyRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type DeleteCompanyResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Your API request was not properly authorized.
	Schema *shared.Schema
}

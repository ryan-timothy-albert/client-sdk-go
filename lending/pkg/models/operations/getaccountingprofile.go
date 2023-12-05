// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/shared"
	"net/http"
)

type GetAccountingProfileRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *GetAccountingProfileRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type GetAccountingProfileResponse struct {
	// Success
	AccountingCompanyInfo *shared.AccountingCompanyInfo
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAccountingProfileResponse) GetAccountingCompanyInfo() *shared.AccountingCompanyInfo {
	if o == nil {
		return nil
	}
	return o.AccountingCompanyInfo
}

func (o *GetAccountingProfileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountingProfileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountingProfileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

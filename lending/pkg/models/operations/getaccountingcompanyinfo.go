// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/pkg/models/shared"
	"net/http"
)

type GetAccountingCompanyInfoRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *GetAccountingCompanyInfoRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type GetAccountingCompanyInfoResponse struct {
	// Success
	AccountingCompanyInfo *shared.AccountingCompanyInfo
	ContentType           string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetAccountingCompanyInfoResponse) GetAccountingCompanyInfo() *shared.AccountingCompanyInfo {
	if o == nil {
		return nil
	}
	return o.AccountingCompanyInfo
}

func (o *GetAccountingCompanyInfoResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountingCompanyInfoResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetAccountingCompanyInfoResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountingCompanyInfoResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

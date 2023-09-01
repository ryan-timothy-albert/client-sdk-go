// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"net/http"
)

type RefreshCompanyInfoRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *RefreshCompanyInfoRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type RefreshCompanyInfoResponse struct {
	ContentType string
	// Success
	Dataset *shared.Dataset
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *RefreshCompanyInfoResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RefreshCompanyInfoResponse) GetDataset() *shared.Dataset {
	if o == nil {
		return nil
	}
	return o.Dataset
}

func (o *RefreshCompanyInfoResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *RefreshCompanyInfoResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RefreshCompanyInfoResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

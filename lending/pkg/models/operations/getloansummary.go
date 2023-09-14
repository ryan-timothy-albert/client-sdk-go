// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v3/pkg/models/shared"
	"net/http"
)

type GetLoanSummaryRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *GetLoanSummaryRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type GetLoanSummaryResponse struct {
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	// OK
	LoanSummary *shared.LoanSummary
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetLoanSummaryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLoanSummaryResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetLoanSummaryResponse) GetLoanSummary() *shared.LoanSummary {
	if o == nil {
		return nil
	}
	return o.LoanSummary
}

func (o *GetLoanSummaryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLoanSummaryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"net/http"
)

type GetAccountsForEnhancedBalanceSheetRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// The number of periods to return. If not provided, 12 periods will be used as the default value.
	NumberOfPeriods *int64 `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	// The date in which the report is created up to. Users must specify a specific date, however the response will be provided for the full month.
	ReportDate string `queryParam:"style=form,explode=true,name=reportDate"`
}

func (o *GetAccountsForEnhancedBalanceSheetRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetAccountsForEnhancedBalanceSheetRequest) GetNumberOfPeriods() *int64 {
	if o == nil {
		return nil
	}
	return o.NumberOfPeriods
}

func (o *GetAccountsForEnhancedBalanceSheetRequest) GetReportDate() string {
	if o == nil {
		return ""
	}
	return o.ReportDate
}

type GetAccountsForEnhancedBalanceSheetResponse struct {
	ContentType string
	// OK
	EnhancedReport *shared.EnhancedReport
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetAccountsForEnhancedBalanceSheetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountsForEnhancedBalanceSheetResponse) GetEnhancedReport() *shared.EnhancedReport {
	if o == nil {
		return nil
	}
	return o.EnhancedReport
}

func (o *GetAccountsForEnhancedBalanceSheetResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetAccountsForEnhancedBalanceSheetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountsForEnhancedBalanceSheetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

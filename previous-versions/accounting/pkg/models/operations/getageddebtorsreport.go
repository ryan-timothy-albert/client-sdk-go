// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/types"
	"net/http"
)

type GetAgedDebtorsReportRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Number of periods to include in the report.
	NumberOfPeriods *int `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	// The length of period in days.
	PeriodLengthDays *int `queryParam:"style=form,explode=true,name=periodLengthDays"`
	// Date the report is generated up to.
	ReportDate *types.Date `queryParam:"style=form,explode=true,name=reportDate"`
}

func (o *GetAgedDebtorsReportRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetAgedDebtorsReportRequest) GetNumberOfPeriods() *int {
	if o == nil {
		return nil
	}
	return o.NumberOfPeriods
}

func (o *GetAgedDebtorsReportRequest) GetPeriodLengthDays() *int {
	if o == nil {
		return nil
	}
	return o.PeriodLengthDays
}

func (o *GetAgedDebtorsReportRequest) GetReportDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.ReportDate
}

type GetAgedDebtorsReportResponse struct {
	// OK
	AgedDebtorReport *shared.AgedDebtorReport
	ContentType      string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetAgedDebtorsReportResponse) GetAgedDebtorReport() *shared.AgedDebtorReport {
	if o == nil {
		return nil
	}
	return o.AgedDebtorReport
}

func (o *GetAgedDebtorsReportResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAgedDebtorsReportResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetAgedDebtorsReportResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAgedDebtorsReportResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

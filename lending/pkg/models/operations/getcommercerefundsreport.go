// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v3/pkg/models/shared"
	"net/http"
)

type GetCommerceRefundsReportRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Shows the dimensionDisplayName and itemDisplayName in measures to make the report data human-readable.
	IncludeDisplayNames *bool `queryParam:"style=form,explode=true,name=includeDisplayNames"`
	// The number of periods to return. There will be no pagination as a query parameter.
	NumberOfPeriods int64 `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	// The number of months per period. E.g. 2 = 2 months per period.
	PeriodLength int64 `queryParam:"style=form,explode=true,name=periodLength"`
	// The period unit of time returned.
	PeriodUnit shared.PeriodUnit `queryParam:"style=form,explode=true,name=periodUnit"`
	// The date in which the report is created up to. Users must specify a specific date, however the response will be provided for the full month.
	ReportDate string `queryParam:"style=form,explode=true,name=reportDate"`
}

func (o *GetCommerceRefundsReportRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetCommerceRefundsReportRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetCommerceRefundsReportRequest) GetIncludeDisplayNames() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeDisplayNames
}

func (o *GetCommerceRefundsReportRequest) GetNumberOfPeriods() int64 {
	if o == nil {
		return 0
	}
	return o.NumberOfPeriods
}

func (o *GetCommerceRefundsReportRequest) GetPeriodLength() int64 {
	if o == nil {
		return 0
	}
	return o.PeriodLength
}

func (o *GetCommerceRefundsReportRequest) GetPeriodUnit() shared.PeriodUnit {
	if o == nil {
		return shared.PeriodUnit("")
	}
	return o.PeriodUnit
}

func (o *GetCommerceRefundsReportRequest) GetReportDate() string {
	if o == nil {
		return ""
	}
	return o.ReportDate
}

type GetCommerceRefundsReportResponse struct {
	// OK
	CommerceReport *shared.CommerceReport
	ContentType    string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetCommerceRefundsReportResponse) GetCommerceReport() *shared.CommerceReport {
	if o == nil {
		return nil
	}
	return o.CommerceReport
}

func (o *GetCommerceRefundsReportResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCommerceRefundsReportResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetCommerceRefundsReportResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCommerceRefundsReportResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

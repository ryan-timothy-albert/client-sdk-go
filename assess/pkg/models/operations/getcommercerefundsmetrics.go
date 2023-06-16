// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
	"net/http"
)

type GetCommerceRefundsMetricsRequest struct {
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

type GetCommerceRefundsMetricsResponse struct {
	ContentType string
	// OK
	Report      *shared.Report
	StatusCode  int
	RawResponse *http.Response
	// Your API request was not properly authorized.
	Schema *shared.Schema
}

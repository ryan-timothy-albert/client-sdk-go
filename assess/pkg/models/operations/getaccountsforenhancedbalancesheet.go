// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetAccountsForEnhancedBalanceSheetRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// The number of periods to return.  There will be no pagination as a query parameter, however Codat will limit the number of periods to request to 12 periods.
	NumberOfPeriods int64 `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	// The date in which the report is created up to. Users must specify a specific date, however the response will be provided for the full month.
	ReportDate string `queryParam:"style=form,explode=true,name=reportDate"`
}

type GetAccountsForEnhancedBalanceSheetEnhancedReportReportInfo struct {
	// Company name the report was generated for.
	CompanyName *string `json:"companyName,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
	// The date the report was generated.
	GeneratedDate *string `json:"generatedDate,omitempty"`
	// The name of the report.
	ReportName *string `json:"reportName,omitempty"`
}

type GetAccountsForEnhancedBalanceSheetEnhancedReportReportItemsAccountCategoryLevels struct {
	Confidence *float64 `json:"confidence,omitempty"`
	LevelName  *string  `json:"levelName,omitempty"`
}

type GetAccountsForEnhancedBalanceSheetEnhancedReportReportItemsAccountCategory struct {
	Levels []GetAccountsForEnhancedBalanceSheetEnhancedReportReportItemsAccountCategoryLevels `json:"levels,omitempty"`
	Status *string                                                                            `json:"status,omitempty"`
}

type GetAccountsForEnhancedBalanceSheetEnhancedReportReportItems struct {
	AccountCategory *GetAccountsForEnhancedBalanceSheetEnhancedReportReportItemsAccountCategory `json:"accountCategory,omitempty"`
	// The unique account ID.
	AccountID   *string  `json:"accountId,omitempty"`
	AccountName *string  `json:"accountName,omitempty"`
	Balance     *float64 `json:"balance,omitempty"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	Date *string `json:"date,omitempty"`
}

// GetAccountsForEnhancedBalanceSheetEnhancedReport - OK
type GetAccountsForEnhancedBalanceSheetEnhancedReport struct {
	ReportInfo *GetAccountsForEnhancedBalanceSheetEnhancedReportReportInfo `json:"reportInfo,omitempty"`
	// An array of report items.
	ReportItems []GetAccountsForEnhancedBalanceSheetEnhancedReportReportItems `json:"reportItems,omitempty"`
}

type GetAccountsForEnhancedBalanceSheetResponse struct {
	ContentType string
	// OK
	EnhancedReport *GetAccountsForEnhancedBalanceSheetEnhancedReport
	StatusCode     int
	RawResponse    *http.Response
}

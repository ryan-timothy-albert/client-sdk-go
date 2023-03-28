// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/types"
	"net/http"
)

type GetAgedCreditorsReportRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Number of periods to include in the report.
	NumberOfPeriods *int `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	// The length of period in days.
	PeriodLengthDays *int `queryParam:"style=form,explode=true,name=periodLengthDays"`
	// Date the report is generated up to.
	ReportDate *types.Date `queryParam:"style=form,explode=true,name=reportDate"`
}

type GetAgedCreditorsReportAgedCreditorsReportAgedCreditorAgedCurrencyOutstandingAgedOutstandingAmountAmountsOutstandingByDataType struct {
	// The amount outstanding.
	Amount *float64 `json:"amount,omitempty"`
	// Name of data type with outstanding amount for given period.
	Name *string `json:"name,omitempty"`
}

type GetAgedCreditorsReportAgedCreditorsReportAgedCreditorAgedCurrencyOutstandingAgedOutstandingAmount struct {
	// The amount outstanding.
	Amount *float64 `json:"amount,omitempty"`
	// Array of details.
	Details []GetAgedCreditorsReportAgedCreditorsReportAgedCreditorAgedCurrencyOutstandingAgedOutstandingAmountAmountsOutstandingByDataType `json:"details,omitempty"`
	// Start date of period.
	FromDate *string `json:"fromDate,omitempty"`
	// End date of period.
	ToDate *string `json:"toDate,omitempty"`
}

type GetAgedCreditorsReportAgedCreditorsReportAgedCreditorAgedCurrencyOutstanding struct {
	// Array of outstanding amounts by period.
	AgedOutstandingAmounts []GetAgedCreditorsReportAgedCreditorsReportAgedCreditorAgedCurrencyOutstandingAgedOutstandingAmount `json:"agedOutstandingAmounts,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code. e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
}

type GetAgedCreditorsReportAgedCreditorsReportAgedCreditor struct {
	// Array of aged creditors by currency.
	AgedCurrencyOutstanding []GetAgedCreditorsReportAgedCreditorsReportAgedCreditorAgedCurrencyOutstanding `json:"agedCurrencyOutstanding,omitempty"`
	// Supplier ID of the aged creditor.
	SupplierID *string `json:"supplierId,omitempty"`
	// Supplier name of the aged creditor.
	SupplierName *string `json:"supplierName,omitempty"`
}

// GetAgedCreditorsReportAgedCreditorsReport - OK
type GetAgedCreditorsReportAgedCreditorsReport struct {
	// Array of aged debtors.
	Data []GetAgedCreditorsReportAgedCreditorsReportAgedCreditor `json:"data,omitempty"`
	// Date and time the report was generated.
	Generated *string `json:"generated,omitempty"`
	// Date the report is generated up to.
	ReportDate *string `json:"reportDate,omitempty"`
}

type GetAgedCreditorsReportResponse struct {
	// OK
	AgedCreditorsReport *GetAgedCreditorsReportAgedCreditorsReport
	ContentType         string
	StatusCode          int
	RawResponse         *http.Response
}

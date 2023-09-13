// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v2/pkg/types"
)

type LoanSummaryReportItem struct {
	// The loan outstanding balance.  This may not equal totalDrawdowns - totalRepayments due to interest which has been accrued.
	Balance *types.Decimal `json:"balance,omitempty"`
	// The description of the object being referred to. E.g. the account.
	Description *string               `json:"description,omitempty"`
	RecordRef   *LoanSummaryRecordRef `json:"recordRef,omitempty"`
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
	StartDate *string `json:"startDate,omitempty"`
	// The total loan drawdowns.
	TotalDrawdowns *types.Decimal `json:"totalDrawdowns,omitempty"`
	// The total loan repayments which includes capital plus any interest.
	TotalRepayments *types.Decimal `json:"totalRepayments,omitempty"`
}

func (o *LoanSummaryReportItem) GetBalance() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.Balance
}

func (o *LoanSummaryReportItem) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *LoanSummaryReportItem) GetRecordRef() *LoanSummaryRecordRef {
	if o == nil {
		return nil
	}
	return o.RecordRef
}

func (o *LoanSummaryReportItem) GetStartDate() *string {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *LoanSummaryReportItem) GetTotalDrawdowns() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.TotalDrawdowns
}

func (o *LoanSummaryReportItem) GetTotalRepayments() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.TotalRepayments
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AgedCreditorAgedCurrencyOutstandingAgedOutstandingAmountAmountsOutstandingByDataType struct {
	// The amount outstanding.
	Amount *float64 `json:"amount,omitempty"`
	// Name of data type with outstanding amount for given period.
	Name *string `json:"name,omitempty"`
}

type AgedCreditorAgedCurrencyOutstandingAgedOutstandingAmount struct {
	// The amount outstanding.
	Amount *float64 `json:"amount,omitempty"`
	// Array of details.
	Details []AgedCreditorAgedCurrencyOutstandingAgedOutstandingAmountAmountsOutstandingByDataType `json:"details,omitempty"`
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
	FromDate *string `json:"fromDate,omitempty"`
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
	ToDate *string `json:"toDate,omitempty"`
}

type AgedCreditorAgedCurrencyOutstanding struct {
	// Array of outstanding amounts by period.
	AgedOutstandingAmounts []AgedCreditorAgedCurrencyOutstandingAgedOutstandingAmount `json:"agedOutstandingAmounts,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
}

type AgedCreditor struct {
	// Array of aged creditors by currency.
	AgedCurrencyOutstanding []AgedCreditorAgedCurrencyOutstanding `json:"agedCurrencyOutstanding,omitempty"`
	// Supplier ID of the aged creditor.
	SupplierID *string `json:"supplierId,omitempty"`
	// Supplier name of the aged creditor.
	SupplierName *string `json:"supplierName,omitempty"`
}

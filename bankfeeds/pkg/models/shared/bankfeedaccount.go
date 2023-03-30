// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// BankFeedAccount - The target bank account in a supported accounting package for ingestion into a bank feed.
type BankFeedAccount struct {
	// The bank account name
	AccountName *string `json:"accountName,omitempty"`
	// The account number
	AccountNumber *string `json:"accountNumber,omitempty"`
	// The type of bank account e.g. Credit
	AccountType *string `json:"accountType,omitempty"`
	// The latest balance for the bank account
	Balance *float64 `json:"balance,omitempty"`
	// The currency e.g. USD
	Currency *string `json:"currency,omitempty"`
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
	FeedStartDate *string `json:"feedStartDate,omitempty"`
	// Unique ID for the bank feed account
	ID string `json:"id"`
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
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// The sort code
	SortCode *string `json:"sortCode,omitempty"`
	Status   *string `json:"status,omitempty"`
}

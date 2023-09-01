// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DataIntegrityByAmount struct {
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
	// The percentage of the absolute value of transactions of the type specified in the route which have a match.
	MatchPercentage *float64 `json:"matchPercentage,omitempty"`
	// The sum of the absolute value of transactions of the type specified in the route which have a match.
	Matched *float64 `json:"matched,omitempty"`
	// The total of unmatched and matched.
	Total *float64 `json:"total,omitempty"`
	// The sum of the absolute value of transactions of the type specified in the route which don't have a match.
	Unmatched *float64 `json:"unmatched,omitempty"`
}
// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DataIntegrityMatch struct {
	// The transaction value.
	Amount *string `json:"amount,omitempty"`
	// ID GUID representing the connection of the accounting or banking platform.
	ConnectionID *string `json:"connectionId,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
	// The date of the transaction.
	Date *string `json:"date,omitempty"`
	// The transaction description.
	Description *string `json:"description,omitempty"`
	// ID GUID of the transaction.
	ID *string `json:"id,omitempty"`
	// The data type which the data type in the URL has been matched against. For example, if you've matched accountTransactions and banking-transactions, and you call this endpoint with accountTransactions in the URL, this property would be banking-transactions.
	Type *string `json:"type,omitempty"`
}
// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AgedCurrencyOutstanding struct {
	// Array of outstanding amounts by period.
	AgedOutstandingAmounts []AgedOutstandingAmount `json:"agedOutstandingAmounts,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
}

func (o *AgedCurrencyOutstanding) GetAgedOutstandingAmounts() []AgedOutstandingAmount {
	if o == nil {
		return nil
	}
	return o.AgedOutstandingAmounts
}

func (o *AgedCurrencyOutstanding) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

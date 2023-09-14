// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v3/pkg/types"
)

type ProductPrice struct {
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency  *string        `json:"currency,omitempty"`
	UnitPrice *types.Decimal `json:"unitPrice,omitempty"`
}

func (o *ProductPrice) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *ProductPrice) GetUnitPrice() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.UnitPrice
}

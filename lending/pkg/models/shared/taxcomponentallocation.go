// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v2/pkg/types"
)

type TaxComponentAllocation struct {
	// Tax amount on order line sale as available from source commerce platform.
	Rate *types.Decimal `json:"rate,omitempty"`
	// Taxes rates reference object depending on the rates being available on source commerce package.
	TaxComponentRef *TaxComponentRef `json:"taxComponentRef,omitempty"`
}

func (o *TaxComponentAllocation) GetRate() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.Rate
}

func (o *TaxComponentAllocation) GetTaxComponentRef() *TaxComponentRef {
	if o == nil {
		return nil
	}
	return o.TaxComponentRef
}

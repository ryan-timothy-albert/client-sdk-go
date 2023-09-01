// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type TaxRateMappingInfoValidTransactionTypes string

const (
	TaxRateMappingInfoValidTransactionTypesPayment       TaxRateMappingInfoValidTransactionTypes = "Payment"
	TaxRateMappingInfoValidTransactionTypesRefund        TaxRateMappingInfoValidTransactionTypes = "Refund"
	TaxRateMappingInfoValidTransactionTypesReward        TaxRateMappingInfoValidTransactionTypes = "Reward"
	TaxRateMappingInfoValidTransactionTypesChargeback    TaxRateMappingInfoValidTransactionTypes = "Chargeback"
	TaxRateMappingInfoValidTransactionTypesTransferIn    TaxRateMappingInfoValidTransactionTypes = "TransferIn"
	TaxRateMappingInfoValidTransactionTypesTransferOut   TaxRateMappingInfoValidTransactionTypes = "TransferOut"
	TaxRateMappingInfoValidTransactionTypesAdjustmentIn  TaxRateMappingInfoValidTransactionTypes = "AdjustmentIn"
	TaxRateMappingInfoValidTransactionTypesAdjustmentOut TaxRateMappingInfoValidTransactionTypes = "AdjustmentOut"
)

func (e TaxRateMappingInfoValidTransactionTypes) ToPointer() *TaxRateMappingInfoValidTransactionTypes {
	return &e
}

func (e *TaxRateMappingInfoValidTransactionTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Payment":
		fallthrough
	case "Refund":
		fallthrough
	case "Reward":
		fallthrough
	case "Chargeback":
		fallthrough
	case "TransferIn":
		fallthrough
	case "TransferOut":
		fallthrough
	case "AdjustmentIn":
		fallthrough
	case "AdjustmentOut":
		*e = TaxRateMappingInfoValidTransactionTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaxRateMappingInfoValidTransactionTypes: %v", v)
	}
}

type TaxRateMappingInfo struct {
	// Code for the tax rate from the accounting platform.
	Code *string `json:"code,omitempty"`
	// Effective tax rate.
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	// Unique identifier of tax rate.
	ID *string `json:"id,omitempty"`
	// Name of the tax rate in the accounting platform.
	Name *string `json:"name,omitempty"`
	// Total (not compounded) sum of the components of a tax rate.
	TotalTaxRate *float64 `json:"totalTaxRate,omitempty"`
	// Supported transaction types for the account.
	ValidTransactionTypes []TaxRateMappingInfoValidTransactionTypes `json:"validTransactionTypes,omitempty"`
}

func (o *TaxRateMappingInfo) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *TaxRateMappingInfo) GetEffectiveTaxRate() *float64 {
	if o == nil {
		return nil
	}
	return o.EffectiveTaxRate
}

func (o *TaxRateMappingInfo) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TaxRateMappingInfo) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TaxRateMappingInfo) GetTotalTaxRate() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalTaxRate
}

func (o *TaxRateMappingInfo) GetValidTransactionTypes() []TaxRateMappingInfoValidTransactionTypes {
	if o == nil {
		return nil
	}
	return o.ValidTransactionTypes
}

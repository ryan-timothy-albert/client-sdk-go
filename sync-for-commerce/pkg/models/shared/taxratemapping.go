// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TaxRateMapping struct {
	// Selected tax rate id from the list of tax rates on the accounting software.
	SelectedAccountingTaxRateID *string `json:"selectedAccountingTaxRateId,omitempty"`
	// Selected tax component id from the list of tax components on the commerce software.
	SelectedCommerceTaxRateIds []string `json:"selectedCommerceTaxRateIds,omitempty"`
}
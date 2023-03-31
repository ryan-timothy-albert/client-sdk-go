// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TaxRateComponent - A tax rate can be made up of multiple sub taxes, often called components of the tax.
type TaxRateComponent struct {
	// A flag to indicate with the tax is calculated using the principle of compounding.
	IsCompound bool `json:"isCompound"`
	// Name of the tax rate component.
	Name *string `json:"name,omitempty"`
	// The rate of the tax rate component, usually a percentage.
	Rate *float64 `json:"rate,omitempty"`
}

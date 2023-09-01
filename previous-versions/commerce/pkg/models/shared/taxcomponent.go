// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TaxComponent - The Tax Components endpoints return tax rates data from the commerce platform, including tax rate names and values. This is to support the mapping of tax rates from the commerce platform to those in the accounting platform.
//
// Explore our [data coverage](https://knowledge.codat.io/supported-features/commerce?view=tab-by-data-type&dataType=commerce-disputes) for this data type.
type TaxComponent struct {
	// A unique, persistent identifier for this record
	ID string `json:"id"`
	// The Boolean flag to indicate when a Tax Rate Component compounds on a sale.
	IsCompound   *bool   `json:"isCompound,omitempty"`
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// Name of the Tax Rate Component in the source commerce platform.
	Name *string `json:"name,omitempty"`
	// Rate of taxation represented as a fraction of the net price (typically in the range 0.00 - 1.00).
	Rate               *float32 `json:"rate,omitempty"`
	SourceModifiedDate *string  `json:"sourceModifiedDate,omitempty"`
}
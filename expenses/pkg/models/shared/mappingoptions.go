// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// MappingOptions - Success
type MappingOptions struct {
	// Array of available accounts for mapping.
	Accounts []AccountMappingInfo `json:"accounts,omitempty"`
	// Name of the expense integration.
	ExpenseProvider *string `json:"expenseProvider,omitempty"`
	// Array of available tax rates for mapping.
	TaxRates []TaxRateMappingInfo `json:"taxRates,omitempty"`
	// Array of available tracking categories for mapping.
	TrackingCategories []TrackingCategoryMappingInfo `json:"trackingCategories,omitempty"`
}
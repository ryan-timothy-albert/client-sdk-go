// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type FeesSupplier struct {
	// Selected supplier id from the list of supplier records on the accounting software.
	SelectedSupplierID *string `json:"selectedSupplierId,omitempty"`
	// List of supplier options from the list of supplier records on the accounting software.
	SupplierOptions []Option `json:"supplierOptions,omitempty"`
}

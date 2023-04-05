// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Customer struct {
	// List of customer options from the list of customer records on the accounting software.
	CustomerOptions []Option `json:"customerOptions,omitempty"`
	// Selected customer id from the list of customer records on the accounting software.
	SelectedCustomerID *string `json:"selectedCustomerId,omitempty"`
}

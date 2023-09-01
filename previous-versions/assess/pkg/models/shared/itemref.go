// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ItemRef struct {
	// The data connection id being referenced.
	DataConnectionID *string `json:"dataConnectionId,omitempty"`
	// The id of the object, e.g. the Journal entry.
	ID *string `json:"id,omitempty"`
	// The data type the loan transaction entry was extracted from.
	Type *string `json:"type,omitempty"`
}
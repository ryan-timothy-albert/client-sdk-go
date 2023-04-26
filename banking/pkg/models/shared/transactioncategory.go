// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TransactionCategory - Status of the bank transaction category.
type TransactionCategory struct {
	// A Boolean indicating whether there are other bank transaction categories beneath this one in the hierarchy.
	HasChildren *bool `json:"hasChildren,omitempty"`
	// The unique identifier of the bank transaction category.
	ID           string  `json:"id"`
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// The name of the bank transaction category.
	Name string `json:"name"`
	// The unique identifier of the parent bank transaction category.
	ParentID           *string                        `json:"parentId,omitempty"`
	SourceModifiedDate *string                        `json:"sourceModifiedDate,omitempty"`
	Status             *TransactionCategoryStatusEnum `json:"status,omitempty"`
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Propertiestracking1 - Categories, and a project and customer, against which the item is tracked.
type Propertiestracking1 struct {
	CategoryRefs []TrackingCategoryRef `json:"categoryRefs"`
	CustomerRef  *CustomerRef          `json:"customerRef,omitempty"`
	IsBilledTo   BilledToType1         `json:"isBilledTo"`
	IsRebilledTo BilledToType1         `json:"isRebilledTo"`
	ProjectRef   *ProjectRef           `json:"projectRef,omitempty"`
}
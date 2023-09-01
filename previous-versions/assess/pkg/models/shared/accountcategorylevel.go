// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AccountCategoryLevel - An object containing an ordered list of account category levels.
type AccountCategoryLevel struct {
	// Confidence level of the category. This will only be populated where `status` is `Suggested`.
	Confidence *float64 `json:"confidence,omitempty"`
	// Account category name.
	LevelName *string `json:"levelName,omitempty"`
}
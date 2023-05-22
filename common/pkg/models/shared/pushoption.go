// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PushOption - OK
type PushOption struct {
	Description *string                       `json:"description,omitempty"`
	DisplayName string                        `json:"displayName"`
	Options     []PushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]PushOptionProperty `json:"properties,omitempty"`
	Required    bool                          `json:"required"`
	Type        PushOptionType                `json:"type"`
	Validation  *PushValidationInfo           `json:"validation,omitempty"`
}

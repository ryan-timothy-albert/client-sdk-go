// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

// GetConfigTextSyncFlow200ApplicationJSONAdditionalProp1 - Placeholder for the property name.
type GetConfigTextSyncFlow200ApplicationJSONAdditionalProp1 struct {
	Required *bool `json:"required,omitempty"`
	// Value of the property.
	Text *string `json:"text,omitempty"`
}

// GetConfigTextSyncFlow200ApplicationJSONAdditionalProp2 - Placeholder for the property name.
type GetConfigTextSyncFlow200ApplicationJSONAdditionalProp2 struct {
	Required *bool `json:"required,omitempty"`
	// Value of the property.
	Text *string `json:"text,omitempty"`
}

// GetConfigTextSyncFlow200ApplicationJSONAdditionalProp3 - Placeholder for the property name.
type GetConfigTextSyncFlow200ApplicationJSONAdditionalProp3 struct {
	Required *bool `json:"required,omitempty"`
	// Value of the property.
	Text *string `json:"text,omitempty"`
}

// GetConfigTextSyncFlow200ApplicationJSON - Success
type GetConfigTextSyncFlow200ApplicationJSON struct {
	// Placeholder for the property name.
	AdditionalProp1 *GetConfigTextSyncFlow200ApplicationJSONAdditionalProp1 `json:"additionalProp1,omitempty"`
	// Placeholder for the property name.
	AdditionalProp2 *GetConfigTextSyncFlow200ApplicationJSONAdditionalProp2 `json:"additionalProp2,omitempty"`
	// Placeholder for the property name.
	AdditionalProp3 *GetConfigTextSyncFlow200ApplicationJSONAdditionalProp3 `json:"additionalProp3,omitempty"`
}

type GetConfigTextSyncFlowResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Success
	GetConfigTextSyncFlow200ApplicationJSONObject *GetConfigTextSyncFlow200ApplicationJSON
}
// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
	"net/http"
)

// UpdateProfileSyncSettingsRequestBody - Include a `syncSetting` object for each data type.
// `syncFromWindow`, `syncFromUTC` & `monthsToSync` only need to be included if you wish to set a value for them.
type UpdateProfileSyncSettingsRequestBody struct {
	ClientID          string               `json:"clientId"`
	OverridesDefaults bool                 `json:"overridesDefaults"`
	Settings          []shared.SyncSetting `json:"settings"`
}

type UpdateProfileSyncSettingsResponse struct {
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}
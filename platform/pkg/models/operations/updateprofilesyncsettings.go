// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/platform/pkg/models/shared"
	"github.com/codatio/client-sdk-go/platform/pkg/utils"
	"net/http"
)

// UpdateProfileSyncSettingsRequestBody - Include a `syncSetting` object for each data type.
// `syncFromWindow`, `syncFromUTC` & `monthsToSync` only need to be included if you wish to set a value for them.
type UpdateProfileSyncSettingsRequestBody struct {
	// Unique identifier for your client in Codat.
	ClientID string `json:"clientId"`
	// Set to `True` if you want to override default [sync settings](https://docs.codat.io/knowledge-base/advanced-sync-settings).
	OverridesDefaults *bool                `default:"true" json:"overridesDefaults"`
	Settings          []shared.SyncSetting `json:"settings"`
}

func (u UpdateProfileSyncSettingsRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateProfileSyncSettingsRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateProfileSyncSettingsRequestBody) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *UpdateProfileSyncSettingsRequestBody) GetOverridesDefaults() *bool {
	if o == nil {
		return nil
	}
	return o.OverridesDefaults
}

func (o *UpdateProfileSyncSettingsRequestBody) GetSettings() []shared.SyncSetting {
	if o == nil {
		return []shared.SyncSetting{}
	}
	return o.Settings
}

type UpdateProfileSyncSettingsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateProfileSyncSettingsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateProfileSyncSettingsResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *UpdateProfileSyncSettingsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateProfileSyncSettingsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

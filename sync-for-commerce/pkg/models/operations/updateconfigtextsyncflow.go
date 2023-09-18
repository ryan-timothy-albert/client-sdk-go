// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
	"net/http"
)

type UpdateConfigTextSyncFlowResponse struct {
	ContentType string
	// Success
	LocalizationInfo map[string]shared.Localization
	StatusCode       int
	RawResponse      *http.Response
}

func (o *UpdateConfigTextSyncFlowResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateConfigTextSyncFlowResponse) GetLocalizationInfo() map[string]shared.Localization {
	if o == nil {
		return nil
	}
	return o.LocalizationInfo
}

func (o *UpdateConfigTextSyncFlowResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateConfigTextSyncFlowResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

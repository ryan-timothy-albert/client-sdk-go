// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetIntegrationsPlatformKeyBrandingRequest struct {
	PlatformKey string `pathParam:"style=simple,explode=false,name=platformKey"`
}

// GetIntegrationsPlatformKeyBrandingBrandingLogoBrandingImageImage - Image reference.
type GetIntegrationsPlatformKeyBrandingBrandingLogoBrandingImageImage struct {
	// Alternative text when image is not available.
	Alt *string `json:"alt,omitempty"`
	// Source URL for image.
	Src *string `json:"src,omitempty"`
}

type GetIntegrationsPlatformKeyBrandingBrandingLogoBrandingImage struct {
	// Image reference.
	Image *GetIntegrationsPlatformKeyBrandingBrandingLogoBrandingImageImage `json:"image,omitempty"`
}

// GetIntegrationsPlatformKeyBrandingBrandingLogo - Logo branding references.
type GetIntegrationsPlatformKeyBrandingBrandingLogo struct {
	Full   *GetIntegrationsPlatformKeyBrandingBrandingLogoBrandingImage `json:"full,omitempty"`
	Square *GetIntegrationsPlatformKeyBrandingBrandingLogoBrandingImage `json:"square,omitempty"`
}

// GetIntegrationsPlatformKeyBrandingBranding - OK
type GetIntegrationsPlatformKeyBrandingBranding struct {
	// Button branding references.
	Button interface{} `json:"button,omitempty"`
	// Logo branding references.
	Logo *GetIntegrationsPlatformKeyBrandingBrandingLogo `json:"logo,omitempty"`
	// A source-specific ID used to distinguish between different sources originating from the same data connection. In general, a data connection is a single data source. However, for TrueLayer, `sourceId` is associated with a specific bank and has a many-to-one relationship with the `integrationId`.
	SourceID *string `json:"sourceId,omitempty"`
}

type GetIntegrationsPlatformKeyBrandingResponse struct {
	// OK
	Branding    *GetIntegrationsPlatformKeyBrandingBranding
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

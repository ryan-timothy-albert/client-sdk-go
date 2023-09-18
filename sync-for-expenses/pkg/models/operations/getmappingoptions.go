// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/shared"
	"net/http"
)

type GetMappingOptionsRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *GetMappingOptionsRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type GetMappingOptionsResponse struct {
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	// Success
	MappingOptions *shared.MappingOptions
	StatusCode     int
	RawResponse    *http.Response
}

func (o *GetMappingOptionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMappingOptionsResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetMappingOptionsResponse) GetMappingOptions() *shared.MappingOptions {
	if o == nil {
		return nil
	}
	return o.MappingOptions
}

func (o *GetMappingOptionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMappingOptionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

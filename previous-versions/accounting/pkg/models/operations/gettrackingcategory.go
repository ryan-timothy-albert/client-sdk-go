// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"net/http"
)

type GetTrackingCategoryRequest struct {
	CompanyID          string `pathParam:"style=simple,explode=false,name=companyId"`
	TrackingCategoryID string `pathParam:"style=simple,explode=false,name=trackingCategoryId"`
}

func (o *GetTrackingCategoryRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetTrackingCategoryRequest) GetTrackingCategoryID() string {
	if o == nil {
		return ""
	}
	return o.TrackingCategoryID
}

type GetTrackingCategoryResponse struct {
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
	// Success
	TrackingCategoryTree *shared.TrackingCategoryTree
}

func (o *GetTrackingCategoryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTrackingCategoryResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetTrackingCategoryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTrackingCategoryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTrackingCategoryResponse) GetTrackingCategoryTree() *shared.TrackingCategoryTree {
	if o == nil {
		return nil
	}
	return o.TrackingCategoryTree
}

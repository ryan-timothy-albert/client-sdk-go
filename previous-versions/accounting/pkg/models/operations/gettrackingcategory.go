// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type GetTrackingCategoryRequest struct {
	CompanyID          string `pathParam:"style=simple,explode=false,name=companyId"`
	TrackingCategoryID string `pathParam:"style=simple,explode=false,name=trackingCategoryId"`
}

// GetTrackingCategory409ApplicationJSON - The data type's dataset has not been requested or is still syncing.
type GetTrackingCategory409ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetTrackingCategoryResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Success
	TrackingCategoryTree *shared.TrackingCategoryTree
	// The data type's dataset has not been requested or is still syncing.
	GetTrackingCategory409ApplicationJSONObject *GetTrackingCategory409ApplicationJSON
	// Your API request was not properly authorized.
	Schema *shared.Schema
}
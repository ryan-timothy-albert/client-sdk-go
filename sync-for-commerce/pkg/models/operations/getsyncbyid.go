// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
	"net/http"
)

type GetSyncByIDRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a sync.
	SyncID string `pathParam:"style=simple,explode=false,name=syncId"`
}

func (o *GetSyncByIDRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetSyncByIDRequest) GetSyncID() string {
	if o == nil {
		return ""
	}
	return o.SyncID
}

type GetSyncByIDResponse struct {
	// Success
	CompanySyncStatus *shared.CompanySyncStatus
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetSyncByIDResponse) GetCompanySyncStatus() *shared.CompanySyncStatus {
	if o == nil {
		return nil
	}
	return o.CompanySyncStatus
}

func (o *GetSyncByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSyncByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSyncByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

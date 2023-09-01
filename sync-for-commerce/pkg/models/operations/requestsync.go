// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-commerce/pkg/models/shared"
	"net/http"
)

type RequestSyncRequest struct {
	SyncToLatestArgs *shared.SyncToLatestArgs `request:"mediaType=application/json"`
	CompanyID        string                   `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *RequestSyncRequest) GetSyncToLatestArgs() *shared.SyncToLatestArgs {
	if o == nil {
		return nil
	}
	return o.SyncToLatestArgs
}

func (o *RequestSyncRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type RequestSyncResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Success
	SyncSummary *shared.SyncSummary
}

func (o *RequestSyncResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RequestSyncResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RequestSyncResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RequestSyncResponse) GetSyncSummary() *shared.SyncSummary {
	if o == nil {
		return nil
	}
	return o.SyncSummary
}

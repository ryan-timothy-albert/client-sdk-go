// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/shared"
	"net/http"
)

type GetDataStatusRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *GetDataStatusRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type GetDataStatusResponse struct {
	ContentType string
	// OK
	DataStatusResponse map[string]shared.DataStatus
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetDataStatusResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDataStatusResponse) GetDataStatusResponse() map[string]shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.DataStatusResponse
}

func (o *GetDataStatusResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetDataStatusResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDataStatusResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

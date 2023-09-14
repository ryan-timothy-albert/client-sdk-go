// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v3/pkg/models/shared"
	"net/http"
)

type DownloadFilesRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Only download files uploaded on this date.
	Date *string `queryParam:"style=form,explode=true,name=date"`
}

func (o *DownloadFilesRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *DownloadFilesRequest) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

type DownloadFilesResponse struct {
	ContentType string
	// Success
	Data []byte
	// The request made is not valid.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *DownloadFilesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DownloadFilesResponse) GetData() []byte {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *DownloadFilesResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *DownloadFilesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DownloadFilesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

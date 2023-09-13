// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v2/pkg/models/shared"
	"net/http"
)

type ListFilesRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *ListFilesRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type ListFilesResponse struct {
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	// Success
	Files       []shared.File
	StatusCode  int
	RawResponse *http.Response
}

func (o *ListFilesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListFilesResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *ListFilesResponse) GetFiles() []shared.File {
	if o == nil {
		return nil
	}
	return o.Files
}

func (o *ListFilesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListFilesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

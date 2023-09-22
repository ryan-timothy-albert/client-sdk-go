// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
	"net/http"
)

type GetPullOperationRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique ID of a dataset or pull operation.
	DatasetID string `pathParam:"style=simple,explode=false,name=datasetId"`
}

func (o *GetPullOperationRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetPullOperationRequest) GetDatasetID() string {
	if o == nil {
		return ""
	}
	return o.DatasetID
}

type GetPullOperationResponse struct {
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	// OK
	PullOperation *shared.PullOperation
	StatusCode    int
	RawResponse   *http.Response
}

func (o *GetPullOperationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPullOperationResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetPullOperationResponse) GetPullOperation() *shared.PullOperation {
	if o == nil {
		return nil
	}
	return o.PullOperation
}

func (o *GetPullOperationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPullOperationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

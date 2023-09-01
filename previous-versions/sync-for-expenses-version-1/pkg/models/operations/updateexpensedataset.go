// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/shared"
	"net/http"
)

type UpdateExpenseDatasetRequest struct {
	UpdateExpenseRequest *shared.UpdateExpenseRequest `request:"mediaType=application/json"`
	CompanyID            string                       `pathParam:"style=simple,explode=false,name=companyId"`
	// The unique identifier for your SMB's transaction.
	TransactionID string `pathParam:"style=simple,explode=false,name=transactionId"`
}

func (o *UpdateExpenseDatasetRequest) GetUpdateExpenseRequest() *shared.UpdateExpenseRequest {
	if o == nil {
		return nil
	}
	return o.UpdateExpenseRequest
}

func (o *UpdateExpenseDatasetRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *UpdateExpenseDatasetRequest) GetTransactionID() string {
	if o == nil {
		return ""
	}
	return o.TransactionID
}

// UpdateExpenseDataset202ApplicationJSON - Accepted
type UpdateExpenseDataset202ApplicationJSON struct {
	SyncID *string `json:"syncId,omitempty"`
}

func (o *UpdateExpenseDataset202ApplicationJSON) GetSyncID() *string {
	if o == nil {
		return nil
	}
	return o.SyncID
}

type UpdateExpenseDatasetResponse struct {
	ContentType string
	// The request made is not valid.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
	// Accepted
	UpdateExpenseDataset202ApplicationJSONObject *UpdateExpenseDataset202ApplicationJSON
}

func (o *UpdateExpenseDatasetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateExpenseDatasetResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *UpdateExpenseDatasetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateExpenseDatasetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateExpenseDatasetResponse) GetUpdateExpenseDataset202ApplicationJSONObject() *UpdateExpenseDataset202ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateExpenseDataset202ApplicationJSONObject
}

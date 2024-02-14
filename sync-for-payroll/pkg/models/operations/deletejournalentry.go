// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/utils"
	"net/http"
)

type DeleteJournalEntryRequest struct {
	// Allow a sync upon push completion.
	AllowSyncOnPushComplete *bool `default:"true" queryParam:"style=form,explode=true,name=allowSyncOnPushComplete"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a journal entry.
	JournalEntryID string `pathParam:"style=simple,explode=false,name=journalEntryId"`
	// Time limit for the push operation to complete before it is timed out.
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

func (d DeleteJournalEntryRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DeleteJournalEntryRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DeleteJournalEntryRequest) GetAllowSyncOnPushComplete() *bool {
	if o == nil {
		return nil
	}
	return o.AllowSyncOnPushComplete
}

func (o *DeleteJournalEntryRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *DeleteJournalEntryRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *DeleteJournalEntryRequest) GetJournalEntryID() string {
	if o == nil {
		return ""
	}
	return o.JournalEntryID
}

func (o *DeleteJournalEntryRequest) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

type DeleteJournalEntryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	PushOperation *shared.PushOperation
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteJournalEntryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteJournalEntryResponse) GetPushOperation() *shared.PushOperation {
	if o == nil {
		return nil
	}
	return o.PushOperation
}

func (o *DeleteJournalEntryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteJournalEntryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

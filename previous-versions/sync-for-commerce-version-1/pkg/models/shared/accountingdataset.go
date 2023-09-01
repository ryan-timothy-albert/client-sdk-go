// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type AccountingDatasetStatus string

const (
	AccountingDatasetStatusInitial            AccountingDatasetStatus = "Initial"
	AccountingDatasetStatusQueued             AccountingDatasetStatus = "Queued"
	AccountingDatasetStatusFetching           AccountingDatasetStatus = "Fetching"
	AccountingDatasetStatusMapQueued          AccountingDatasetStatus = "MapQueued"
	AccountingDatasetStatusMapping            AccountingDatasetStatus = "Mapping"
	AccountingDatasetStatusComplete           AccountingDatasetStatus = "Complete"
	AccountingDatasetStatusFetchError         AccountingDatasetStatus = "FetchError"
	AccountingDatasetStatusMapError           AccountingDatasetStatus = "MapError"
	AccountingDatasetStatusInternalError      AccountingDatasetStatus = "InternalError"
	AccountingDatasetStatusProcessingQueued   AccountingDatasetStatus = "ProcessingQueued"
	AccountingDatasetStatusProcessing         AccountingDatasetStatus = "Processing"
	AccountingDatasetStatusProcessingError    AccountingDatasetStatus = "ProcessingError"
	AccountingDatasetStatusValidationQueued   AccountingDatasetStatus = "ValidationQueued"
	AccountingDatasetStatusValidating         AccountingDatasetStatus = "Validating"
	AccountingDatasetStatusValidationError    AccountingDatasetStatus = "ValidationError"
	AccountingDatasetStatusAuthError          AccountingDatasetStatus = "AuthError"
	AccountingDatasetStatusCancelled          AccountingDatasetStatus = "Cancelled"
	AccountingDatasetStatusNotSupported       AccountingDatasetStatus = "NotSupported"
	AccountingDatasetStatusRateLimitError     AccountingDatasetStatus = "RateLimitError"
	AccountingDatasetStatusPermissionsError   AccountingDatasetStatus = "PermissionsError"
	AccountingDatasetStatusPrerequisiteNotMet AccountingDatasetStatus = "PrerequisiteNotMet"
)

func (e AccountingDatasetStatus) ToPointer() *AccountingDatasetStatus {
	return &e
}

func (e *AccountingDatasetStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Initial":
		fallthrough
	case "Queued":
		fallthrough
	case "Fetching":
		fallthrough
	case "MapQueued":
		fallthrough
	case "Mapping":
		fallthrough
	case "Complete":
		fallthrough
	case "FetchError":
		fallthrough
	case "MapError":
		fallthrough
	case "InternalError":
		fallthrough
	case "ProcessingQueued":
		fallthrough
	case "Processing":
		fallthrough
	case "ProcessingError":
		fallthrough
	case "ValidationQueued":
		fallthrough
	case "Validating":
		fallthrough
	case "ValidationError":
		fallthrough
	case "AuthError":
		fallthrough
	case "Cancelled":
		fallthrough
	case "NotSupported":
		fallthrough
	case "RateLimitError":
		fallthrough
	case "PermissionsError":
		fallthrough
	case "PrerequisiteNotMet":
		*e = AccountingDatasetStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountingDatasetStatus: %v", v)
	}
}

// AccountingDataset - Success
type AccountingDataset struct {
	CompanyID string `json:"companyId"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	Completed      *string `json:"completed,omitempty"`
	ConnectionID   string  `json:"connectionId"`
	DataType       *string `json:"dataType,omitempty"`
	DatasetLogsURL *string `json:"datasetLogsUrl,omitempty"`
	ErrorMessage   *string `json:"errorMessage,omitempty"`
	ID             string  `json:"id"`
	IsCompleted    bool    `json:"isCompleted"`
	IsErrored      bool    `json:"isErrored"`
	Progress       int     `json:"progress"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	Requested                string                  `json:"requested"`
	Status                   AccountingDatasetStatus `json:"status"`
	ValidationInformationURL *string                 `json:"validationInformationUrl,omitempty"`
}

func (o *AccountingDataset) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *AccountingDataset) GetCompleted() *string {
	if o == nil {
		return nil
	}
	return o.Completed
}

func (o *AccountingDataset) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *AccountingDataset) GetDataType() *string {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *AccountingDataset) GetDatasetLogsURL() *string {
	if o == nil {
		return nil
	}
	return o.DatasetLogsURL
}

func (o *AccountingDataset) GetErrorMessage() *string {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *AccountingDataset) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AccountingDataset) GetIsCompleted() bool {
	if o == nil {
		return false
	}
	return o.IsCompleted
}

func (o *AccountingDataset) GetIsErrored() bool {
	if o == nil {
		return false
	}
	return o.IsErrored
}

func (o *AccountingDataset) GetProgress() int {
	if o == nil {
		return 0
	}
	return o.Progress
}

func (o *AccountingDataset) GetRequested() string {
	if o == nil {
		return ""
	}
	return o.Requested
}

func (o *AccountingDataset) GetStatus() AccountingDatasetStatus {
	if o == nil {
		return AccountingDatasetStatus("")
	}
	return o.Status
}

func (o *AccountingDataset) GetValidationInformationURL() *string {
	if o == nil {
		return nil
	}
	return o.ValidationInformationURL
}

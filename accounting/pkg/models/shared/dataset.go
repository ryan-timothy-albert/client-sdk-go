// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type DataSetStatusEnum string

const (
	DataSetStatusEnumInitial            DataSetStatusEnum = "Initial"
	DataSetStatusEnumQueued             DataSetStatusEnum = "Queued"
	DataSetStatusEnumFetching           DataSetStatusEnum = "Fetching"
	DataSetStatusEnumMapQueued          DataSetStatusEnum = "MapQueued"
	DataSetStatusEnumMapping            DataSetStatusEnum = "Mapping"
	DataSetStatusEnumComplete           DataSetStatusEnum = "Complete"
	DataSetStatusEnumFetchError         DataSetStatusEnum = "FetchError"
	DataSetStatusEnumMapError           DataSetStatusEnum = "MapError"
	DataSetStatusEnumInternalError      DataSetStatusEnum = "InternalError"
	DataSetStatusEnumProcessingQueued   DataSetStatusEnum = "ProcessingQueued"
	DataSetStatusEnumProcessing         DataSetStatusEnum = "Processing"
	DataSetStatusEnumProcessingError    DataSetStatusEnum = "ProcessingError"
	DataSetStatusEnumValidationQueued   DataSetStatusEnum = "ValidationQueued"
	DataSetStatusEnumValidating         DataSetStatusEnum = "Validating"
	DataSetStatusEnumValidationError    DataSetStatusEnum = "ValidationError"
	DataSetStatusEnumAuthError          DataSetStatusEnum = "AuthError"
	DataSetStatusEnumCancelled          DataSetStatusEnum = "Cancelled"
	DataSetStatusEnumNotSupported       DataSetStatusEnum = "NotSupported"
	DataSetStatusEnumRateLimitError     DataSetStatusEnum = "RateLimitError"
	DataSetStatusEnumPermissionsError   DataSetStatusEnum = "PermissionsError"
	DataSetStatusEnumPrerequisiteNotMet DataSetStatusEnum = "PrerequisiteNotMet"
)

func (e *DataSetStatusEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
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
		*e = DataSetStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DataSetStatusEnum: %s", s)
	}
}

// DataSet - Success
type DataSet struct {
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
	Requested                string            `json:"requested"`
	Status                   DataSetStatusEnum `json:"status"`
	ValidationInformationURL *string           `json:"validationInformationUrl,omitempty"`
}

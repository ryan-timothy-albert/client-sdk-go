// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// MakeRequestToDownloadExcelReportReportTypeEnum - The type of report you want to generate and download.
type MakeRequestToDownloadExcelReportReportTypeEnum string

const (
	MakeRequestToDownloadExcelReportReportTypeEnumAudit MakeRequestToDownloadExcelReportReportTypeEnum = "audit"
)

func (e *MakeRequestToDownloadExcelReportReportTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "audit":
		*e = MakeRequestToDownloadExcelReportReportTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for MakeRequestToDownloadExcelReportReportTypeEnum: %s", s)
	}
}

type MakeRequestToDownloadExcelReportRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// The type of report you want to generate and download.
	ReportType MakeRequestToDownloadExcelReportReportTypeEnum `queryParam:"style=form,explode=true,name=reportType"`
}

// MakeRequestToDownloadExcelReport200ApplicationJSON - OK
type MakeRequestToDownloadExcelReport200ApplicationJSON struct {
	ErrorMessage *string `json:"errorMessage,omitempty"`
	FileSize     *int64  `json:"fileSize,omitempty"`
	InProgress   *bool   `json:"inProgress,omitempty"`
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
	// > 📘 Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	LastGenerated    *string `json:"lastGenerated,omitempty"`
	LastInvocationID *string `json:"lastInvocationId,omitempty"`
	Queued           *string `json:"queued,omitempty"`
	ReportType       *string `json:"reportType,omitempty"`
	Success          *bool   `json:"success,omitempty"`
}

type MakeRequestToDownloadExcelReportResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	MakeRequestToDownloadExcelReport200ApplicationJSONObject *MakeRequestToDownloadExcelReport200ApplicationJSON
}

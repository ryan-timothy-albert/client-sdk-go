// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"time"
)

type GetCompanyRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

// GetCompany401ApplicationJSON - Your API request was not properly authorized.
type GetCompany401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetCompanyCompanyConnectionConnectionInfo struct {
	AdditionalProp1 *string `json:"additionalProp1,omitempty"`
	AdditionalProp2 *string `json:"additionalProp2,omitempty"`
	AdditionalProp3 *string `json:"additionalProp3,omitempty"`
}

type GetCompanyCompanyConnectionDataConnectionErrors struct {
	ErrorMessage *string `json:"errorMessage,omitempty"`
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
	ErroredOnUtc *time.Time `json:"erroredOnUtc,omitempty"`
	StatusCode   *string    `json:"statusCode,omitempty"`
	StatusText   *string    `json:"statusText,omitempty"`
}

// GetCompanyCompanyConnectionSourceTypeEnum - The type of platform of the connection.
type GetCompanyCompanyConnectionSourceTypeEnum string

const (
	GetCompanyCompanyConnectionSourceTypeEnumAccounting GetCompanyCompanyConnectionSourceTypeEnum = "Accounting"
	GetCompanyCompanyConnectionSourceTypeEnumBanking    GetCompanyCompanyConnectionSourceTypeEnum = "Banking"
	GetCompanyCompanyConnectionSourceTypeEnumCommerce   GetCompanyCompanyConnectionSourceTypeEnum = "Commerce"
	GetCompanyCompanyConnectionSourceTypeEnumOther      GetCompanyCompanyConnectionSourceTypeEnum = "Other"
	GetCompanyCompanyConnectionSourceTypeEnumUnknown    GetCompanyCompanyConnectionSourceTypeEnum = "Unknown"
)

// GetCompanyCompanyConnectionDataConnectionStatusEnum - The current authorization status of the data connection.
type GetCompanyCompanyConnectionDataConnectionStatusEnum string

const (
	GetCompanyCompanyConnectionDataConnectionStatusEnumPendingAuth  GetCompanyCompanyConnectionDataConnectionStatusEnum = "PendingAuth"
	GetCompanyCompanyConnectionDataConnectionStatusEnumLinked       GetCompanyCompanyConnectionDataConnectionStatusEnum = "Linked"
	GetCompanyCompanyConnectionDataConnectionStatusEnumUnlinked     GetCompanyCompanyConnectionDataConnectionStatusEnum = "Unlinked"
	GetCompanyCompanyConnectionDataConnectionStatusEnumDeauthorized GetCompanyCompanyConnectionDataConnectionStatusEnum = "Deauthorized"
)

// GetCompanyCompanyConnection - A connection represents the link between a `company` and a source of data.
type GetCompanyCompanyConnection struct {
	ConnectionInfo *GetCompanyCompanyConnectionConnectionInfo `json:"connectionInfo,omitempty"`
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
	Created              time.Time                                         `json:"created"`
	DataConnectionErrors []GetCompanyCompanyConnectionDataConnectionErrors `json:"dataConnectionErrors,omitempty"`
	// Unique identifier for a company's data connection.
	ID string `json:"id"`
	// A Codat ID representing the integration.
	IntegrationID string `json:"integrationId"`
	// A unique four-character ID that identifies the platform of the company's data connection. This ensures continuity if the platform changes its name in the future.
	IntegrationKey string `json:"integrationKey"`
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
	LastSync     *time.Time `json:"lastSync,omitempty"`
	LinkURL      string     `json:"linkUrl"`
	PlatformName string     `json:"platformName"`
	// A source-specific ID used to distinguish between different sources originating from the same data connection. In general, a data connection is a single data source. However, for TrueLayer, `sourceId` is associated with a specific bank and has a many-to-one relationship with the `integrationId`.
	SourceID string `json:"sourceId"`
	// The type of platform of the connection.
	SourceType GetCompanyCompanyConnectionSourceTypeEnum `json:"sourceType"`
	// The current authorization status of the data connection.
	Status GetCompanyCompanyConnectionDataConnectionStatusEnum `json:"status"`
}

// GetCompanyCompany - A company in Codat represent a small or medium sized business, whose data you wish to share
type GetCompanyCompany struct {
	Created           *time.Time                    `json:"created,omitempty"`
	CreatedByUserName *string                       `json:"createdByUserName,omitempty"`
	DataConnections   []GetCompanyCompanyConnection `json:"dataConnections,omitempty"`
	// Additional information about the company. This can be used to store foreign IDs, references, etc.
	Description *string `json:"description,omitempty"`
	// Unique identifier for your SMB in Codat.
	ID       string     `json:"id"`
	LastSync *time.Time `json:"lastSync,omitempty"`
	// The name of the company
	Name     string  `json:"name"`
	Platform *string `json:"platform,omitempty"`
	// The `redirect` [Link URL](https://docs.codat.io/auth-flow/authorize-hosted-link) enabling the customer to start their auth flow journey for the company.
	Redirect string `json:"redirect"`
}

type GetCompanyResponse struct {
	// OK
	Company     *GetCompanyCompany
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Your API request was not properly authorized.
	GetCompany401ApplicationJSONObject *GetCompany401ApplicationJSON
}

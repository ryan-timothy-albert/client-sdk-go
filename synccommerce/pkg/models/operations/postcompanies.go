// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"time"
)

type PostCompaniesRequestBody struct {
	// Name of the company in Codat with a partner-commerce data connection.
	Name string `json:"name"`
}

type PostCompanies200ApplicationJSONDataConnectionsDataConnectionErrors struct {
	// Message on the data connection error.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// In Codat's data model, dates and times are represented using the ISO 8601 standard. Date and time fields are formatted as strings.
	ErroredOnUtc *time.Time `json:"erroredOnUtc,omitempty"`
	// Codat's error status code for the connection error.
	StatusCode *string `json:"statusCode,omitempty"`
	// Descriptive text for the data connection error.
	StatusText *string `json:"statusText,omitempty"`
}

// PostCompanies200ApplicationJSONDataConnectionsSourceTypeEnum - The type of platform of the connection.
type PostCompanies200ApplicationJSONDataConnectionsSourceTypeEnum string

const (
	PostCompanies200ApplicationJSONDataConnectionsSourceTypeEnumAccounting PostCompanies200ApplicationJSONDataConnectionsSourceTypeEnum = "Accounting"
	PostCompanies200ApplicationJSONDataConnectionsSourceTypeEnumBanking    PostCompanies200ApplicationJSONDataConnectionsSourceTypeEnum = "Banking"
	PostCompanies200ApplicationJSONDataConnectionsSourceTypeEnumCommerce   PostCompanies200ApplicationJSONDataConnectionsSourceTypeEnum = "Commerce"
	PostCompanies200ApplicationJSONDataConnectionsSourceTypeEnumOther      PostCompanies200ApplicationJSONDataConnectionsSourceTypeEnum = "Other"
	PostCompanies200ApplicationJSONDataConnectionsSourceTypeEnumUnknown    PostCompanies200ApplicationJSONDataConnectionsSourceTypeEnum = "Unknown"
)

// PostCompanies200ApplicationJSONDataConnectionsStatusEnum - The current authorization status of the data connection.
type PostCompanies200ApplicationJSONDataConnectionsStatusEnum string

const (
	PostCompanies200ApplicationJSONDataConnectionsStatusEnumPendingAuth  PostCompanies200ApplicationJSONDataConnectionsStatusEnum = "PendingAuth"
	PostCompanies200ApplicationJSONDataConnectionsStatusEnumLinked       PostCompanies200ApplicationJSONDataConnectionsStatusEnum = "Linked"
	PostCompanies200ApplicationJSONDataConnectionsStatusEnumUnlinked     PostCompanies200ApplicationJSONDataConnectionsStatusEnum = "Unlinked"
	PostCompanies200ApplicationJSONDataConnectionsStatusEnumDeauthorized PostCompanies200ApplicationJSONDataConnectionsStatusEnum = "Deauthorized"
)

type PostCompanies200ApplicationJSONDataConnections struct {
	// In Codat's data model, dates and times are represented using the ISO 8601 standard. Date and time fields are formatted as string.
	Created *time.Time `json:"created,omitempty"`
	// Array containing errors on data connections.
	DataConnectionErrors []PostCompanies200ApplicationJSONDataConnectionsDataConnectionErrors `json:"dataConnectionErrors,omitempty"`
	// Unique identifier for a company's data connection.
	ID string `json:"id"`
	// A Codat ID representing the integration.
	IntegrationID string `json:"integrationId"`
	// In Codat's data model, dates and times are represented using the ISO 8601 standard. Date and time fields are formatted as string.
	LastSync *time.Time `json:"lastSync,omitempty"`
	// Whitelabelled link site URL for the user link flow.
	LinkURL string `json:"linkUrl"`
	// The name of the platform to which the data connection is established.
	PlatformName string `json:"platformName"`
	// A source-specific ID used to distinguish between different sources originating from the same data connection. In general, a data connection is a single data source. However, for TrueLayer, sourceId is associated with a specific bank and has a many-to-one relationship with the integrationId.
	SourceID string `json:"sourceId"`
	// The type of platform of the connection.
	SourceType *PostCompanies200ApplicationJSONDataConnectionsSourceTypeEnum `json:"sourceType,omitempty"`
	// The current authorization status of the data connection.
	Status *PostCompanies200ApplicationJSONDataConnectionsStatusEnum `json:"status,omitempty"`
}

// PostCompanies200ApplicationJSON - Success
type PostCompanies200ApplicationJSON struct {
	// The date the data connection was established.
	Created *time.Time `json:"created,omitempty"`
	// Name of the Codat user who created the data connection.
	CreatedByUserName *string                                          `json:"createdByUserName,omitempty"`
	DataConnections   []PostCompanies200ApplicationJSONDataConnections `json:"dataConnections"`
	// Unique identifier for your SMB in Codat.
	ID string `json:"id"`
	// The date time for the previous sync operation.
	LastSync *time.Time `json:"lastSync,omitempty"`
	// The name of the company.
	Name string `json:"name"`
	// Additional information about the company. This can be used to store foreign IDs, references, etc.
	Platform string `json:"platform"`
	// The redirect Link URL enabling the customer to start their auth flow journey for the company.
	Redirect string `json:"redirect"`
}

type PostCompaniesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Success
	PostCompanies200ApplicationJSONObject *PostCompanies200ApplicationJSON
}

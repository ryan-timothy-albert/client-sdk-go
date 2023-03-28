// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetCustomerAttachmentRequest struct {
	// Unique identifier for an attachment
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	CustomerID   string `pathParam:"style=simple,explode=false,name=customerId"`
}

type GetCustomerAttachmentAttachmentModifiedDate struct {
	// The date on which this record was last modified in Codat.
	ModifiedDate *string `json:"modifiedDate,omitempty"`
}

type GetCustomerAttachmentAttachmentSourceModifiedDate struct {
	// The date on which this record was last modified in the originating system
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
}

// GetCustomerAttachmentAttachment - Success
type GetCustomerAttachmentAttachment struct {
	ContentType *string `json:"contentType,omitempty"`
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
	DateCreated        *string                                            `json:"dateCreated,omitempty"`
	FileSize           *int                                               `json:"fileSize,omitempty"`
	ID                 *string                                            `json:"id,omitempty"`
	IncludeWhenSent    *bool                                              `json:"includeWhenSent,omitempty"`
	ModifiedDate       *GetCustomerAttachmentAttachmentModifiedDate       `json:"modifiedDate,omitempty"`
	Name               *string                                            `json:"name,omitempty"`
	SourceModifiedDate *GetCustomerAttachmentAttachmentSourceModifiedDate `json:"sourceModifiedDate,omitempty"`
}

type GetCustomerAttachmentResponse struct {
	// Success
	Attachment  *GetCustomerAttachmentAttachment
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

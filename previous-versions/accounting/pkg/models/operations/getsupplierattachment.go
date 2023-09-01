// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"net/http"
)

type GetSupplierAttachmentRequest struct {
	// Unique identifier for an attachment
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a supplier
	SupplierID string `pathParam:"style=simple,explode=false,name=supplierId"`
}

func (o *GetSupplierAttachmentRequest) GetAttachmentID() string {
	if o == nil {
		return ""
	}
	return o.AttachmentID
}

func (o *GetSupplierAttachmentRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetSupplierAttachmentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetSupplierAttachmentRequest) GetSupplierID() string {
	if o == nil {
		return ""
	}
	return o.SupplierID
}

type GetSupplierAttachmentResponse struct {
	// Success
	Attachment  *shared.Attachment
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetSupplierAttachmentResponse) GetAttachment() *shared.Attachment {
	if o == nil {
		return nil
	}
	return o.Attachment
}

func (o *GetSupplierAttachmentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSupplierAttachmentResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetSupplierAttachmentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSupplierAttachmentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

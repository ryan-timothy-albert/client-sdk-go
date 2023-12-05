// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"net/http"
)

type GetBillAttachmentRequest struct {
	// Unique identifier for an attachment.
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	// Unique identifier for a bill.
	BillID string `pathParam:"style=simple,explode=false,name=billId"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

func (o *GetBillAttachmentRequest) GetAttachmentID() string {
	if o == nil {
		return ""
	}
	return o.AttachmentID
}

func (o *GetBillAttachmentRequest) GetBillID() string {
	if o == nil {
		return ""
	}
	return o.BillID
}

func (o *GetBillAttachmentRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetBillAttachmentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type GetBillAttachmentResponse struct {
	// Success
	Attachment *shared.Attachment
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetBillAttachmentResponse) GetAttachment() *shared.Attachment {
	if o == nil {
		return nil
	}
	return o.Attachment
}

func (o *GetBillAttachmentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetBillAttachmentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetBillAttachmentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

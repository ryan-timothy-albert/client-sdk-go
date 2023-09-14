// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v3/pkg/models/shared"
	"net/http"
)

type GetAccountingCustomerAttachmentRequest struct {
	// Unique identifier for an attachment
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	CustomerID   string `pathParam:"style=simple,explode=false,name=customerId"`
}

func (o *GetAccountingCustomerAttachmentRequest) GetAttachmentID() string {
	if o == nil {
		return ""
	}
	return o.AttachmentID
}

func (o *GetAccountingCustomerAttachmentRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetAccountingCustomerAttachmentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetAccountingCustomerAttachmentRequest) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

type GetAccountingCustomerAttachmentResponse struct {
	// Success
	AccountingAttachment *shared.AccountingAttachment
	ContentType          string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetAccountingCustomerAttachmentResponse) GetAccountingAttachment() *shared.AccountingAttachment {
	if o == nil {
		return nil
	}
	return o.AccountingAttachment
}

func (o *GetAccountingCustomerAttachmentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountingCustomerAttachmentResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetAccountingCustomerAttachmentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountingCustomerAttachmentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

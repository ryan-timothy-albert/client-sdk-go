// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/shared"
	"net/http"
)

type UploadExpenseAttachmentRequestBody struct {
	Content     []byte `multipartForm:"content"`
	RequestBody string `multipartForm:"name=requestBody"`
}

func (o *UploadExpenseAttachmentRequestBody) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

func (o *UploadExpenseAttachmentRequestBody) GetRequestBody() string {
	if o == nil {
		return ""
	}
	return o.RequestBody
}

type UploadExpenseAttachmentRequest struct {
	RequestBody *UploadExpenseAttachmentRequestBody `multipartForm:"file" request:"mediaType=multipart/form-data"`
	CompanyID   string                              `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a sync.
	SyncID string `pathParam:"style=simple,explode=false,name=syncId"`
	// The unique identifier for your SMB's transaction.
	TransactionID string `pathParam:"style=simple,explode=false,name=transactionId"`
}

func (o *UploadExpenseAttachmentRequest) GetRequestBody() *UploadExpenseAttachmentRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *UploadExpenseAttachmentRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *UploadExpenseAttachmentRequest) GetSyncID() string {
	if o == nil {
		return ""
	}
	return o.SyncID
}

func (o *UploadExpenseAttachmentRequest) GetTransactionID() string {
	if o == nil {
		return ""
	}
	return o.TransactionID
}

type UploadExpenseAttachmentResponse struct {
	// OK
	Attachment  *shared.Attachment
	ContentType string
	// The request made is not valid.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *UploadExpenseAttachmentResponse) GetAttachment() *shared.Attachment {
	if o == nil {
		return nil
	}
	return o.Attachment
}

func (o *UploadExpenseAttachmentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UploadExpenseAttachmentResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *UploadExpenseAttachmentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UploadExpenseAttachmentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"net/http"
)

type UploadDirectIncomeAttachmentRequestBody struct {
	Content     []byte `multipartForm:"content"`
	RequestBody string `multipartForm:"name=requestBody"`
}

func (o *UploadDirectIncomeAttachmentRequestBody) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

func (o *UploadDirectIncomeAttachmentRequestBody) GetRequestBody() string {
	if o == nil {
		return ""
	}
	return o.RequestBody
}

type UploadDirectIncomeAttachmentRequest struct {
	RequestBody    *UploadDirectIncomeAttachmentRequestBody `multipartForm:"file" request:"mediaType=multipart/form-data"`
	CompanyID      string                                   `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID   string                                   `pathParam:"style=simple,explode=false,name=connectionId"`
	DirectIncomeID string                                   `pathParam:"style=simple,explode=false,name=directIncomeId"`
}

func (o *UploadDirectIncomeAttachmentRequest) GetRequestBody() *UploadDirectIncomeAttachmentRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *UploadDirectIncomeAttachmentRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *UploadDirectIncomeAttachmentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UploadDirectIncomeAttachmentRequest) GetDirectIncomeID() string {
	if o == nil {
		return ""
	}
	return o.DirectIncomeID
}

type UploadDirectIncomeAttachmentResponse struct {
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *UploadDirectIncomeAttachmentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UploadDirectIncomeAttachmentResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *UploadDirectIncomeAttachmentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UploadDirectIncomeAttachmentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

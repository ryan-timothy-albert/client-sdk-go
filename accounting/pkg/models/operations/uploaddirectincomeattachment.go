// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type UploadDirectIncomeAttachmentRequestBody struct {
	Content     []byte `multipartForm:"content"`
	RequestBody string `multipartForm:"name=requestBody"`
}

type UploadDirectIncomeAttachmentRequest struct {
	RequestBody    *UploadDirectIncomeAttachmentRequestBody `multipartForm:"file" request:"mediaType=multipart/form-data"`
	CompanyID      string                                   `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID   string                                   `pathParam:"style=simple,explode=false,name=connectionId"`
	DirectIncomeID string                                   `pathParam:"style=simple,explode=false,name=directIncomeId"`
}

type UploadDirectIncomeAttachmentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Your API request was not properly authorized.
	Schema *shared.Schema
}

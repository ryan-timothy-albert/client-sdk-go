// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type UploadDirectCostAttachmentRequestBody struct {
	Content     []byte `multipartForm:"content"`
	RequestBody string `multipartForm:"name=requestBody"`
}

type UploadDirectCostAttachmentRequest struct {
	RequestBody  *UploadDirectCostAttachmentRequestBody `multipartForm:"file" request:"mediaType=multipart/form-data"`
	CompanyID    string                                 `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string                                 `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a direct cost
	DirectCostID string `pathParam:"style=simple,explode=false,name=directCostId"`
}

type UploadDirectCostAttachmentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/shared"
	"net/http"
)

type UpdateVisibleAccountsSyncFlowRequest struct {
	VisibleAccounts *shared.VisibleAccounts `request:"mediaType=application/json"`
	CommerceKey     string                  `pathParam:"style=simple,explode=false,name=commerceKey"`
}

type UpdateVisibleAccountsSyncFlowResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Success
	VisibleAccounts *shared.VisibleAccounts
}

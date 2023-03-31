// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type ListDirectCostAttachmentsRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a direct cost
	DirectCostID string `pathParam:"style=simple,explode=false,name=directCostId"`
}

type ListDirectCostAttachmentsResponse struct {
	// Success
	AttachmentsDataset *shared.AttachmentsDataset
	ContentType        string
	StatusCode         int
	RawResponse        *http.Response
}

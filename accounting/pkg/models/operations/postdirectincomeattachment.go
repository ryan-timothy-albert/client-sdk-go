package operations

import (
	"net/http"
)

type PostDirectIncomeAttachmentRequest struct {
	CompanyID      string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID   string `pathParam:"style=simple,explode=false,name=connectionId"`
	DirectIncomeID string `pathParam:"style=simple,explode=false,name=directIncomeId"`
}

type PostDirectIncomeAttachmentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

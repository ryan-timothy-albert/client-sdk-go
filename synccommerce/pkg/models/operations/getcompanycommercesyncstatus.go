package operations

import (
	"net/http"
)

type GetCompanyCommerceSyncStatusRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCompanyCommerceSyncStatusResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

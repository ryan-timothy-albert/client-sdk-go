// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type CreateDirectIncomeRequest struct {
	DirectIncome     *shared.DirectIncome `request:"mediaType=application/json"`
	CompanyID        string               `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID     string               `pathParam:"style=simple,explode=false,name=connectionId"`
	TimeoutInMinutes *int                 `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type CreateDirectIncomeResponse struct {
	ContentType string
	// Success
	CreateDirectIncomeResponse *shared.CreateDirectIncomeResponse
	StatusCode                 int
	RawResponse                *http.Response
}

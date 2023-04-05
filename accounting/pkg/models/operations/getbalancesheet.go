// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type GetBalanceSheetRequest struct {
	CompanyID        string  `pathParam:"style=simple,explode=false,name=companyId"`
	PeriodLength     int     `queryParam:"style=form,explode=true,name=periodLength"`
	PeriodsToCompare int     `queryParam:"style=form,explode=true,name=periodsToCompare"`
	StartMonth       *string `queryParam:"style=form,explode=true,name=startMonth"`
}

type GetBalanceSheetResponse struct {
	// Success
	BalanceSheet *shared.BalanceSheet1
	ContentType  string
	StatusCode   int
	RawResponse  *http.Response
}

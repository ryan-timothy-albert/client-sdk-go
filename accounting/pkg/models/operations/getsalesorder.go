// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type GetSalesOrderRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	SalesOrderID string `pathParam:"style=simple,explode=false,name=salesOrderId"`
}

type GetSalesOrderResponse struct {
	ContentType string
	// Success
	SalesOrder  *shared.SalesOrder
	StatusCode  int
	RawResponse *http.Response
}

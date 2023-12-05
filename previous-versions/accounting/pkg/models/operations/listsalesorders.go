// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/utils"
	"net/http"
)

type ListSalesOrdersRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Field to order results by. [Read more](https://docs.codat.io/using-the-api/ordering-results).
	OrderBy *string `queryParam:"style=form,explode=true,name=orderBy"`
	// Page number. [Read more](https://docs.codat.io/using-the-api/paging).
	Page *int `default:"1" queryParam:"style=form,explode=true,name=page"`
	// Number of records to return in a page. [Read more](https://docs.codat.io/using-the-api/paging).
	PageSize *int `default:"100" queryParam:"style=form,explode=true,name=pageSize"`
	// Codat query string. [Read more](https://docs.codat.io/using-the-api/querying).
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

func (l ListSalesOrdersRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListSalesOrdersRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListSalesOrdersRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *ListSalesOrdersRequest) GetOrderBy() *string {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *ListSalesOrdersRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListSalesOrdersRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListSalesOrdersRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

type ListSalesOrdersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	SalesOrders *shared.SalesOrders
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListSalesOrdersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListSalesOrdersResponse) GetSalesOrders() *shared.SalesOrders {
	if o == nil {
		return nil
	}
	return o.SalesOrders
}

func (o *ListSalesOrdersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListSalesOrdersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

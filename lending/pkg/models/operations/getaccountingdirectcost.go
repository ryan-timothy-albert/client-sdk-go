// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v4/pkg/models/shared"
	"net/http"
)

type GetAccountingDirectCostRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a direct cost
	DirectCostID string `pathParam:"style=simple,explode=false,name=directCostId"`
}

func (o *GetAccountingDirectCostRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetAccountingDirectCostRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetAccountingDirectCostRequest) GetDirectCostID() string {
	if o == nil {
		return ""
	}
	return o.DirectCostID
}

type GetAccountingDirectCostResponse struct {
	// Success
	AccountingDirectCost *shared.AccountingDirectCost
	ContentType          string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetAccountingDirectCostResponse) GetAccountingDirectCost() *shared.AccountingDirectCost {
	if o == nil {
		return nil
	}
	return o.AccountingDirectCost
}

func (o *GetAccountingDirectCostResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountingDirectCostResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetAccountingDirectCostResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountingDirectCostResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

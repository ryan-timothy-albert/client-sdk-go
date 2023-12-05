// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// QueryParamSourceType - Data source type.
type QueryParamSourceType string

const (
	QueryParamSourceTypeBanking    QueryParamSourceType = "banking"
	QueryParamSourceTypeCommerce   QueryParamSourceType = "commerce"
	QueryParamSourceTypeAccounting QueryParamSourceType = "accounting"
)

func (e QueryParamSourceType) ToPointer() *QueryParamSourceType {
	return &e
}

func (e *QueryParamSourceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "banking":
		fallthrough
	case "commerce":
		fallthrough
	case "accounting":
		*e = QueryParamSourceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamSourceType: %v", v)
	}
}

type GenerateLoanTransactionsRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Data source type.
	SourceType QueryParamSourceType `queryParam:"style=form,explode=true,name=sourceType"`
}

func (o *GenerateLoanTransactionsRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GenerateLoanTransactionsRequest) GetSourceType() QueryParamSourceType {
	if o == nil {
		return QueryParamSourceType("")
	}
	return o.SourceType
}

type GenerateLoanTransactionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GenerateLoanTransactionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GenerateLoanTransactionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GenerateLoanTransactionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

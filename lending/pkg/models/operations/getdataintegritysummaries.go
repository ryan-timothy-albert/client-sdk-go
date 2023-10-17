// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v4/pkg/models/shared"
	"net/http"
)

type GetDataIntegritySummariesRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// A key for a Codat data type.
	DataType shared.DataIntegrityDataType `pathParam:"style=simple,explode=false,name=dataType"`
	// Codat query string. [Read more](https://docs.codat.io/using-the-api/querying).
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

func (o *GetDataIntegritySummariesRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetDataIntegritySummariesRequest) GetDataType() shared.DataIntegrityDataType {
	if o == nil {
		return shared.DataIntegrityDataType("")
	}
	return o.DataType
}

func (o *GetDataIntegritySummariesRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

type GetDataIntegritySummariesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	DataIntegritySummaries *shared.DataIntegritySummaries
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetDataIntegritySummariesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDataIntegritySummariesResponse) GetDataIntegritySummaries() *shared.DataIntegritySummaries {
	if o == nil {
		return nil
	}
	return o.DataIntegritySummaries
}

func (o *GetDataIntegritySummariesResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetDataIntegritySummariesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDataIntegritySummariesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

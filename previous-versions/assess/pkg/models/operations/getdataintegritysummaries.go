// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
	"net/http"
)

type GetDataIntegritySummariesRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// A key for a Codat data type.
	DataType shared.DataIntegrityDataType `pathParam:"style=simple,explode=false,name=dataType"`
	// Codat query string. [Read more](https://docs.codat.io/using-the-api/querying).
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

type GetDataIntegritySummariesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	Summaries *shared.Summaries
	// Your API request was not properly authorized.
	Schema *shared.Schema
}
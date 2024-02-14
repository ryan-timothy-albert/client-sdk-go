// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/shared"
	"net/http"
)

type UpdateCompanyRequest struct {
	CompanyRequestBody *shared.CompanyRequestBody `request:"mediaType=application/json"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *UpdateCompanyRequest) GetCompanyRequestBody() *shared.CompanyRequestBody {
	if o == nil {
		return nil
	}
	return o.CompanyRequestBody
}

func (o *UpdateCompanyRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type UpdateCompanyResponse struct {
	// OK
	Company *shared.Company
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateCompanyResponse) GetCompany() *shared.Company {
	if o == nil {
		return nil
	}
	return o.Company
}

func (o *UpdateCompanyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateCompanyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateCompanyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

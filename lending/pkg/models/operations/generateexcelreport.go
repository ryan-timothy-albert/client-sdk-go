// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/pkg/models/shared"
	"net/http"
)

type GenerateExcelReportRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// The type of report you want to generate and download.
	ReportType shared.ExcelReportTypes `queryParam:"style=form,explode=true,name=reportType"`
}

func (o *GenerateExcelReportRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GenerateExcelReportRequest) GetReportType() shared.ExcelReportTypes {
	if o == nil {
		return shared.ExcelReportTypes("")
	}
	return o.ReportType
}

type GenerateExcelReportResponse struct {
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	// OK
	ExcelStatus *shared.ExcelStatus
	StatusCode  int
	RawResponse *http.Response
}

func (o *GenerateExcelReportResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GenerateExcelReportResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GenerateExcelReportResponse) GetExcelStatus() *shared.ExcelStatus {
	if o == nil {
		return nil
	}
	return o.ExcelStatus
}

func (o *GenerateExcelReportResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GenerateExcelReportResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

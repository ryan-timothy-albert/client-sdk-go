// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type CreatePaymentRequest struct {
	Payment          *shared.Payment `request:"mediaType=application/json"`
	CompanyID        string          `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID     string          `pathParam:"style=simple,explode=false,name=connectionId"`
	TimeoutInMinutes *int            `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type CreatePaymentResponse struct {
	ContentType string
	// Success
	CreatePaymentResponse *shared.CreatePaymentResponse
	StatusCode            int
	RawResponse           *http.Response
}

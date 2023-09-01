// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type GetPaymentRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	PaymentID string `pathParam:"style=simple,explode=false,name=paymentId"`
}

// GetPayment409ApplicationJSON - The data type's dataset has not been requested or is still syncing.
type GetPayment409ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetPaymentResponse struct {
	ContentType string
	// Success
	Payment     *shared.Payment
	StatusCode  int
	RawResponse *http.Response
	// The data type's dataset has not been requested or is still syncing.
	GetPayment409ApplicationJSONObject *GetPayment409ApplicationJSON
	// Your API request was not properly authorized.
	Schema *shared.Schema
}
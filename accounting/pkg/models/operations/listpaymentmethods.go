// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ListPaymentMethodsRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Field to order results by. [Read more](https://docs.codat.io/using-the-api/ordering-results).
	OrderBy *string `queryParam:"style=form,explode=true,name=orderBy"`
	// Page number. [Read more](https://docs.codat.io/using-the-api/paging).
	Page int `queryParam:"style=form,explode=true,name=page"`
	// Number of records to return in a page. [Read more](https://docs.codat.io/using-the-api/paging).
	PageSize *int `queryParam:"style=form,explode=true,name=pageSize"`
	// Codat query string. [Read more](https://docs.codat.io/using-the-api/querying).
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

type ListPaymentMethods200ApplicationJSONLinksHypertextReference struct {
	Href *string `json:"href,omitempty"`
}

type ListPaymentMethods200ApplicationJSONLinks struct {
	Current  ListPaymentMethods200ApplicationJSONLinksHypertextReference  `json:"current"`
	Next     *ListPaymentMethods200ApplicationJSONLinksHypertextReference `json:"next,omitempty"`
	Previous *ListPaymentMethods200ApplicationJSONLinksHypertextReference `json:"previous,omitempty"`
	Self     ListPaymentMethods200ApplicationJSONLinksHypertextReference  `json:"self"`
}

type ListPaymentMethods200ApplicationJSONSourceModifiedDateMetadata struct {
	// Indicates whether the record has been deleted in the third-party system this record originated from.
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// ListPaymentMethods200ApplicationJSONSourceModifiedDateStatusEnum - Status of the Payment Method.
type ListPaymentMethods200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	ListPaymentMethods200ApplicationJSONSourceModifiedDateStatusEnumUnknown  ListPaymentMethods200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	ListPaymentMethods200ApplicationJSONSourceModifiedDateStatusEnumActive   ListPaymentMethods200ApplicationJSONSourceModifiedDateStatusEnum = "Active"
	ListPaymentMethods200ApplicationJSONSourceModifiedDateStatusEnumArchived ListPaymentMethods200ApplicationJSONSourceModifiedDateStatusEnum = "Archived"
)

func (e *ListPaymentMethods200ApplicationJSONSourceModifiedDateStatusEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Unknown":
		fallthrough
	case "Active":
		fallthrough
	case "Archived":
		*e = ListPaymentMethods200ApplicationJSONSourceModifiedDateStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for ListPaymentMethods200ApplicationJSONSourceModifiedDateStatusEnum: %s", s)
	}
}

// ListPaymentMethods200ApplicationJSONSourceModifiedDateTypeEnum - Method of payment.
type ListPaymentMethods200ApplicationJSONSourceModifiedDateTypeEnum string

const (
	ListPaymentMethods200ApplicationJSONSourceModifiedDateTypeEnumUnknown      ListPaymentMethods200ApplicationJSONSourceModifiedDateTypeEnum = "Unknown"
	ListPaymentMethods200ApplicationJSONSourceModifiedDateTypeEnumCash         ListPaymentMethods200ApplicationJSONSourceModifiedDateTypeEnum = "Cash"
	ListPaymentMethods200ApplicationJSONSourceModifiedDateTypeEnumCheck        ListPaymentMethods200ApplicationJSONSourceModifiedDateTypeEnum = "Check"
	ListPaymentMethods200ApplicationJSONSourceModifiedDateTypeEnumCreditCard   ListPaymentMethods200ApplicationJSONSourceModifiedDateTypeEnum = "CreditCard"
	ListPaymentMethods200ApplicationJSONSourceModifiedDateTypeEnumDebitCard    ListPaymentMethods200ApplicationJSONSourceModifiedDateTypeEnum = "DebitCard"
	ListPaymentMethods200ApplicationJSONSourceModifiedDateTypeEnumBankTransfer ListPaymentMethods200ApplicationJSONSourceModifiedDateTypeEnum = "BankTransfer"
	ListPaymentMethods200ApplicationJSONSourceModifiedDateTypeEnumOther        ListPaymentMethods200ApplicationJSONSourceModifiedDateTypeEnum = "Other"
)

func (e *ListPaymentMethods200ApplicationJSONSourceModifiedDateTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Unknown":
		fallthrough
	case "Cash":
		fallthrough
	case "Check":
		fallthrough
	case "CreditCard":
		fallthrough
	case "DebitCard":
		fallthrough
	case "BankTransfer":
		fallthrough
	case "Other":
		*e = ListPaymentMethods200ApplicationJSONSourceModifiedDateTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for ListPaymentMethods200ApplicationJSONSourceModifiedDateTypeEnum: %s", s)
	}
}

// ListPaymentMethods200ApplicationJSONSourceModifiedDate - > View the coverage for payment methods in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=paymentMethods" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A Payment Method represents the payment method(s) used to pay a Bill. Payment Methods are referenced on [Bill Payments](https://docs.codat.io/accounting-api#/schemas/BillPayment) and [Payments](https://docs.codat.io/accounting-api#/schemas/Payment).
//
// From the Payment Methods endpoints you can retrieve:
//
//   - A list of all the Payment Methods used by a company: `GET/companies/{companyId}/data/paymentMethods`.
//   - The details of an individual Payment Method:
//     `GET /companies/{companyId}/data/paymentMethods/{paymentMethodId}`.
type ListPaymentMethods200ApplicationJSONSourceModifiedDate struct {
	// Unique identifier for the payment method.
	ID       *string                                                         `json:"id,omitempty"`
	Metadata *ListPaymentMethods200ApplicationJSONSourceModifiedDateMetadata `json:"metadata,omitempty"`
	// The date on which this record was last modified in Codat.
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// Name of the payment method.
	Name *string `json:"name,omitempty"`
	// The date on which this record was last modified in the originating system
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Status of the Payment Method.
	Status *ListPaymentMethods200ApplicationJSONSourceModifiedDateStatusEnum `json:"status,omitempty"`
	// Method of payment.
	Type *ListPaymentMethods200ApplicationJSONSourceModifiedDateTypeEnum `json:"type,omitempty"`
}

// ListPaymentMethods200ApplicationJSON - Success
type ListPaymentMethods200ApplicationJSON struct {
	Links        ListPaymentMethods200ApplicationJSONLinks                `json:"_links"`
	PageNumber   int64                                                    `json:"pageNumber"`
	PageSize     int64                                                    `json:"pageSize"`
	Results      []ListPaymentMethods200ApplicationJSONSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                                    `json:"totalResults"`
}

type ListPaymentMethodsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Success
	ListPaymentMethods200ApplicationJSONObject *ListPaymentMethods200ApplicationJSON
}

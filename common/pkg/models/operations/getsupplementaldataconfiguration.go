// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
	"net/http"
)

// GetSupplementalDataConfigurationDataType - Data types that support supplemental data
type GetSupplementalDataConfigurationDataType string

const (
	GetSupplementalDataConfigurationDataTypeChartOfAccounts           GetSupplementalDataConfigurationDataType = "chartOfAccounts"
	GetSupplementalDataConfigurationDataTypeBills                     GetSupplementalDataConfigurationDataType = "bills"
	GetSupplementalDataConfigurationDataTypeCompany                   GetSupplementalDataConfigurationDataType = "company"
	GetSupplementalDataConfigurationDataTypeCreditNotes               GetSupplementalDataConfigurationDataType = "creditNotes"
	GetSupplementalDataConfigurationDataTypeCustomers                 GetSupplementalDataConfigurationDataType = "customers"
	GetSupplementalDataConfigurationDataTypeInvoices                  GetSupplementalDataConfigurationDataType = "invoices"
	GetSupplementalDataConfigurationDataTypeItems                     GetSupplementalDataConfigurationDataType = "items"
	GetSupplementalDataConfigurationDataTypeJournalEntries            GetSupplementalDataConfigurationDataType = "journalEntries"
	GetSupplementalDataConfigurationDataTypeSuppliers                 GetSupplementalDataConfigurationDataType = "suppliers"
	GetSupplementalDataConfigurationDataTypeTaxRates                  GetSupplementalDataConfigurationDataType = "taxRates"
	GetSupplementalDataConfigurationDataTypeCommerceCompanyInfo       GetSupplementalDataConfigurationDataType = "commerce-companyInfo"
	GetSupplementalDataConfigurationDataTypeCommerceCustomers         GetSupplementalDataConfigurationDataType = "commerce-customers"
	GetSupplementalDataConfigurationDataTypeCommerceDisputes          GetSupplementalDataConfigurationDataType = "commerce-disputes"
	GetSupplementalDataConfigurationDataTypeCommerceLocations         GetSupplementalDataConfigurationDataType = "commerce-locations"
	GetSupplementalDataConfigurationDataTypeCommerceOrders            GetSupplementalDataConfigurationDataType = "commerce-orders"
	GetSupplementalDataConfigurationDataTypeCommercePayments          GetSupplementalDataConfigurationDataType = "commerce-payments"
	GetSupplementalDataConfigurationDataTypeCommercePaymentMethods    GetSupplementalDataConfigurationDataType = "commerce-paymentMethods"
	GetSupplementalDataConfigurationDataTypeCommerceProducts          GetSupplementalDataConfigurationDataType = "commerce-products"
	GetSupplementalDataConfigurationDataTypeCommerceProductCategories GetSupplementalDataConfigurationDataType = "commerce-productCategories"
	GetSupplementalDataConfigurationDataTypeCommerceTaxComponents     GetSupplementalDataConfigurationDataType = "commerce-taxComponents"
	GetSupplementalDataConfigurationDataTypeCommerceTransactions      GetSupplementalDataConfigurationDataType = "commerce-transactions"
)

func (e GetSupplementalDataConfigurationDataType) ToPointer() *GetSupplementalDataConfigurationDataType {
	return &e
}

func (e *GetSupplementalDataConfigurationDataType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "chartOfAccounts":
		fallthrough
	case "bills":
		fallthrough
	case "company":
		fallthrough
	case "creditNotes":
		fallthrough
	case "customers":
		fallthrough
	case "invoices":
		fallthrough
	case "items":
		fallthrough
	case "journalEntries":
		fallthrough
	case "suppliers":
		fallthrough
	case "taxRates":
		fallthrough
	case "commerce-companyInfo":
		fallthrough
	case "commerce-customers":
		fallthrough
	case "commerce-disputes":
		fallthrough
	case "commerce-locations":
		fallthrough
	case "commerce-orders":
		fallthrough
	case "commerce-payments":
		fallthrough
	case "commerce-paymentMethods":
		fallthrough
	case "commerce-products":
		fallthrough
	case "commerce-productCategories":
		fallthrough
	case "commerce-taxComponents":
		fallthrough
	case "commerce-transactions":
		*e = GetSupplementalDataConfigurationDataType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSupplementalDataConfigurationDataType: %v", v)
	}
}

type GetSupplementalDataConfigurationRequest struct {
	// Data types that support supplemental data
	DataType    GetSupplementalDataConfigurationDataType `pathParam:"style=simple,explode=false,name=dataType"`
	PlatformKey string                                   `pathParam:"style=simple,explode=false,name=platformKey"`
}

type GetSupplementalDataConfigurationResponse struct {
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
	// OK
	SupplementalDataConfiguration map[string]shared.SupplementalDataConfiguration
}

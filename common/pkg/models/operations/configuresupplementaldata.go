// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
	"net/http"
)

// ConfigureSupplementalDataDataType - Data types that support supplemental data
type ConfigureSupplementalDataDataType string

const (
	ConfigureSupplementalDataDataTypeChartOfAccounts           ConfigureSupplementalDataDataType = "chartOfAccounts"
	ConfigureSupplementalDataDataTypeBills                     ConfigureSupplementalDataDataType = "bills"
	ConfigureSupplementalDataDataTypeCompany                   ConfigureSupplementalDataDataType = "company"
	ConfigureSupplementalDataDataTypeCreditNotes               ConfigureSupplementalDataDataType = "creditNotes"
	ConfigureSupplementalDataDataTypeCustomers                 ConfigureSupplementalDataDataType = "customers"
	ConfigureSupplementalDataDataTypeInvoices                  ConfigureSupplementalDataDataType = "invoices"
	ConfigureSupplementalDataDataTypeItems                     ConfigureSupplementalDataDataType = "items"
	ConfigureSupplementalDataDataTypeJournalEntries            ConfigureSupplementalDataDataType = "journalEntries"
	ConfigureSupplementalDataDataTypeSuppliers                 ConfigureSupplementalDataDataType = "suppliers"
	ConfigureSupplementalDataDataTypeTaxRates                  ConfigureSupplementalDataDataType = "taxRates"
	ConfigureSupplementalDataDataTypeCommerceCompanyInfo       ConfigureSupplementalDataDataType = "commerce-companyInfo"
	ConfigureSupplementalDataDataTypeCommerceCustomers         ConfigureSupplementalDataDataType = "commerce-customers"
	ConfigureSupplementalDataDataTypeCommerceDisputes          ConfigureSupplementalDataDataType = "commerce-disputes"
	ConfigureSupplementalDataDataTypeCommerceLocations         ConfigureSupplementalDataDataType = "commerce-locations"
	ConfigureSupplementalDataDataTypeCommerceOrders            ConfigureSupplementalDataDataType = "commerce-orders"
	ConfigureSupplementalDataDataTypeCommercePayments          ConfigureSupplementalDataDataType = "commerce-payments"
	ConfigureSupplementalDataDataTypeCommercePaymentMethods    ConfigureSupplementalDataDataType = "commerce-paymentMethods"
	ConfigureSupplementalDataDataTypeCommerceProducts          ConfigureSupplementalDataDataType = "commerce-products"
	ConfigureSupplementalDataDataTypeCommerceProductCategories ConfigureSupplementalDataDataType = "commerce-productCategories"
	ConfigureSupplementalDataDataTypeCommerceTaxComponents     ConfigureSupplementalDataDataType = "commerce-taxComponents"
	ConfigureSupplementalDataDataTypeCommerceTransactions      ConfigureSupplementalDataDataType = "commerce-transactions"
)

func (e ConfigureSupplementalDataDataType) ToPointer() *ConfigureSupplementalDataDataType {
	return &e
}

func (e *ConfigureSupplementalDataDataType) UnmarshalJSON(data []byte) error {
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
		*e = ConfigureSupplementalDataDataType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConfigureSupplementalDataDataType: %v", v)
	}
}

type ConfigureSupplementalDataRequest struct {
	// The configuration for the specified platform and data type.
	SupplementalDataConfiguration *shared.SupplementalDataConfiguration `request:"mediaType=application/json"`
	// Data types that support supplemental data
	DataType    ConfigureSupplementalDataDataType `pathParam:"style=simple,explode=false,name=dataType"`
	PlatformKey string                            `pathParam:"style=simple,explode=false,name=platformKey"`
}

type ConfigureSupplementalDataResponse struct {
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

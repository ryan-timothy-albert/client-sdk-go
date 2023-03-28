// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateSuppliersSourceModifiedDateAddressesTypeEnum - Type of the address.
type CreateSuppliersSourceModifiedDateAddressesTypeEnum string

const (
	CreateSuppliersSourceModifiedDateAddressesTypeEnumUnknown  CreateSuppliersSourceModifiedDateAddressesTypeEnum = "Unknown"
	CreateSuppliersSourceModifiedDateAddressesTypeEnumBilling  CreateSuppliersSourceModifiedDateAddressesTypeEnum = "Billing"
	CreateSuppliersSourceModifiedDateAddressesTypeEnumDelivery CreateSuppliersSourceModifiedDateAddressesTypeEnum = "Delivery"
)

func (e *CreateSuppliersSourceModifiedDateAddressesTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Unknown":
		fallthrough
	case "Billing":
		fallthrough
	case "Delivery":
		*e = CreateSuppliersSourceModifiedDateAddressesTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSuppliersSourceModifiedDateAddressesTypeEnum: %s", s)
	}
}

type CreateSuppliersSourceModifiedDateAddresses struct {
	// City of the customer address.
	City *string `json:"city,omitempty"`
	// Country of the customer address.
	Country *string `json:"country,omitempty"`
	// Line 1 of the customer address.
	Line1 *string `json:"line1,omitempty"`
	// Line 2 of the customer address.
	Line2 *string `json:"line2,omitempty"`
	// Postal code or zip code.
	PostalCode *string `json:"postalCode,omitempty"`
	// Region of the customer address.
	Region *string `json:"region,omitempty"`
	// Type of the address.
	Type CreateSuppliersSourceModifiedDateAddressesTypeEnum `json:"type"`
}

type CreateSuppliersSourceModifiedDateMetadata struct {
	// Indicates whether the record has been deleted in the third-party system this record originated from.
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// CreateSuppliersSourceModifiedDateStatusEnum - Status of the supplier.
type CreateSuppliersSourceModifiedDateStatusEnum string

const (
	CreateSuppliersSourceModifiedDateStatusEnumUnknown  CreateSuppliersSourceModifiedDateStatusEnum = "Unknown"
	CreateSuppliersSourceModifiedDateStatusEnumActive   CreateSuppliersSourceModifiedDateStatusEnum = "Active"
	CreateSuppliersSourceModifiedDateStatusEnumArchived CreateSuppliersSourceModifiedDateStatusEnum = "Archived"
)

func (e *CreateSuppliersSourceModifiedDateStatusEnum) UnmarshalJSON(data []byte) error {
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
		*e = CreateSuppliersSourceModifiedDateStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSuppliersSourceModifiedDateStatusEnum: %s", s)
	}
}

// CreateSuppliersSourceModifiedDateSupplementalData - Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreateSuppliersSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// CreateSuppliersSourceModifiedDate - > View the coverage for suppliers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// From the **Suppliers** endpoints, you can retrieve a list of [all the suppliers for a company](https://api.codat.io/swagger/index.html#/Suppliers/get_companies__companyId__data_suppliers). Suppliers' data links to accounts payable [bills](https://docs.codat.io/accounting-api#/schemas/Bill).
type CreateSuppliersSourceModifiedDate struct {
	// An array of Addresses.
	Addresses []CreateSuppliersSourceModifiedDateAddresses `json:"addresses,omitempty"`
	// Name of the main contact for the supplier.
	ContactName *string `json:"contactName,omitempty"`
	// Default currency the supplier's transactional data is recorded in.
	DefaultCurrency *string `json:"defaultCurrency,omitempty"`
	// Email address that the supplier may be contacted on.
	EmailAddress *string `json:"emailAddress,omitempty"`
	// Identifier for the supplier, unique to the company in the accounting platform.
	ID       *string                                    `json:"id,omitempty"`
	Metadata *CreateSuppliersSourceModifiedDateMetadata `json:"metadata,omitempty"`
	// The date on which this record was last modified in Codat.
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// Phone number that the supplier may be contacted on.
	Phone *string `json:"phone,omitempty"`
	// Company number of the supplier. In the UK, this is typically the company registration number issued by Companies House.
	RegistrationNumber *string `json:"registrationNumber,omitempty"`
	// The date on which this record was last modified in the originating system
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Status of the supplier.
	Status CreateSuppliersSourceModifiedDateStatusEnum `json:"status"`
	// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
	SupplementalData *CreateSuppliersSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	//
	// Name of the supplier as recorded in the accounting system, typically the company name.
	SupplierName *string `json:"supplierName,omitempty"`
	// Supplier's company tax number.
	TaxNumber *string `json:"taxNumber,omitempty"`
}

type CreateSuppliersRequest struct {
	RequestBody      *CreateSuppliersSourceModifiedDate `request:"mediaType=application/json"`
	CompanyID        string                             `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID     string                             `pathParam:"style=simple,explode=false,name=connectionId"`
	TimeoutInMinutes *int                               `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type CreateSuppliers200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type CreateSuppliers200ApplicationJSONChangesTypeEnum string

const (
	CreateSuppliers200ApplicationJSONChangesTypeEnumUnknown            CreateSuppliers200ApplicationJSONChangesTypeEnum = "Unknown"
	CreateSuppliers200ApplicationJSONChangesTypeEnumCreated            CreateSuppliers200ApplicationJSONChangesTypeEnum = "Created"
	CreateSuppliers200ApplicationJSONChangesTypeEnumModified           CreateSuppliers200ApplicationJSONChangesTypeEnum = "Modified"
	CreateSuppliers200ApplicationJSONChangesTypeEnumDeleted            CreateSuppliers200ApplicationJSONChangesTypeEnum = "Deleted"
	CreateSuppliers200ApplicationJSONChangesTypeEnumAttachmentUploaded CreateSuppliers200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

func (e *CreateSuppliers200ApplicationJSONChangesTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Unknown":
		fallthrough
	case "Created":
		fallthrough
	case "Modified":
		fallthrough
	case "Deleted":
		fallthrough
	case "AttachmentUploaded":
		*e = CreateSuppliers200ApplicationJSONChangesTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSuppliers200ApplicationJSONChangesTypeEnum: %s", s)
	}
}

type CreateSuppliers200ApplicationJSONChanges struct {
	AttachmentID *string                                                         `json:"attachmentId,omitempty"`
	RecordRef    *CreateSuppliers200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *CreateSuppliers200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// CreateSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnum - Type of the address.
type CreateSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnum string

const (
	CreateSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnumUnknown  CreateSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnum = "Unknown"
	CreateSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnumBilling  CreateSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnum = "Billing"
	CreateSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnumDelivery CreateSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnum = "Delivery"
)

func (e *CreateSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Unknown":
		fallthrough
	case "Billing":
		fallthrough
	case "Delivery":
		*e = CreateSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnum: %s", s)
	}
}

type CreateSuppliers200ApplicationJSONSourceModifiedDateAddresses struct {
	// City of the customer address.
	City *string `json:"city,omitempty"`
	// Country of the customer address.
	Country *string `json:"country,omitempty"`
	// Line 1 of the customer address.
	Line1 *string `json:"line1,omitempty"`
	// Line 2 of the customer address.
	Line2 *string `json:"line2,omitempty"`
	// Postal code or zip code.
	PostalCode *string `json:"postalCode,omitempty"`
	// Region of the customer address.
	Region *string `json:"region,omitempty"`
	// Type of the address.
	Type CreateSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnum `json:"type"`
}

type CreateSuppliers200ApplicationJSONSourceModifiedDateMetadata struct {
	// Indicates whether the record has been deleted in the third-party system this record originated from.
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// CreateSuppliers200ApplicationJSONSourceModifiedDateStatusEnum - Status of the supplier.
type CreateSuppliers200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	CreateSuppliers200ApplicationJSONSourceModifiedDateStatusEnumUnknown  CreateSuppliers200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	CreateSuppliers200ApplicationJSONSourceModifiedDateStatusEnumActive   CreateSuppliers200ApplicationJSONSourceModifiedDateStatusEnum = "Active"
	CreateSuppliers200ApplicationJSONSourceModifiedDateStatusEnumArchived CreateSuppliers200ApplicationJSONSourceModifiedDateStatusEnum = "Archived"
)

func (e *CreateSuppliers200ApplicationJSONSourceModifiedDateStatusEnum) UnmarshalJSON(data []byte) error {
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
		*e = CreateSuppliers200ApplicationJSONSourceModifiedDateStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSuppliers200ApplicationJSONSourceModifiedDateStatusEnum: %s", s)
	}
}

// CreateSuppliers200ApplicationJSONSourceModifiedDateSupplementalData - Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreateSuppliers200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// CreateSuppliers200ApplicationJSONSourceModifiedDate - > View the coverage for suppliers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// From the **Suppliers** endpoints, you can retrieve a list of [all the suppliers for a company](https://api.codat.io/swagger/index.html#/Suppliers/get_companies__companyId__data_suppliers). Suppliers' data links to accounts payable [bills](https://docs.codat.io/accounting-api#/schemas/Bill).
type CreateSuppliers200ApplicationJSONSourceModifiedDate struct {
	// An array of Addresses.
	Addresses []CreateSuppliers200ApplicationJSONSourceModifiedDateAddresses `json:"addresses,omitempty"`
	// Name of the main contact for the supplier.
	ContactName *string `json:"contactName,omitempty"`
	// Default currency the supplier's transactional data is recorded in.
	DefaultCurrency *string `json:"defaultCurrency,omitempty"`
	// Email address that the supplier may be contacted on.
	EmailAddress *string `json:"emailAddress,omitempty"`
	// Identifier for the supplier, unique to the company in the accounting platform.
	ID       *string                                                      `json:"id,omitempty"`
	Metadata *CreateSuppliers200ApplicationJSONSourceModifiedDateMetadata `json:"metadata,omitempty"`
	// The date on which this record was last modified in Codat.
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// Phone number that the supplier may be contacted on.
	Phone *string `json:"phone,omitempty"`
	// Company number of the supplier. In the UK, this is typically the company registration number issued by Companies House.
	RegistrationNumber *string `json:"registrationNumber,omitempty"`
	// The date on which this record was last modified in the originating system
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Status of the supplier.
	Status CreateSuppliers200ApplicationJSONSourceModifiedDateStatusEnum `json:"status"`
	// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
	SupplementalData *CreateSuppliers200ApplicationJSONSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	//
	// Name of the supplier as recorded in the accounting system, typically the company name.
	SupplierName *string `json:"supplierName,omitempty"`
	// Supplier's company tax number.
	TaxNumber *string `json:"taxNumber,omitempty"`
}

// CreateSuppliers200ApplicationJSONStatusEnum - The status of the push operation.
type CreateSuppliers200ApplicationJSONStatusEnum string

const (
	CreateSuppliers200ApplicationJSONStatusEnumPending  CreateSuppliers200ApplicationJSONStatusEnum = "Pending"
	CreateSuppliers200ApplicationJSONStatusEnumFailed   CreateSuppliers200ApplicationJSONStatusEnum = "Failed"
	CreateSuppliers200ApplicationJSONStatusEnumSuccess  CreateSuppliers200ApplicationJSONStatusEnum = "Success"
	CreateSuppliers200ApplicationJSONStatusEnumTimedOut CreateSuppliers200ApplicationJSONStatusEnum = "TimedOut"
)

func (e *CreateSuppliers200ApplicationJSONStatusEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Pending":
		fallthrough
	case "Failed":
		fallthrough
	case "Success":
		fallthrough
	case "TimedOut":
		*e = CreateSuppliers200ApplicationJSONStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSuppliers200ApplicationJSONStatusEnum: %s", s)
	}
}

type CreateSuppliers200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// CreateSuppliers200ApplicationJSONValidation - A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type CreateSuppliers200ApplicationJSONValidation struct {
	Errors   []CreateSuppliers200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []CreateSuppliers200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

// CreateSuppliers200ApplicationJSON - Success
type CreateSuppliers200ApplicationJSON struct {
	Changes []CreateSuppliers200ApplicationJSONChanges `json:"changes,omitempty"`
	// Unique identifier for your SMB in Codat.
	CompanyID string `json:"companyId"`
	// The datetime when the push was completed, null if Pending.
	CompletedOnUtc *string `json:"completedOnUtc,omitempty"`
	// > View the coverage for suppliers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers" target="_blank">Data coverage explorer</a>.
	//
	// ## Overview
	//
	// From the **Suppliers** endpoints, you can retrieve a list of [all the suppliers for a company](https://api.codat.io/swagger/index.html#/Suppliers/get_companies__companyId__data_suppliers). Suppliers' data links to accounts payable [bills](https://docs.codat.io/accounting-api#/schemas/Bill).
	Data *CreateSuppliers200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	// Unique identifier for a company's data connection.
	DataConnectionKey string `json:"dataConnectionKey"`
	// The type of data being pushed, eg invoices, customers.
	DataType     *string `json:"dataType,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// A unique identifier generated by Codat to represent this single push operation. This identifier can be used to track the status of the push, and should be persisted.
	PushOperationKey string `json:"pushOperationKey"`
	// The datetime when the push was requested.
	RequestedOnUtc string `json:"requestedOnUtc"`
	// The status of the push operation.
	Status           CreateSuppliers200ApplicationJSONStatusEnum `json:"status"`
	StatusCode       int                                         `json:"statusCode"`
	TimeoutInMinutes *int                                        `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds *int                                        `json:"timeoutInSeconds,omitempty"`
	// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
	Validation *CreateSuppliers200ApplicationJSONValidation `json:"validation,omitempty"`
}

type CreateSuppliersResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Success
	CreateSuppliers200ApplicationJSONObject *CreateSuppliers200ApplicationJSON
}

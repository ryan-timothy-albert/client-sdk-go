// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetCustomerRequest struct {
	CompanyID  string `pathParam:"style=simple,explode=false,name=companyId"`
	CustomerID string `pathParam:"style=simple,explode=false,name=customerId"`
}

// GetCustomerSourceModifiedDateAddressesTypeEnum - Type of the address.
type GetCustomerSourceModifiedDateAddressesTypeEnum string

const (
	GetCustomerSourceModifiedDateAddressesTypeEnumUnknown  GetCustomerSourceModifiedDateAddressesTypeEnum = "Unknown"
	GetCustomerSourceModifiedDateAddressesTypeEnumBilling  GetCustomerSourceModifiedDateAddressesTypeEnum = "Billing"
	GetCustomerSourceModifiedDateAddressesTypeEnumDelivery GetCustomerSourceModifiedDateAddressesTypeEnum = "Delivery"
)

func (e *GetCustomerSourceModifiedDateAddressesTypeEnum) UnmarshalJSON(data []byte) error {
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
		*e = GetCustomerSourceModifiedDateAddressesTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCustomerSourceModifiedDateAddressesTypeEnum: %s", s)
	}
}

type GetCustomerSourceModifiedDateAddresses struct {
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
	Type GetCustomerSourceModifiedDateAddressesTypeEnum `json:"type"`
}

// GetCustomerSourceModifiedDateContactsAddressTypeEnum - Type of the address.
type GetCustomerSourceModifiedDateContactsAddressTypeEnum string

const (
	GetCustomerSourceModifiedDateContactsAddressTypeEnumUnknown  GetCustomerSourceModifiedDateContactsAddressTypeEnum = "Unknown"
	GetCustomerSourceModifiedDateContactsAddressTypeEnumBilling  GetCustomerSourceModifiedDateContactsAddressTypeEnum = "Billing"
	GetCustomerSourceModifiedDateContactsAddressTypeEnumDelivery GetCustomerSourceModifiedDateContactsAddressTypeEnum = "Delivery"
)

func (e *GetCustomerSourceModifiedDateContactsAddressTypeEnum) UnmarshalJSON(data []byte) error {
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
		*e = GetCustomerSourceModifiedDateContactsAddressTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCustomerSourceModifiedDateContactsAddressTypeEnum: %s", s)
	}
}

// GetCustomerSourceModifiedDateContactsAddress - An object of Address information.
type GetCustomerSourceModifiedDateContactsAddress struct {
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
	Type GetCustomerSourceModifiedDateContactsAddressTypeEnum `json:"type"`
}

// GetCustomerSourceModifiedDateContactsPhoneTypeEnum - Type of phone number.
type GetCustomerSourceModifiedDateContactsPhoneTypeEnum string

const (
	GetCustomerSourceModifiedDateContactsPhoneTypeEnumUnknown  GetCustomerSourceModifiedDateContactsPhoneTypeEnum = "Unknown"
	GetCustomerSourceModifiedDateContactsPhoneTypeEnumPrimary  GetCustomerSourceModifiedDateContactsPhoneTypeEnum = "Primary"
	GetCustomerSourceModifiedDateContactsPhoneTypeEnumLandline GetCustomerSourceModifiedDateContactsPhoneTypeEnum = "Landline"
	GetCustomerSourceModifiedDateContactsPhoneTypeEnumMobile   GetCustomerSourceModifiedDateContactsPhoneTypeEnum = "Mobile"
	GetCustomerSourceModifiedDateContactsPhoneTypeEnumFax      GetCustomerSourceModifiedDateContactsPhoneTypeEnum = "Fax"
)

func (e *GetCustomerSourceModifiedDateContactsPhoneTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Unknown":
		fallthrough
	case "Primary":
		fallthrough
	case "Landline":
		fallthrough
	case "Mobile":
		fallthrough
	case "Fax":
		*e = GetCustomerSourceModifiedDateContactsPhoneTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCustomerSourceModifiedDateContactsPhoneTypeEnum: %s", s)
	}
}

type GetCustomerSourceModifiedDateContactsPhone struct {
	// Phone number for a customer contact.
	Number *string `json:"number,omitempty"`
	// Type of phone number.
	Type GetCustomerSourceModifiedDateContactsPhoneTypeEnum `json:"type"`
}

// GetCustomerSourceModifiedDateContactsStatusEnum - Status of customer contacts.
//
// Customers can have multiple contacts.
type GetCustomerSourceModifiedDateContactsStatusEnum string

const (
	GetCustomerSourceModifiedDateContactsStatusEnumUnknown  GetCustomerSourceModifiedDateContactsStatusEnum = "Unknown"
	GetCustomerSourceModifiedDateContactsStatusEnumActive   GetCustomerSourceModifiedDateContactsStatusEnum = "Active"
	GetCustomerSourceModifiedDateContactsStatusEnumArchived GetCustomerSourceModifiedDateContactsStatusEnum = "Archived"
)

func (e *GetCustomerSourceModifiedDateContactsStatusEnum) UnmarshalJSON(data []byte) error {
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
		*e = GetCustomerSourceModifiedDateContactsStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCustomerSourceModifiedDateContactsStatusEnum: %s", s)
	}
}

type GetCustomerSourceModifiedDateContacts struct {
	// An object of Address information.
	Address *GetCustomerSourceModifiedDateContactsAddress `json:"address,omitempty"`
	// Email of a contact for a customer.
	Email *string `json:"email,omitempty"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > 📘 Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// Name of a contact for a customer.
	Name *string `json:"name,omitempty"`
	// An array of Phone numbers.
	Phone []GetCustomerSourceModifiedDateContactsPhone `json:"phone,omitempty"`
	// Status of customer contacts.
	//
	// Customers can have multiple contacts.
	Status GetCustomerSourceModifiedDateContactsStatusEnum `json:"status"`
}

type GetCustomerSourceModifiedDateMetadata struct {
	// Indicates whether the record has been deleted in the third-party system this record originated from.
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// GetCustomerSourceModifiedDateStatusEnum - Current state of the customer.
type GetCustomerSourceModifiedDateStatusEnum string

const (
	GetCustomerSourceModifiedDateStatusEnumUnknown  GetCustomerSourceModifiedDateStatusEnum = "Unknown"
	GetCustomerSourceModifiedDateStatusEnumActive   GetCustomerSourceModifiedDateStatusEnum = "Active"
	GetCustomerSourceModifiedDateStatusEnumArchived GetCustomerSourceModifiedDateStatusEnum = "Archived"
)

func (e *GetCustomerSourceModifiedDateStatusEnum) UnmarshalJSON(data []byte) error {
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
		*e = GetCustomerSourceModifiedDateStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCustomerSourceModifiedDateStatusEnum: %s", s)
	}
}

// GetCustomerSourceModifiedDateSupplementalData - Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type GetCustomerSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// GetCustomerSourceModifiedDate - > View the coverage for customers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A customer is a person or organisation that buys goods or services. From the Customers endpoints, you can retrieve a [list of all the customers of a company](https://api.codat.io/swagger/index.html#/Customers/get_companies__companyId__data_customers).
//
// Customers' data links to accounts receivable [invoices](https://docs.codat.io/accounting-api#/schemas/Invoice).
type GetCustomerSourceModifiedDate struct {
	// An array of Addresses.
	Addresses []GetCustomerSourceModifiedDateAddresses `json:"addresses,omitempty"`
	// Name of the main contact for the identified customer.
	ContactName *string `json:"contactName,omitempty"`
	// An array of Contacts.
	Contacts []GetCustomerSourceModifiedDateContacts `json:"contacts,omitempty"`
	// Name of the customer as recorded in the accounting system, typically the company name.
	CustomerName *string `json:"customerName,omitempty"`
	// Default currency the transactional data of the customer is recorded in.
	DefaultCurrency *string `json:"defaultCurrency,omitempty"`
	// Email address the customer can be contacted by.
	EmailAddress *string `json:"emailAddress,omitempty"`
	// Identifier for the customer, unique to the company in the accounting platform.
	ID       *string                                `json:"id,omitempty"`
	Metadata *GetCustomerSourceModifiedDateMetadata `json:"metadata,omitempty"`
	// The date on which this record was last modified in Codat.
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// Phone number the customer can be contacted by.
	Phone *string `json:"phone,omitempty"`
	// Company number. In the UK, this is typically the Companies House company registration number.
	RegistrationNumber *string `json:"registrationNumber,omitempty"`
	// The date on which this record was last modified in the originating system
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Current state of the customer.
	Status GetCustomerSourceModifiedDateStatusEnum `json:"status"`
	// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
	SupplementalData *GetCustomerSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	// Company tax number.
	TaxNumber *string `json:"taxNumber,omitempty"`
}

type GetCustomerResponse struct {
	ContentType string
	// Success
	SourceModifiedDate *GetCustomerSourceModifiedDate
	StatusCode         int
	RawResponse        *http.Response
}

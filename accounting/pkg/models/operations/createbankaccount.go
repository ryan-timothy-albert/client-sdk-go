// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateBankAccountSourceModifiedDateAccountTypeEnum - The type of the account.
type CreateBankAccountSourceModifiedDateAccountTypeEnum string

const (
	CreateBankAccountSourceModifiedDateAccountTypeEnumUnknown CreateBankAccountSourceModifiedDateAccountTypeEnum = "Unknown"
	CreateBankAccountSourceModifiedDateAccountTypeEnumCredit  CreateBankAccountSourceModifiedDateAccountTypeEnum = "Credit"
	CreateBankAccountSourceModifiedDateAccountTypeEnumDebit   CreateBankAccountSourceModifiedDateAccountTypeEnum = "Debit"
)

func (e *CreateBankAccountSourceModifiedDateAccountTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Unknown":
		fallthrough
	case "Credit":
		fallthrough
	case "Debit":
		*e = CreateBankAccountSourceModifiedDateAccountTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateBankAccountSourceModifiedDateAccountTypeEnum: %s", s)
	}
}

type CreateBankAccountSourceModifiedDateMetadata struct {
	// Indicates whether the record has been deleted in the third-party system this record originated from.
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// CreateBankAccountSourceModifiedDate - > **Accessing Bank Accounts through Banking API**
// >
// > This datatype was originally used for accessing bank account data both in accounting integrations and open banking aggregators.
// >
// > To view bank account data through the Banking API, please refer to the new datatype [here](https://docs.codat.io/banking-api#/schemas/Account)
//
// > View the coverage for bank accounts in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankAccounts" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A list of bank accounts associated with a company and a specific [data connection](https://api.codat.io/swagger/index.html#/Connection/get_companies__companyId__connections__connectionId_).
//
// Bank accounts data includes:
// * The name and ID of the account in the accounting platform.
// * The currency and balance of the account.
// * The sort code and account number.
type CreateBankAccountSourceModifiedDate struct {
	// Name of the bank account in the accounting platform.
	AccountName *string `json:"accountName,omitempty"`
	// Account number for the bank account.
	//
	// Xero integrations
	// Only a UK account number shows for bank accounts with GBP currency and a combined total of sort code and account number that equals 14 digits, For non-GBP accounts, the full bank account number is populated.
	//
	// FreeAgent integrations
	// For Credit accounts, only the last four digits are required. For other types, the field is optional.
	AccountNumber *string `json:"accountNumber,omitempty"`
	// The type of the account.
	AccountType *CreateBankAccountSourceModifiedDateAccountTypeEnum `json:"accountType,omitempty"`
	// Total available balance of the bank account as reported by the underlying data source. This may take into account overdrafts or pending transactions for example.
	AvailableBalance *float64 `json:"availableBalance,omitempty"`
	// Balance of the bank account.
	Balance *float64 `json:"balance,omitempty"`
	// Base currency of the bank account.
	Currency *string `json:"currency,omitempty"`
	// International bank account number of the account. Often used when making or receiving international payments.
	IBan *string `json:"iBan,omitempty"`
	// Identifier for the account, unique for the company in the accounting platform.
	ID *string `json:"id,omitempty"`
	// The institution of the bank account.
	Institution *string                                      `json:"institution,omitempty"`
	Metadata    *CreateBankAccountSourceModifiedDateMetadata `json:"metadata,omitempty"`
	// The date on which this record was last modified in Codat.
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// Code used to identify each nominal account for a business.
	NominalCode *string `json:"nominalCode,omitempty"`
	// Pre-arranged overdraft limit of the account.
	//
	// The value is always positive. For example, an overdraftLimit of `1000` means that the balance of the account can go down to `-1000`.
	OverdraftLimit *float64 `json:"overdraftLimit,omitempty"`
	// Sort code for the bank account.
	//
	// Xero integrations
	// The sort code is only displayed when the currency = GBP and the sort code and account number sum to 14 digits. For non-GBP accounts, this field is not populated.
	SortCode *string `json:"sortCode,omitempty"`
	// The date on which this record was last modified in the originating system
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
}

type CreateBankAccountRequest struct {
	RequestBody             *CreateBankAccountSourceModifiedDate `request:"mediaType=application/json"`
	AllowSyncOnPushComplete *bool                                `queryParam:"style=form,explode=true,name=allowSyncOnPushComplete"`
	CompanyID               string                               `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID            string                               `pathParam:"style=simple,explode=false,name=connectionId"`
	TimeoutInMinutes        *int                                 `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type CreateBankAccount200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type CreateBankAccount200ApplicationJSONChangesTypeEnum string

const (
	CreateBankAccount200ApplicationJSONChangesTypeEnumUnknown            CreateBankAccount200ApplicationJSONChangesTypeEnum = "Unknown"
	CreateBankAccount200ApplicationJSONChangesTypeEnumCreated            CreateBankAccount200ApplicationJSONChangesTypeEnum = "Created"
	CreateBankAccount200ApplicationJSONChangesTypeEnumModified           CreateBankAccount200ApplicationJSONChangesTypeEnum = "Modified"
	CreateBankAccount200ApplicationJSONChangesTypeEnumDeleted            CreateBankAccount200ApplicationJSONChangesTypeEnum = "Deleted"
	CreateBankAccount200ApplicationJSONChangesTypeEnumAttachmentUploaded CreateBankAccount200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

func (e *CreateBankAccount200ApplicationJSONChangesTypeEnum) UnmarshalJSON(data []byte) error {
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
		*e = CreateBankAccount200ApplicationJSONChangesTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateBankAccount200ApplicationJSONChangesTypeEnum: %s", s)
	}
}

type CreateBankAccount200ApplicationJSONChanges struct {
	AttachmentID *string                                                           `json:"attachmentId,omitempty"`
	RecordRef    *CreateBankAccount200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *CreateBankAccount200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// CreateBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnum - The type of the account.
type CreateBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnum string

const (
	CreateBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnumUnknown CreateBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnum = "Unknown"
	CreateBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnumCredit  CreateBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnum = "Credit"
	CreateBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnumDebit   CreateBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnum = "Debit"
)

func (e *CreateBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Unknown":
		fallthrough
	case "Credit":
		fallthrough
	case "Debit":
		*e = CreateBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnum: %s", s)
	}
}

type CreateBankAccount200ApplicationJSONSourceModifiedDateMetadata struct {
	// Indicates whether the record has been deleted in the third-party system this record originated from.
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// CreateBankAccount200ApplicationJSONSourceModifiedDate - > **Accessing Bank Accounts through Banking API**
// >
// > This datatype was originally used for accessing bank account data both in accounting integrations and open banking aggregators.
// >
// > To view bank account data through the Banking API, please refer to the new datatype [here](https://docs.codat.io/banking-api#/schemas/Account)
//
// > View the coverage for bank accounts in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankAccounts" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A list of bank accounts associated with a company and a specific [data connection](https://api.codat.io/swagger/index.html#/Connection/get_companies__companyId__connections__connectionId_).
//
// Bank accounts data includes:
// * The name and ID of the account in the accounting platform.
// * The currency and balance of the account.
// * The sort code and account number.
type CreateBankAccount200ApplicationJSONSourceModifiedDate struct {
	// Name of the bank account in the accounting platform.
	AccountName *string `json:"accountName,omitempty"`
	// Account number for the bank account.
	//
	// Xero integrations
	// Only a UK account number shows for bank accounts with GBP currency and a combined total of sort code and account number that equals 14 digits, For non-GBP accounts, the full bank account number is populated.
	//
	// FreeAgent integrations
	// For Credit accounts, only the last four digits are required. For other types, the field is optional.
	AccountNumber *string `json:"accountNumber,omitempty"`
	// The type of the account.
	AccountType *CreateBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnum `json:"accountType,omitempty"`
	// Total available balance of the bank account as reported by the underlying data source. This may take into account overdrafts or pending transactions for example.
	AvailableBalance *float64 `json:"availableBalance,omitempty"`
	// Balance of the bank account.
	Balance *float64 `json:"balance,omitempty"`
	// Base currency of the bank account.
	Currency *string `json:"currency,omitempty"`
	// International bank account number of the account. Often used when making or receiving international payments.
	IBan *string `json:"iBan,omitempty"`
	// Identifier for the account, unique for the company in the accounting platform.
	ID *string `json:"id,omitempty"`
	// The institution of the bank account.
	Institution *string                                                        `json:"institution,omitempty"`
	Metadata    *CreateBankAccount200ApplicationJSONSourceModifiedDateMetadata `json:"metadata,omitempty"`
	// The date on which this record was last modified in Codat.
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// Code used to identify each nominal account for a business.
	NominalCode *string `json:"nominalCode,omitempty"`
	// Pre-arranged overdraft limit of the account.
	//
	// The value is always positive. For example, an overdraftLimit of `1000` means that the balance of the account can go down to `-1000`.
	OverdraftLimit *float64 `json:"overdraftLimit,omitempty"`
	// Sort code for the bank account.
	//
	// Xero integrations
	// The sort code is only displayed when the currency = GBP and the sort code and account number sum to 14 digits. For non-GBP accounts, this field is not populated.
	SortCode *string `json:"sortCode,omitempty"`
	// The date on which this record was last modified in the originating system
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
}

// CreateBankAccount200ApplicationJSONStatusEnum - The status of the push operation.
type CreateBankAccount200ApplicationJSONStatusEnum string

const (
	CreateBankAccount200ApplicationJSONStatusEnumPending  CreateBankAccount200ApplicationJSONStatusEnum = "Pending"
	CreateBankAccount200ApplicationJSONStatusEnumFailed   CreateBankAccount200ApplicationJSONStatusEnum = "Failed"
	CreateBankAccount200ApplicationJSONStatusEnumSuccess  CreateBankAccount200ApplicationJSONStatusEnum = "Success"
	CreateBankAccount200ApplicationJSONStatusEnumTimedOut CreateBankAccount200ApplicationJSONStatusEnum = "TimedOut"
)

func (e *CreateBankAccount200ApplicationJSONStatusEnum) UnmarshalJSON(data []byte) error {
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
		*e = CreateBankAccount200ApplicationJSONStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateBankAccount200ApplicationJSONStatusEnum: %s", s)
	}
}

type CreateBankAccount200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// CreateBankAccount200ApplicationJSONValidation - A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type CreateBankAccount200ApplicationJSONValidation struct {
	Errors   []CreateBankAccount200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []CreateBankAccount200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

// CreateBankAccount200ApplicationJSON - Success
type CreateBankAccount200ApplicationJSON struct {
	Changes []CreateBankAccount200ApplicationJSONChanges `json:"changes,omitempty"`
	// Unique identifier for your SMB in Codat.
	CompanyID string `json:"companyId"`
	// The datetime when the push was completed, null if Pending.
	CompletedOnUtc *string `json:"completedOnUtc,omitempty"`
	// > **Accessing Bank Accounts through Banking API**
	// >
	// > This datatype was originally used for accessing bank account data both in accounting integrations and open banking aggregators.
	// >
	// > To view bank account data through the Banking API, please refer to the new datatype [here](https://docs.codat.io/banking-api#/schemas/Account)
	//
	// > View the coverage for bank accounts in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankAccounts" target="_blank">Data coverage explorer</a>.
	//
	// ## Overview
	//
	// A list of bank accounts associated with a company and a specific [data connection](https://api.codat.io/swagger/index.html#/Connection/get_companies__companyId__connections__connectionId_).
	//
	// Bank accounts data includes:
	// * The name and ID of the account in the accounting platform.
	// * The currency and balance of the account.
	// * The sort code and account number.
	Data *CreateBankAccount200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
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
	Status           CreateBankAccount200ApplicationJSONStatusEnum `json:"status"`
	StatusCode       int                                           `json:"statusCode"`
	TimeoutInMinutes *int                                          `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds *int                                          `json:"timeoutInSeconds,omitempty"`
	// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
	Validation *CreateBankAccount200ApplicationJSONValidation `json:"validation,omitempty"`
}

type CreateBankAccountResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Success
	CreateBankAccount200ApplicationJSONObject *CreateBankAccount200ApplicationJSON
}

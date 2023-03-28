// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateTransferSourceModifiedDateContactRef - The customer or supplier for the transfer, if available.
type CreateTransferSourceModifiedDateContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// CreateTransferSourceModifiedDateTransferAccountAccountRef - The account that the transfer is moving from or to.
type CreateTransferSourceModifiedDateTransferAccountAccountRef struct {
	// 'id' from the Accounts data type.
	ID *string `json:"id,omitempty"`
	// 'name' from the Accounts data type.
	Name *string `json:"name,omitempty"`
}

// CreateTransferSourceModifiedDateTransferAccount - The details of the accounts the transfer is moving from.
type CreateTransferSourceModifiedDateTransferAccount struct {
	// The account that the transfer is moving from or to.
	AccountRef *CreateTransferSourceModifiedDateTransferAccountAccountRef `json:"accountRef,omitempty"`
	// The amount transferred between accounts.
	Amount *float64 `json:"amount,omitempty"`
	// ISO currency code recorded for the transfer in the accounting platform.
	Currency *string `json:"currency,omitempty"`
}

type CreateTransferSourceModifiedDateMetadata struct {
	// Indicates whether the record has been deleted in the third-party system this record originated from.
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// CreateTransferSourceModifiedDateSupplementalData - Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreateTransferSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// CreateTransferSourceModifiedDateTrackingCategoryRefs - References a category against which the item is tracked.
type CreateTransferSourceModifiedDateTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreateTransferSourceModifiedDate - > View the coverage for transfers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=transfers" target="_blank">Data coverage explorer</a>.
//
// From the **Transfers** endpoints, you can:
//
// - [Retrieve a list of all transfers for a specified company](https://api.codat.io/swagger/index.html#/Transfers/get_companies__companyId__connections__connectionId__data_transfers)
// - [Retrieve a single transfer for a specified company](https://api.codat.io/swagger/index.html#/Transfers/get_companies__companyId__connections__connectionId__data_transfers__transferId_)
// - [Add a new transfer for a specified company](https://api.codat.io/swagger/index.html#/Transfers/post_companies__companyId__connections__connectionId__push_transfers)
//
// **Transfers** is a child data type of [account transactions](https://docs.codat.io/accounting-api#/schemas/AccountTransaction).
type CreateTransferSourceModifiedDate struct {
	// The customer or supplier for the transfer, if available.
	ContactRef *CreateTransferSourceModifiedDateContactRef `json:"contactRef,omitempty"`
	// The day on which the transfer was made.
	Date                *string  `json:"date,omitempty"`
	DepositedRecordRefs []string `json:"depositedRecordRefs,omitempty"`
	// Description of the transfer.
	Description *string `json:"description,omitempty"`
	// The details of the accounts the transfer is moving from.
	From *CreateTransferSourceModifiedDateTransferAccount `json:"from,omitempty"`
	// Unique identifier for the transfer.
	ID       *string                                   `json:"id,omitempty"`
	Metadata *CreateTransferSourceModifiedDateMetadata `json:"metadata,omitempty"`
	// The date on which this record was last modified in Codat.
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// The date on which this record was last modified in the originating system
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
	SupplementalData *CreateTransferSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	// The details of the accounts the transfer is moving to.
	To *CreateTransferSourceModifiedDateTransferAccount `json:"to,omitempty"`
	// Reference to the tracking categories this transfer is being tracked against.
	TrackingCategoryRefs []CreateTransferSourceModifiedDateTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
}

type CreateTransferRequest struct {
	RequestBody  *CreateTransferSourceModifiedDate `request:"mediaType=application/json"`
	CompanyID    string                            `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string                            `pathParam:"style=simple,explode=false,name=connectionId"`
}

type CreateTransfer200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type CreateTransfer200ApplicationJSONChangesTypeEnum string

const (
	CreateTransfer200ApplicationJSONChangesTypeEnumUnknown            CreateTransfer200ApplicationJSONChangesTypeEnum = "Unknown"
	CreateTransfer200ApplicationJSONChangesTypeEnumCreated            CreateTransfer200ApplicationJSONChangesTypeEnum = "Created"
	CreateTransfer200ApplicationJSONChangesTypeEnumModified           CreateTransfer200ApplicationJSONChangesTypeEnum = "Modified"
	CreateTransfer200ApplicationJSONChangesTypeEnumDeleted            CreateTransfer200ApplicationJSONChangesTypeEnum = "Deleted"
	CreateTransfer200ApplicationJSONChangesTypeEnumAttachmentUploaded CreateTransfer200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

func (e *CreateTransfer200ApplicationJSONChangesTypeEnum) UnmarshalJSON(data []byte) error {
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
		*e = CreateTransfer200ApplicationJSONChangesTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateTransfer200ApplicationJSONChangesTypeEnum: %s", s)
	}
}

type CreateTransfer200ApplicationJSONChanges struct {
	AttachmentID *string                                                        `json:"attachmentId,omitempty"`
	RecordRef    *CreateTransfer200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *CreateTransfer200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// CreateTransfer200ApplicationJSONSourceModifiedDateContactRef - The customer or supplier for the transfer, if available.
type CreateTransfer200ApplicationJSONSourceModifiedDateContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// CreateTransfer200ApplicationJSONSourceModifiedDateTransferAccountAccountRef - The account that the transfer is moving from or to.
type CreateTransfer200ApplicationJSONSourceModifiedDateTransferAccountAccountRef struct {
	// 'id' from the Accounts data type.
	ID *string `json:"id,omitempty"`
	// 'name' from the Accounts data type.
	Name *string `json:"name,omitempty"`
}

// CreateTransfer200ApplicationJSONSourceModifiedDateTransferAccount - The details of the accounts the transfer is moving from.
type CreateTransfer200ApplicationJSONSourceModifiedDateTransferAccount struct {
	// The account that the transfer is moving from or to.
	AccountRef *CreateTransfer200ApplicationJSONSourceModifiedDateTransferAccountAccountRef `json:"accountRef,omitempty"`
	// The amount transferred between accounts.
	Amount *float64 `json:"amount,omitempty"`
	// ISO currency code recorded for the transfer in the accounting platform.
	Currency *string `json:"currency,omitempty"`
}

type CreateTransfer200ApplicationJSONSourceModifiedDateMetadata struct {
	// Indicates whether the record has been deleted in the third-party system this record originated from.
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// CreateTransfer200ApplicationJSONSourceModifiedDateSupplementalData - Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreateTransfer200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// CreateTransfer200ApplicationJSONSourceModifiedDateTrackingCategoryRefs - References a category against which the item is tracked.
type CreateTransfer200ApplicationJSONSourceModifiedDateTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreateTransfer200ApplicationJSONSourceModifiedDate - > View the coverage for transfers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=transfers" target="_blank">Data coverage explorer</a>.
//
// From the **Transfers** endpoints, you can:
//
// - [Retrieve a list of all transfers for a specified company](https://api.codat.io/swagger/index.html#/Transfers/get_companies__companyId__connections__connectionId__data_transfers)
// - [Retrieve a single transfer for a specified company](https://api.codat.io/swagger/index.html#/Transfers/get_companies__companyId__connections__connectionId__data_transfers__transferId_)
// - [Add a new transfer for a specified company](https://api.codat.io/swagger/index.html#/Transfers/post_companies__companyId__connections__connectionId__push_transfers)
//
// **Transfers** is a child data type of [account transactions](https://docs.codat.io/accounting-api#/schemas/AccountTransaction).
type CreateTransfer200ApplicationJSONSourceModifiedDate struct {
	// The customer or supplier for the transfer, if available.
	ContactRef *CreateTransfer200ApplicationJSONSourceModifiedDateContactRef `json:"contactRef,omitempty"`
	// The day on which the transfer was made.
	Date                *string  `json:"date,omitempty"`
	DepositedRecordRefs []string `json:"depositedRecordRefs,omitempty"`
	// Description of the transfer.
	Description *string `json:"description,omitempty"`
	// The details of the accounts the transfer is moving from.
	From *CreateTransfer200ApplicationJSONSourceModifiedDateTransferAccount `json:"from,omitempty"`
	// Unique identifier for the transfer.
	ID       *string                                                     `json:"id,omitempty"`
	Metadata *CreateTransfer200ApplicationJSONSourceModifiedDateMetadata `json:"metadata,omitempty"`
	// The date on which this record was last modified in Codat.
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// The date on which this record was last modified in the originating system
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
	SupplementalData *CreateTransfer200ApplicationJSONSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	// The details of the accounts the transfer is moving to.
	To *CreateTransfer200ApplicationJSONSourceModifiedDateTransferAccount `json:"to,omitempty"`
	// Reference to the tracking categories this transfer is being tracked against.
	TrackingCategoryRefs []CreateTransfer200ApplicationJSONSourceModifiedDateTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
}

// CreateTransfer200ApplicationJSONStatusEnum - The status of the push operation.
type CreateTransfer200ApplicationJSONStatusEnum string

const (
	CreateTransfer200ApplicationJSONStatusEnumPending  CreateTransfer200ApplicationJSONStatusEnum = "Pending"
	CreateTransfer200ApplicationJSONStatusEnumFailed   CreateTransfer200ApplicationJSONStatusEnum = "Failed"
	CreateTransfer200ApplicationJSONStatusEnumSuccess  CreateTransfer200ApplicationJSONStatusEnum = "Success"
	CreateTransfer200ApplicationJSONStatusEnumTimedOut CreateTransfer200ApplicationJSONStatusEnum = "TimedOut"
)

func (e *CreateTransfer200ApplicationJSONStatusEnum) UnmarshalJSON(data []byte) error {
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
		*e = CreateTransfer200ApplicationJSONStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateTransfer200ApplicationJSONStatusEnum: %s", s)
	}
}

type CreateTransfer200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// CreateTransfer200ApplicationJSONValidation - A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type CreateTransfer200ApplicationJSONValidation struct {
	Errors   []CreateTransfer200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []CreateTransfer200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

// CreateTransfer200ApplicationJSON - Success
type CreateTransfer200ApplicationJSON struct {
	Changes []CreateTransfer200ApplicationJSONChanges `json:"changes,omitempty"`
	// Unique identifier for your SMB in Codat.
	CompanyID string `json:"companyId"`
	// The datetime when the push was completed, null if Pending.
	CompletedOnUtc *string `json:"completedOnUtc,omitempty"`
	// > View the coverage for transfers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=transfers" target="_blank">Data coverage explorer</a>.
	//
	// From the **Transfers** endpoints, you can:
	//
	// - [Retrieve a list of all transfers for a specified company](https://api.codat.io/swagger/index.html#/Transfers/get_companies__companyId__connections__connectionId__data_transfers)
	// - [Retrieve a single transfer for a specified company](https://api.codat.io/swagger/index.html#/Transfers/get_companies__companyId__connections__connectionId__data_transfers__transferId_)
	// - [Add a new transfer for a specified company](https://api.codat.io/swagger/index.html#/Transfers/post_companies__companyId__connections__connectionId__push_transfers)
	//
	// **Transfers** is a child data type of [account transactions](https://docs.codat.io/accounting-api#/schemas/AccountTransaction).
	Data *CreateTransfer200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
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
	Status           CreateTransfer200ApplicationJSONStatusEnum `json:"status"`
	StatusCode       int                                        `json:"statusCode"`
	TimeoutInMinutes *int                                       `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds *int                                       `json:"timeoutInSeconds,omitempty"`
	// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
	Validation *CreateTransfer200ApplicationJSONValidation `json:"validation,omitempty"`
}

type CreateTransferResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Success
	CreateTransfer200ApplicationJSONObject *CreateTransfer200ApplicationJSON
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentIDRequest struct {
	BillPaymentID string `pathParam:"style=simple,explode=false,name=billPaymentId"`
	CompanyID     string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID  string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnum string

const (
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnumUnknown            DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnum = "Unknown"
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnumCreated            DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnum = "Created"
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnumModified           DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnum = "Modified"
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnumDeleted            DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnum = "Deleted"
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnumAttachmentUploaded DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

func (e *DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnum) UnmarshalJSON(data []byte) error {
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
		*e = DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnum: %s", s)
	}
}

type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChanges struct {
	AttachmentID *string                                                                                                                      `json:"attachmentId,omitempty"`
	RecordRef    *DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONStatusEnum - The status of the push operation.
type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONStatusEnum string

const (
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONStatusEnumPending  DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONStatusEnum = "Pending"
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONStatusEnumFailed   DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONStatusEnum = "Failed"
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONStatusEnumSuccess  DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONStatusEnum = "Success"
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONStatusEnumTimedOut DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONStatusEnum = "TimedOut"
)

func (e *DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONStatusEnum) UnmarshalJSON(data []byte) error {
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
		*e = DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONStatusEnum: %s", s)
	}
}

type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONValidation - A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONValidation struct {
	Errors   []DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

// DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSON - OK
type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSON struct {
	Changes []DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChanges `json:"changes,omitempty"`
	// Unique identifier for your SMB in Codat.
	CompanyID string `json:"companyId"`
	// The datetime when the push was completed, null if Pending.
	CompletedOnUtc *string `json:"completedOnUtc,omitempty"`
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
	Status           DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONStatusEnum `json:"status"`
	StatusCode       int                                                                                                      `json:"statusCode"`
	TimeoutInMinutes *int                                                                                                     `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds *int                                                                                                     `json:"timeoutInSeconds,omitempty"`
	// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
	Validation *DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONValidation `json:"validation,omitempty"`
}

type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONObject *DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSON
}

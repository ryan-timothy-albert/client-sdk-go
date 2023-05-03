// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// PaymentStatusEnum - Status of the payment
type PaymentStatusEnum string

const (
	PaymentStatusEnumPending    PaymentStatusEnum = "Pending"
	PaymentStatusEnumAuthorized PaymentStatusEnum = "Authorized"
	PaymentStatusEnumPaid       PaymentStatusEnum = "Paid"
	PaymentStatusEnumFailed     PaymentStatusEnum = "Failed"
	PaymentStatusEnumCancelled  PaymentStatusEnum = "Cancelled"
	PaymentStatusEnumUnknown    PaymentStatusEnum = "Unknown"
)

func (e PaymentStatusEnum) ToPointer() *PaymentStatusEnum {
	return &e
}

func (e *PaymentStatusEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Pending":
		fallthrough
	case "Authorized":
		fallthrough
	case "Paid":
		fallthrough
	case "Failed":
		fallthrough
	case "Cancelled":
		fallthrough
	case "Unknown":
		*e = PaymentStatusEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentStatusEnum: %v", v)
	}
}

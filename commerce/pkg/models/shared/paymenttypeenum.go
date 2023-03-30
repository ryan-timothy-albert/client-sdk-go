// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// PaymentTypeEnum - Status of the payment
type PaymentTypeEnum string

const (
	PaymentTypeEnumCash        PaymentTypeEnum = "Cash"
	PaymentTypeEnumCard        PaymentTypeEnum = "Card"
	PaymentTypeEnumInvoice     PaymentTypeEnum = "Invoice"
	PaymentTypeEnumOnlineCard  PaymentTypeEnum = "OnlineCard"
	PaymentTypeEnumSwish       PaymentTypeEnum = "Swish"
	PaymentTypeEnumVipps       PaymentTypeEnum = "Vipps"
	PaymentTypeEnumMobile      PaymentTypeEnum = "Mobile"
	PaymentTypeEnumStoreCredit PaymentTypeEnum = "StoreCredit"
	PaymentTypeEnumPaypal      PaymentTypeEnum = "Paypal"
	PaymentTypeEnumCustom      PaymentTypeEnum = "Custom"
	PaymentTypeEnumPrepaid     PaymentTypeEnum = "Prepaid"
	PaymentTypeEnumUnknown     PaymentTypeEnum = "Unknown"
)

func (e *PaymentTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Cash":
		fallthrough
	case "Card":
		fallthrough
	case "Invoice":
		fallthrough
	case "OnlineCard":
		fallthrough
	case "Swish":
		fallthrough
	case "Vipps":
		fallthrough
	case "Mobile":
		fallthrough
	case "StoreCredit":
		fallthrough
	case "Paypal":
		fallthrough
	case "Custom":
		fallthrough
	case "Prepaid":
		fallthrough
	case "Unknown":
		*e = PaymentTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentTypeEnum: %s", s)
	}
}

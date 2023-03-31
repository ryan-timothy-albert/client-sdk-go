// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// BillPaymentLineLinkTypeEnum - Types of links to bill payment lines.
type BillPaymentLineLinkTypeEnum string

const (
	BillPaymentLineLinkTypeEnumUnknown          BillPaymentLineLinkTypeEnum = "Unknown"
	BillPaymentLineLinkTypeEnumUnlinked         BillPaymentLineLinkTypeEnum = "Unlinked"
	BillPaymentLineLinkTypeEnumBill             BillPaymentLineLinkTypeEnum = "Bill"
	BillPaymentLineLinkTypeEnumOther            BillPaymentLineLinkTypeEnum = "Other"
	BillPaymentLineLinkTypeEnumCreditNote       BillPaymentLineLinkTypeEnum = "CreditNote"
	BillPaymentLineLinkTypeEnumBillPayment      BillPaymentLineLinkTypeEnum = "BillPayment"
	BillPaymentLineLinkTypeEnumPaymentOnAccount BillPaymentLineLinkTypeEnum = "PaymentOnAccount"
	BillPaymentLineLinkTypeEnumRefund           BillPaymentLineLinkTypeEnum = "Refund"
	BillPaymentLineLinkTypeEnumManualJournal    BillPaymentLineLinkTypeEnum = "ManualJournal"
	BillPaymentLineLinkTypeEnumDiscount         BillPaymentLineLinkTypeEnum = "Discount"
)

func (e *BillPaymentLineLinkTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Unknown":
		fallthrough
	case "Unlinked":
		fallthrough
	case "Bill":
		fallthrough
	case "Other":
		fallthrough
	case "CreditNote":
		fallthrough
	case "BillPayment":
		fallthrough
	case "PaymentOnAccount":
		fallthrough
	case "Refund":
		fallthrough
	case "ManualJournal":
		fallthrough
	case "Discount":
		*e = BillPaymentLineLinkTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for BillPaymentLineLinkTypeEnum: %s", s)
	}
}

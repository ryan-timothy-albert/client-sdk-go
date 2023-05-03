// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DisputeStatusEnum - Current status of the dispute
type DisputeStatusEnum string

const (
	DisputeStatusEnumWon                     DisputeStatusEnum = "Won"
	DisputeStatusEnumLost                    DisputeStatusEnum = "Lost"
	DisputeStatusEnumAccepted                DisputeStatusEnum = "Accepted"
	DisputeStatusEnumProcessing              DisputeStatusEnum = "Processing"
	DisputeStatusEnumChargeRefunded          DisputeStatusEnum = "ChargeRefunded"
	DisputeStatusEnumEvidenceRequired        DisputeStatusEnum = "EvidenceRequired"
	DisputeStatusEnumInquiryEvidenceRequired DisputeStatusEnum = "InquiryEvidenceRequired"
	DisputeStatusEnumInquiryProcessing       DisputeStatusEnum = "InquiryProcessing"
	DisputeStatusEnumInquiryClosed           DisputeStatusEnum = "InquiryClosed"
	DisputeStatusEnumWaitingThirdParty       DisputeStatusEnum = "WaitingThirdParty"
	DisputeStatusEnumUnknown                 DisputeStatusEnum = "Unknown"
)

func (e DisputeStatusEnum) ToPointer() *DisputeStatusEnum {
	return &e
}

func (e *DisputeStatusEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Won":
		fallthrough
	case "Lost":
		fallthrough
	case "Accepted":
		fallthrough
	case "Processing":
		fallthrough
	case "ChargeRefunded":
		fallthrough
	case "EvidenceRequired":
		fallthrough
	case "InquiryEvidenceRequired":
		fallthrough
	case "InquiryProcessing":
		fallthrough
	case "InquiryClosed":
		fallthrough
	case "WaitingThirdParty":
		fallthrough
	case "Unknown":
		*e = DisputeStatusEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DisputeStatusEnum: %v", v)
	}
}

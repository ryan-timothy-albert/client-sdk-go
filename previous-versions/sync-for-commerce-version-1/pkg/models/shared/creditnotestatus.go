// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CreditNoteStatus string

const (
	CreditNoteStatusUnknown       CreditNoteStatus = "Unknown"
	CreditNoteStatusDraft         CreditNoteStatus = "Draft"
	CreditNoteStatusSubmitted     CreditNoteStatus = "Submitted"
	CreditNoteStatusPaid          CreditNoteStatus = "Paid"
	CreditNoteStatusVoid          CreditNoteStatus = "Void"
	CreditNoteStatusPartiallyPaid CreditNoteStatus = "PartiallyPaid"
)

func (e CreditNoteStatus) ToPointer() *CreditNoteStatus {
	return &e
}

func (e *CreditNoteStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unknown":
		fallthrough
	case "Draft":
		fallthrough
	case "Submitted":
		fallthrough
	case "Paid":
		fallthrough
	case "Void":
		fallthrough
	case "PartiallyPaid":
		*e = CreditNoteStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreditNoteStatus: %v", v)
	}
}
// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// PurchaseOrderStatusEnum - Current state of the purchase order
type PurchaseOrderStatusEnum string

const (
	PurchaseOrderStatusEnumUnknown PurchaseOrderStatusEnum = "Unknown"
	PurchaseOrderStatusEnumDraft   PurchaseOrderStatusEnum = "Draft"
	PurchaseOrderStatusEnumOpen    PurchaseOrderStatusEnum = "Open"
	PurchaseOrderStatusEnumClosed  PurchaseOrderStatusEnum = "Closed"
	PurchaseOrderStatusEnumVoid    PurchaseOrderStatusEnum = "Void"
)

func (e *PurchaseOrderStatusEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Unknown":
		fallthrough
	case "Draft":
		fallthrough
	case "Open":
		fallthrough
	case "Closed":
		fallthrough
	case "Void":
		*e = PurchaseOrderStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for PurchaseOrderStatusEnum: %s", s)
	}
}

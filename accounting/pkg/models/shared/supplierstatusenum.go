// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// SupplierStatusEnum - Status of the supplier.
type SupplierStatusEnum string

const (
	SupplierStatusEnumUnknown  SupplierStatusEnum = "Unknown"
	SupplierStatusEnumActive   SupplierStatusEnum = "Active"
	SupplierStatusEnumArchived SupplierStatusEnum = "Archived"
)

func (e *SupplierStatusEnum) UnmarshalJSON(data []byte) error {
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
		*e = SupplierStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SupplierStatusEnum: %s", s)
	}
}

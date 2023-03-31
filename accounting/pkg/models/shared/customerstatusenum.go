// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CustomerStatusEnum - Status of customer.
type CustomerStatusEnum string

const (
	CustomerStatusEnumUnknown  CustomerStatusEnum = "Unknown"
	CustomerStatusEnumActive   CustomerStatusEnum = "Active"
	CustomerStatusEnumArchived CustomerStatusEnum = "Archived"
)

func (e *CustomerStatusEnum) UnmarshalJSON(data []byte) error {
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
		*e = CustomerStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for CustomerStatusEnum: %s", s)
	}
}

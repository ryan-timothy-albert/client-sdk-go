// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// IntegrityStatusEnum - The current status of the most recently run matching algorithm.
type IntegrityStatusEnum string

const (
	IntegrityStatusEnumUnknown      IntegrityStatusEnum = "Unknown"
	IntegrityStatusEnumDoesNotExist IntegrityStatusEnum = "DoesNotExist"
	IntegrityStatusEnumError        IntegrityStatusEnum = "Error"
	IntegrityStatusEnumComplete     IntegrityStatusEnum = "Complete"
)

func (e *IntegrityStatusEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Unknown":
		fallthrough
	case "DoesNotExist":
		fallthrough
	case "Error":
		fallthrough
	case "Complete":
		*e = IntegrityStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for IntegrityStatusEnum: %s", s)
	}
}

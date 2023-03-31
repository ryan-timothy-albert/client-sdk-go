// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// JournalStatusEnum - Current journal status.
type JournalStatusEnum string

const (
	JournalStatusEnumUnknown  JournalStatusEnum = "Unknown"
	JournalStatusEnumActive   JournalStatusEnum = "Active"
	JournalStatusEnumArchived JournalStatusEnum = "Archived"
)

func (e *JournalStatusEnum) UnmarshalJSON(data []byte) error {
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
		*e = JournalStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for JournalStatusEnum: %s", s)
	}
}

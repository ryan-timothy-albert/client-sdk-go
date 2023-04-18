// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type PushChangeTypeEnum string

const (
	PushChangeTypeEnumUnknown            PushChangeTypeEnum = "Unknown"
	PushChangeTypeEnumCreated            PushChangeTypeEnum = "Created"
	PushChangeTypeEnumModified           PushChangeTypeEnum = "Modified"
	PushChangeTypeEnumDeleted            PushChangeTypeEnum = "Deleted"
	PushChangeTypeEnumAttachmentUploaded PushChangeTypeEnum = "AttachmentUploaded"
)

func (e *PushChangeTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Unknown":
		fallthrough
	case "Created":
		fallthrough
	case "Modified":
		fallthrough
	case "Deleted":
		fallthrough
	case "AttachmentUploaded":
		*e = PushChangeTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for PushChangeTypeEnum: %s", s)
	}
}
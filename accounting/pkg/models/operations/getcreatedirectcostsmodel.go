// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetCreateDirectCostsModelRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

func (e *GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Array":
		fallthrough
	case "Object":
		fallthrough
	case "String":
		fallthrough
	case "Number":
		fallthrough
	case "Boolean":
		fallthrough
	case "DateTime":
		fallthrough
	case "File":
		fallthrough
	case "MultiPart":
		*e = GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum: %s", s)
	}
}

type GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                              `json:"description,omitempty"`
	DisplayName *string                                                                              `json:"displayName,omitempty"`
	Required    *bool                                                                                `json:"required,omitempty"`
	Type        *GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	// Allowed value for field.
	Value *string `json:"value,omitempty"`
}

type GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

func (e *GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Array":
		fallthrough
	case "Object":
		fallthrough
	case "String":
		fallthrough
	case "Number":
		fallthrough
	case "Boolean":
		fallthrough
	case "DateTime":
		fallthrough
	case "File":
		fallthrough
	case "MultiPart":
		*e = GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum: %s", s)
	}
}

type GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                `json:"description,omitempty"`
	DisplayName *string                                                                                                `json:"displayName,omitempty"`
	Required    *bool                                                                                                  `json:"required,omitempty"`
	Type        *GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	// Allowed value for field.
	Value *string `json:"value,omitempty"`
}

type GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

func (e *GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Array":
		fallthrough
	case "Object":
		fallthrough
	case "String":
		fallthrough
	case "Number":
		fallthrough
	case "Boolean":
		fallthrough
	case "DateTime":
		fallthrough
	case "File":
		fallthrough
	case "MultiPart":
		*e = GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum: %s", s)
	}
}

type GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                                  `json:"description,omitempty"`
	DisplayName *string                                                                                                                  `json:"displayName,omitempty"`
	Required    *bool                                                                                                                    `json:"required,omitempty"`
	Type        *GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	// Allowed value for field.
	Value *string `json:"value,omitempty"`
}

type GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

func (e *GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Array":
		fallthrough
	case "Object":
		fallthrough
	case "String":
		fallthrough
	case "Number":
		fallthrough
	case "Boolean":
		fallthrough
	case "DateTime":
		fallthrough
	case "File":
		fallthrough
	case "MultiPart":
		*e = GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum: %s", s)
	}
}

type GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                                                    `json:"description,omitempty"`
	DisplayName *string                                                                                                                                    `json:"displayName,omitempty"`
	Required    *bool                                                                                                                                      `json:"required,omitempty"`
	Type        *GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	// Allowed value for field.
	Value *string `json:"value,omitempty"`
}

type GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

func (e *GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Array":
		fallthrough
	case "Object":
		fallthrough
	case "String":
		fallthrough
	case "Number":
		fallthrough
	case "Boolean":
		fallthrough
	case "DateTime":
		fallthrough
	case "File":
		fallthrough
	case "MultiPart":
		*e = GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum: %s", s)
	}
}

type GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                                         `json:"description"`
	DisplayName string                                                                                                                         `json:"displayName"`
	Options     []GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice  `json:"options,omitempty"`
	Required    bool                                                                                                                           `json:"required"`
	Type        GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum      `json:"type"`
	Validation  *GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo `json:"validation,omitempty"`
}

type GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

func (e *GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Array":
		fallthrough
	case "Object":
		fallthrough
	case "String":
		fallthrough
	case "Number":
		fallthrough
	case "Boolean":
		fallthrough
	case "DateTime":
		fallthrough
	case "File":
		fallthrough
	case "MultiPart":
		*e = GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum: %s", s)
	}
}

type GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                                 `json:"description"`
	DisplayName string                                                                                                                 `json:"displayName"`
	Options     []GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                                                   `json:"required"`
	Type        GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

func (e *GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Array":
		fallthrough
	case "Object":
		fallthrough
	case "String":
		fallthrough
	case "Number":
		fallthrough
	case "Boolean":
		fallthrough
	case "DateTime":
		fallthrough
	case "File":
		fallthrough
	case "MultiPart":
		*e = GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum: %s", s)
	}
}

type GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                               `json:"description"`
	DisplayName string                                                                                               `json:"displayName"`
	Options     []GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                                 `json:"required"`
	Type        GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreateDirectCostsModelPushOptionPushOptionPropertyOptionTypeEnum string

const (
	GetCreateDirectCostsModelPushOptionPushOptionPropertyOptionTypeEnumArray     GetCreateDirectCostsModelPushOptionPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyOptionTypeEnumObject    GetCreateDirectCostsModelPushOptionPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyOptionTypeEnumString    GetCreateDirectCostsModelPushOptionPushOptionPropertyOptionTypeEnum = "String"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyOptionTypeEnumNumber    GetCreateDirectCostsModelPushOptionPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyOptionTypeEnumBoolean   GetCreateDirectCostsModelPushOptionPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyOptionTypeEnumDateTime  GetCreateDirectCostsModelPushOptionPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyOptionTypeEnumFile      GetCreateDirectCostsModelPushOptionPushOptionPropertyOptionTypeEnum = "File"
	GetCreateDirectCostsModelPushOptionPushOptionPropertyOptionTypeEnumMultiPart GetCreateDirectCostsModelPushOptionPushOptionPropertyOptionTypeEnum = "MultiPart"
)

func (e *GetCreateDirectCostsModelPushOptionPushOptionPropertyOptionTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Array":
		fallthrough
	case "Object":
		fallthrough
	case "String":
		fallthrough
	case "Number":
		fallthrough
	case "Boolean":
		fallthrough
	case "DateTime":
		fallthrough
	case "File":
		fallthrough
	case "MultiPart":
		*e = GetCreateDirectCostsModelPushOptionPushOptionPropertyOptionTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCreateDirectCostsModelPushOptionPushOptionPropertyOptionTypeEnum: %s", s)
	}
}

type GetCreateDirectCostsModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateDirectCostsModelPushOptionPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateDirectCostsModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateDirectCostsModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateDirectCostsModelPushOptionPushOptionProperty struct {
	Description string                                                                             `json:"description"`
	DisplayName string                                                                             `json:"displayName"`
	Options     []GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreateDirectCostsModelPushOptionPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                               `json:"required"`
	Type        GetCreateDirectCostsModelPushOptionPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreateDirectCostsModelPushOptionPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreateDirectCostsModelPushOptionOptionTypeEnum string

const (
	GetCreateDirectCostsModelPushOptionOptionTypeEnumArray     GetCreateDirectCostsModelPushOptionOptionTypeEnum = "Array"
	GetCreateDirectCostsModelPushOptionOptionTypeEnumObject    GetCreateDirectCostsModelPushOptionOptionTypeEnum = "Object"
	GetCreateDirectCostsModelPushOptionOptionTypeEnumString    GetCreateDirectCostsModelPushOptionOptionTypeEnum = "String"
	GetCreateDirectCostsModelPushOptionOptionTypeEnumNumber    GetCreateDirectCostsModelPushOptionOptionTypeEnum = "Number"
	GetCreateDirectCostsModelPushOptionOptionTypeEnumBoolean   GetCreateDirectCostsModelPushOptionOptionTypeEnum = "Boolean"
	GetCreateDirectCostsModelPushOptionOptionTypeEnumDateTime  GetCreateDirectCostsModelPushOptionOptionTypeEnum = "DateTime"
	GetCreateDirectCostsModelPushOptionOptionTypeEnumFile      GetCreateDirectCostsModelPushOptionOptionTypeEnum = "File"
	GetCreateDirectCostsModelPushOptionOptionTypeEnumMultiPart GetCreateDirectCostsModelPushOptionOptionTypeEnum = "MultiPart"
)

func (e *GetCreateDirectCostsModelPushOptionOptionTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Array":
		fallthrough
	case "Object":
		fallthrough
	case "String":
		fallthrough
	case "Number":
		fallthrough
	case "Boolean":
		fallthrough
	case "DateTime":
		fallthrough
	case "File":
		fallthrough
	case "MultiPart":
		*e = GetCreateDirectCostsModelPushOptionOptionTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCreateDirectCostsModelPushOptionOptionTypeEnum: %s", s)
	}
}

// GetCreateDirectCostsModelPushOption - OK
type GetCreateDirectCostsModelPushOption struct {
	Description *string                                                          `json:"description,omitempty"`
	DisplayName string                                                           `json:"displayName"`
	Properties  map[string]GetCreateDirectCostsModelPushOptionPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                             `json:"required"`
	Type        GetCreateDirectCostsModelPushOptionOptionTypeEnum                `json:"type"`
}

type GetCreateDirectCostsModelResponse struct {
	ContentType string
	// OK
	PushOption  *GetCreateDirectCostsModelPushOption
	StatusCode  int
	RawResponse *http.Response
}

package operations

import (
	"net/http"
)

type GetCreateJournalEntriesModelRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                 `json:"description,omitempty"`
	DisplayName *string                                                                                 `json:"displayName,omitempty"`
	Required    *bool                                                                                   `json:"required,omitempty"`
	Type        *GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                 `json:"value,omitempty"`
}

type GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                   `json:"description,omitempty"`
	DisplayName *string                                                                                                   `json:"displayName,omitempty"`
	Required    *bool                                                                                                     `json:"required,omitempty"`
	Type        *GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                   `json:"value,omitempty"`
}

type GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                                     `json:"description,omitempty"`
	DisplayName *string                                                                                                                     `json:"displayName,omitempty"`
	Required    *bool                                                                                                                       `json:"required,omitempty"`
	Type        *GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                                     `json:"value,omitempty"`
}

type GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                                                       `json:"description,omitempty"`
	DisplayName *string                                                                                                                                       `json:"displayName,omitempty"`
	Required    *bool                                                                                                                                         `json:"required,omitempty"`
	Type        *GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                                                       `json:"value,omitempty"`
}

type GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                                            `json:"description"`
	DisplayName string                                                                                                                            `json:"displayName"`
	Options     []GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice  `json:"options,omitempty"`
	Required    bool                                                                                                                              `json:"required"`
	Type        GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum      `json:"type"`
	Validation  *GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo `json:"validation,omitempty"`
}

type GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                                    `json:"description"`
	DisplayName string                                                                                                                    `json:"displayName"`
	Options     []GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                                                      `json:"required"`
	Type        GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                  `json:"description"`
	DisplayName string                                                                                                  `json:"displayName"`
	Options     []GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                                    `json:"required"`
	Type        GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreateJournalEntriesModelPushOptionPushOptionPropertyOptionTypeEnum string

const (
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyOptionTypeEnumArray     GetCreateJournalEntriesModelPushOptionPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyOptionTypeEnumObject    GetCreateJournalEntriesModelPushOptionPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyOptionTypeEnumString    GetCreateJournalEntriesModelPushOptionPushOptionPropertyOptionTypeEnum = "String"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyOptionTypeEnumNumber    GetCreateJournalEntriesModelPushOptionPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyOptionTypeEnumBoolean   GetCreateJournalEntriesModelPushOptionPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyOptionTypeEnumDateTime  GetCreateJournalEntriesModelPushOptionPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyOptionTypeEnumFile      GetCreateJournalEntriesModelPushOptionPushOptionPropertyOptionTypeEnum = "File"
	GetCreateJournalEntriesModelPushOptionPushOptionPropertyOptionTypeEnumMultiPart GetCreateJournalEntriesModelPushOptionPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateJournalEntriesModelPushOptionPushOptionProperty struct {
	Description string                                                                                `json:"description"`
	DisplayName string                                                                                `json:"displayName"`
	Options     []GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                  `json:"required"`
	Type        GetCreateJournalEntriesModelPushOptionPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreateJournalEntriesModelPushOptionPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreateJournalEntriesModelPushOptionOptionTypeEnum string

const (
	GetCreateJournalEntriesModelPushOptionOptionTypeEnumArray     GetCreateJournalEntriesModelPushOptionOptionTypeEnum = "Array"
	GetCreateJournalEntriesModelPushOptionOptionTypeEnumObject    GetCreateJournalEntriesModelPushOptionOptionTypeEnum = "Object"
	GetCreateJournalEntriesModelPushOptionOptionTypeEnumString    GetCreateJournalEntriesModelPushOptionOptionTypeEnum = "String"
	GetCreateJournalEntriesModelPushOptionOptionTypeEnumNumber    GetCreateJournalEntriesModelPushOptionOptionTypeEnum = "Number"
	GetCreateJournalEntriesModelPushOptionOptionTypeEnumBoolean   GetCreateJournalEntriesModelPushOptionOptionTypeEnum = "Boolean"
	GetCreateJournalEntriesModelPushOptionOptionTypeEnumDateTime  GetCreateJournalEntriesModelPushOptionOptionTypeEnum = "DateTime"
	GetCreateJournalEntriesModelPushOptionOptionTypeEnumFile      GetCreateJournalEntriesModelPushOptionOptionTypeEnum = "File"
	GetCreateJournalEntriesModelPushOptionOptionTypeEnumMultiPart GetCreateJournalEntriesModelPushOptionOptionTypeEnum = "MultiPart"
)

type GetCreateJournalEntriesModelPushOption struct {
	Description *string                                                             `json:"description,omitempty"`
	DisplayName string                                                              `json:"displayName"`
	Properties  map[string]GetCreateJournalEntriesModelPushOptionPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                `json:"required"`
	Type        GetCreateJournalEntriesModelPushOptionOptionTypeEnum                `json:"type"`
}

type GetCreateJournalEntriesModelResponse struct {
	ContentType string
	PushOption  *GetCreateJournalEntriesModelPushOption
	StatusCode  int
	RawResponse *http.Response
}
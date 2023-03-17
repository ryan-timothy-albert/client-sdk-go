package operations

import (
	"net/http"
)

type GetCreateChartOfAccountsModelRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                  `json:"description,omitempty"`
	DisplayName *string                                                                                  `json:"displayName,omitempty"`
	Required    *bool                                                                                    `json:"required,omitempty"`
	Type        *GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                  `json:"value,omitempty"`
}

type GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                    `json:"description,omitempty"`
	DisplayName *string                                                                                                    `json:"displayName,omitempty"`
	Required    *bool                                                                                                      `json:"required,omitempty"`
	Type        *GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                    `json:"value,omitempty"`
}

type GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                                      `json:"description,omitempty"`
	DisplayName *string                                                                                                                      `json:"displayName,omitempty"`
	Required    *bool                                                                                                                        `json:"required,omitempty"`
	Type        *GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                                      `json:"value,omitempty"`
}

type GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                                                        `json:"description,omitempty"`
	DisplayName *string                                                                                                                                        `json:"displayName,omitempty"`
	Required    *bool                                                                                                                                          `json:"required,omitempty"`
	Type        *GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                                                        `json:"value,omitempty"`
}

type GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                                             `json:"description"`
	DisplayName string                                                                                                                             `json:"displayName"`
	Options     []GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice  `json:"options,omitempty"`
	Required    bool                                                                                                                               `json:"required"`
	Type        GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum      `json:"type"`
	Validation  *GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo `json:"validation,omitempty"`
}

type GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                                     `json:"description"`
	DisplayName string                                                                                                                     `json:"displayName"`
	Options     []GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                                                       `json:"required"`
	Type        GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                   `json:"description"`
	DisplayName string                                                                                                   `json:"displayName"`
	Options     []GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                                     `json:"required"`
	Type        GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreateChartOfAccountsModelPushOptionPushOptionPropertyOptionTypeEnum string

const (
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyOptionTypeEnumArray     GetCreateChartOfAccountsModelPushOptionPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyOptionTypeEnumObject    GetCreateChartOfAccountsModelPushOptionPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyOptionTypeEnumString    GetCreateChartOfAccountsModelPushOptionPushOptionPropertyOptionTypeEnum = "String"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyOptionTypeEnumNumber    GetCreateChartOfAccountsModelPushOptionPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyOptionTypeEnumBoolean   GetCreateChartOfAccountsModelPushOptionPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyOptionTypeEnumDateTime  GetCreateChartOfAccountsModelPushOptionPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyOptionTypeEnumFile      GetCreateChartOfAccountsModelPushOptionPushOptionPropertyOptionTypeEnum = "File"
	GetCreateChartOfAccountsModelPushOptionPushOptionPropertyOptionTypeEnumMultiPart GetCreateChartOfAccountsModelPushOptionPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateChartOfAccountsModelPushOptionPushOptionProperty struct {
	Description string                                                                                 `json:"description"`
	DisplayName string                                                                                 `json:"displayName"`
	Options     []GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                   `json:"required"`
	Type        GetCreateChartOfAccountsModelPushOptionPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreateChartOfAccountsModelPushOptionPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreateChartOfAccountsModelPushOptionOptionTypeEnum string

const (
	GetCreateChartOfAccountsModelPushOptionOptionTypeEnumArray     GetCreateChartOfAccountsModelPushOptionOptionTypeEnum = "Array"
	GetCreateChartOfAccountsModelPushOptionOptionTypeEnumObject    GetCreateChartOfAccountsModelPushOptionOptionTypeEnum = "Object"
	GetCreateChartOfAccountsModelPushOptionOptionTypeEnumString    GetCreateChartOfAccountsModelPushOptionOptionTypeEnum = "String"
	GetCreateChartOfAccountsModelPushOptionOptionTypeEnumNumber    GetCreateChartOfAccountsModelPushOptionOptionTypeEnum = "Number"
	GetCreateChartOfAccountsModelPushOptionOptionTypeEnumBoolean   GetCreateChartOfAccountsModelPushOptionOptionTypeEnum = "Boolean"
	GetCreateChartOfAccountsModelPushOptionOptionTypeEnumDateTime  GetCreateChartOfAccountsModelPushOptionOptionTypeEnum = "DateTime"
	GetCreateChartOfAccountsModelPushOptionOptionTypeEnumFile      GetCreateChartOfAccountsModelPushOptionOptionTypeEnum = "File"
	GetCreateChartOfAccountsModelPushOptionOptionTypeEnumMultiPart GetCreateChartOfAccountsModelPushOptionOptionTypeEnum = "MultiPart"
)

type GetCreateChartOfAccountsModelPushOption struct {
	Description *string                                                              `json:"description,omitempty"`
	DisplayName string                                                               `json:"displayName"`
	Properties  map[string]GetCreateChartOfAccountsModelPushOptionPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                 `json:"required"`
	Type        GetCreateChartOfAccountsModelPushOptionOptionTypeEnum                `json:"type"`
}

type GetCreateChartOfAccountsModelResponse struct {
	ContentType string
	PushOption  *GetCreateChartOfAccountsModelPushOption
	StatusCode  int
	RawResponse *http.Response
}

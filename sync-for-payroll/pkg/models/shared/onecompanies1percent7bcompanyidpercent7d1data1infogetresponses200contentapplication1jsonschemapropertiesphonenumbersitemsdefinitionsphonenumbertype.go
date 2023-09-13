// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Onecompanies1Percent7BcompanyIDPercent7D1data1infoGetResponses200ContentApplication1jsonSchemaPropertiesPhoneNumbersItemsDefinitionsPhoneNumberType - The type of phone number
type Onecompanies1Percent7BcompanyIDPercent7D1data1infoGetResponses200ContentApplication1jsonSchemaPropertiesPhoneNumbersItemsDefinitionsPhoneNumberType string

const (
	Onecompanies1Percent7BcompanyIDPercent7D1data1infoGetResponses200ContentApplication1jsonSchemaPropertiesPhoneNumbersItemsDefinitionsPhoneNumberTypePrimary  Onecompanies1Percent7BcompanyIDPercent7D1data1infoGetResponses200ContentApplication1jsonSchemaPropertiesPhoneNumbersItemsDefinitionsPhoneNumberType = "Primary"
	Onecompanies1Percent7BcompanyIDPercent7D1data1infoGetResponses200ContentApplication1jsonSchemaPropertiesPhoneNumbersItemsDefinitionsPhoneNumberTypeLandline Onecompanies1Percent7BcompanyIDPercent7D1data1infoGetResponses200ContentApplication1jsonSchemaPropertiesPhoneNumbersItemsDefinitionsPhoneNumberType = "Landline"
	Onecompanies1Percent7BcompanyIDPercent7D1data1infoGetResponses200ContentApplication1jsonSchemaPropertiesPhoneNumbersItemsDefinitionsPhoneNumberTypeMobile   Onecompanies1Percent7BcompanyIDPercent7D1data1infoGetResponses200ContentApplication1jsonSchemaPropertiesPhoneNumbersItemsDefinitionsPhoneNumberType = "Mobile"
	Onecompanies1Percent7BcompanyIDPercent7D1data1infoGetResponses200ContentApplication1jsonSchemaPropertiesPhoneNumbersItemsDefinitionsPhoneNumberTypeFax      Onecompanies1Percent7BcompanyIDPercent7D1data1infoGetResponses200ContentApplication1jsonSchemaPropertiesPhoneNumbersItemsDefinitionsPhoneNumberType = "Fax"
	Onecompanies1Percent7BcompanyIDPercent7D1data1infoGetResponses200ContentApplication1jsonSchemaPropertiesPhoneNumbersItemsDefinitionsPhoneNumberTypeUnknown  Onecompanies1Percent7BcompanyIDPercent7D1data1infoGetResponses200ContentApplication1jsonSchemaPropertiesPhoneNumbersItemsDefinitionsPhoneNumberType = "Unknown"
)

func (e Onecompanies1Percent7BcompanyIDPercent7D1data1infoGetResponses200ContentApplication1jsonSchemaPropertiesPhoneNumbersItemsDefinitionsPhoneNumberType) ToPointer() *Onecompanies1Percent7BcompanyIDPercent7D1data1infoGetResponses200ContentApplication1jsonSchemaPropertiesPhoneNumbersItemsDefinitionsPhoneNumberType {
	return &e
}

func (e *Onecompanies1Percent7BcompanyIDPercent7D1data1infoGetResponses200ContentApplication1jsonSchemaPropertiesPhoneNumbersItemsDefinitionsPhoneNumberType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Primary":
		fallthrough
	case "Landline":
		fallthrough
	case "Mobile":
		fallthrough
	case "Fax":
		fallthrough
	case "Unknown":
		*e = Onecompanies1Percent7BcompanyIDPercent7D1data1infoGetResponses200ContentApplication1jsonSchemaPropertiesPhoneNumbersItemsDefinitionsPhoneNumberType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Onecompanies1Percent7BcompanyIDPercent7D1data1infoGetResponses200ContentApplication1jsonSchemaPropertiesPhoneNumbersItemsDefinitionsPhoneNumberType: %v", v)
	}
}

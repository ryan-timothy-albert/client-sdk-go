// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// AccountIdentifierTypeEnum - Type of account
type AccountIdentifierTypeEnum string

const (
	AccountIdentifierTypeEnumAccount    AccountIdentifierTypeEnum = "Account"
	AccountIdentifierTypeEnumCard       AccountIdentifierTypeEnum = "Card"
	AccountIdentifierTypeEnumCredit     AccountIdentifierTypeEnum = "Credit"
	AccountIdentifierTypeEnumDepository AccountIdentifierTypeEnum = "Depository"
	AccountIdentifierTypeEnumInvestment AccountIdentifierTypeEnum = "Investment"
	AccountIdentifierTypeEnumLoan       AccountIdentifierTypeEnum = "Loan"
	AccountIdentifierTypeEnumOther      AccountIdentifierTypeEnum = "Other"
)

func (e AccountIdentifierTypeEnum) ToPointer() *AccountIdentifierTypeEnum {
	return &e
}

func (e *AccountIdentifierTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Account":
		fallthrough
	case "Card":
		fallthrough
	case "Credit":
		fallthrough
	case "Depository":
		fallthrough
	case "Investment":
		fallthrough
	case "Loan":
		fallthrough
	case "Other":
		*e = AccountIdentifierTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountIdentifierTypeEnum: %s", s)
	}
}

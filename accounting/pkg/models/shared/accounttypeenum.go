// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// AccountTypeEnum - Type of account
type AccountTypeEnum string

const (
	AccountTypeEnumUnknown   AccountTypeEnum = "Unknown"
	AccountTypeEnumAsset     AccountTypeEnum = "Asset"
	AccountTypeEnumExpense   AccountTypeEnum = "Expense"
	AccountTypeEnumIncome    AccountTypeEnum = "Income"
	AccountTypeEnumLiability AccountTypeEnum = "Liability"
	AccountTypeEnumEquity    AccountTypeEnum = "Equity"
)

func (e *AccountTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Unknown":
		fallthrough
	case "Asset":
		fallthrough
	case "Expense":
		fallthrough
	case "Income":
		fallthrough
	case "Liability":
		fallthrough
	case "Equity":
		*e = AccountTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountTypeEnum: %s", s)
	}
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/codatio/client-sdk-go/bank-feeds/v5/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// BankAccountType - The type of transactions and balances on the account.
// For Credit accounts, positive balances are liabilities, and positive transactions **reduce** liabilities.
// For Debit accounts, positive balances are assets, and positive transactions **increase** assets.
type BankAccountType string

const (
	BankAccountTypeUnknown BankAccountType = "Unknown"
	BankAccountTypeCredit  BankAccountType = "Credit"
	BankAccountTypeDebit   BankAccountType = "Debit"
)

func (e BankAccountType) ToPointer() *BankAccountType {
	return &e
}

func (e *BankAccountType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unknown":
		fallthrough
	case "Credit":
		fallthrough
	case "Debit":
		*e = BankAccountType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BankAccountType: %v", v)
	}
}

type BankAccountPrototype struct {
	// Name of the bank account in the accounting platform.
	AccountName *string `json:"accountName,omitempty"`
	// Account number for the bank account.
	//
	// Xero integrations
	// Only a UK account number shows for bank accounts with GBP currency and a combined total of sort code and account number that equals 14 digits, For non-GBP accounts, the full bank account number is populated.
	//
	// FreeAgent integrations
	// For Credit accounts, only the last four digits are required. For other types, the field is optional.
	AccountNumber *string `json:"accountNumber,omitempty"`
	// The type of transactions and balances on the account.
	// For Credit accounts, positive balances are liabilities, and positive transactions **reduce** liabilities.
	// For Debit accounts, positive balances are assets, and positive transactions **increase** assets.
	AccountType *BankAccountType `json:"accountType,omitempty"`
	// Total available balance of the bank account as reported by the underlying data source. This may take into account overdrafts or pending transactions for example.
	AvailableBalance *decimal.Big `decimal:"number" json:"availableBalance,omitempty"`
	// Balance of the bank account.
	Balance *decimal.Big `decimal:"number" json:"balance,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
	// International bank account number of the account. Often used when making or receiving international payments.
	IBan *string `json:"iBan,omitempty"`
	// The institution of the bank account.
	Institution *string `json:"institution,omitempty"`
	// Code used to identify each nominal account for a business.
	NominalCode *string `json:"nominalCode,omitempty"`
	// Pre-arranged overdraft limit of the account.
	//
	// The value is always positive. For example, an overdraftLimit of `1000` means that the balance of the account can go down to `-1000`.
	OverdraftLimit *decimal.Big `decimal:"number" json:"overdraftLimit,omitempty"`
	// Sort code for the bank account.
	//
	// Xero integrations
	// The sort code is only displayed when the currency = GBP and the sort code and account number sum to 14 digits. For non-GBP accounts, this field is not populated.
	SortCode *string `json:"sortCode,omitempty"`
}

func (b BankAccountPrototype) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BankAccountPrototype) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BankAccountPrototype) GetAccountName() *string {
	if o == nil {
		return nil
	}
	return o.AccountName
}

func (o *BankAccountPrototype) GetAccountNumber() *string {
	if o == nil {
		return nil
	}
	return o.AccountNumber
}

func (o *BankAccountPrototype) GetAccountType() *BankAccountType {
	if o == nil {
		return nil
	}
	return o.AccountType
}

func (o *BankAccountPrototype) GetAvailableBalance() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.AvailableBalance
}

func (o *BankAccountPrototype) GetBalance() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Balance
}

func (o *BankAccountPrototype) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *BankAccountPrototype) GetIBan() *string {
	if o == nil {
		return nil
	}
	return o.IBan
}

func (o *BankAccountPrototype) GetInstitution() *string {
	if o == nil {
		return nil
	}
	return o.Institution
}

func (o *BankAccountPrototype) GetNominalCode() *string {
	if o == nil {
		return nil
	}
	return o.NominalCode
}

func (o *BankAccountPrototype) GetOverdraftLimit() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.OverdraftLimit
}

func (o *BankAccountPrototype) GetSortCode() *string {
	if o == nil {
		return nil
	}
	return o.SortCode
}

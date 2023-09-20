// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v4/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// BankingTransaction - The Banking Transactions data type provides an immutable source of up-to-date information on income and expenditure.
//
// Responses are paged, so you should provide `page` and `pageSize` query parameters in your request.
//
// View the coverage for banking transactions in the [Data Coverage Explorer](https://knowledge.codat.io/supported-features/banking?view=tab-by-data-type&dataType=banking-transactions).
type BankingTransaction struct {
	// The unique identifier of the bank account.
	AccountID string `json:"accountId"`
	// The amount of the bank transaction.
	Amount *decimal.Big `decimal:"number" json:"amount,omitempty"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	AuthorizedDate *string `json:"authorizedDate,omitempty"`
	// Code to identify the underlying transaction.
	Code *TransactionCode `json:"code,omitempty"`
	// The currency of the bank transaction.
	Currency string `json:"currency"`
	// The description of the bank transaction.
	Description *string `json:"description,omitempty"`
	// The unique identifier of the bank transaction.
	ID string `json:"id"`
	// The name of the merchant.
	MerchantName *string `json:"merchantName,omitempty"`
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	PostedDate         *string `json:"postedDate,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// An object of bank transaction category reference data.
	TransactionCategoryRef *TransactionCategoryRef `json:"transactionCategoryRef,omitempty"`
}

func (b BankingTransaction) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BankingTransaction) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BankingTransaction) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *BankingTransaction) GetAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *BankingTransaction) GetAuthorizedDate() *string {
	if o == nil {
		return nil
	}
	return o.AuthorizedDate
}

func (o *BankingTransaction) GetCode() *TransactionCode {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *BankingTransaction) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *BankingTransaction) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *BankingTransaction) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *BankingTransaction) GetMerchantName() *string {
	if o == nil {
		return nil
	}
	return o.MerchantName
}

func (o *BankingTransaction) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *BankingTransaction) GetPostedDate() *string {
	if o == nil {
		return nil
	}
	return o.PostedDate
}

func (o *BankingTransaction) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *BankingTransaction) GetTransactionCategoryRef() *TransactionCategoryRef {
	if o == nil {
		return nil
	}
	return o.TransactionCategoryRef
}

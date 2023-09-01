// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Account - > **Language tip:** Accounts are also referred to as **chart of accounts**, **nominal accounts**, and **general ledger**.
//
// View the coverage for accounts in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=chartOfAccounts" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Accounts are the categories a business uses to record accounting transactions. From the Accounts endpoints, you can retrieve a list of all accounts for a specified company.
//
// The categories for an account include:
//   - Asset
//   - Expense
//   - Income
//   - Liability
//   - Equity.
//
// The same account may have a different category based on the integration it is used in. For example, a current account (known as checking in the US) should be categorized as `Asset.Current` for Xero, and `Asset.Bank.Checking` for QuickBooks Online.
//
// At the same time, each integration may have its own requirements to the categories. For example, a Paypal account in Xero is of the `Asset.Bank` category and therefore requires additional properties to be provided.
//
// To determine the list of allowed categories for a specific integration, you can:
// - Follow our [Create, update, delete data](https://docs.codat.io/using-the-api/push) guide and use the [Get create account model](https://docs.codat.io/accounting-api#/operations/get-create-chartOfAccounts-model).
// - Refer to the integration's own documentation.
//
// > **Accounts with no category**
// >
// > If an account is pulled from the chart of accounts and its nominal code does not lie within the category layout for the company's accounts, then the **type** is `Unknown`. The **fullyQualifiedCategory** and **fullyQualifiedName** fields return `null`.
// >
// > This approach gives a true representation of the company's accounts whilst preventing distorting financials such as a company's profit and loss and balance sheet reports.
type Account struct {
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
	// Current balance in the account.
	CurrentBalance *float64 `json:"currentBalance,omitempty"`
	// Description for the account.
	Description *string `json:"description,omitempty"`
	// Full category of the account.
	//
	// For example, `Liability.Current` or `Income.Revenue`. To determine a list of possible categories for each integration, see our examples, follow our [Create, update, delete data](https://docs.codat.io/using-the-api/push) guide, or refer to the integration's own documentation.
	FullyQualifiedCategory *string `json:"fullyQualifiedCategory,omitempty"`
	// Full name of the account, for example:
	// - `Cash On Hand`
	// - `Rents Held In Trust`
	// - `Fixed Asset`
	FullyQualifiedName *string `json:"fullyQualifiedName,omitempty"`
	// Identifier for the account, unique for the company.
	ID *string `json:"id,omitempty"`
	// Confirms whether the account is a bank account or not.
	IsBankAccount *bool     `json:"isBankAccount,omitempty"`
	Metadata      *Metadata `json:"metadata,omitempty"`
	ModifiedDate  *string   `json:"modifiedDate,omitempty"`
	// Name of the account.
	Name *string `json:"name,omitempty"`
	// Reference given to each nominal account for a business. It ensures money is allocated to the correct account. This code isn't a unique identifier in the Codat system.
	NominalCode        *string `json:"nominalCode,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Status of the account
	Status *AccountStatus `json:"status,omitempty"`
	// Type of account
	Type *AccountType `json:"type,omitempty"`
	// The validDatatypeLinks can be used to determine whether an account can be correctly mapped to another object; for example, accounts with a `type` of `income` might only support being used on an Invoice and Direct Income. For more information, see [Valid Data Type Links](/accounting-api#/schemas/ValidDataTypeLinks).
	ValidDatatypeLinks []ValidDataTypeLinks `json:"validDatatypeLinks,omitempty"`
}

func (o *Account) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *Account) GetCurrentBalance() *float64 {
	if o == nil {
		return nil
	}
	return o.CurrentBalance
}

func (o *Account) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Account) GetFullyQualifiedCategory() *string {
	if o == nil {
		return nil
	}
	return o.FullyQualifiedCategory
}

func (o *Account) GetFullyQualifiedName() *string {
	if o == nil {
		return nil
	}
	return o.FullyQualifiedName
}

func (o *Account) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Account) GetIsBankAccount() *bool {
	if o == nil {
		return nil
	}
	return o.IsBankAccount
}

func (o *Account) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *Account) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *Account) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Account) GetNominalCode() *string {
	if o == nil {
		return nil
	}
	return o.NominalCode
}

func (o *Account) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *Account) GetStatus() *AccountStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *Account) GetType() *AccountType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Account) GetValidDatatypeLinks() []ValidDataTypeLinks {
	if o == nil {
		return nil
	}
	return o.ValidDatatypeLinks
}

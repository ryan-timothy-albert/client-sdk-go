// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AccountingCreateBankTransactions struct {
	AccountID    *string                        `json:"accountId,omitempty"`
	Transactions []CreateBankAccountTransaction `json:"transactions,omitempty"`
}

func (o *AccountingCreateBankTransactions) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *AccountingCreateBankTransactions) GetTransactions() []CreateBankAccountTransaction {
	if o == nil {
		return nil
	}
	return o.Transactions
}

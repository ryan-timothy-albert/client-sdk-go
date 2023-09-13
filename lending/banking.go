// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package codatlending

type banking struct {
	AccountBalances       *bankingAccountBalances
	Accounts              *bankingAccounts
	CategorizedStatement  *bankingCategorizedStatement
	TransactionCategories *bankingTransactionCategories
	Transactions          *bankingTransactions

	sdkConfiguration sdkConfiguration
}

func newBanking(sdkConfig sdkConfiguration) *banking {
	return &banking{
		sdkConfiguration:      sdkConfig,
		AccountBalances:       newBankingAccountBalances(sdkConfig),
		Accounts:              newBankingAccounts(sdkConfig),
		CategorizedStatement:  newBankingCategorizedStatement(sdkConfig),
		TransactionCategories: newBankingTransactionCategories(sdkConfig),
		Transactions:          newBankingTransactions(sdkConfig),
	}
}
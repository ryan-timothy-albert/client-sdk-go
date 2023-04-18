// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// BankAccountTransactions - Success
type BankAccountTransactions struct {
	Links        Links                 `json:"_links"`
	PageNumber   int64                 `json:"pageNumber"`
	PageSize     int64                 `json:"pageSize"`
	Results      []BankTransactionLine `json:"results,omitempty"`
	TotalResults int64                 `json:"totalResults"`
}
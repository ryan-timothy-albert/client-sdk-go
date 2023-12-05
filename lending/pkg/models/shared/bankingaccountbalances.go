// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type BankingAccountBalances struct {
	Links Links `json:"_links"`
	// Current page number.
	PageNumber int64 `json:"pageNumber"`
	// Number of items to return in results array.
	PageSize int64                   `json:"pageSize"`
	Results  []BankingAccountBalance `json:"results,omitempty"`
	// Total number of items.
	TotalResults int64 `json:"totalResults"`
}

func (o *BankingAccountBalances) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *BankingAccountBalances) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *BankingAccountBalances) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *BankingAccountBalances) GetResults() []BankingAccountBalance {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *BankingAccountBalances) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}

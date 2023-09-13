// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AccountingAccountTransactions struct {
	Links        Links                          `json:"_links"`
	PageNumber   int64                          `json:"pageNumber"`
	PageSize     int64                          `json:"pageSize"`
	Results      []AccountingAccountTransaction `json:"results,omitempty"`
	TotalResults int64                          `json:"totalResults"`
}

func (o *AccountingAccountTransactions) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *AccountingAccountTransactions) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *AccountingAccountTransactions) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *AccountingAccountTransactions) GetResults() []AccountingAccountTransaction {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *AccountingAccountTransactions) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}

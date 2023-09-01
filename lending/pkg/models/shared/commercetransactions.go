// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CommerceTransactions - OK
type CommerceTransactions struct {
	Links        Links                 `json:"_links"`
	PageNumber   int64                 `json:"pageNumber"`
	PageSize     int64                 `json:"pageSize"`
	Results      []CommerceTransaction `json:"results,omitempty"`
	TotalResults int64                 `json:"totalResults"`
}

func (o *CommerceTransactions) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *CommerceTransactions) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *CommerceTransactions) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *CommerceTransactions) GetResults() []CommerceTransaction {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *CommerceTransactions) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}
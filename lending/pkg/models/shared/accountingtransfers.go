// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AccountingTransfers struct {
	Links        Links                `json:"_links"`
	PageNumber   int64                `json:"pageNumber"`
	PageSize     int64                `json:"pageSize"`
	Results      []AccountingTransfer `json:"results,omitempty"`
	TotalResults int64                `json:"totalResults"`
}

func (o *AccountingTransfers) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *AccountingTransfers) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *AccountingTransfers) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *AccountingTransfers) GetResults() []AccountingTransfer {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *AccountingTransfers) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}

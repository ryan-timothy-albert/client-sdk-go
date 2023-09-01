// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AccountingPayments - Success
type AccountingPayments struct {
	Links        Links               `json:"_links"`
	PageNumber   int64               `json:"pageNumber"`
	PageSize     int64               `json:"pageSize"`
	Results      []AccountingPayment `json:"results,omitempty"`
	TotalResults int64               `json:"totalResults"`
}

func (o *AccountingPayments) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *AccountingPayments) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *AccountingPayments) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *AccountingPayments) GetResults() []AccountingPayment {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *AccountingPayments) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}
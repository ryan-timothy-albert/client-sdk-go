// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PaymentMethods - Success
type PaymentMethods struct {
	Links        Links           `json:"_links"`
	PageNumber   int64           `json:"pageNumber"`
	PageSize     int64           `json:"pageSize"`
	Results      []PaymentMethod `json:"results,omitempty"`
	TotalResults int64           `json:"totalResults"`
}

func (o *PaymentMethods) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *PaymentMethods) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *PaymentMethods) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *PaymentMethods) GetResults() []PaymentMethod {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *PaymentMethods) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}

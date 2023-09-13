// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Accounts struct {
	Links        Links     `json:"_links"`
	PageNumber   int64     `json:"pageNumber"`
	PageSize     int64     `json:"pageSize"`
	Results      []Account `json:"results,omitempty"`
	TotalResults int64     `json:"totalResults"`
}

func (o *Accounts) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *Accounts) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *Accounts) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *Accounts) GetResults() []Account {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *Accounts) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Integrations - Success
type Integrations struct {
	Links        Links         `json:"_links"`
	PageNumber   int64         `json:"pageNumber"`
	PageSize     int64         `json:"pageSize"`
	Results      []Integration `json:"results,omitempty"`
	TotalResults int64         `json:"totalResults"`
}

func (o *Integrations) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *Integrations) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *Integrations) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *Integrations) GetResults() []Integration {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *Integrations) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}
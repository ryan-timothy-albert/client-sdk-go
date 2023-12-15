// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ItemReceipts struct {
	Links Links `json:"_links"`
	// Current page number.
	PageNumber int64 `json:"pageNumber"`
	// Number of items to return in results array.
	PageSize int64         `json:"pageSize"`
	Results  []ItemReceipt `json:"results,omitempty"`
	// Total number of items.
	TotalResults int64 `json:"totalResults"`
}

func (o *ItemReceipts) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *ItemReceipts) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *ItemReceipts) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *ItemReceipts) GetResults() []ItemReceipt {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *ItemReceipts) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}
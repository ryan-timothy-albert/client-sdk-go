// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Suppliers - Success
type Suppliers struct {
	Links        Links      `json:"_links"`
	PageNumber   int64      `json:"pageNumber"`
	PageSize     int64      `json:"pageSize"`
	Results      []Supplier `json:"results,omitempty"`
	TotalResults int64      `json:"totalResults"`
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreditNotes - Success
type CreditNotes struct {
	Links        Links        `json:"_links"`
	PageNumber   int64        `json:"pageNumber"`
	PageSize     int64        `json:"pageSize"`
	Results      []CreditNote `json:"results,omitempty"`
	TotalResults int64        `json:"totalResults"`
}
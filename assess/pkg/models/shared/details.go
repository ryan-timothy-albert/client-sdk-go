// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Details - OK
type Details struct {
	Links        Links                  `json:"_links"`
	PageNumber   int64                  `json:"pageNumber"`
	PageSize     int64                  `json:"pageSize"`
	Results      []DataIntegrityDetails `json:"results,omitempty"`
	TotalResults int64                  `json:"totalResults"`
}
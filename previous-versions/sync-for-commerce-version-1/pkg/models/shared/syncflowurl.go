// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SyncFlowURL - Success
type SyncFlowURL struct {
	// Sync flow URL.
	URL *string `json:"url,omitempty"`
}

func (o *SyncFlowURL) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}
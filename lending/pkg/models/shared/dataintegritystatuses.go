// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// DataIntegrityStatuses - OK
type DataIntegrityStatuses struct {
	Metadata []DataIntegrityStatus `json:"metadata,omitempty"`
}

func (o *DataIntegrityStatuses) GetMetadata() []DataIntegrityStatus {
	if o == nil {
		return nil
	}
	return o.Metadata
}
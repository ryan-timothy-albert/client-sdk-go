// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RecordRef struct {
	// identifier of linked reference from mapping options.
	ID *string `json:"id,omitempty"`
}

func (o *RecordRef) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
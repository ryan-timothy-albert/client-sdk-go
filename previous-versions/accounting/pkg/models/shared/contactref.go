// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ContactRef - A customer or supplier associated with the direct cost.
type ContactRef struct {
	// Available Data types
	DataType *DataType `json:"dataType,omitempty"`
	// Unique identifier for a customer or supplier.
	ID string `json:"id"`
}

func (o *ContactRef) GetDataType() *DataType {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *ContactRef) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

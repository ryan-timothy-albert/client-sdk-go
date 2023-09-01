// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PushOperationRef struct {
	// Available Data types
	DataType *DataType `json:"dataType,omitempty"`
	ID       *string   `json:"id,omitempty"`
}

func (o *PushOperationRef) GetDataType() *DataType {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *PushOperationRef) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

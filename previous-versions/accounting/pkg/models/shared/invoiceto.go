// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// InvoiceTo - Links the current record to the underlying record or data type that created it.
//
// For example, if a journal entry is generated based on an invoice, this property allows you to connect the journal entry to the underlying invoice in our data model.
type InvoiceTo struct {
	// Allowed name of the 'dataType'.
	DataType *string `json:"dataType,omitempty"`
	// 'id' of the underlying record or data type.
	ID *string `json:"id,omitempty"`
}

func (o *InvoiceTo) GetDataType() *string {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *InvoiceTo) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

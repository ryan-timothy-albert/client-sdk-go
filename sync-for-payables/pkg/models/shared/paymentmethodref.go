// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PaymentMethodRef - The Payment Method to which the payment is linked in the accounting platform.
type PaymentMethodRef struct {
	// The unique identifier of the location being referenced.
	ID string `json:"id"`
	// Name of the location being referenced.
	Name *string `json:"name,omitempty"`
}

func (o *PaymentMethodRef) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PaymentMethodRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
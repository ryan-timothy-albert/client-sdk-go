// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PhoneNumberItems struct {
	// A phone number.
	Number *string `json:"number"`
	// The type of phone number
	Type PhoneNumberType `json:"type"`
}

func (o *PhoneNumberItems) GetNumber() *string {
	if o == nil {
		return nil
	}
	return o.Number
}

func (o *PhoneNumberItems) GetType() PhoneNumberType {
	if o == nil {
		return PhoneNumberType("")
	}
	return o.Type
}
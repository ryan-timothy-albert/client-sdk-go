// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Items struct {
	// A phone number.
	Number string `json:"number"`
	// The type of phone number
	Type PhoneNumberType `json:"type"`
}

func (o *Items) GetNumber() string {
	if o == nil {
		return ""
	}
	return o.Number
}

func (o *Items) GetType() PhoneNumberType {
	if o == nil {
		return PhoneNumberType("")
	}
	return o.Type
}
// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type WithholdingTaxitems struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

func (o *WithholdingTaxitems) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}

func (o *WithholdingTaxitems) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ExpenseTransactionLine struct {
	AccountRef RecordRef `json:"accountRef"`
	// Amount of the line, exclusive of tax.
	NetAmount float64 `json:"netAmount"`
	// Amount of tax for the line.
	TaxAmount    float64     `json:"taxAmount"`
	TaxRateRef   *RecordRef  `json:"taxRateRef,omitempty"`
	TrackingRefs []RecordRef `json:"trackingRefs,omitempty"`
}

func (o *ExpenseTransactionLine) GetAccountRef() RecordRef {
	if o == nil {
		return RecordRef{}
	}
	return o.AccountRef
}

func (o *ExpenseTransactionLine) GetNetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.NetAmount
}

func (o *ExpenseTransactionLine) GetTaxAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.TaxAmount
}

func (o *ExpenseTransactionLine) GetTaxRateRef() *RecordRef {
	if o == nil {
		return nil
	}
	return o.TaxRateRef
}

func (o *ExpenseTransactionLine) GetTrackingRefs() []RecordRef {
	if o == nil {
		return nil
	}
	return o.TrackingRefs
}
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
// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ReportLine struct {
	// Identifier for the account, unique for the company in the accounting platform.
	AccountID *string `json:"accountId,omitempty"`
	// An array of ReportLine items.
	Items []ReportLine `json:"items,omitempty"`
	// Name of the report line item.
	Name *string `json:"name,omitempty"`
	// Numerical value of the line item.
	Value float64 `json:"value"`
}

func (o *ReportLine) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *ReportLine) GetItems() []ReportLine {
	if o == nil {
		return nil
	}
	return o.Items
}

func (o *ReportLine) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ReportLine) GetValue() float64 {
	if o == nil {
		return 0.0
	}
	return o.Value
}

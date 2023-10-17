// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type ReportLine struct {
	// Identifier for the account, unique for the company in the accounting platform.
	AccountID *string `json:"accountId,omitempty"`
	// An array of ReportLine items.
	Items []ReportLine `json:"items,omitempty"`
	// Name of the report line item.
	Name *string `json:"name,omitempty"`
	// Numerical value of the line item.
	Value *decimal.Big `decimal:"number" json:"value"`
}

func (r ReportLine) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *ReportLine) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
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

func (o *ReportLine) GetValue() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.Value
}

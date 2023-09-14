// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/codatio/client-sdk-go/lending/v3/pkg/types"
)

type ReportItemsLoanTransactionType string

const (
	ReportItemsLoanTransactionTypeInvestment      ReportItemsLoanTransactionType = "Investment"
	ReportItemsLoanTransactionTypeRepayment       ReportItemsLoanTransactionType = "Repayment"
	ReportItemsLoanTransactionTypeInterest        ReportItemsLoanTransactionType = "Interest"
	ReportItemsLoanTransactionTypeAccuredInterest ReportItemsLoanTransactionType = "AccuredInterest"
)

func (e ReportItemsLoanTransactionType) ToPointer() *ReportItemsLoanTransactionType {
	return &e
}

func (e *ReportItemsLoanTransactionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Investment":
		fallthrough
	case "Repayment":
		fallthrough
	case "Interest":
		fallthrough
	case "AccuredInterest":
		*e = ReportItemsLoanTransactionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReportItemsLoanTransactionType: %v", v)
	}
}

type ReportItems struct {
	// The loan transaction amount.
	Amount *types.Decimal `json:"amount,omitempty"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	Date                *string                         `json:"date,omitempty"`
	ItemRef             *DefinitionsitemRef             `json:"itemRef,omitempty"`
	LoanRef             *LoanRef                        `json:"loanRef,omitempty"`
	LoanTransactionType *ReportItemsLoanTransactionType `json:"loanTransactionType,omitempty"`
}

func (o *ReportItems) GetAmount() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *ReportItems) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *ReportItems) GetItemRef() *DefinitionsitemRef {
	if o == nil {
		return nil
	}
	return o.ItemRef
}

func (o *ReportItems) GetLoanRef() *LoanRef {
	if o == nil {
		return nil
	}
	return o.LoanRef
}

func (o *ReportItems) GetLoanTransactionType() *ReportItemsLoanTransactionType {
	if o == nil {
		return nil
	}
	return o.LoanTransactionType
}

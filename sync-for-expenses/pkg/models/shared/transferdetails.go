// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/sync-for-expenses/v4/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type TransferDetails struct {
	// Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.
	AccountRef *AccountRef `json:"accountRef,omitempty"`
	// Amount of the transfer.
	Amount *decimal.Big `decimal:"number" json:"amount,omitempty"`
}

func (t TransferDetails) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TransferDetails) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TransferDetails) GetAccountRef() *AccountRef {
	if o == nil {
		return nil
	}
	return o.AccountRef
}

func (o *TransferDetails) GetAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Amount
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreateReimbursableExpenseRequest struct {
	Items [][]ReimbursableExpenseTransaction `json:"items,omitempty"`
}

func (o *CreateReimbursableExpenseRequest) GetItems() [][]ReimbursableExpenseTransaction {
	if o == nil {
		return nil
	}
	return o.Items
}

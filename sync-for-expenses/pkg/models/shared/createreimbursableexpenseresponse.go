// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreateReimbursableExpenseResponse struct {
	// Unique id of sync created
	SyncID *string `json:"syncId,omitempty"`
}

func (o *CreateReimbursableExpenseResponse) GetSyncID() *string {
	if o == nil {
		return nil
	}
	return o.SyncID
}

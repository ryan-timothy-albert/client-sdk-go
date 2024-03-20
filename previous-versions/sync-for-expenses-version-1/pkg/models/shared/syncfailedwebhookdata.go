// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SyncFailedWebhookData struct {
	// The stage of the job the sync failed.
	FailureStage *string `json:"FailureStage,omitempty"`
	// Unique identifier for the failed sync.
	SyncID *string `json:"syncId,omitempty"`
	// The type of sync being performed.
	SyncType *string `json:"syncType,omitempty"`
}

func (o *SyncFailedWebhookData) GetFailureStage() *string {
	if o == nil {
		return nil
	}
	return o.FailureStage
}

func (o *SyncFailedWebhookData) GetSyncID() *string {
	if o == nil {
		return nil
	}
	return o.SyncID
}

func (o *SyncFailedWebhookData) GetSyncType() *string {
	if o == nil {
		return nil
	}
	return o.SyncType
}

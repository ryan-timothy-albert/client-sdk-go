// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SyncConfiguration struct {
	SyncAsBankFeeds *SyncAsBankFeeds `json:"syncAsBankFeeds,omitempty"`
	SyncAsExpenses  *SyncAsExpenses  `json:"syncAsExpenses,omitempty"`
}

func (o *SyncConfiguration) GetSyncAsBankFeeds() *SyncAsBankFeeds {
	if o == nil {
		return nil
	}
	return o.SyncAsBankFeeds
}

func (o *SyncConfiguration) GetSyncAsExpenses() *SyncAsExpenses {
	if o == nil {
		return nil
	}
	return o.SyncAsExpenses
}

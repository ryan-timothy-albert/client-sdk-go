// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// JournalRef - Links journal entries to the relevant journal in accounting integrations that use multi-book accounting (multiple journals).
type JournalRef struct {
	// GUID of the underlying journal.
	ID string `json:"id"`
	// Name of journal
	Name *string `json:"name,omitempty"`
}

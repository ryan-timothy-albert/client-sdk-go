// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PushOperationTimedOutWebhookData struct {
	// Data type used in the push operation.
	DataType *string `json:"dataType,omitempty"`
	// Unique identifier for the push operation.
	PushOperationKey *string `json:"pushOperationKey,omitempty"`
}

// PushOperationTimedOutWebhook - Webhook request body notifying that a push push operation has timed out.
type PushOperationTimedOutWebhook struct {
	// Unique identifier of the alert.
	AlertID *string `json:"alertId,omitempty"`
	// Unique identifier for your SMB in Codat.
	CompanyID *string                           `json:"companyId,omitempty"`
	Data      *PushOperationTimedOutWebhookData `json:"data,omitempty"`
	// A human readable message about the webhook.
	Message *string `json:"message,omitempty"`
	// Unique identifier for the rule.
	RuleID *string `json:"ruleId,omitempty"`
	// The type of rule.
	Type *string `json:"type,omitempty"`
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DataSyncCompletedWebhookData struct {
	// Data type the sync completed for.
	DataType *string `json:"dataType,omitempty"`
	// Unique identifier for the dataset that completed its sync.
	DatasetID *string `json:"datasetId,omitempty"`
}

// DataSyncCompletedWebhook - Webhook request body to notify the completion of a data sync.
type DataSyncCompletedWebhook struct {
	// Unique identifier of the alert.
	AlertID *string `json:"alertId,omitempty"`
	// Unique identifier for your client in Codat.
	ClientID *string `json:"clientId,omitempty"`
	// Name of your client in Codat.
	ClientName *string `json:"clientName,omitempty"`
	// Unique identifier for your SMB in Codat.
	CompanyID *string                       `json:"companyId,omitempty"`
	Data      *DataSyncCompletedWebhookData `json:"data,omitempty"`
	// Unique identifier for a company's data connection.
	DataConnectionID *string `json:"dataConnectionId,omitempty"`
	// A human readable message about the webhook.
	Message *string `json:"message,omitempty"`
	// Unique identifier for the rule.
	RuleID *string `json:"ruleId,omitempty"`
	// The type of rule.
	Type *string `json:"type,omitempty"`
}

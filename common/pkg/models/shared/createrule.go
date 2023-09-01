// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreateRuleNotifiers struct {
	Emails  []string `json:"emails,omitempty"`
	Webhook *string  `json:"webhook,omitempty"`
}

// CreateRule - Create an event notification to a URL or list of email addresses based on the given type or condition.
type CreateRule struct {
	CompanyID *string             `json:"companyId,omitempty"`
	Notifiers CreateRuleNotifiers `json:"notifiers"`
	Type      string              `json:"type"`
}
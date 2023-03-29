// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package codatio

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
	"github.com/codatio/client-sdk-go/common/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// Production
	"https://api.codat.io",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Codatio - Common API
// An API for the common components of all of Codat's products.
//
// These end points cover creating and managing your companies, data connections, and integrations.
//
// [Read about the building blocks of Codat...](https://docs.codat.io/core-concepts/companies)
//
// [See our OpenAPI spec](https://github.com/codatio/oas)
type Codatio struct {
	// Companies - Create and manage your Codat companies.
	Companies *companies
	// Connections - Manage your companies' data connections.
	Connections *connections
	// DataStatus - Understand the state of data within Codat.
	DataStatus *dataStatus
	// Integrations - View and manage your available integrations in Codat.
	Integrations *integrations
	// PushData - View push options and get push statuses.
	PushData *pushData
	// RefreshData - Queue pull operations to refresh data in Codat.
	RefreshData *refreshData
	// Settings - Manage your Codat instance.
	Settings *settings
	// Webhooks - Manage webhooks, rules and alerts.
	Webhooks *webhooks

	// Non-idiomatic field names below are to namespace fields from the fields names above to avoid name conflicts
	_defaultClient  HTTPClient
	_securityClient HTTPClient
	_security       *shared.Security
	_serverURL      string
	_language       string
	_sdkVersion     string
	_genVersion     string
}

type SDKOption func(*Codatio)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Codatio) {
		sdk._serverURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Codatio) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk._serverURL = serverURL
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Codatio) {
		sdk._defaultClient = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *Codatio) {
		sdk._security = &security
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Codatio {
	sdk := &Codatio{
		_language:   "go",
		_sdkVersion: "0.7.0",
		_genVersion: "2.16.5",
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk._defaultClient == nil {
		sdk._defaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk._securityClient == nil {
		if sdk._security != nil {
			sdk._securityClient = utils.ConfigureSecurityClient(sdk._defaultClient, sdk._security)
		} else {
			sdk._securityClient = sdk._defaultClient
		}
	}

	if sdk._serverURL == "" {
		sdk._serverURL = ServerList[0]
	}

	sdk.Companies = newCompanies(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Connections = newConnections(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.DataStatus = newDataStatus(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Integrations = newIntegrations(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.PushData = newPushData(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.RefreshData = newRefreshData(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Settings = newSettings(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Webhooks = newWebhooks(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	return sdk
}

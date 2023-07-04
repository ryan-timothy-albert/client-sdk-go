// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package codatsyncexpenses

import (
	"fmt"
	"github.com/codatio/client-sdk-go/expenses/pkg/models/shared"
	"github.com/codatio/client-sdk-go/expenses/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https://api.codat.io",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          *shared.Security
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// CodatSyncExpenses - Sync for Expenses API: The API for Sync for Expenses.
// Sync for Expenses is an API and a set of supporting tools. It has been built to enable corporate card and expense management platforms to provide high-quality integrations with multiple accounting platforms through a standardized API.
//
// [Read more...](https://docs.codat.io/sync-for-expenses/overview)
//
// [See our OpenAPI spec](https://github.com/codatio/oas)
type CodatSyncExpenses struct {
	// Configuration - Companies sync configuration.
	Configuration *configuration
	// Connections - Create and manage partner expense connection.
	Connections *connections
	// Expenses - Create expense datasets and upload receipts.
	Expenses *expenses
	// MappingOptions - Mapping options for a companies expenses.
	MappingOptions *mappingOptions
	// Sync - Triggering a new sync of expenses to accounting software.
	Sync *sync
	// SyncStatus - Check the status of ongoing or previous expense syncs.
	SyncStatus *syncStatus
	// TransactionStatus - Retrieve the status of transactions within a sync.
	TransactionStatus *transactionStatus

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*CodatSyncExpenses)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *CodatSyncExpenses) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *CodatSyncExpenses) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *CodatSyncExpenses) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *CodatSyncExpenses) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *CodatSyncExpenses) {
		sdk.sdkConfiguration.Security = &security
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *CodatSyncExpenses {
	sdk := &CodatSyncExpenses{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "prealpha",
			SDKVersion:        "0.22.0",
			GenVersion:        "2.55.0",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.Configuration = newConfiguration(sdk.sdkConfiguration)

	sdk.Connections = newConnections(sdk.sdkConfiguration)

	sdk.Expenses = newExpenses(sdk.sdkConfiguration)

	sdk.MappingOptions = newMappingOptions(sdk.sdkConfiguration)

	sdk.Sync = newSync(sdk.sdkConfiguration)

	sdk.SyncStatus = newSyncStatus(sdk.sdkConfiguration)

	sdk.TransactionStatus = newTransactionStatus(sdk.sdkConfiguration)

	return sdk
}

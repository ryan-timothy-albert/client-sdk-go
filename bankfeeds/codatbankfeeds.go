// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package codatbankfeeds

import (
	"fmt"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/utils"
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
	DefaultClient  HTTPClient
	SecurityClient HTTPClient
	Security       *shared.Security
	ServerURL      string
	ServerIndex    int
	Language       string
	SDKVersion     string
	GenVersion     string
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// CodatBankFeeds - Bank Feeds API: Bank Feeds API enables your SMB users to set up bank feeds from accounts in your application to supported accounting platforms.
//
// A bank feed is a connection between a source bank account—in your application—and a target bank account in a supported accounting package.
//
// [Read more...](https://docs.codat.io/bank-feeds-api/overview)
//
// [See our OpenAPI spec](https://github.com/codatio/oas)
type CodatBankFeeds struct {
	// BankAccountTransactions - Bank feed bank accounts
	BankAccountTransactions *bankAccountTransactions
	// BankFeedAccounts - Bank feed bank accounts
	BankFeedAccounts *bankFeedAccounts
	// Companies - Create and manage your Codat companies.
	Companies *companies
	// Connections - Manage your companies' data connections.
	Connections *connections

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*CodatBankFeeds)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *CodatBankFeeds) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *CodatBankFeeds) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *CodatBankFeeds) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *CodatBankFeeds) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *CodatBankFeeds) {
		sdk.sdkConfiguration.Security = &security
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *CodatBankFeeds {
	sdk := &CodatBankFeeds{
		sdkConfiguration: sdkConfiguration{
			Language:   "go",
			SDKVersion: "0.18.0",
			GenVersion: "2.35.9",
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

	sdk.BankAccountTransactions = newBankAccountTransactions(sdk.sdkConfiguration)

	sdk.BankFeedAccounts = newBankFeedAccounts(sdk.sdkConfiguration)

	sdk.Companies = newCompanies(sdk.sdkConfiguration)

	sdk.Connections = newConnections(sdk.sdkConfiguration)

	return sdk
}

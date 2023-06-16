// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package codatsynccommerce

import (
	"fmt"
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/shared"
	"github.com/codatio/client-sdk-go/synccommerce/pkg/utils"
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

// CodatSyncCommerce - Sync for Commerce API: The API for Sync for Commerce. Sync for Commerce is an API and a set of supporting tools. It has been built to enable e-commerce, point of sale platforms to provide high-quality integrations with numerous accounting platform through standardized API, seamlessly transforming business sale's data into accounting artefacts.
// [Read More...](https://docs.codat.io/sfc/overview)
type CodatSyncCommerce struct {
	// CompanyManagement - Create new and manage existing Sync for Commerce companies.
	CompanyManagement *companyManagement
	// Configuration - Expressively configure preferences for any given Sync for Commerce company.
	Configuration *configuration
	// Integrations - View useful information about codat's integrations.
	Integrations *integrations
	// Sync - Initiate a sync of Sync for Commerce company data into their respective accounting software.
	Sync *sync
	// SyncFlowPreferences - Configure preferences for any given Sync for Commerce company using sync flow.
	SyncFlowPreferences *syncFlowPreferences

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*CodatSyncCommerce)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *CodatSyncCommerce) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *CodatSyncCommerce) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *CodatSyncCommerce) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *CodatSyncCommerce) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *CodatSyncCommerce) {
		sdk.sdkConfiguration.Security = &security
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *CodatSyncCommerce {
	sdk := &CodatSyncCommerce{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.1",
			SDKVersion:        "0.18.0",
			GenVersion:        "2.41.1",
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

	sdk.CompanyManagement = newCompanyManagement(sdk.sdkConfiguration)

	sdk.Configuration = newConfiguration(sdk.sdkConfiguration)

	sdk.Integrations = newIntegrations(sdk.sdkConfiguration)

	sdk.Sync = newSync(sdk.sdkConfiguration)

	sdk.SyncFlowPreferences = newSyncFlowPreferences(sdk.sdkConfiguration)

	return sdk
}

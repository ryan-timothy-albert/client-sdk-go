// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package codataccounting

import (
	"fmt"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/utils"
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

// CodatAccounting - Accounting API: A flexible API for pulling accounting data, normalized and aggregated from 20 accounting integrations.
//
// Standardize how you connect to your customers’ accounting software. View, create, update, and delete data in the same way for all the leading accounting platforms.
//
// [Read more...](https://docs.codat.io/accounting-api/overview)
//
// [See our OpenAPI spec](https://github.com/codatio/oas)
type CodatAccounting struct {
	// AccountTransactions - Account transactions
	AccountTransactions *accountTransactions
	// Accounts - Accounts
	Accounts *accounts
	// BankAccountTransactions - Bank transactions for bank accounts
	BankAccountTransactions *bankAccountTransactions
	// BankAccounts - Bank accounts
	BankAccounts *bankAccounts
	// BillCreditNotes - Bill credit notes
	BillCreditNotes *billCreditNotes
	// BillPayments - Bill payments
	BillPayments *billPayments
	// Bills - Bills
	Bills *bills
	// CompanyInfo - Company info
	CompanyInfo *companyInfo
	// CreditNotes - Credit notes
	CreditNotes *creditNotes
	// Customers - Customers
	Customers *customers
	// DirectCosts - Direct costs
	DirectCosts *directCosts
	// DirectIncomes - Direct incomes
	DirectIncomes *directIncomes
	// Invoices - Invoices
	Invoices *invoices
	// Items - Items
	Items *items
	// JournalEntries - Journal entries
	JournalEntries *journalEntries
	// Journals - Journals
	Journals *journals
	// PaymentMethods - Payment methods
	PaymentMethods *paymentMethods
	// Payments - Payments
	Payments *payments
	// PurchaseOrders - Purchase orders
	PurchaseOrders *purchaseOrders
	// Reports - Reports
	Reports *reports
	// SalesOrders - Sales orders
	SalesOrders *salesOrders
	// Suppliers - Suppliers
	Suppliers *suppliers
	// TaxRates - Tax rates
	TaxRates *taxRates
	// TrackingCategories - Tracking categories
	TrackingCategories *trackingCategories
	// Transfers - Transfers
	Transfers *transfers

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*CodatAccounting)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *CodatAccounting) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *CodatAccounting) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *CodatAccounting) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *CodatAccounting) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *CodatAccounting) {
		sdk.sdkConfiguration.Security = &security
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *CodatAccounting {
	sdk := &CodatAccounting{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "2.1.0",
			SDKVersion:        "0.25.0",
			GenVersion:        "2.58.0",
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

	sdk.AccountTransactions = newAccountTransactions(sdk.sdkConfiguration)

	sdk.Accounts = newAccounts(sdk.sdkConfiguration)

	sdk.BankAccountTransactions = newBankAccountTransactions(sdk.sdkConfiguration)

	sdk.BankAccounts = newBankAccounts(sdk.sdkConfiguration)

	sdk.BillCreditNotes = newBillCreditNotes(sdk.sdkConfiguration)

	sdk.BillPayments = newBillPayments(sdk.sdkConfiguration)

	sdk.Bills = newBills(sdk.sdkConfiguration)

	sdk.CompanyInfo = newCompanyInfo(sdk.sdkConfiguration)

	sdk.CreditNotes = newCreditNotes(sdk.sdkConfiguration)

	sdk.Customers = newCustomers(sdk.sdkConfiguration)

	sdk.DirectCosts = newDirectCosts(sdk.sdkConfiguration)

	sdk.DirectIncomes = newDirectIncomes(sdk.sdkConfiguration)

	sdk.Invoices = newInvoices(sdk.sdkConfiguration)

	sdk.Items = newItems(sdk.sdkConfiguration)

	sdk.JournalEntries = newJournalEntries(sdk.sdkConfiguration)

	sdk.Journals = newJournals(sdk.sdkConfiguration)

	sdk.PaymentMethods = newPaymentMethods(sdk.sdkConfiguration)

	sdk.Payments = newPayments(sdk.sdkConfiguration)

	sdk.PurchaseOrders = newPurchaseOrders(sdk.sdkConfiguration)

	sdk.Reports = newReports(sdk.sdkConfiguration)

	sdk.SalesOrders = newSalesOrders(sdk.sdkConfiguration)

	sdk.Suppliers = newSuppliers(sdk.sdkConfiguration)

	sdk.TaxRates = newTaxRates(sdk.sdkConfiguration)

	sdk.TrackingCategories = newTrackingCategories(sdk.sdkConfiguration)

	sdk.Transfers = newTransfers(sdk.sdkConfiguration)

	return sdk
}

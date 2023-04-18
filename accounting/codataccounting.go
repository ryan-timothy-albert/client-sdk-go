// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package codataccounting

import (
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

// CodatAccounting - A flexible API for pulling accounting data, normalized and aggregated from 20 accounting integrations.
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
	// Financials - Financials
	Financials *financials
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

	// Non-idiomatic field names below are to namespace fields from the fields names above to avoid name conflicts
	_defaultClient  HTTPClient
	_securityClient HTTPClient
	_security       *shared.Security
	_serverURL      string
	_language       string
	_sdkVersion     string
	_genVersion     string
}

type SDKOption func(*CodatAccounting)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *CodatAccounting) {
		sdk._serverURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *CodatAccounting) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk._serverURL = serverURL
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *CodatAccounting) {
		sdk._defaultClient = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *CodatAccounting) {
		sdk._security = &security
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *CodatAccounting {
	sdk := &CodatAccounting{
		_language:   "go",
		_sdkVersion: "0.9.1",
		_genVersion: "2.18.1",
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

	sdk.AccountTransactions = newAccountTransactions(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Accounts = newAccounts(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.BankAccountTransactions = newBankAccountTransactions(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.BankAccounts = newBankAccounts(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.BillCreditNotes = newBillCreditNotes(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.BillPayments = newBillPayments(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Bills = newBills(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.CompanyInfo = newCompanyInfo(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.CreditNotes = newCreditNotes(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Customers = newCustomers(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.DirectCosts = newDirectCosts(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.DirectIncomes = newDirectIncomes(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Financials = newFinancials(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Invoices = newInvoices(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Items = newItems(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.JournalEntries = newJournalEntries(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Journals = newJournals(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.PaymentMethods = newPaymentMethods(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Payments = newPayments(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.PurchaseOrders = newPurchaseOrders(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Reports = newReports(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.SalesOrders = newSalesOrders(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Suppliers = newSuppliers(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.TaxRates = newTaxRates(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.TrackingCategories = newTrackingCategories(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Transfers = newTransfers(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	return sdk
}

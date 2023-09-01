// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type BillPaymentAllocationAllocation struct {
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	AllocatedOnDate *string `json:"allocatedOnDate,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
	// Rate to convert the total amount of the payment into the base currency for the company at the time of the payment.
	//
	// Currency rates in Codat are implemented as the multiple of foreign currency units to each base currency unit.
	//
	// It is not possible to perform the currency conversion with two or more non-base currencies participating in the transaction. For example, if a company's base currency is USD, and it has a bill issued in EUR, then the bill payment must happen in USD or EUR.
	//
	// Where the currency rate is provided by the underlying accounting platform, it will be available from Codat with the same precision (up to a maximum of 9 decimal places).
	//
	// For accounting platforms which do not provide an explicit currency rate, it is calculated as `baseCurrency / foreignCurrency` and will be returned to 9 decimal places.
	//
	// ## Examples with base currency of GBP
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (GBP) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **USD**          | $20            | 0.781         | £15.62                     |
	// | **EUR**          | €20            | 0.885         | £17.70                     |
	// | **RUB**          | ₽20            | 0.011         | £0.22                      |
	//
	// ## Examples with base currency of USD
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (USD) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **GBP**          | £20            | 1.277         | $25.54                     |
	// | **EUR**          | €20            | 1.134         | $22.68                     |
	// | **RUB**          | ₽20            | 0.015         | $0.30                      |
	CurrencyRate *float64 `json:"currencyRate,omitempty"`
	// The total amount that has been allocated.
	TotalAmount *float64 `json:"totalAmount,omitempty"`
}

func (o *BillPaymentAllocationAllocation) GetAllocatedOnDate() *string {
	if o == nil {
		return nil
	}
	return o.AllocatedOnDate
}

func (o *BillPaymentAllocationAllocation) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *BillPaymentAllocationAllocation) GetCurrencyRate() *float64 {
	if o == nil {
		return nil
	}
	return o.CurrencyRate
}

func (o *BillPaymentAllocationAllocation) GetTotalAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

type BillPaymentAllocation struct {
	Allocation BillPaymentAllocationAllocation `json:"allocation"`
	Payment    PaymentAllocationPayment        `json:"payment"`
}

func (o *BillPaymentAllocation) GetAllocation() BillPaymentAllocationAllocation {
	if o == nil {
		return BillPaymentAllocationAllocation{}
	}
	return o.Allocation
}

func (o *BillPaymentAllocation) GetPayment() PaymentAllocationPayment {
	if o == nil {
		return PaymentAllocationPayment{}
	}
	return o.Payment
}

// BillSupplementalData - Supplemental data is additional data you can include in our standard data types.
//
// It is referenced as a configured dynamic key value pair that is unique to the accounting platform. [Learn more](https://docs.codat.io/using-the-api/additional-data) about supplemental data.
type BillSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

func (o *BillSupplementalData) GetContent() map[string]map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Content
}

type BillWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

func (o *BillWithholdingTax) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}

func (o *BillWithholdingTax) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// Bill - > **Invoices or bills?**
// >
// > We distinguish between invoices where the company *owes money* vs. *is owed money*. If the company has received an invoice, and owes money to someone else (accounts payable) we call this a Bill.
// >
// > See [Invoices](https://docs.codat.io/accounting-api#/schemas/Invoice) for the accounts receivable equivalent of bills.
//
// View the coverage for bills in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// In Codat, a bill contains details of:
// * When the bill was recorded in the accounting system.
// * How much the bill is for and the currency of the amount.
// * Who the bill was received from — the *supplier*.
// * What the bill is for — the *line items*.
//
// Some accounting platforms give a separate name to purchases where the payment is made immediately, such as something bought with a credit card or online payment. One example of this would be QuickBooks Online's *expenses*.
//
// You can find these types of transactions in our [Direct costs](https://docs.codat.io/accounting-api#/schemas/DirectCost) data model.
type Bill struct {
	// Amount outstanding on the bill.
	AmountDue *float64 `json:"amountDue,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
	// Rate to convert the total amount of the payment into the base currency for the company at the time of the payment.
	//
	// Currency rates in Codat are implemented as the multiple of foreign currency units to each base currency unit.
	//
	// It is not possible to perform the currency conversion with two or more non-base currencies participating in the transaction. For example, if a company's base currency is USD, and it has a bill issued in EUR, then the bill payment must happen in USD or EUR.
	//
	// Where the currency rate is provided by the underlying accounting platform, it will be available from Codat with the same precision (up to a maximum of 9 decimal places).
	//
	// For accounting platforms which do not provide an explicit currency rate, it is calculated as `baseCurrency / foreignCurrency` and will be returned to 9 decimal places.
	//
	// ## Examples with base currency of GBP
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (GBP) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **USD**          | $20            | 0.781         | £15.62                     |
	// | **EUR**          | €20            | 0.885         | £17.70                     |
	// | **RUB**          | ₽20            | 0.011         | £0.22                      |
	//
	// ## Examples with base currency of USD
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (USD) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **GBP**          | £20            | 1.277         | $25.54                     |
	// | **EUR**          | €20            | 1.134         | $22.68                     |
	// | **RUB**          | ₽20            | 0.015         | $0.30                      |
	CurrencyRate *float64 `json:"currencyRate,omitempty"`
	DueDate      *string  `json:"dueDate,omitempty"`
	// Identifier for the bill, unique for the company in the accounting platform.
	ID        *string `json:"id,omitempty"`
	IssueDate string  `json:"issueDate"`
	// Array of Bill line items.
	LineItems    []BillLineItem `json:"lineItems,omitempty"`
	Metadata     *Metadata      `json:"metadata,omitempty"`
	ModifiedDate *string        `json:"modifiedDate,omitempty"`
	// Any private, company notes about the bill, such as payment information.
	Note *string `json:"note,omitempty"`
	// An array of payment allocations.
	PaymentAllocations []BillPaymentAllocation `json:"paymentAllocations,omitempty"`
	PurchaseOrderRefs  []PurchaseOrderRef      `json:"purchaseOrderRefs,omitempty"`
	// User-friendly reference for the bill.
	Reference          *string `json:"reference,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Current state of the bill.
	Status BillStatus `json:"status"`
	// Total amount of the bill, excluding any taxes.
	SubTotal float64 `json:"subTotal"`
	// Supplemental data is additional data you can include in our standard data types.
	//
	// It is referenced as a configured dynamic key value pair that is unique to the accounting platform. [Learn more](https://docs.codat.io/using-the-api/additional-data) about supplemental data.
	SupplementalData *BillSupplementalData `json:"supplementalData,omitempty"`
	// Reference to the supplier the record relates to.
	SupplierRef *SupplierRef `json:"supplierRef,omitempty"`
	// Amount of tax on the bill.
	TaxAmount float64 `json:"taxAmount"`
	// Amount of the bill, including tax.
	TotalAmount    float64              `json:"totalAmount"`
	WithholdingTax []BillWithholdingTax `json:"withholdingTax,omitempty"`
}

func (o *Bill) GetAmountDue() *float64 {
	if o == nil {
		return nil
	}
	return o.AmountDue
}

func (o *Bill) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *Bill) GetCurrencyRate() *float64 {
	if o == nil {
		return nil
	}
	return o.CurrencyRate
}

func (o *Bill) GetDueDate() *string {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *Bill) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Bill) GetIssueDate() string {
	if o == nil {
		return ""
	}
	return o.IssueDate
}

func (o *Bill) GetLineItems() []BillLineItem {
	if o == nil {
		return nil
	}
	return o.LineItems
}

func (o *Bill) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *Bill) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *Bill) GetNote() *string {
	if o == nil {
		return nil
	}
	return o.Note
}

func (o *Bill) GetPaymentAllocations() []BillPaymentAllocation {
	if o == nil {
		return nil
	}
	return o.PaymentAllocations
}

func (o *Bill) GetPurchaseOrderRefs() []PurchaseOrderRef {
	if o == nil {
		return nil
	}
	return o.PurchaseOrderRefs
}

func (o *Bill) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *Bill) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *Bill) GetStatus() BillStatus {
	if o == nil {
		return BillStatus("")
	}
	return o.Status
}

func (o *Bill) GetSubTotal() float64 {
	if o == nil {
		return 0.0
	}
	return o.SubTotal
}

func (o *Bill) GetSupplementalData() *BillSupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}

func (o *Bill) GetSupplierRef() *SupplierRef {
	if o == nil {
		return nil
	}
	return o.SupplierRef
}

func (o *Bill) GetTaxAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.TaxAmount
}

func (o *Bill) GetTotalAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.TotalAmount
}

func (o *Bill) GetWithholdingTax() []BillWithholdingTax {
	if o == nil {
		return nil
	}
	return o.WithholdingTax
}

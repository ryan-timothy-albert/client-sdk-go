// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AccountingCreditNote - > View the coverage for credit notes in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=creditNotes" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Think of a credit note as a voucher issued to a customer. It is a reduction that can be applied against one or multiple invoices. A credit note can either reduce the amount owed or cancel out an invoice entirely.
//
// In the Codat system a credit note is issued to a [customer's](https://docs.codat.io/accounting-api#/schemas/Customer) accounts receivable.
//
// It contains details of:
// * The amount of credit remaining and its status.
// * Payment allocations against the payments type, in this case an invoice.
// * Which customers the credit notes have been issued to.
type AccountingCreditNote struct {
	AdditionalTaxAmount     *float64 `json:"additionalTaxAmount,omitempty"`
	AdditionalTaxPercentage *float64 `json:"additionalTaxPercentage,omitempty"`
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
	// Friendly reference for the credit note.
	CreditNoteNumber *string `json:"creditNoteNumber,omitempty"`
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
	CurrencyRate *float64               `json:"currencyRate,omitempty"`
	CustomerRef  *AccountingCustomerRef `json:"customerRef,omitempty"`
	// Percentage rate (from 0 to 100) of discounts applied to the credit note.
	DiscountPercentage float64 `json:"discountPercentage"`
	// Identifier for the credit note, unique to the company in the accounting platform.
	ID *string `json:"id,omitempty"`
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
	IssueDate    *string              `json:"issueDate,omitempty"`
	LineItems    []CreditNoteLineItem `json:"lineItems,omitempty"`
	Metadata     *Metadata            `json:"metadata,omitempty"`
	ModifiedDate *string              `json:"modifiedDate,omitempty"`
	// Any additional information about the credit note. Where possible, Codat links to a data field in the accounting platform that is publicly available. This means that the contents of the note field are included when a credit note is emailed from the accounting platform to the customer.
	Note *string `json:"note,omitempty"`
	// An array of payment allocations.
	PaymentAllocations []PaymentAllocationsitems `json:"paymentAllocations,omitempty"`
	// Unused balance of totalAmount originally raised.
	RemainingCredit    float64          `json:"remainingCredit"`
	SourceModifiedDate *string          `json:"sourceModifiedDate,omitempty"`
	Status             CreditNoteStatus `json:"status"`
	// Value of the credit note, including discounts and excluding tax.
	SubTotal float64 `json:"subTotal"`
	// Supplemental data is additional data you can include in our standard data types.
	//
	// It is referenced as a configured dynamic key value pair that is unique to the accounting platform. [Learn more](https://docs.codat.io/using-the-api/additional-data) about supplemental data.
	SupplementalData *SupplementalData `json:"supplementalData,omitempty"`
	// Total amount of credit that has been applied to the customer's accounts receivable
	TotalAmount float64 `json:"totalAmount"`
	// Any discounts applied to the credit note amount.
	TotalDiscount float64 `json:"totalDiscount"`
	// Any tax applied to the credit note amount.
	TotalTaxAmount float64               `json:"totalTaxAmount"`
	WithholdingTax []WithholdingTaxitems `json:"withholdingTax,omitempty"`
}

func (o *AccountingCreditNote) GetAdditionalTaxAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.AdditionalTaxAmount
}

func (o *AccountingCreditNote) GetAdditionalTaxPercentage() *float64 {
	if o == nil {
		return nil
	}
	return o.AdditionalTaxPercentage
}

func (o *AccountingCreditNote) GetAllocatedOnDate() *string {
	if o == nil {
		return nil
	}
	return o.AllocatedOnDate
}

func (o *AccountingCreditNote) GetCreditNoteNumber() *string {
	if o == nil {
		return nil
	}
	return o.CreditNoteNumber
}

func (o *AccountingCreditNote) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *AccountingCreditNote) GetCurrencyRate() *float64 {
	if o == nil {
		return nil
	}
	return o.CurrencyRate
}

func (o *AccountingCreditNote) GetCustomerRef() *AccountingCustomerRef {
	if o == nil {
		return nil
	}
	return o.CustomerRef
}

func (o *AccountingCreditNote) GetDiscountPercentage() float64 {
	if o == nil {
		return 0.0
	}
	return o.DiscountPercentage
}

func (o *AccountingCreditNote) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingCreditNote) GetIssueDate() *string {
	if o == nil {
		return nil
	}
	return o.IssueDate
}

func (o *AccountingCreditNote) GetLineItems() []CreditNoteLineItem {
	if o == nil {
		return nil
	}
	return o.LineItems
}

func (o *AccountingCreditNote) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *AccountingCreditNote) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *AccountingCreditNote) GetNote() *string {
	if o == nil {
		return nil
	}
	return o.Note
}

func (o *AccountingCreditNote) GetPaymentAllocations() []PaymentAllocationsitems {
	if o == nil {
		return nil
	}
	return o.PaymentAllocations
}

func (o *AccountingCreditNote) GetRemainingCredit() float64 {
	if o == nil {
		return 0.0
	}
	return o.RemainingCredit
}

func (o *AccountingCreditNote) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *AccountingCreditNote) GetStatus() CreditNoteStatus {
	if o == nil {
		return CreditNoteStatus("")
	}
	return o.Status
}

func (o *AccountingCreditNote) GetSubTotal() float64 {
	if o == nil {
		return 0.0
	}
	return o.SubTotal
}

func (o *AccountingCreditNote) GetSupplementalData() *SupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}

func (o *AccountingCreditNote) GetTotalAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.TotalAmount
}

func (o *AccountingCreditNote) GetTotalDiscount() float64 {
	if o == nil {
		return 0.0
	}
	return o.TotalDiscount
}

func (o *AccountingCreditNote) GetTotalTaxAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.TotalTaxAmount
}

func (o *AccountingCreditNote) GetWithholdingTax() []WithholdingTaxitems {
	if o == nil {
		return nil
	}
	return o.WithholdingTax
}

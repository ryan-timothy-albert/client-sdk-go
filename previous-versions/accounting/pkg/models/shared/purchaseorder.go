// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// PurchaseOrder - > View the coverage for purchase orders in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=purchaseOrders" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Purchase orders represent a business's intent to purchase goods or services from a supplier and normally include information such as expected delivery dates and shipping details.
//
// This information can be used to provide visibility on a business's expected payables and to track a purchase through the full procurement process.
type PurchaseOrder struct {
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
	//
	//
	// ### Integration-specific details
	//
	// | Integration       | Scenario                                        | System behavior                                                                                                                                                      |
	// |-------------------|-------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	// | QuickBooks Online | Transaction currency differs from base currency | If currency rate value is left `null`, a rate of 1 will be used by QBO by default. To override this, include the required currency rate in the expense transaction.  |
	CurrencyRate *decimal.Big `decimal:"number" json:"currencyRate,omitempty"`
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
	DeliveryDate *string `json:"deliveryDate,omitempty"`
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
	ExpectedDeliveryDate *string `json:"expectedDeliveryDate,omitempty"`
	// Identifier for the purchase order, unique for the company in the accounting platform.
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
	IssueDate *string `json:"issueDate,omitempty"`
	// Array of line items.
	LineItems    []PurchaseOrderLineItem `json:"lineItems,omitempty"`
	Metadata     *Metadata               `json:"metadata,omitempty"`
	ModifiedDate *string                 `json:"modifiedDate,omitempty"`
	// Any additional information associated with the purchase order.
	Note *string `json:"note,omitempty"`
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
	PaymentDueDate *string `json:"paymentDueDate,omitempty"`
	// Friendly reference for the purchase order, commonly generated by the accounting platform.
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
	// Delivery details for any goods that have been ordered.
	ShipTo             *ShipTo `json:"shipTo,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Current state of the purchase order
	Status *PurchaseOrderStatus `json:"status,omitempty"`
	// Total amount of the purchase order, including discounts but excluding tax.
	SubTotal *decimal.Big `decimal:"number" json:"subTotal,omitempty"`
	// Reference to the supplier the record relates to.
	SupplierRef *SupplierRef `json:"supplierRef,omitempty"`
	// Total amount of the purchase order, including discounts and tax.
	TotalAmount *decimal.Big `decimal:"number" json:"totalAmount,omitempty"`
	// Total value of any discounts applied to the purchase order.
	TotalDiscount *decimal.Big `decimal:"number" json:"totalDiscount,omitempty"`
	//
	// Total amount of tax included in the purchase order.
	TotalTaxAmount *decimal.Big `decimal:"number" json:"totalTaxAmount,omitempty"`
}

func (p PurchaseOrder) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PurchaseOrder) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PurchaseOrder) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *PurchaseOrder) GetCurrencyRate() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.CurrencyRate
}

func (o *PurchaseOrder) GetDeliveryDate() *string {
	if o == nil {
		return nil
	}
	return o.DeliveryDate
}

func (o *PurchaseOrder) GetExpectedDeliveryDate() *string {
	if o == nil {
		return nil
	}
	return o.ExpectedDeliveryDate
}

func (o *PurchaseOrder) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PurchaseOrder) GetIssueDate() *string {
	if o == nil {
		return nil
	}
	return o.IssueDate
}

func (o *PurchaseOrder) GetLineItems() []PurchaseOrderLineItem {
	if o == nil {
		return nil
	}
	return o.LineItems
}

func (o *PurchaseOrder) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *PurchaseOrder) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *PurchaseOrder) GetNote() *string {
	if o == nil {
		return nil
	}
	return o.Note
}

func (o *PurchaseOrder) GetPaymentDueDate() *string {
	if o == nil {
		return nil
	}
	return o.PaymentDueDate
}

func (o *PurchaseOrder) GetPurchaseOrderNumber() *string {
	if o == nil {
		return nil
	}
	return o.PurchaseOrderNumber
}

func (o *PurchaseOrder) GetShipTo() *ShipTo {
	if o == nil {
		return nil
	}
	return o.ShipTo
}

func (o *PurchaseOrder) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *PurchaseOrder) GetStatus() *PurchaseOrderStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *PurchaseOrder) GetSubTotal() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.SubTotal
}

func (o *PurchaseOrder) GetSupplierRef() *SupplierRef {
	if o == nil {
		return nil
	}
	return o.SupplierRef
}

func (o *PurchaseOrder) GetTotalAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *PurchaseOrder) GetTotalDiscount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TotalDiscount
}

func (o *PurchaseOrder) GetTotalTaxAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TotalTaxAmount
}

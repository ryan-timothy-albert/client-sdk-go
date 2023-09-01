// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// BillCreditNoteLineItemItemReference - Reference to the item the line is linked to.
type BillCreditNoteLineItemItemReference struct {
	// Unique identifier for the item in the accounting platform.
	ID string `json:"id"`
	// Name of the item in the accounting platform.
	Name *string `json:"name,omitempty"`
}

func (o *BillCreditNoteLineItemItemReference) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *BillCreditNoteLineItemItemReference) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// BillCreditNoteLineItemTaxRateReference - Data types that reference a tax rate, for example invoice and bill line items, use a taxRateRef that includes the ID and name of the linked tax rate.
//
// Found on:
//
// - Bill line items
// - Bill Credit Note line items
// - Credit Note line items
// - Direct incomes line items
// - Invoice line items
// - Items
type BillCreditNoteLineItemTaxRateReference struct {
	// Applicable tax rate.
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	// Unique identifier for the tax rate in the accounting platform.
	ID *string `json:"id,omitempty"`
	// Name of the tax rate in the accounting platform.
	Name *string `json:"name,omitempty"`
}

func (o *BillCreditNoteLineItemTaxRateReference) GetEffectiveTaxRate() *float64 {
	if o == nil {
		return nil
	}
	return o.EffectiveTaxRate
}

func (o *BillCreditNoteLineItemTaxRateReference) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BillCreditNoteLineItemTaxRateReference) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type BillCreditNoteLineItem struct {
	// Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.
	AccountRef *AccountRef `json:"accountRef,omitempty"`
	// Friendly name of each line item. For example, the goods or service for which credit has been received.
	Description *string `json:"description,omitempty"`
	// Value of any discounts applied.
	DiscountAmount *float64 `json:"discountAmount,omitempty"`
	// Percentage rate of any discount applied to the line item.
	DiscountPercentage *float64 `json:"discountPercentage,omitempty"`
	// Reference to the item the line is linked to.
	ItemRef *BillCreditNoteLineItemItemReference `json:"itemRef,omitempty"`
	// Number of units of the goods or service for which credit has been received.
	Quantity float64 `json:"quantity"`
	// Amount of credit associated with the line item, including discounts but excluding tax.
	SubTotal *float64 `json:"subTotal,omitempty"`
	// Amount of tax associated with the line item.
	TaxAmount *float64 `json:"taxAmount,omitempty"`
	// Data types that reference a tax rate, for example invoice and bill line items, use a taxRateRef that includes the ID and name of the linked tax rate.
	//
	// Found on:
	//
	// - Bill line items
	// - Bill Credit Note line items
	// - Credit Note line items
	// - Direct incomes line items
	// - Invoice line items
	// - Items
	TaxRateRef *BillCreditNoteLineItemTaxRateReference `json:"taxRateRef,omitempty"`
	// Total amount of the line item, including discounts and tax.
	TotalAmount *float64 `json:"totalAmount,omitempty"`
	// Categories, and a project and customer, against which the item is tracked.
	Tracking *AccountsPayableTracking `json:"tracking,omitempty"`
	// Reference to the tracking categories to which the line item is linked.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	TrackingCategoryRefs []TrackingCategoryRef `json:"trackingCategoryRefs,omitempty"`
	// Unit price of the goods or service.
	UnitAmount float64 `json:"unitAmount"`
}

func (o *BillCreditNoteLineItem) GetAccountRef() *AccountRef {
	if o == nil {
		return nil
	}
	return o.AccountRef
}

func (o *BillCreditNoteLineItem) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *BillCreditNoteLineItem) GetDiscountAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.DiscountAmount
}

func (o *BillCreditNoteLineItem) GetDiscountPercentage() *float64 {
	if o == nil {
		return nil
	}
	return o.DiscountPercentage
}

func (o *BillCreditNoteLineItem) GetItemRef() *BillCreditNoteLineItemItemReference {
	if o == nil {
		return nil
	}
	return o.ItemRef
}

func (o *BillCreditNoteLineItem) GetQuantity() float64 {
	if o == nil {
		return 0.0
	}
	return o.Quantity
}

func (o *BillCreditNoteLineItem) GetSubTotal() *float64 {
	if o == nil {
		return nil
	}
	return o.SubTotal
}

func (o *BillCreditNoteLineItem) GetTaxAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *BillCreditNoteLineItem) GetTaxRateRef() *BillCreditNoteLineItemTaxRateReference {
	if o == nil {
		return nil
	}
	return o.TaxRateRef
}

func (o *BillCreditNoteLineItem) GetTotalAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *BillCreditNoteLineItem) GetTracking() *AccountsPayableTracking {
	if o == nil {
		return nil
	}
	return o.Tracking
}

func (o *BillCreditNoteLineItem) GetTrackingCategoryRefs() []TrackingCategoryRef {
	if o == nil {
		return nil
	}
	return o.TrackingCategoryRefs
}

func (o *BillCreditNoteLineItem) GetUnitAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.UnitAmount
}
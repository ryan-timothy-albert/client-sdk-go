// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type InvoiceLineItem struct {
	// Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.
	AccountRef *AccountRef `json:"accountRef,omitempty"`
	// Friendly name of the goods or services provided.
	Description *string `json:"description,omitempty"`
	// Numerical value of any discounts applied.
	DiscountAmount *float64 `json:"discountAmount,omitempty"`
	// Percentage rate (from 0 to 100) of any discounts applied to the unit amount.
	DiscountPercentage *float64 `json:"discountPercentage,omitempty"`
	IsDirectIncome     *bool    `json:"isDirectIncome,omitempty"`
	// Reference to the product, service type, or inventory item to which the direct cost is linked.
	ItemRef *ItemRef `json:"itemRef,omitempty"`
	// Number of units of goods or services provided.
	Quantity float64 `json:"quantity"`
	// Amount of the line, inclusive of discounts but exclusive of tax.
	SubTotal *float64 `json:"subTotal,omitempty"`
	// Amount of tax for the line.
	TaxAmount *float64 `json:"taxAmount,omitempty"`
	// Reference to the tax rate to which the line item is linked.
	TaxRateRef *TaxRateRef `json:"taxRateRef,omitempty"`
	// Total amount of the line, including tax. When pushing invoices to Xero, the total amount is exclusive of tax to allow automatic calculations if a tax rate or tax amount is not specified.
	TotalAmount *float64 `json:"totalAmount,omitempty"`
	// Categories, and a project and customer, against which the item is tracked.
	Tracking *Tracking `json:"tracking,omitempty"`
	// Reference to the tracking categories to which the line item is linked.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	TrackingCategoryRefs []TrackingCategoryRefsitems `json:"trackingCategoryRefs,omitempty"`
	// Price of each unit of goods or services.
	UnitAmount float64 `json:"unitAmount"`
}

func (o *InvoiceLineItem) GetAccountRef() *AccountRef {
	if o == nil {
		return nil
	}
	return o.AccountRef
}

func (o *InvoiceLineItem) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *InvoiceLineItem) GetDiscountAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.DiscountAmount
}

func (o *InvoiceLineItem) GetDiscountPercentage() *float64 {
	if o == nil {
		return nil
	}
	return o.DiscountPercentage
}

func (o *InvoiceLineItem) GetIsDirectIncome() *bool {
	if o == nil {
		return nil
	}
	return o.IsDirectIncome
}

func (o *InvoiceLineItem) GetItemRef() *ItemRef {
	if o == nil {
		return nil
	}
	return o.ItemRef
}

func (o *InvoiceLineItem) GetQuantity() float64 {
	if o == nil {
		return 0.0
	}
	return o.Quantity
}

func (o *InvoiceLineItem) GetSubTotal() *float64 {
	if o == nil {
		return nil
	}
	return o.SubTotal
}

func (o *InvoiceLineItem) GetTaxAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *InvoiceLineItem) GetTaxRateRef() *TaxRateRef {
	if o == nil {
		return nil
	}
	return o.TaxRateRef
}

func (o *InvoiceLineItem) GetTotalAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *InvoiceLineItem) GetTracking() *Tracking {
	if o == nil {
		return nil
	}
	return o.Tracking
}

func (o *InvoiceLineItem) GetTrackingCategoryRefs() []TrackingCategoryRefsitems {
	if o == nil {
		return nil
	}
	return o.TrackingCategoryRefs
}

func (o *InvoiceLineItem) GetUnitAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.UnitAmount
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// DirectIncomeLineItemItemReference - Reference to the product, service type, or inventory item to which the direct cost is linked.
type DirectIncomeLineItemItemReference struct {
	// Unique identifier for the item in the accounting platform.
	ID string `json:"id"`
	// Name of the item in the accounting platform.
	Name *string `json:"name,omitempty"`
}

func (o *DirectIncomeLineItemItemReference) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DirectIncomeLineItemItemReference) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// DirectIncomeLineItemTaxRateReference - Reference to the tax rate to which the line item is linked.
type DirectIncomeLineItemTaxRateReference struct {
	// Applicable tax rate.
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	// Unique identifier for the tax rate in the accounting platform.
	ID *string `json:"id,omitempty"`
	// Name of the tax rate in the accounting platform.
	Name *string `json:"name,omitempty"`
}

func (o *DirectIncomeLineItemTaxRateReference) GetEffectiveTaxRate() *float64 {
	if o == nil {
		return nil
	}
	return o.EffectiveTaxRate
}

func (o *DirectIncomeLineItemTaxRateReference) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DirectIncomeLineItemTaxRateReference) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// DirectIncomeLineItemTrackingCategoryRefs - References a category against which the item is tracked.
//
// Deprecated type: This will be removed in a future release, please migrate away from it as soon as possible.
type DirectIncomeLineItemTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

func (o *DirectIncomeLineItemTrackingCategoryRefs) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DirectIncomeLineItemTrackingCategoryRefs) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type DirectIncomeLineItem struct {
	// Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.
	AccountRef *AccountRef `json:"accountRef,omitempty"`
	// A user-friendly name of the goods or services.
	Description *string `json:"description,omitempty"`
	// Discount amount for the line before tax.
	DiscountAmount *float64 `json:"discountAmount,omitempty"`
	// Discount percentage for the line before tax.
	DiscountPercentage *float64 `json:"discountPercentage,omitempty"`
	// Reference to the product, service type, or inventory item to which the direct cost is linked.
	ItemRef *DirectIncomeLineItemItemReference `json:"itemRef,omitempty"`
	// The number of units of goods or services received.
	//
	// Note: If the platform does not provide this information, the quantity will be mapped as 1.
	Quantity float64 `json:"quantity"`
	// The amount of the line, inclusive of discounts, but exclusive of tax.
	SubTotal *float64 `json:"subTotal,omitempty"`
	// The amount of tax for the line.
	// Note: If the platform does not provide this information, the quantity will be mapped as 0.00.
	TaxAmount *float64 `json:"taxAmount,omitempty"`
	// Reference to the tax rate to which the line item is linked.
	TaxRateRef *DirectIncomeLineItemTaxRateReference `json:"taxRateRef,omitempty"`
	// The total amount of the line, including tax.
	TotalAmount *float64 `json:"totalAmount,omitempty"`
	// An array of categories against which this direct cost is tracked.
	TrackingCategoryRefs []DirectIncomeLineItemTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	// The price of each unit of goods or services.
	// Note: If the platform does not provide this information, the unit amount will be mapped to the total amount.
	UnitAmount float64 `json:"unitAmount"`
}

func (o *DirectIncomeLineItem) GetAccountRef() *AccountRef {
	if o == nil {
		return nil
	}
	return o.AccountRef
}

func (o *DirectIncomeLineItem) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *DirectIncomeLineItem) GetDiscountAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.DiscountAmount
}

func (o *DirectIncomeLineItem) GetDiscountPercentage() *float64 {
	if o == nil {
		return nil
	}
	return o.DiscountPercentage
}

func (o *DirectIncomeLineItem) GetItemRef() *DirectIncomeLineItemItemReference {
	if o == nil {
		return nil
	}
	return o.ItemRef
}

func (o *DirectIncomeLineItem) GetQuantity() float64 {
	if o == nil {
		return 0.0
	}
	return o.Quantity
}

func (o *DirectIncomeLineItem) GetSubTotal() *float64 {
	if o == nil {
		return nil
	}
	return o.SubTotal
}

func (o *DirectIncomeLineItem) GetTaxAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *DirectIncomeLineItem) GetTaxRateRef() *DirectIncomeLineItemTaxRateReference {
	if o == nil {
		return nil
	}
	return o.TaxRateRef
}

func (o *DirectIncomeLineItem) GetTotalAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *DirectIncomeLineItem) GetTrackingCategoryRefs() []DirectIncomeLineItemTrackingCategoryRefs {
	if o == nil {
		return nil
	}
	return o.TrackingCategoryRefs
}

func (o *DirectIncomeLineItem) GetUnitAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.UnitAmount
}
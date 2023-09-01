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

type BillCreditNoteLineItemTrackingCustomerRef struct {
	// `customerName` from the Customer data type
	CompanyName *string `json:"companyName,omitempty"`
	// `id` from the Customers data type
	ID string `json:"id"`
}

func (o *BillCreditNoteLineItemTrackingCustomerRef) GetCompanyName() *string {
	if o == nil {
		return nil
	}
	return o.CompanyName
}

func (o *BillCreditNoteLineItemTrackingCustomerRef) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type BillCreditNoteLineItemTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

func (o *BillCreditNoteLineItemTrackingProjectRef) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *BillCreditNoteLineItemTrackingProjectRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// BillCreditNoteLineItemTracking - Categories, and a project and customer, against which the item is tracked.
type BillCreditNoteLineItemTracking struct {
	CategoryRefs []TrackingCategoryRef                      `json:"categoryRefs"`
	CustomerRef  *BillCreditNoteLineItemTrackingCustomerRef `json:"customerRef,omitempty"`
	IsBilledTo   BilledToType                               `json:"isBilledTo"`
	IsRebilledTo BilledToType                               `json:"isRebilledTo"`
	ProjectRef   *BillCreditNoteLineItemTrackingProjectRef  `json:"projectRef,omitempty"`
}

func (o *BillCreditNoteLineItemTracking) GetCategoryRefs() []TrackingCategoryRef {
	if o == nil {
		return []TrackingCategoryRef{}
	}
	return o.CategoryRefs
}

func (o *BillCreditNoteLineItemTracking) GetCustomerRef() *BillCreditNoteLineItemTrackingCustomerRef {
	if o == nil {
		return nil
	}
	return o.CustomerRef
}

func (o *BillCreditNoteLineItemTracking) GetIsBilledTo() BilledToType {
	if o == nil {
		return BilledToType("")
	}
	return o.IsBilledTo
}

func (o *BillCreditNoteLineItemTracking) GetIsRebilledTo() BilledToType {
	if o == nil {
		return BilledToType("")
	}
	return o.IsRebilledTo
}

func (o *BillCreditNoteLineItemTracking) GetProjectRef() *BillCreditNoteLineItemTrackingProjectRef {
	if o == nil {
		return nil
	}
	return o.ProjectRef
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
	TaxRateRef *TaxRateRef `json:"taxRateRef,omitempty"`
	// Total amount of the line item, including discounts and tax.
	TotalAmount *float64 `json:"totalAmount,omitempty"`
	// Categories, and a project and customer, against which the item is tracked.
	Tracking *BillCreditNoteLineItemTracking `json:"tracking,omitempty"`
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

func (o *BillCreditNoteLineItem) GetTaxRateRef() *TaxRateRef {
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

func (o *BillCreditNoteLineItem) GetTracking() *BillCreditNoteLineItemTracking {
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

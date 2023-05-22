// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreditNoteLineItemTracking - Categories, and a project and customer, against which the item is tracked.
type CreditNoteLineItemTracking struct {
	CategoryRefs []TrackingCategoryRef `json:"categoryRefs"`
	CustomerRef  *CustomerRef          `json:"customerRef,omitempty"`
	IsBilledTo   BilledToType1         `json:"isBilledTo"`
	IsRebilledTo BilledToType1         `json:"isRebilledTo"`
	ProjectRef   *ProjectRef           `json:"projectRef,omitempty"`
}

type CreditNoteLineItem struct {
	// Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.
	AccountRef *AccountRef `json:"accountRef,omitempty"`
	// Friendly name of each line item. For example, the goods or service for which credit has been issued.
	Description *string `json:"description,omitempty"`
	// Value of any discounts applied.
	DiscountAmount *float64 `json:"discountAmount,omitempty"`
	// Percentage rate of any discount applied to the line item.
	DiscountPercentage *float64 `json:"discountPercentage,omitempty"`
	IsDirectIncome     *bool    `json:"isDirectIncome,omitempty"`
	ItemRef            *ItemRef `json:"itemRef,omitempty"`
	// Number of units of the goods or service for which credit has been issued.
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
	Tracking *CreditNoteLineItemTracking `json:"tracking,omitempty"`
	// Reference to the tracking categories to which the line item is linked.
	TrackingCategoryRefs []TrackingCategoryRef `json:"trackingCategoryRefs,omitempty"`
	// Unit price of the goods or service.
	UnitAmount float64 `json:"unitAmount"`
}

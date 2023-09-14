// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v3/pkg/types"
)

type OrderLineItem struct {
	DiscountAllocations []OrderDiscountAllocation `json:"discountAllocations,omitempty"`
	// A unique, persistent identifier for this record
	ID string `json:"id"`
	// Reference that links the line item to the correct product details.
	ProductRef *ProductRef `json:"productRef,omitempty"`
	// Reference that links the line item to the specific version of product that has been ordered.
	ProductVariantRef *ProductVariantRef `json:"productVariantRef,omitempty"`
	// Number of units of the product sold.
	// For refunds, quantity is a negative value.
	//
	Quantity *types.Decimal `json:"quantity,omitempty"`
	// Percentage rate (from 0 to 100) of any sale tax applied to the unit amount.
	TaxPercentage *types.Decimal `json:"taxPercentage,omitempty"`
	// Taxes breakdown as applied to order lines.
	Taxes []TaxComponentAllocation `json:"taxes,omitempty"`
	// Total price of the line item, including discounts, tax and minus any refunds.
	TotalAmount *types.Decimal `json:"totalAmount,omitempty"`
	// Total amount of tax applied to the line item.
	TotalTaxAmount *types.Decimal `json:"totalTaxAmount,omitempty"`
	// Price per unit of goods or service.
	UnitPrice *types.Decimal `json:"unitPrice,omitempty"`
}

func (o *OrderLineItem) GetDiscountAllocations() []OrderDiscountAllocation {
	if o == nil {
		return nil
	}
	return o.DiscountAllocations
}

func (o *OrderLineItem) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *OrderLineItem) GetProductRef() *ProductRef {
	if o == nil {
		return nil
	}
	return o.ProductRef
}

func (o *OrderLineItem) GetProductVariantRef() *ProductVariantRef {
	if o == nil {
		return nil
	}
	return o.ProductVariantRef
}

func (o *OrderLineItem) GetQuantity() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.Quantity
}

func (o *OrderLineItem) GetTaxPercentage() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.TaxPercentage
}

func (o *OrderLineItem) GetTaxes() []TaxComponentAllocation {
	if o == nil {
		return nil
	}
	return o.Taxes
}

func (o *OrderLineItem) GetTotalAmount() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *OrderLineItem) GetTotalTaxAmount() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.TotalTaxAmount
}

func (o *OrderLineItem) GetUnitPrice() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.UnitPrice
}

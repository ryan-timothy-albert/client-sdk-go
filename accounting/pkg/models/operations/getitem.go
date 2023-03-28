// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetItemRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	ItemID    string `pathParam:"style=simple,explode=false,name=itemId"`
}

// GetItemSourceModifiedDateBillItemAccountRef - Reference of the account to which the item is linked.
type GetItemSourceModifiedDateBillItemAccountRef struct {
	// 'id' from the Accounts data type.
	ID *string `json:"id,omitempty"`
	// 'name' from the Accounts data type.
	Name *string `json:"name,omitempty"`
}

// GetItemSourceModifiedDateBillItemTaxRateRef - Reference of the tax rate to which the item is linked.
type GetItemSourceModifiedDateBillItemTaxRateRef struct {
	// Applicable tax rate.
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	// 'id' from the 'taxRates' data type.
	ID *string `json:"id,omitempty"`
	// 'name' from the 'taxRates' data type.
	Name *string `json:"name,omitempty"`
}

// GetItemSourceModifiedDateBillItem - Item details that are only for bills.
type GetItemSourceModifiedDateBillItem struct {
	// Reference of the account to which the item is linked.
	AccountRef *GetItemSourceModifiedDateBillItemAccountRef `json:"accountRef,omitempty"`
	// Short description of the product or service that has been bought by the customer.
	Description *string `json:"description,omitempty"`
	// Reference of the tax rate to which the item is linked.
	TaxRateRef *GetItemSourceModifiedDateBillItemTaxRateRef `json:"taxRateRef,omitempty"`
	// Unit price of the product or service.
	UnitPrice *float64 `json:"unitPrice,omitempty"`
}

// GetItemSourceModifiedDateInvoiceItemAccountRef - Reference of the account to which the item is linked.
type GetItemSourceModifiedDateInvoiceItemAccountRef struct {
	// 'id' from the Accounts data type.
	ID *string `json:"id,omitempty"`
	// 'name' from the Accounts data type.
	Name *string `json:"name,omitempty"`
}

// GetItemSourceModifiedDateInvoiceItemTaxRateRef - Reference of the tax rate to which the item is linked.
type GetItemSourceModifiedDateInvoiceItemTaxRateRef struct {
	// Applicable tax rate.
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	// 'id' from the 'taxRates' data type.
	ID *string `json:"id,omitempty"`
	// 'name' from the 'taxRates' data type.
	Name *string `json:"name,omitempty"`
}

// GetItemSourceModifiedDateInvoiceItem - Item details that are only for bills.
type GetItemSourceModifiedDateInvoiceItem struct {
	// Reference of the account to which the item is linked.
	AccountRef *GetItemSourceModifiedDateInvoiceItemAccountRef `json:"accountRef,omitempty"`
	// Short description of the product or service that has been bought by the customer.
	Description *string `json:"description,omitempty"`
	// Reference of the tax rate to which the item is linked.
	TaxRateRef *GetItemSourceModifiedDateInvoiceItemTaxRateRef `json:"taxRateRef,omitempty"`
	// Unit price of the product or service.
	UnitPrice *float64 `json:"unitPrice,omitempty"`
}

// GetItemSourceModifiedDateItemStatusEnum - Current state of the item, either:
//
// - `Active`: Available for use
// - `Archived`: Unavailable
// - `Unknown`
//
// Due to a [limitation in Xero's API](https://docs.codat.io/integrations/accounting/xero/xero-faq#why-do-all-of-my-items-from-xero-have-their-status-as-unknown), all items from Xero are mapped as `Unknown`.
type GetItemSourceModifiedDateItemStatusEnum string

const (
	GetItemSourceModifiedDateItemStatusEnumUnknown  GetItemSourceModifiedDateItemStatusEnum = "Unknown"
	GetItemSourceModifiedDateItemStatusEnumActive   GetItemSourceModifiedDateItemStatusEnum = "Active"
	GetItemSourceModifiedDateItemStatusEnumArchived GetItemSourceModifiedDateItemStatusEnum = "Archived"
)

func (e *GetItemSourceModifiedDateItemStatusEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Unknown":
		fallthrough
	case "Active":
		fallthrough
	case "Archived":
		*e = GetItemSourceModifiedDateItemStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for GetItemSourceModifiedDateItemStatusEnum: %s", s)
	}
}

type GetItemSourceModifiedDateMetadata struct {
	// Indicates whether the record has been deleted in the third-party system this record originated from.
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// GetItemSourceModifiedDateTypeEnum - Type of the item.
type GetItemSourceModifiedDateTypeEnum string

const (
	GetItemSourceModifiedDateTypeEnumUnknown      GetItemSourceModifiedDateTypeEnum = "Unknown"
	GetItemSourceModifiedDateTypeEnumInventory    GetItemSourceModifiedDateTypeEnum = "Inventory"
	GetItemSourceModifiedDateTypeEnumNonInventory GetItemSourceModifiedDateTypeEnum = "NonInventory"
	GetItemSourceModifiedDateTypeEnumService      GetItemSourceModifiedDateTypeEnum = "Service"
)

func (e *GetItemSourceModifiedDateTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Unknown":
		fallthrough
	case "Inventory":
		fallthrough
	case "NonInventory":
		fallthrough
	case "Service":
		*e = GetItemSourceModifiedDateTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for GetItemSourceModifiedDateTypeEnum: %s", s)
	}
}

// GetItemSourceModifiedDate - > View the coverage for items in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=items" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// **Items** allow your customers to save and track details of the products and services that they buy and sell.
type GetItemSourceModifiedDate struct {
	// Item details that are only for bills.
	BillItem *GetItemSourceModifiedDateBillItem `json:"billItem,omitempty"`
	// Friendly reference for the item.
	Code *string `json:"code,omitempty"`
	// Identifier for the item that is unique to a company in the accounting platform.
	ID *string `json:"id,omitempty"`
	// Item details that are only for bills.
	InvoiceItem *GetItemSourceModifiedDateInvoiceItem `json:"invoiceItem,omitempty"`
	// Whether you can use this item for bills.
	IsBillItem bool `json:"isBillItem"`
	// Whether you can use this item for invoices.
	IsInvoiceItem bool `json:"isInvoiceItem"`
	// Current state of the item, either:
	//
	// - `Active`: Available for use
	// - `Archived`: Unavailable
	// - `Unknown`
	//
	// Due to a [limitation in Xero's API](https://docs.codat.io/integrations/accounting/xero/xero-faq#why-do-all-of-my-items-from-xero-have-their-status-as-unknown), all items from Xero are mapped as `Unknown`.
	ItemStatus GetItemSourceModifiedDateItemStatusEnum `json:"itemStatus"`
	Metadata   *GetItemSourceModifiedDateMetadata      `json:"metadata,omitempty"`
	// The date on which this record was last modified in Codat.
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// Name of the item in the accounting platform.
	Name *string `json:"name,omitempty"`
	// The date on which this record was last modified in the originating system
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Type of the item.
	Type GetItemSourceModifiedDateTypeEnum `json:"type"`
}

type GetItemResponse struct {
	ContentType string
	// Success
	SourceModifiedDate *GetItemSourceModifiedDate
	StatusCode         int
	RawResponse        *http.Response
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/types"
)

// TaxRate - > View the coverage for tax rates in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=taxRates" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Accounting systems typically store a set of taxes and associated rates within the accounting package. This means that users don't have to look up or remember the rates for each type of tax. For example, applying the tax "UK sales VAT" to line items of an invoice adds the correct rate of 20%.
//
// ### Tax components
//
// In some cases, a tax is made up of multiple sub taxes, often called _components_ of the tax.  For example, you may have an item that is charged a tax rate called "City import tax (8%)" that has two components:
//
// - A city tax of 5%
// - An import tax of 3%
//
// > **Effective tax rates**
// > - Where there are multiple components of a tax, each component may be calculated on the original amount and added together. Alternatively, one tax may be calculated on the sub-total of the original amount plus another tax, which is referred to as _compounding_. When there is compounding, the effective tax rate is the rate that, if applied to the original amount, would result in the total amount of tax with compounding.
// >
// > **Example:**
// > A tax has two components. Both components have a rate of 10%, and one component is compound. In this case, there is a total tax rate of 20% but an effective tax rate of 21%.
// >
// > - For QuickBooks Online, Codat doesn't use compound rates. Instead, the calculated effective tax rate for each component is shown. This means that the effective and total rates are the same because the total tax rate is a sum of the component rates.
type TaxRate struct {
	// Code for the tax rate from the accounting platform.
	Code       *string            `json:"code,omitempty"`
	Components []TaxRateComponent `json:"components,omitempty"`
	// See Effective tax rates description.
	EffectiveTaxRate *types.Decimal `json:"effectiveTaxRate,omitempty"`
	// Identifier for the tax rate, unique for the company in the accounting platform.
	ID           *string   `json:"id,omitempty"`
	Metadata     *Metadata `json:"metadata,omitempty"`
	ModifiedDate *string   `json:"modifiedDate,omitempty"`
	// Codat-augmented name of the tax rate in the accounting platform.
	Name               *string `json:"name,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Status of the tax rate in the accounting platform.
	// - `Active` - An active tax rate in use by a company.
	// - `Archived` - A tax rate that has been archived or is inactive in the accounting platform.
	// - `Unknown` - Where the status of the tax rate cannot be determined from the underlying platform.
	Status *TaxRateStatus `json:"status,omitempty"`
	// Total (not compounded) sum of the components of a tax rate.
	TotalTaxRate       *types.Decimal            `json:"totalTaxRate,omitempty"`
	ValidDatatypeLinks []ValidDatatypeLinksitems `json:"validDatatypeLinks,omitempty"`
}

func (o *TaxRate) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *TaxRate) GetComponents() []TaxRateComponent {
	if o == nil {
		return nil
	}
	return o.Components
}

func (o *TaxRate) GetEffectiveTaxRate() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.EffectiveTaxRate
}

func (o *TaxRate) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TaxRate) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *TaxRate) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *TaxRate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TaxRate) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *TaxRate) GetStatus() *TaxRateStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *TaxRate) GetTotalTaxRate() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.TotalTaxRate
}

func (o *TaxRate) GetValidDatatypeLinks() []ValidDatatypeLinksitems {
	if o == nil {
		return nil
	}
	return o.ValidDatatypeLinks
}

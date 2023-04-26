// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PaymentMethod - > View the coverage for payment methods in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=paymentMethods" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A Payment Method represents the payment method(s) used to pay a Bill. Payment Methods are referenced on [Bill Payments](https://docs.codat.io/accounting-api#/schemas/BillPayment) and [Payments](https://docs.codat.io/accounting-api#/schemas/Payment).
type PaymentMethod struct {
	// Unique identifier for the payment method.
	ID           *string   `json:"id,omitempty"`
	Metadata     *Metadata `json:"metadata,omitempty"`
	ModifiedDate *string   `json:"modifiedDate,omitempty"`
	// Name of the payment method.
	Name               *string `json:"name,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Status of the Payment Method.
	Status *PaymentMethodStatusEnum `json:"status,omitempty"`
	// Method of payment.
	Type *PaymentMethodTypeEnum `json:"type,omitempty"`
}

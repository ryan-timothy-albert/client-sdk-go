// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ProductVariantRef - Reference that links the line item to the specific version of product that has been ordered.
type ProductVariantRef struct {
	// The unique identitifer of the product variant being referenced.
	ID string `json:"id"`
	// Name of the product variant being referenced.
	Name *string `json:"name,omitempty"`
}

func (o *ProductVariantRef) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ProductVariantRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
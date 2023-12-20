// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Group - A container you can use to organize companies together according to a shared characteristic of your choice.
type Group struct {
	// Unique identifier for the group.
	ID *string `json:"id,omitempty"`
	// Descriptive name of the group.
	Name *string `json:"name,omitempty"`
}

func (o *Group) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Group) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

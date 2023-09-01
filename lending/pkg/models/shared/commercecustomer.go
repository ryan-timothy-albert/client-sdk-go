// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CommerceCustomer - When a customer places an order with the connected commerce store their details are added to the Customers dataset. You can use the data from the Customers endpoints to calculate key metrics, such as customer churn.
//
// Explore our [data coverage](https://knowledge.codat.io/supported-features/commerce?view=tab-by-data-type&dataType=commerce-customers) for this data type.
type CommerceCustomer struct {
	// Addresses of the customer
	Addresses []CommerceAddress `json:"addresses,omitempty"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	CreatedDate *string `json:"createdDate,omitempty"`
	// Name of the customer
	CustomerName    *string `json:"customerName,omitempty"`
	DefaultCurrency *string `json:"defaultCurrency,omitempty"`
	// Email address of the customer
	EmailAddress *string `json:"emailAddress,omitempty"`
	// A unique, persistent identifier for this record
	ID           string  `json:"id"`
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// Any additional information about the customer
	Note *string `json:"note,omitempty"`
	// A phone number.
	Phone              *string `json:"phone,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
}

func (o *CommerceCustomer) GetAddresses() []CommerceAddress {
	if o == nil {
		return nil
	}
	return o.Addresses
}

func (o *CommerceCustomer) GetCreatedDate() *string {
	if o == nil {
		return nil
	}
	return o.CreatedDate
}

func (o *CommerceCustomer) GetCustomerName() *string {
	if o == nil {
		return nil
	}
	return o.CustomerName
}

func (o *CommerceCustomer) GetDefaultCurrency() *string {
	if o == nil {
		return nil
	}
	return o.DefaultCurrency
}

func (o *CommerceCustomer) GetEmailAddress() *string {
	if o == nil {
		return nil
	}
	return o.EmailAddress
}

func (o *CommerceCustomer) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CommerceCustomer) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *CommerceCustomer) GetNote() *string {
	if o == nil {
		return nil
	}
	return o.Note
}

func (o *CommerceCustomer) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *CommerceCustomer) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

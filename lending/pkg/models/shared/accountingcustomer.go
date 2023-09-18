// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AccountingCustomer - > View the coverage for customers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A customer is a person or organisation that buys goods or services. From the Customers endpoints, you can retrieve a [list of all the customers of a company](https://api.codat.io/swagger/index.html#/Customers/get_companies__companyId__data_customers).
//
// Customers' data links to accounts receivable [invoices](https://docs.codat.io/accounting-api#/schemas/Invoice).
type AccountingCustomer struct {
	// An array of Addresses.
	Addresses []AccountingAddress `json:"addresses,omitempty"`
	// Name of the main contact for the identified customer.
	ContactName *string `json:"contactName,omitempty"`
	// An array of Contacts.
	Contacts []Contact `json:"contacts,omitempty"`
	// Name of the customer as recorded in the accounting system, typically the company name.
	CustomerName *string `json:"customerName,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	DefaultCurrency *string `json:"defaultCurrency,omitempty"`
	// Email address the customer can be contacted by.
	EmailAddress *string `json:"emailAddress,omitempty"`
	// Identifier for the customer, unique to the company in the accounting platform.
	ID           *string   `json:"id,omitempty"`
	Metadata     *Metadata `json:"metadata,omitempty"`
	ModifiedDate *string   `json:"modifiedDate,omitempty"`
	// Phone number the customer can be contacted by.
	Phone *string `json:"phone,omitempty"`
	// Company number. In the UK, this is typically the Companies House company registration number.
	RegistrationNumber *string `json:"registrationNumber,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Status of customer.
	Status CustomerStatus `json:"status"`
	// Supplemental data is additional data you can include in our standard data types.
	//
	// It is referenced as a configured dynamic key value pair that is unique to the accounting platform. [Learn more](https://docs.codat.io/using-the-api/supplemental-data/overview) about supplemental data.
	SupplementalData *SupplementalData `json:"supplementalData,omitempty"`
	// Company tax number.
	TaxNumber *string `json:"taxNumber,omitempty"`
}

func (o *AccountingCustomer) GetAddresses() []AccountingAddress {
	if o == nil {
		return nil
	}
	return o.Addresses
}

func (o *AccountingCustomer) GetContactName() *string {
	if o == nil {
		return nil
	}
	return o.ContactName
}

func (o *AccountingCustomer) GetContacts() []Contact {
	if o == nil {
		return nil
	}
	return o.Contacts
}

func (o *AccountingCustomer) GetCustomerName() *string {
	if o == nil {
		return nil
	}
	return o.CustomerName
}

func (o *AccountingCustomer) GetDefaultCurrency() *string {
	if o == nil {
		return nil
	}
	return o.DefaultCurrency
}

func (o *AccountingCustomer) GetEmailAddress() *string {
	if o == nil {
		return nil
	}
	return o.EmailAddress
}

func (o *AccountingCustomer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingCustomer) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *AccountingCustomer) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *AccountingCustomer) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *AccountingCustomer) GetRegistrationNumber() *string {
	if o == nil {
		return nil
	}
	return o.RegistrationNumber
}

func (o *AccountingCustomer) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *AccountingCustomer) GetStatus() CustomerStatus {
	if o == nil {
		return CustomerStatus("")
	}
	return o.Status
}

func (o *AccountingCustomer) GetSupplementalData() *SupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}

func (o *AccountingCustomer) GetTaxNumber() *string {
	if o == nil {
		return nil
	}
	return o.TaxNumber
}

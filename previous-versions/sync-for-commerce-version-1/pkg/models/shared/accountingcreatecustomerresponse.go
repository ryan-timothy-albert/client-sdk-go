// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AccountingCreateCustomerResponse - Success
type AccountingCreateCustomerResponse struct {
	// Contains a single entry that communicates which record has changed and the manner in which it changed.
	Changes []PushOperationChange `json:"changes,omitempty"`
	// Unique identifier for your SMB in Codat.
	CompanyID string `json:"companyId"`
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
	CompletedOnUtc *string `json:"completedOnUtc,omitempty"`
	// > View the coverage for customers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers" target="_blank">Data coverage explorer</a>.
	//
	// ## Overview
	//
	// A customer is a person or organisation that buys goods or services. From the Customers endpoints, you can retrieve a [list of all the customers of a company](https://api.codat.io/swagger/index.html#/Customers/get_companies__companyId__data_customers).
	//
	// Customers' data links to accounts receivable [invoices](https://docs.codat.io/accounting-api#/schemas/Invoice).
	//
	Data *AccountingCustomer `json:"data,omitempty"`
	// Unique identifier for a company's data connection.
	DataConnectionKey string `json:"dataConnectionKey"`
	// Available Data types
	DataType     *DataType `json:"dataType,omitempty"`
	ErrorMessage *string   `json:"errorMessage,omitempty"`
	// A unique identifier generated by Codat to represent this single push operation. This identifier can be used to track the status of the push, and should be persisted.
	PushOperationKey string `json:"pushOperationKey"`
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
	RequestedOnUtc string `json:"requestedOnUtc"`
	// The status of the push operation.
	Status           PushOperationStatus `json:"status"`
	StatusCode       int64               `json:"statusCode"`
	TimeoutInMinutes *int                `json:"timeoutInMinutes,omitempty"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	TimeoutInSeconds *int `json:"timeoutInSeconds,omitempty"`
	// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
	Validation *Validation `json:"validation,omitempty"`
}

func (o *AccountingCreateCustomerResponse) GetChanges() []PushOperationChange {
	if o == nil {
		return nil
	}
	return o.Changes
}

func (o *AccountingCreateCustomerResponse) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *AccountingCreateCustomerResponse) GetCompletedOnUtc() *string {
	if o == nil {
		return nil
	}
	return o.CompletedOnUtc
}

func (o *AccountingCreateCustomerResponse) GetData() *AccountingCustomer {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *AccountingCreateCustomerResponse) GetDataConnectionKey() string {
	if o == nil {
		return ""
	}
	return o.DataConnectionKey
}

func (o *AccountingCreateCustomerResponse) GetDataType() *DataType {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *AccountingCreateCustomerResponse) GetErrorMessage() *string {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *AccountingCreateCustomerResponse) GetPushOperationKey() string {
	if o == nil {
		return ""
	}
	return o.PushOperationKey
}

func (o *AccountingCreateCustomerResponse) GetRequestedOnUtc() string {
	if o == nil {
		return ""
	}
	return o.RequestedOnUtc
}

func (o *AccountingCreateCustomerResponse) GetStatus() PushOperationStatus {
	if o == nil {
		return PushOperationStatus("")
	}
	return o.Status
}

func (o *AccountingCreateCustomerResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AccountingCreateCustomerResponse) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

func (o *AccountingCreateCustomerResponse) GetTimeoutInSeconds() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInSeconds
}

func (o *AccountingCreateCustomerResponse) GetValidation() *Validation {
	if o == nil {
		return nil
	}
	return o.Validation
}

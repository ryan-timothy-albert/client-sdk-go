// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type BillPaymentLine struct {
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
	AllocatedOnDate *string `json:"allocatedOnDate,omitempty"`
	// Amount in the bill payment currency.
	Amount float64               `json:"amount"`
	Links  []BillPaymentLineLink `json:"links,omitempty"`
}

func (o *BillPaymentLine) GetAllocatedOnDate() *string {
	if o == nil {
		return nil
	}
	return o.AllocatedOnDate
}

func (o *BillPaymentLine) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}

func (o *BillPaymentLine) GetLinks() []BillPaymentLineLink {
	if o == nil {
		return nil
	}
	return o.Links
}

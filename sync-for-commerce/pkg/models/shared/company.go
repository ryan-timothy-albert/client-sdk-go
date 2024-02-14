// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type GroupReference struct {
	// Unique identifier for the group.
	ID *string `json:"id,omitempty"`
}

func (o *GroupReference) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// Company - In Codat, a company represents a business sharing access to their data. Each company can have multiple [connections](https://docs.codat.io/sync-for-sync-for-commerce-api#/schemas/Connection) to different data sources such as one connection to [Xero](https://docs.codat.io/integrations/accounting/xero/accounting-xero) for accounting data, two connections to [Plaid](https://docs.codat.io/integrations/banking/plaid/banking-plaid) for two bank accounts and a connection to [Zettle](https://docs.codat.io/integrations/commerce/zettle/commerce-zettle) for POS data.
//
// Typically each company is one of your customers.
//
// When you create a company, you can specify a `name` and we will automatically generate a unique `id` for the company. You can also add a `description` to store any additional information about the company.
type Company struct {
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
	Created *string `json:"created,omitempty"`
	// Name of user that created the company in Codat.
	CreatedByUserName *string      `json:"createdByUserName,omitempty"`
	DataConnections   []Connection `json:"dataConnections,omitempty"`
	// Additional information about the company. This can be used to store foreign IDs, references, etc.
	Description *string `json:"description,omitempty"`
	// An array of groups the company has been assigned to.
	Groups []GroupReference `json:"groups,omitempty"`
	// Unique identifier for your SMB in Codat.
	ID string `json:"id"`
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
	LastSync *string `json:"lastSync,omitempty"`
	// The name of the company
	Name string `json:"name"`
	// `platformKeys` name used when creating the company.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Platform *string `json:"platform,omitempty"`
	// The `redirect` [Link URL](https://docs.codat.io/auth-flow/authorize-hosted-link) enabling the customer to start their auth flow journey for the company.
	Redirect string `json:"redirect"`
}

func (o *Company) GetCreated() *string {
	if o == nil {
		return nil
	}
	return o.Created
}

func (o *Company) GetCreatedByUserName() *string {
	if o == nil {
		return nil
	}
	return o.CreatedByUserName
}

func (o *Company) GetDataConnections() []Connection {
	if o == nil {
		return nil
	}
	return o.DataConnections
}

func (o *Company) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Company) GetGroups() []GroupReference {
	if o == nil {
		return nil
	}
	return o.Groups
}

func (o *Company) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Company) GetLastSync() *string {
	if o == nil {
		return nil
	}
	return o.LastSync
}

func (o *Company) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Company) GetPlatform() *string {
	if o == nil {
		return nil
	}
	return o.Platform
}

func (o *Company) GetRedirect() string {
	if o == nil {
		return ""
	}
	return o.Redirect
}

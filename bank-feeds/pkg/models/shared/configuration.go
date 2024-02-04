// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Configuration struct {
	// The company name defined in the accounting platform.
	AccountingSoftwareCompanyName *string `json:"accountingSoftwareCompanyName,omitempty"`
	// Unique identifier for your SMB in Codat.
	CompanyID     *string            `json:"companyId,omitempty"`
	Configuration *SyncConfiguration `json:"configuration,omitempty"`
	// True if the company has been configured.
	Configured *bool `json:"configured,omitempty"`
	// Enabled or disable bank feeds.
	Enabled  *bool                  `json:"enabled,omitempty"`
	Schedule *ConfigurationSchedule `json:"schedule,omitempty"`
}

func (o *Configuration) GetAccountingSoftwareCompanyName() *string {
	if o == nil {
		return nil
	}
	return o.AccountingSoftwareCompanyName
}

func (o *Configuration) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *Configuration) GetConfiguration() *SyncConfiguration {
	if o == nil {
		return nil
	}
	return o.Configuration
}

func (o *Configuration) GetConfigured() *bool {
	if o == nil {
		return nil
	}
	return o.Configured
}

func (o *Configuration) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *Configuration) GetSchedule() *ConfigurationSchedule {
	if o == nil {
		return nil
	}
	return o.Schedule
}
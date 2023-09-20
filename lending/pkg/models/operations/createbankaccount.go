// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v4/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v4/pkg/utils"
	"net/http"
)

type CreateBankAccountRequest struct {
	AccountingBankAccount   *shared.AccountingBankAccount `request:"mediaType=application/json"`
	AllowSyncOnPushComplete *bool                         `default:"true" queryParam:"style=form,explode=true,name=allowSyncOnPushComplete"`
	CompanyID               string                        `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID            string                        `pathParam:"style=simple,explode=false,name=connectionId"`
	// When updating data in the destination platform Codat checks the `sourceModifiedDate` against the `lastupdated` date from the accounting platform, if they're different Codat will return an error suggesting you should initiate another pull of the data. If this is set to `true` then the update will override this check.
	ForceUpdate      *bool `default:"false" queryParam:"style=form,explode=true,name=forceUpdate"`
	TimeoutInMinutes *int  `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

func (c CreateBankAccountRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateBankAccountRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateBankAccountRequest) GetAccountingBankAccount() *shared.AccountingBankAccount {
	if o == nil {
		return nil
	}
	return o.AccountingBankAccount
}

func (o *CreateBankAccountRequest) GetAllowSyncOnPushComplete() *bool {
	if o == nil {
		return nil
	}
	return o.AllowSyncOnPushComplete
}

func (o *CreateBankAccountRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateBankAccountRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateBankAccountRequest) GetForceUpdate() *bool {
	if o == nil {
		return nil
	}
	return o.ForceUpdate
}

func (o *CreateBankAccountRequest) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

type CreateBankAccountResponse struct {
	// Success
	AccountingCreateBankAccountResponse *shared.AccountingCreateBankAccountResponse
	ContentType                         string
	// The request made is not valid.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *CreateBankAccountResponse) GetAccountingCreateBankAccountResponse() *shared.AccountingCreateBankAccountResponse {
	if o == nil {
		return nil
	}
	return o.AccountingCreateBankAccountResponse
}

func (o *CreateBankAccountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateBankAccountResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *CreateBankAccountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateBankAccountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

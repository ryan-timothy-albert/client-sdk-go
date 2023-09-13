// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ErrorMessage struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

func (o *ErrorMessage) GetCanBeRetried() *string {
	if o == nil {
		return nil
	}
	return o.CanBeRetried
}

func (o *ErrorMessage) GetCorrelationID() *string {
	if o == nil {
		return nil
	}
	return o.CorrelationID
}

func (o *ErrorMessage) GetDetailedErrorCode() *int64 {
	if o == nil {
		return nil
	}
	return o.DetailedErrorCode
}

func (o *ErrorMessage) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *ErrorMessage) GetService() *string {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *ErrorMessage) GetStatusCode() *int64 {
	if o == nil {
		return nil
	}
	return o.StatusCode
}
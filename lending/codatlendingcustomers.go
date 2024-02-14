// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package lending

import (
	"bytes"
	"context"
	"fmt"
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/operations"
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/sdkerrors"
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v5/pkg/utils"
	"io"
	"net/http"
)

type CodatLendingCustomers struct {
	sdkConfiguration sdkConfiguration
}

func newCodatLendingCustomers(sdkConfig sdkConfiguration) *CodatLendingCustomers {
	return &CodatLendingCustomers{
		sdkConfiguration: sdkConfig,
	}
}

// Get customer
// The *Get customer* endpoint returns a single customer for a given customerId.
//
// [Customers](https://docs.codat.io/lending-api#/schemas/Customer) are people or organizations that place orders, make payments and receive goods and/or services from the SMB.
//
// Check out our [coverage explorer](https://knowledge.codat.io/supported-features/commerce?view=tab-by-data-type&dataType=commerce-customers) for integrations that support getting a specific customer.
//
// Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/lending-api#/operations/refresh-company-data).
func (s *CodatLendingCustomers) Get(ctx context.Context, request operations.GetCommerceCustomerRequest, opts ...operations.Option) (*operations.GetCommerceCustomerResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/companies/{companyId}/connections/{connectionId}/data/commerce-customers/{customerId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	client := s.sdkConfiguration.SecurityClient

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig == nil {
			retryConfig = &utils.RetryConfig{
				Strategy: "backoff",
				Backoff: &utils.BackoffStrategy{
					InitialInterval: 500,
					MaxInterval:     60000,
					Exponent:        1.5,
					MaxElapsedTime:  3600000,
				},
				RetryConnectionErrors: true,
			}
		} else {
			retryConfig = globalRetryConfig
		}
	}

	httpRes, err := utils.Retry(ctx, utils.Retries{
		Config: retryConfig,
		StatusCodes: []string{
			"408",
			"429",
			"5XX",
		},
	}, func() (*http.Response, error) {
		if req.Body != nil {
			copyBody, err := req.GetBody()
			if err != nil {
				return nil, err
			}
			req.Body = copyBody
		}
		return client.Do(req)
	})
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetCommerceCustomerResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.CommerceCustomer
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.CommerceCustomer = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 402:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 409:
		fallthrough
	case httpRes.StatusCode == 429:
		fallthrough
	case httpRes.StatusCode == 500:
		fallthrough
	case httpRes.StatusCode == 503:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out sdkerrors.ErrorMessage
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}
			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}

// List customers
// The *List customers* endpoint returns a list of [customers](https://docs.codat.io/lending-api#/schemas/Customer) for a given company's connection.
//
// [Customers](https://docs.codat.io/lending-api#/schemas/Customer) are people or organizations that place orders, make payments and recieve goods and/or services from the SMB.
//
// Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/lending-api#/operations/refresh-company-data).
func (s *CodatLendingCustomers) List(ctx context.Context, request operations.ListCommerceCustomersRequest, opts ...operations.Option) (*operations.ListCommerceCustomersResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/companies/{companyId}/connections/{connectionId}/data/commerce-customers", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig == nil {
			retryConfig = &utils.RetryConfig{
				Strategy: "backoff",
				Backoff: &utils.BackoffStrategy{
					InitialInterval: 500,
					MaxInterval:     60000,
					Exponent:        1.5,
					MaxElapsedTime:  3600000,
				},
				RetryConnectionErrors: true,
			}
		} else {
			retryConfig = globalRetryConfig
		}
	}

	httpRes, err := utils.Retry(ctx, utils.Retries{
		Config: retryConfig,
		StatusCodes: []string{
			"408",
			"429",
			"5XX",
		},
	}, func() (*http.Response, error) {
		if req.Body != nil {
			copyBody, err := req.GetBody()
			if err != nil {
				return nil, err
			}
			req.Body = copyBody
		}
		return client.Do(req)
	})
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListCommerceCustomersResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.CommerceCustomers
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.CommerceCustomers = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 402:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 409:
		fallthrough
	case httpRes.StatusCode == 429:
		fallthrough
	case httpRes.StatusCode == 500:
		fallthrough
	case httpRes.StatusCode == 503:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out sdkerrors.ErrorMessage
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}
			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}

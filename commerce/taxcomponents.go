// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package codatcommerce

import (
	"bytes"
	"context"
	"fmt"
	"github.com/codatio/client-sdk-go/commerce/pkg/models/operations"
	"github.com/codatio/client-sdk-go/commerce/pkg/models/shared"
	"github.com/codatio/client-sdk-go/commerce/pkg/utils"
	"io"
	"net/http"
)

// taxComponents - Retrieve standardized data from linked commerce platforms.
type taxComponents struct {
	sdkConfiguration sdkConfiguration
}

func newTaxComponents(sdkConfig sdkConfiguration) *taxComponents {
	return &taxComponents{
		sdkConfiguration: sdkConfig,
	}
}

// Get - Get tax component
// This endpoint returns a specific tax rate from the commerce platform, including tax rate names and values. This supports the mapping of tax rates from the commerce platform to the accounting platform.
func (s *taxComponents) Get(ctx context.Context, request operations.GetTaxComponentRequest, opts ...operations.Option) (*operations.GetTaxComponentResponse, error) {
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
	url, err := utils.GenerateURL(ctx, baseURL, "/companies/{companyId}/connections/{connectionId}/data/commerce-taxComponents/{taxId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/json;q=0.7, application/json;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	client := s.sdkConfiguration.SecurityClient

	retryConfig := o.Retries
	if retryConfig == nil {
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
	}

	httpRes, err := utils.Retry(ctx, utils.Retries{
		Config: retryConfig,
		StatusCodes: []string{
			"408",
			"429",
			"5XX",
		},
	}, func() (*http.Response, error) {
		return client.Do(req)
	})
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetTaxComponentResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.TaxComponent
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.TaxComponent = out
		}
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Schema
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.Schema = out
		}
	case httpRes.StatusCode == 409:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetTaxComponent409ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.GetTaxComponent409ApplicationJSONObject = out
		}
	}

	return res, nil
}

// List - List tax components
// This endpoint returns a lists of tax rates from the commerce platform, including tax rate names and values. This supports the mapping of tax rates from the commerce platform to the accounting platform.
func (s *taxComponents) List(ctx context.Context, request operations.ListTaxComponentsRequest, opts ...operations.Option) (*operations.ListTaxComponentsResponse, error) {
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
	url, err := utils.GenerateURL(ctx, baseURL, "/companies/{companyId}/connections/{connectionId}/data/commerce-taxComponents", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/json;q=0.7, application/json;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

	retryConfig := o.Retries
	if retryConfig == nil {
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
	}

	httpRes, err := utils.Retry(ctx, utils.Retries{
		Config: retryConfig,
		StatusCodes: []string{
			"408",
			"429",
			"5XX",
		},
	}, func() (*http.Response, error) {
		return client.Do(req)
	})
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListTaxComponentsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.TaxComponents
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.TaxComponents = out
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Schema
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.Schema = out
		}
	case httpRes.StatusCode == 409:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListTaxComponents409ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.ListTaxComponents409ApplicationJSONObject = out
		}
	}

	return res, nil
}

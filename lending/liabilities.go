// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package codatlending

import (
	"bytes"
	"context"
	"fmt"
	"github.com/codatio/client-sdk-go/lending/pkg/models/operations"
	"github.com/codatio/client-sdk-go/lending/pkg/models/sdkerrors"
	"github.com/codatio/client-sdk-go/lending/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/pkg/utils"
	"io"
	"net/http"
)

// liabilities - Debt and other liabilities.
type liabilities struct {
	sdkConfiguration sdkConfiguration
}

func newLiabilities(sdkConfig sdkConfiguration) *liabilities {
	return &liabilities{
		sdkConfiguration: sdkConfig,
	}
}

// GetLoanSummary - Get enhanced loan summaries
// The *Get enhanced loan summaries* endpoint returns a summary by integration type of all loans identified in each integration.
//
// The endpoint returns a list of a company's [loan summaries](https://docs.codat.io/lending-api#/schemas/LoanSummary) for each valid data connection.
//
// Make sure you have [synced a company](https://docs.codat.io/lending-api#/operations/refresh-company-data) recently before calling the endpoint.
func (s *liabilities) GetLoanSummary(ctx context.Context, request operations.GetLoanSummaryRequest, opts ...operations.Option) (*operations.GetLoanSummaryResponse, error) {
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
	url, err := utils.GenerateURL(ctx, baseURL, "/companies/{companyId}/reports/enhancedLiabilities/loan", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
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

	res := &operations.GetLoanSummaryResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.LoanSummary
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.LoanSummary = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ErrorMessage
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.ErrorMessage = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// ListLoanTransactions - List enhanced loan transactions
// The *List enhanced loan transactions* endpoint returns all [loan transactions](https://docs.codat.io/lending-api#/schemas/LoanTransactions) identified from a company's accounting, banking, and commerce integrations.
//
// This detail gives analysts a better idea of the loan obligations a company may have.
//
// Make sure you have [synced a company](https://docs.codat.io/lending-api#/operations/refresh-company-data) recently before calling the endpoint.
func (s *liabilities) ListLoanTransactions(ctx context.Context, request operations.ListLoanTransactionsRequest, opts ...operations.Option) (*operations.ListLoanTransactionsResponse, error) {
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
	url, err := utils.GenerateURL(ctx, baseURL, "/companies/{companyId}/reports/enhancedLiabilities/loan/transactions", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
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

	res := &operations.ListLoanTransactionsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.LoanTransactions
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.LoanTransactions = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ErrorMessage
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.ErrorMessage = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

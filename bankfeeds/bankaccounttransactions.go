// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package codatbankfeeds

import (
	"context"
	"fmt"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/operations"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/utils"
	"net/http"
)

// bankAccountTransactions - Bank feed bank accounts
type bankAccountTransactions struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newBankAccountTransactions(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *bankAccountTransactions {
	return &bankAccountTransactions{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// CreateBankTransactions - Create bank transactions
// Posts bank transactions to the accounting package for a given company.
//
// > **Supported Integrations**
// >
// > Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankTransactions) for integrations that support POST methods.
func (s *bankAccountTransactions) CreateBankTransactions(ctx context.Context, request operations.CreateBankTransactionsRequest) (*operations.CreateBankTransactionsResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/companies/{companyId}/connections/{connectionId}/push/bankAccounts/{accountId}/bankTransactions", request, nil)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "BankTransactions", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.CreateBankTransactionsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.CreateBankTransactionsResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateBankTransactionsResponse = out
		}
	}

	return res, nil
}

// GetCreateBankAccountModel - List push options for bank account bank transactions
// Gets the options of pushing bank account transactions.
func (s *bankAccountTransactions) GetCreateBankAccountModel(ctx context.Context, request operations.GetCreateBankAccountModelRequest) (*operations.GetCreateBankAccountModelResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/companies/{companyId}/connections/{connectionId}/options/bankAccounts/{accountId}/bankTransactions", request, nil)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetCreateBankAccountModelResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.PushOption
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PushOption = out
		}
	}

	return res, nil
}

// ListBankAccountTransactions - List bank transactions for bank account
// Gets bank transactions for a given bank account ID
func (s *bankAccountTransactions) ListBankAccountTransactions(ctx context.Context, request operations.ListBankAccountTransactionsRequest) (*operations.ListBankAccountTransactionsResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/companies/{companyId}/connections/{connectionId}/data/bankAccounts/{accountId}/bankTransactions", request, nil)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListBankAccountTransactionsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.BankTransactionsResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.BankTransactionsResponse = out
		}
	}

	return res, nil
}

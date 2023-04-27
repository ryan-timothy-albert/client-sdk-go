# Accounts

## Overview

Accounts

### Available Operations

* [Create](#create) - Create account
* [Get](#get) - Get account
* [GetCreateModel](#getcreatemodel) - Get create account model
* [List](#list) - List accounts

## Create

Creates a new account for a given company.

Required data may vary by integration. To see what data to post, first call [Get create account model](https://docs.codat.io/accounting-api#/operations/get-create-chartOfAccounts-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=chartOfAccounts) for integrations that support creating an account.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.CreateAccountRequest{
        Account: &shared.Account{
            Currency: codataccounting.String("quibusdam"),
            CurrentBalance: codataccounting.Float64(6027.63),
            Description: codataccounting.String("Invoices the business has issued but has not yet collected payment on."),
            FullyQualifiedCategory: codataccounting.String("Asset.Current"),
            FullyQualifiedName: codataccounting.String("Asset.Current.Accounts Receivable"),
            ID: codataccounting.String("1b6266d1-1e44-46c5-8eb5-a8f98e03124e"),
            IsBankAccount: false,
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("nulla"),
            Name: codataccounting.String("Accounts Receivable"),
            NominalCode: codataccounting.String("610"),
            SourceModifiedDate: codataccounting.String("corrupti"),
            Status: shared.AccountStatusEnumActive,
            Type: shared.AccountTypeEnumAsset,
            ValidDatatypeLinks: []shared.ValidDataTypeLinks{
                shared.ValidDataTypeLinks{
                    Links: []string{
                        "error",
                        "deserunt",
                    },
                    Property: codataccounting.String("suscipit"),
                },
                shared.ValidDataTypeLinks{
                    Links: []string{
                        "magnam",
                        "debitis",
                    },
                    Property: codataccounting.String("ipsa"),
                },
                shared.ValidDataTypeLinks{
                    Links: []string{
                        "tempora",
                        "suscipit",
                        "molestiae",
                        "minus",
                    },
                    Property: codataccounting.String("placeat"),
                },
                shared.ValidDataTypeLinks{
                    Links: []string{
                        "iusto",
                        "excepturi",
                        "nisi",
                    },
                    Property: codataccounting.String("recusandae"),
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(836079),
    }

    res, err := s.Accounts.Create(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateAccountResponse != nil {
        // handle response
    }
}
```

## Get

Gets a single account corresponding to the given ID.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetAccountRequest{
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.Accounts.Get(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Account != nil {
        // handle response
    }
}
```

## GetCreateModel

Get create account model. Returns the expected data for the request payload.

See the examples for integration-specific indicative models.

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=chartOfAccounts) for integrations that support creating an account.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetCreateChartOfAccountsModelRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Accounts.GetCreateModel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

## List

Gets the latest accounts for a company

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListAccountsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("ab"),
    }

    res, err := s.Accounts.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Accounts != nil {
        // handle response
    }
}
```
# Transactions

## Overview

Transactions represent debits and credits from a source account.

### Available Operations

* [Create](#create) - Create bank account transactions
* [GetCreateOperation](#getcreateoperation) - Get create operation
* [ListCreateOperations](#listcreateoperations) - List create operations

## Create

﻿The *Create bank account transactions* endpoint creates new [bank account transactions](https://docs.codat.io/bank-feeds-api#/schemas/BankTransactions) for a given company's connection.

[Bank account transactions](https://docs.codat.io/bank-feeds-api#/schemas/BankTransactions) are records of monetary amounts that have moved in and out of an SMB's bank account.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create bank transaction model](https://docs.codat.io/bank-feeds-api#/operations/get-create-bankTransactions-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankTransactions) for integrations that support creating a bank account transactions.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/bank-feeds"
	"github.com/codatio/client-sdk-go/bank-feeds/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bank-feeds/pkg/models/operations"
)

func main() {
    s := codatbankfeeds.New(
        codatbankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Transactions.Create(ctx, operations.CreateBankTransactionsRequest{
        CreateBankTransactions: &shared.CreateBankTransactions{
            AccountID: codatbankfeeds.String("corporis"),
            Transactions: []shared.CreateBankAccountTransaction{
                shared.CreateBankAccountTransaction{
                    Amount: codatbankfeeds.Float64(7506.86),
                    Balance: codatbankfeeds.Float64(3154.28),
                    Date: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
                    Description: codatbankfeeds.String("nemo"),
                    ID: codatbankfeeds.String("5907aff1-a3a2-4fa9-8677-39251aa52c3f"),
                },
            },
        },
        AccountID: "9wg4lep4ush5cxs79pl8sozmsndbaukll3ind4g7buqbm1h2",
        AllowSyncOnPushComplete: codatbankfeeds.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatbankfeeds.Int(662527),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBankTransactionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateBankTransactionsRequest](../../models/operations/createbanktransactionsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |


### Response

**[*operations.CreateBankTransactionsResponse](../../models/operations/createbanktransactionsresponse.md), error**


## GetCreateOperation

Retrieve push operation.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/bank-feeds"
	"github.com/codatio/client-sdk-go/bank-feeds/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bank-feeds/pkg/models/operations"
)

func main() {
    s := codatbankfeeds.New(
        codatbankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Transactions.GetCreateOperation(ctx, operations.GetCreateOperationRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PushOperationKey: "d019da1f-fe78-4f09-bb00-74f15471b5e6",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOperation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetCreateOperationRequest](../../models/operations/getcreateoperationrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.GetCreateOperationResponse](../../models/operations/getcreateoperationresponse.md), error**


## ListCreateOperations

List create operations.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/bank-feeds"
	"github.com/codatio/client-sdk-go/bank-feeds/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bank-feeds/pkg/models/operations"
)

func main() {
    s := codatbankfeeds.New(
        codatbankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Transactions.ListCreateOperations(ctx, operations.ListCreateOperationsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codatbankfeeds.String("-modifiedDate"),
        Page: codatbankfeeds.Int(1),
        PageSize: codatbankfeeds.Int(100),
        Query: codatbankfeeds.String("repudiandae"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListPushOperations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListCreateOperationsRequest](../../models/operations/listcreateoperationsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |


### Response

**[*operations.ListCreateOperationsResponse](../../models/operations/listcreateoperationsresponse.md), error**

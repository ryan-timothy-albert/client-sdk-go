# BankAccountTransactions

## Overview

Bank transactions for bank accounts

### Available Operations

* [Create](#create) - Create bank transactions
* [GetCreateModel](#getcreatemodel) - List push options for bank account bank transactions
* [List](#list) - List bank transactions for bank account
* [ListTransactions](#listtransactions) - List all bank transactions

## Create

Posts bank transactions to the accounting package for a given company.

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankTransactions) for integrations that support POST methods.

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
    req := operations.CreateBankTransactionsRequest{
        BankTransactions: &shared.BankTransactions{
            AccountID: codataccounting.String("quis"),
            Transactions: []shared.BankTransactionLine{
                shared.BankTransactionLine{
                    Amount: 6481.72,
                    Balance: 202.18,
                    ClearedOnDate: codataccounting.String("ipsam"),
                    Counterparty: codataccounting.String("repellendus"),
                    Description: codataccounting.String("sapiente"),
                    ID: codataccounting.String("c2ddf7cc-78ca-41ba-928f-c816742cb739"),
                    ModifiedDate: codataccounting.String("aspernatur"),
                    Reconciled: false,
                    Reference: codataccounting.String("perferendis"),
                    SourceModifiedDate: codataccounting.String("ad"),
                    TransactionType: shared.BankTransactionTypeEnumCheck,
                },
            },
        },
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        AllowSyncOnPushComplete: codataccounting.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(149675),
    }

    res, err := s.BankAccountTransactions.Create(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBankTransactionsResponse != nil {
        // handle response
    }
}
```

## GetCreateModel

Gets the options of pushing bank account transactions.

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
    req := operations.GetCreateBankAccountModelRequest{
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.BankAccountTransactions.GetCreateModel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

## List

Gets bank transactions for a given bank account ID

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
    req := operations.ListBankAccountTransactionsRequest{
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("iste"),
    }

    res, err := s.BankAccountTransactions.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BankTransactionsResponse != nil {
        // handle response
    }
}
```

## ListTransactions

Gets the latest bank transactions for given account ID and company. Doesn't require connection ID.

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
    req := operations.ListBankTransactionsRequest{
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("dolor"),
    }

    res, err := s.BankAccountTransactions.ListTransactions(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BankAccountTransactions != nil {
        // handle response
    }
}
```
# Accounts

## Overview

Where payments are made or received, and bank transactions are recorded.

### Available Operations

* [Get](#get) - Get account
* [List](#list) - List accounts

## Get

Gets a specified bank account for a given company

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/banking"
	"github.com/codatio/client-sdk-go/banking/pkg/models/operations"
)

func main() {
    s := codatbanking.New(
        codatbanking.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetAccountRequest{
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
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

## List

Gets a list of all bank accounts of the SMB, with rich data like balances, account numbers and institutions holdingthe accounts.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/banking"
	"github.com/codatio/client-sdk-go/banking/pkg/models/operations"
)

func main() {
    s := codatbanking.New(
        codatbanking.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListAccountsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codatbanking.String("-modifiedDate"),
        Page: 1,
        PageSize: codatbanking.Int(100),
        Query: codatbanking.String("distinctio"),
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
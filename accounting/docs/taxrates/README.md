# TaxRates

## Overview

Tax rates

### Available Operations

* [Get](#get) - Get tax rate
* [List](#list) - List all tax rates

## Get

Gets the specified tax rate for a given company.

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
    req := operations.GetTaxRateRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        TaxRateID: "inventore",
    }

    res, err := s.TaxRates.Get(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.TaxRate != nil {
        // handle response
    }
}
```

## List

Gets the latest tax rates for a given company.

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
    req := operations.ListTaxRatesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("eligendi"),
    }

    res, err := s.TaxRates.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.TaxRates != nil {
        // handle response
    }
}
```
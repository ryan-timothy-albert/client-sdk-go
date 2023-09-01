# CompanyInfo

## Overview

View company information fetched from the source platform.

### Available Operations

* [FromAccountingPlatform](#fromaccountingplatform) - Get company info
* [FromCommercePlatform](#fromcommerceplatform) - Get company info

## FromAccountingPlatform

Gets the latest basic info for a company.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/lending"
	"github.com/codatio/client-sdk-go/lending/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/pkg/models/operations"
)

func main() {
    s := codatlending.New(
        codatlending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.CompanyInfo.FromAccountingPlatform(ctx, operations.GetAccountingCompanyInfoRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountingCompanyInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetAccountingCompanyInfoRequest](../../models/operations/getaccountingcompanyinforequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |


### Response

**[*operations.GetAccountingCompanyInfoResponse](../../models/operations/getaccountingcompanyinforesponse.md), error**


## FromCommercePlatform

Retrieve information about the company, as seen in the commerce platform.

This may include information like addresses, tax registration details and social media or website information.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/lending"
	"github.com/codatio/client-sdk-go/lending/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/pkg/models/operations"
)

func main() {
    s := codatlending.New(
        codatlending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.CompanyInfo.FromCommercePlatform(ctx, operations.GetCommerceCompanyInfoRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CommerceCompanyInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetCommerceCompanyInfoRequest](../../models/operations/getcommercecompanyinforequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |


### Response

**[*operations.GetCommerceCompanyInfoResponse](../../models/operations/getcommercecompanyinforesponse.md), error**

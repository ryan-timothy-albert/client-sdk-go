# SyncFlowSettings

## Overview

Configure preferences for any given Sync for Commerce company using sync flow.

### Available Operations

* [GetConfigTextSyncFlow](#getconfigtextsyncflow) - Get preferences for text fields
* [GetVisibleAccounts](#getvisibleaccounts) - List visible accounts
* [UpdateConfigTextSyncFlow](#updateconfigtextsyncflow) - Update preferences for text fields
* [UpdateVisibleAccountsSyncFlow](#updatevisibleaccountssyncflow) - Update visible accounts

## GetConfigTextSyncFlow

Return preferences set for the text fields on sync flow.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-commerce"
	"github.com/codatio/client-sdk-go/sync-for-commerce/pkg/models/shared"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.SyncFlowSettings.GetConfigTextSyncFlow(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.LocalizationInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.GetConfigTextSyncFlowResponse](../../models/operations/getconfigtextsyncflowresponse.md), error**


## GetVisibleAccounts

Return accounts which are visible on sync flow.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-commerce"
	"github.com/codatio/client-sdk-go/sync-for-commerce/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-commerce/pkg/models/operations"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.SyncFlowSettings.GetVisibleAccounts(ctx, operations.GetVisibleAccountsRequest{
        ClientID: "e0f467cc-8796-4ed1-91a0-5dfc2ddf7cc7",
        PlatformKey: "8ca1ba92-8fc8-4167-82cb-739205929396",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VisibleAccounts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetVisibleAccountsRequest](../../models/operations/getvisibleaccountsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.GetVisibleAccountsResponse](../../models/operations/getvisibleaccountsresponse.md), error**


## UpdateConfigTextSyncFlow

Set preferences for the text fields on sync flow.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-commerce"
	"github.com/codatio/client-sdk-go/sync-for-commerce/pkg/models/shared"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.SyncFlowSettings.UpdateConfigTextSyncFlow(ctx, map[string]shared.Localization{
        "hic": shared.Localization{
            Required: codatsynccommerce.Bool(false),
            Text: codatsynccommerce.String("saepe"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LocalizationInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [map[string]shared.Localization](../../models//.md)      | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.UpdateConfigTextSyncFlowResponse](../../models/operations/updateconfigtextsyncflowresponse.md), error**


## UpdateVisibleAccountsSyncFlow

Update which accounts are visible on sync flow.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-commerce"
	"github.com/codatio/client-sdk-go/sync-for-commerce/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-commerce/pkg/models/operations"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.SyncFlowSettings.UpdateVisibleAccountsSyncFlow(ctx, operations.UpdateVisibleAccountsSyncFlowRequest{
        VisibleAccounts: &shared.VisibleAccounts{
            VisibleAccounts: []string{
                "fuga",
            },
        },
        PlatformKey: "7596eb10-faaa-4235-ac59-55907aff1a3a",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VisibleAccounts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.UpdateVisibleAccountsSyncFlowRequest](../../models/operations/updatevisibleaccountssyncflowrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |


### Response

**[*operations.UpdateVisibleAccountsSyncFlowResponse](../../models/operations/updatevisibleaccountssyncflowresponse.md), error**

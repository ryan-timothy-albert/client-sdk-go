# Sync

## Overview

Initiate a sync of Sync for Commerce company data into their respective accounting software.

### Available Operations

* [RequestSync](#requestsync) - Run a Commerce sync from the last successful sync
* [RequestSyncForDateRange](#requestsyncfordaterange) - Run a Commerce sync from a given date range

## RequestSync

Run a Commerce sync from the last successful sync up to the date provided (optional), otherwise UtcNow is used.
If there was no previously successful sync, the start date in the config is used.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/synccommerce"
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/operations"
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/shared"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Sync.RequestSync(ctx, operations.RequestSyncRequest{
        SyncToLatestArgs: &shared.SyncToLatestArgs{
            SyncTo: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SyncSummary != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.RequestSyncRequest](../../models/operations/requestsyncrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |


### Response

**[*operations.RequestSyncResponse](../../models/operations/requestsyncresponse.md), error**


## RequestSyncForDateRange

Run a Commerce sync from the specified start date to the specified finish date in the request payload.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/synccommerce"
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/operations"
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/shared"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Sync.RequestSyncForDateRange(ctx, operations.RequestSyncForDateRangeRequest{
        DateRange: &shared.DateRange{
            Finish: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            Start: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SyncSummary != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.RequestSyncForDateRangeRequest](../../models/operations/requestsyncfordaterangerequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


### Response

**[*operations.RequestSyncForDateRangeResponse](../../models/operations/requestsyncfordaterangeresponse.md), error**

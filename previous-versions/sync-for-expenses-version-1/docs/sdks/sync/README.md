# Sync

## Overview

Triggering a new sync of expenses to accounting software.

### Available Operations

* [IntiateSync](#intiatesync) - Initiate sync

## IntiateSync

Initiate sync of pending transactions.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/operations"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Sync.IntiateSync(ctx, operations.IntiateSyncRequest{
        PostSync: &shared.PostSync{
            DatasetIds: []string{
                "c2ddf7cc-78ca-41ba-928f-c816742cb739",
                "20592939-6fea-4759-aeb1-0faaa2352c59",
                "55907aff-1a3a-42fa-9467-739251aa52c3",
                "f5ad019d-a1ff-4e78-b097-b0074f15471b",
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SyncInitiated != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.IntiateSyncRequest](../../models/operations/intiatesyncrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |


### Response

**[*operations.IntiateSyncResponse](../../models/operations/intiatesyncresponse.md), error**

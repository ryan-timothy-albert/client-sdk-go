# ManageData

## Overview

Control how data is retrieved from an integration.

### Available Operations

* [Get](#get) - Get data status
* [GetPullOperation](#getpulloperation) - Get pull operation
* [ListPullOperations](#listpulloperations) - List pull operations
* [RefreshAllDataTypes](#refreshalldatatypes) - Refresh all data
* [RefreshDataType](#refreshdatatype) - Refresh data type

## Get

Get the state of each data type for a company

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
    res, err := s.ManageData.Get(ctx, operations.GetDataStatusRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DataStatusResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetDataStatusRequest](../../models/operations/getdatastatusrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.GetDataStatusResponse](../../models/operations/getdatastatusresponse.md), error**


## GetPullOperation

Retrieve information about a single dataset or pull operation.

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
    res, err := s.ManageData.GetPullOperation(ctx, operations.GetPullOperationRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        DatasetID: "eaed9f0f-e77b-4bc9-a58f-ab8b4b99ab18",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PullOperation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetPullOperationRequest](../../models/operations/getpulloperationrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |


### Response

**[*operations.GetPullOperationResponse](../../models/operations/getpulloperationresponse.md), error**


## ListPullOperations

Gets the pull operation history (datasets) for a given company.

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
    res, err := s.ManageData.ListPullOperations(ctx, operations.ListPullOperationsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codatlending.String("-modifiedDate"),
        Page: codatlending.Int(1),
        PageSize: codatlending.Int(100),
        Query: codatlending.String("ad"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PullOperations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListPullOperationsRequest](../../models/operations/listpulloperationsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.ListPullOperationsResponse](../../models/operations/listpulloperationsresponse.md), error**


## RefreshAllDataTypes

Refreshes all data types with `fetch on first link` set to `true` for a given company.

This is an asynchronous operation, and will bring updated data into Codat from the linked integration for you to view.

[Read more](https://docs.codat.io/core-concepts/data-type-settings) about data type settings and `fetch on first link`.

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
    res, err := s.ManageData.RefreshAllDataTypes(ctx, operations.RefreshAllDataTypesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.RefreshAllDataTypesRequest](../../models/operations/refreshalldatatypesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |


### Response

**[*operations.RefreshAllDataTypesResponse](../../models/operations/refreshalldatatypesresponse.md), error**


## RefreshDataType

Refreshes a given data type for a given company.

This is an asynchronous operation, and will bring updated data into Codat from the linked integration for you to view.

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
    res, err := s.ManageData.RefreshDataType(ctx, operations.RefreshDataTypeRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: codatlending.String("929396fe-a759-46eb-90fa-aa2352c59559"),
        DataType: shared.DataTypeInvoices,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PullOperation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.RefreshDataTypeRequest](../../models/operations/refreshdatatyperequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |


### Response

**[*operations.RefreshDataTypeResponse](../../models/operations/refreshdatatyperesponse.md), error**

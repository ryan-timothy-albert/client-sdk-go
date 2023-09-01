# ListTransactionCategoriesResponse


## Fields

| Field                                                                                                                  | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ContentType`                                                                                                          | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |
| `StatusCode`                                                                                                           | *int*                                                                                                                  | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |
| `RawResponse`                                                                                                          | [*http.Response](https://pkg.go.dev/net/http#Response)                                                                 | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `TransactionCategories`                                                                                                | [*shared.TransactionCategories](../../models/shared/transactioncategories.md)                                          | :heavy_minus_sign:                                                                                                     | Success                                                                                                                |
| `ListTransactionCategories409ApplicationJSONObject`                                                                    | [*ListTransactionCategories409ApplicationJSON](../../models/operations/listtransactioncategories409applicationjson.md) | :heavy_minus_sign:                                                                                                     | The data type's dataset has not been requested or is still syncing.                                                    |
| `Schema`                                                                                                               | [*shared.Schema](../../models/shared/schema.md)                                                                        | :heavy_minus_sign:                                                                                                     | Your `query` parameter was not correctly formed                                                                        |
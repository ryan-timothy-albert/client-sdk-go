# GetCommerceLocationResponse


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `CommerceLocation`                                                  | [*shared.CommerceLocation](../../models/shared/commercelocation.md) | :heavy_minus_sign:                                                  | OK                                                                  |
| `ContentType`                                                       | *string*                                                            | :heavy_check_mark:                                                  | N/A                                                                 |
| `ErrorMessage`                                                      | [*shared.ErrorMessage](../../models/shared/errormessage.md)         | :heavy_minus_sign:                                                  | Your API request was not properly authorized.                       |
| `StatusCode`                                                        | *int*                                                               | :heavy_check_mark:                                                  | N/A                                                                 |
| `RawResponse`                                                       | [*http.Response](https://pkg.go.dev/net/http#Response)              | :heavy_minus_sign:                                                  | N/A                                                                 |
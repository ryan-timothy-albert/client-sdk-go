# ListFilesResponse


## Fields

| Field                                                                                                                                                      | Type                                                                                                                                                       | Required                                                                                                                                                   | Description                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ContentType`                                                                                                                                              | *string*                                                                                                                                                   | :heavy_check_mark:                                                                                                                                         | N/A                                                                                                                                                        |
| `Files`                                                                                                                                                    | [][shared.File](../../models/shared/file.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                         | Success                                                                                                                                                    |
| `StatusCode`                                                                                                                                               | *int*                                                                                                                                                      | :heavy_check_mark:                                                                                                                                         | N/A                                                                                                                                                        |
| `RawResponse`                                                                                                                                              | [*http.Response](https://pkg.go.dev/net/http#Response)                                                                                                     | :heavy_minus_sign:                                                                                                                                         | N/A                                                                                                                                                        |
| `ListFiles404ApplicationJSONObject`                                                                                                                        | [*ListFiles404ApplicationJSON](../../models/operations/listfiles404applicationjson.md)                                                                     | :heavy_minus_sign:                                                                                                                                         | One or more of the resources you referenced could not be found.<br/>This might be because your company or data connection id is wrong, or was already deleted. |
| `Schema`                                                                                                                                                   | [*shared.Schema](../../models/shared/schema.md)                                                                                                            | :heavy_minus_sign:                                                                                                                                         | Your API request was not properly authorized.                                                                                                              |
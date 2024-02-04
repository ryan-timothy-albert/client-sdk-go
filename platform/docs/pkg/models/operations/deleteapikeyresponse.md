# DeleteAPIKeyResponse


## Fields

| Field                                                                                          | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ContentType`                                                                                  | *string*                                                                                       | :heavy_check_mark:                                                                             | HTTP response content type for this operation                                                  |
| `ErrorMessage`                                                                                 | **sdkerrors.ErrorMessage*                                                                      | :heavy_minus_sign:                                                                             | Too many requests were made in a given amount of time. Wait a short period and then try again. |
| `StatusCode`                                                                                   | *int*                                                                                          | :heavy_check_mark:                                                                             | HTTP response status code for this operation                                                   |
| `RawResponse`                                                                                  | [*http.Response](https://pkg.go.dev/net/http#Response)                                         | :heavy_check_mark:                                                                             | Raw HTTP response; suitable for custom response parsing                                        |
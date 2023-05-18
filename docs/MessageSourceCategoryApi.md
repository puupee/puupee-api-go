# \MessageSourceCategoryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppMessageSourceCategoryGet**](MessageSourceCategoryApi.md#ApiAppMessageSourceCategoryGet) | **Get** /api/app/message-source-category | 



## ApiAppMessageSourceCategoryGet

> []MessageSourceCategoryDto ApiAppMessageSourceCategoryGet(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/puupee/puupee-api-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSourceCategoryApi.ApiAppMessageSourceCategoryGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSourceCategoryApi.ApiAppMessageSourceCategoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppMessageSourceCategoryGet`: []MessageSourceCategoryDto
    fmt.Fprintf(os.Stdout, "Response from `MessageSourceCategoryApi.ApiAppMessageSourceCategoryGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppMessageSourceCategoryGetRequest struct via the builder pattern


### Return type

[**[]MessageSourceCategoryDto**](MessageSourceCategoryDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


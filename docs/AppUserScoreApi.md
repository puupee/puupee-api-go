# \AppUserScoreApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppAppUserScorePost**](AppUserScoreApi.md#ApiAppAppUserScorePost) | **Post** /api/app/app-user-score | 



## ApiAppAppUserScorePost

> AppUserScoreDto ApiAppAppUserScorePost(ctx).CreateOrUpdateAppUserScoreDto(createOrUpdateAppUserScoreDto).Execute()



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
    createOrUpdateAppUserScoreDto := *openapiclient.NewCreateOrUpdateAppUserScoreDto() // CreateOrUpdateAppUserScoreDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppUserScoreApi.ApiAppAppUserScorePost(context.Background()).CreateOrUpdateAppUserScoreDto(createOrUpdateAppUserScoreDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppUserScoreApi.ApiAppAppUserScorePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppAppUserScorePost`: AppUserScoreDto
    fmt.Fprintf(os.Stdout, "Response from `AppUserScoreApi.ApiAppAppUserScorePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppAppUserScorePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrUpdateAppUserScoreDto** | [**CreateOrUpdateAppUserScoreDto**](CreateOrUpdateAppUserScoreDto.md) |  | 

### Return type

[**AppUserScoreDto**](AppUserScoreDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


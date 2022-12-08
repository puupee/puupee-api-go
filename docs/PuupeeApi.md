# \PuupeeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppPuupeePullGet**](PuupeeApi.md#ApiAppPuupeePullGet) | **Get** /api/app/puupee/pull | 
[**ApiAppPuupeePushPost**](PuupeeApi.md#ApiAppPuupeePushPost) | **Post** /api/app/puupee/push | 



## ApiAppPuupeePullGet

> PuupeeDtoPagedResultDto ApiAppPuupeePullGet(ctx).AfterVersion(afterVersion).SkipCount(skipCount).MaxResultCount(maxResultCount).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    afterVersion := int64(789) // int64 |  (optional)
    skipCount := int32(56) // int32 |  (optional)
    maxResultCount := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PuupeeApi.ApiAppPuupeePullGet(context.Background()).AfterVersion(afterVersion).SkipCount(skipCount).MaxResultCount(maxResultCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PuupeeApi.ApiAppPuupeePullGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppPuupeePullGet`: PuupeeDtoPagedResultDto
    fmt.Fprintf(os.Stdout, "Response from `PuupeeApi.ApiAppPuupeePullGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppPuupeePullGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **afterVersion** | **int64** |  | 
 **skipCount** | **int32** |  | 
 **maxResultCount** | **int32** |  | 

### Return type

[**PuupeeDtoPagedResultDto**](PuupeeDtoPagedResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppPuupeePushPost

> PuupeeDto ApiAppPuupeePushPost(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewCreateOrUpdatePuupeeDto("Id_example", "Name_example") // CreateOrUpdatePuupeeDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PuupeeApi.ApiAppPuupeePushPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PuupeeApi.ApiAppPuupeePushPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppPuupeePushPost`: PuupeeDto
    fmt.Fprintf(os.Stdout, "Response from `PuupeeApi.ApiAppPuupeePushPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppPuupeePushPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateOrUpdatePuupeeDto**](CreateOrUpdatePuupeeDto.md) |  | 

### Return type

[**PuupeeDto**](PuupeeDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


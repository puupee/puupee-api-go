# \ThumbApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppThumbGet**](ThumbApi.md#ApiAppThumbGet) | **Get** /api/app/thumb | 
[**ApiAppThumbIdDelete**](ThumbApi.md#ApiAppThumbIdDelete) | **Delete** /api/app/thumb/{id} | 
[**ApiAppThumbIdGet**](ThumbApi.md#ApiAppThumbIdGet) | **Get** /api/app/thumb/{id} | 
[**ApiAppThumbPost**](ThumbApi.md#ApiAppThumbPost) | **Post** /api/app/thumb | 



## ApiAppThumbGet

> ThumbDtoPagedResultDto ApiAppThumbGet(ctx).SearchKey(searchKey).MaxResultCount(maxResultCount).Sorting(sorting).SkipCount(skipCount).Execute()



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
    searchKey := "searchKey_example" // string |  (optional)
    maxResultCount := int32(56) // int32 |  (optional)
    sorting := "sorting_example" // string |  (optional)
    skipCount := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ThumbApi.ApiAppThumbGet(context.Background()).SearchKey(searchKey).MaxResultCount(maxResultCount).Sorting(sorting).SkipCount(skipCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbApi.ApiAppThumbGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppThumbGet`: ThumbDtoPagedResultDto
    fmt.Fprintf(os.Stdout, "Response from `ThumbApi.ApiAppThumbGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppThumbGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchKey** | **string** |  | 
 **maxResultCount** | **int32** |  | 
 **sorting** | **string** |  | 
 **skipCount** | **int32** |  | 

### Return type

[**ThumbDtoPagedResultDto**](ThumbDtoPagedResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppThumbIdDelete

> ApiAppThumbIdDelete(ctx, id).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ThumbApi.ApiAppThumbIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbApi.ApiAppThumbIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppThumbIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppThumbIdGet

> ThumbDto ApiAppThumbIdGet(ctx, id).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ThumbApi.ApiAppThumbIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbApi.ApiAppThumbIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppThumbIdGet`: ThumbDto
    fmt.Fprintf(os.Stdout, "Response from `ThumbApi.ApiAppThumbIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppThumbIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ThumbDto**](ThumbDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppThumbPost

> ThumbDto ApiAppThumbPost(ctx).CreateUpdateThumbDto(createUpdateThumbDto).Execute()



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
    createUpdateThumbDto := *openapiclient.NewCreateUpdateThumbDto() // CreateUpdateThumbDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ThumbApi.ApiAppThumbPost(context.Background()).CreateUpdateThumbDto(createUpdateThumbDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbApi.ApiAppThumbPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppThumbPost`: ThumbDto
    fmt.Fprintf(os.Stdout, "Response from `ThumbApi.ApiAppThumbPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppThumbPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUpdateThumbDto** | [**CreateUpdateThumbDto**](CreateUpdateThumbDto.md) |  | 

### Return type

[**ThumbDto**](ThumbDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


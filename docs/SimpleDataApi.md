# \SimpleDataApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppSimpleDataGet**](SimpleDataApi.md#ApiAppSimpleDataGet) | **Get** /api/app/simple-data | 
[**ApiAppSimpleDataIdDelete**](SimpleDataApi.md#ApiAppSimpleDataIdDelete) | **Delete** /api/app/simple-data/{id} | 
[**ApiAppSimpleDataIdGet**](SimpleDataApi.md#ApiAppSimpleDataIdGet) | **Get** /api/app/simple-data/{id} | 
[**ApiAppSimpleDataSavePost**](SimpleDataApi.md#ApiAppSimpleDataSavePost) | **Post** /api/app/simple-data/save | 



## ApiAppSimpleDataGet

> SimpleDataDtoPagedResultDto ApiAppSimpleDataGet(ctx).Collection(collection).Sorting(sorting).SkipCount(skipCount).MaxResultCount(maxResultCount).Execute()



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
    collection := "collection_example" // string |  (optional)
    sorting := "sorting_example" // string |  (optional)
    skipCount := int32(56) // int32 |  (optional)
    maxResultCount := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimpleDataApi.ApiAppSimpleDataGet(context.Background()).Collection(collection).Sorting(sorting).SkipCount(skipCount).MaxResultCount(maxResultCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimpleDataApi.ApiAppSimpleDataGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppSimpleDataGet`: SimpleDataDtoPagedResultDto
    fmt.Fprintf(os.Stdout, "Response from `SimpleDataApi.ApiAppSimpleDataGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppSimpleDataGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collection** | **string** |  | 
 **sorting** | **string** |  | 
 **skipCount** | **int32** |  | 
 **maxResultCount** | **int32** |  | 

### Return type

[**SimpleDataDtoPagedResultDto**](SimpleDataDtoPagedResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppSimpleDataIdDelete

> ApiAppSimpleDataIdDelete(ctx, id).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SimpleDataApi.ApiAppSimpleDataIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimpleDataApi.ApiAppSimpleDataIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiAppSimpleDataIdDeleteRequest struct via the builder pattern


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


## ApiAppSimpleDataIdGet

> SimpleDataDto ApiAppSimpleDataIdGet(ctx, id).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimpleDataApi.ApiAppSimpleDataIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimpleDataApi.ApiAppSimpleDataIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppSimpleDataIdGet`: SimpleDataDto
    fmt.Fprintf(os.Stdout, "Response from `SimpleDataApi.ApiAppSimpleDataIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppSimpleDataIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleDataDto**](SimpleDataDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppSimpleDataSavePost

> SimpleDataDto ApiAppSimpleDataSavePost(ctx).SimpleDataDto(simpleDataDto).Execute()



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
    simpleDataDto := *openapiclient.NewSimpleDataDto() // SimpleDataDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimpleDataApi.ApiAppSimpleDataSavePost(context.Background()).SimpleDataDto(simpleDataDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimpleDataApi.ApiAppSimpleDataSavePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppSimpleDataSavePost`: SimpleDataDto
    fmt.Fprintf(os.Stdout, "Response from `SimpleDataApi.ApiAppSimpleDataSavePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppSimpleDataSavePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **simpleDataDto** | [**SimpleDataDto**](SimpleDataDto.md) |  | 

### Return type

[**SimpleDataDto**](SimpleDataDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


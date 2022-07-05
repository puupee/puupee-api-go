# \TodoApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppTodoGet**](TodoApi.md#ApiAppTodoGet) | **Get** /api/app/todo | 
[**ApiAppTodoIdDelete**](TodoApi.md#ApiAppTodoIdDelete) | **Delete** /api/app/todo/{id} | 
[**ApiAppTodoIdGet**](TodoApi.md#ApiAppTodoIdGet) | **Get** /api/app/todo/{id} | 
[**ApiAppTodoIdPut**](TodoApi.md#ApiAppTodoIdPut) | **Put** /api/app/todo/{id} | 
[**ApiAppTodoPost**](TodoApi.md#ApiAppTodoPost) | **Post** /api/app/todo | 
[**ApiAppTodoSyncGet**](TodoApi.md#ApiAppTodoSyncGet) | **Get** /api/app/todo/sync | 



## ApiAppTodoGet

> TodoDtoPagedResultDto ApiAppTodoGet(ctx).SearchKey(searchKey).IsDone(isDone).TagId(tagId).MaxResultCount(maxResultCount).ParentId(parentId).IsExpired(isExpired).Sorting(sorting).SkipCount(skipCount).Execute()



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
    isDone := true // bool |  (optional)
    tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    maxResultCount := int32(56) // int32 |  (optional)
    parentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    isExpired := true // bool |  (optional)
    sorting := "sorting_example" // string |  (optional)
    skipCount := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TodoApi.ApiAppTodoGet(context.Background()).SearchKey(searchKey).IsDone(isDone).TagId(tagId).MaxResultCount(maxResultCount).ParentId(parentId).IsExpired(isExpired).Sorting(sorting).SkipCount(skipCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TodoApi.ApiAppTodoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppTodoGet`: TodoDtoPagedResultDto
    fmt.Fprintf(os.Stdout, "Response from `TodoApi.ApiAppTodoGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppTodoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchKey** | **string** |  | 
 **isDone** | **bool** |  | 
 **tagId** | **string** |  | 
 **maxResultCount** | **int32** |  | 
 **parentId** | **string** |  | 
 **isExpired** | **bool** |  | 
 **sorting** | **string** |  | 
 **skipCount** | **int32** |  | 

### Return type

[**TodoDtoPagedResultDto**](TodoDtoPagedResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppTodoIdDelete

> ApiAppTodoIdDelete(ctx, id).Execute()



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
    resp, r, err := api_client.TodoApi.ApiAppTodoIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TodoApi.ApiAppTodoIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiAppTodoIdDeleteRequest struct via the builder pattern


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


## ApiAppTodoIdGet

> TodoDto ApiAppTodoIdGet(ctx, id).Execute()



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
    resp, r, err := api_client.TodoApi.ApiAppTodoIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TodoApi.ApiAppTodoIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppTodoIdGet`: TodoDto
    fmt.Fprintf(os.Stdout, "Response from `TodoApi.ApiAppTodoIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppTodoIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TodoDto**](TodoDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppTodoIdPut

> TodoDto ApiAppTodoIdPut(ctx, id).CreateUpdateTodoDto(createUpdateTodoDto).Execute()



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
    createUpdateTodoDto := *openapiclient.NewCreateUpdateTodoDto() // CreateUpdateTodoDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TodoApi.ApiAppTodoIdPut(context.Background(), id).CreateUpdateTodoDto(createUpdateTodoDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TodoApi.ApiAppTodoIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppTodoIdPut`: TodoDto
    fmt.Fprintf(os.Stdout, "Response from `TodoApi.ApiAppTodoIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppTodoIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createUpdateTodoDto** | [**CreateUpdateTodoDto**](CreateUpdateTodoDto.md) |  | 

### Return type

[**TodoDto**](TodoDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppTodoPost

> TodoDto ApiAppTodoPost(ctx).CreateUpdateTodoDto(createUpdateTodoDto).Execute()



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
    createUpdateTodoDto := *openapiclient.NewCreateUpdateTodoDto() // CreateUpdateTodoDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TodoApi.ApiAppTodoPost(context.Background()).CreateUpdateTodoDto(createUpdateTodoDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TodoApi.ApiAppTodoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppTodoPost`: TodoDto
    fmt.Fprintf(os.Stdout, "Response from `TodoApi.ApiAppTodoPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppTodoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUpdateTodoDto** | [**CreateUpdateTodoDto**](CreateUpdateTodoDto.md) |  | 

### Return type

[**TodoDto**](TodoDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppTodoSyncGet

> TodoDtoPagedResultDto ApiAppTodoSyncGet(ctx).AfterVersion(afterVersion).SkipCount(skipCount).MaxResultCount(maxResultCount).Execute()



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
    skipCount := int32(56) // int32 |  (optional) (default to 0)
    maxResultCount := int32(56) // int32 |  (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TodoApi.ApiAppTodoSyncGet(context.Background()).AfterVersion(afterVersion).SkipCount(skipCount).MaxResultCount(maxResultCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TodoApi.ApiAppTodoSyncGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppTodoSyncGet`: TodoDtoPagedResultDto
    fmt.Fprintf(os.Stdout, "Response from `TodoApi.ApiAppTodoSyncGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppTodoSyncGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **afterVersion** | **int64** |  | 
 **skipCount** | **int32** |  | [default to 0]
 **maxResultCount** | **int32** |  | [default to 100]

### Return type

[**TodoDtoPagedResultDto**](TodoDtoPagedResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


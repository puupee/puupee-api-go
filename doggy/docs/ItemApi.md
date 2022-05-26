# \ItemApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppItemGet**](ItemApi.md#ApiAppItemGet) | **Get** /api/app/item | 
[**ApiAppItemIdDelete**](ItemApi.md#ApiAppItemIdDelete) | **Delete** /api/app/item/{id} | 
[**ApiAppItemIdGet**](ItemApi.md#ApiAppItemIdGet) | **Get** /api/app/item/{id} | 
[**ApiAppItemIdPut**](ItemApi.md#ApiAppItemIdPut) | **Put** /api/app/item/{id} | 
[**ApiAppItemPost**](ItemApi.md#ApiAppItemPost) | **Post** /api/app/item | 
[**ApiAppItemSpecialItemsGet**](ItemApi.md#ApiAppItemSpecialItemsGet) | **Get** /api/app/item/special-items | 
[**ApiAppItemSyncGet**](ItemApi.md#ApiAppItemSyncGet) | **Get** /api/app/item/sync | 



## ApiAppItemGet

> ItemDtoPagedResultDto ApiAppItemGet(ctx).ParentItemId(parentItemId).SearchKey(searchKey).Name(name).Types(types).Extension(extension).ContentType(contentType).TagId(tagId).MaxResultCount(maxResultCount).Sorting(sorting).SkipCount(skipCount).Execute()



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
    parentItemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    searchKey := "searchKey_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    types := []openapiclient.ItemType{openapiclient.ItemType("Folder")} // []ItemType |  (optional)
    extension := "extension_example" // string |  (optional)
    contentType := "contentType_example" // string |  (optional)
    tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    maxResultCount := int32(56) // int32 |  (optional)
    sorting := "sorting_example" // string |  (optional)
    skipCount := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ItemApi.ApiAppItemGet(context.Background()).ParentItemId(parentItemId).SearchKey(searchKey).Name(name).Types(types).Extension(extension).ContentType(contentType).TagId(tagId).MaxResultCount(maxResultCount).Sorting(sorting).SkipCount(skipCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.ApiAppItemGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppItemGet`: ItemDtoPagedResultDto
    fmt.Fprintf(os.Stdout, "Response from `ItemApi.ApiAppItemGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppItemGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parentItemId** | **string** |  | 
 **searchKey** | **string** |  | 
 **name** | **string** |  | 
 **types** | [**[]ItemType**](ItemType.md) |  | 
 **extension** | **string** |  | 
 **contentType** | **string** |  | 
 **tagId** | **string** |  | 
 **maxResultCount** | **int32** |  | 
 **sorting** | **string** |  | 
 **skipCount** | **int32** |  | 

### Return type

[**ItemDtoPagedResultDto**](ItemDtoPagedResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppItemIdDelete

> ApiAppItemIdDelete(ctx, id).Execute()



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
    resp, r, err := api_client.ItemApi.ApiAppItemIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.ApiAppItemIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiAppItemIdDeleteRequest struct via the builder pattern


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


## ApiAppItemIdGet

> ItemDto ApiAppItemIdGet(ctx, id).Execute()



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
    resp, r, err := api_client.ItemApi.ApiAppItemIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.ApiAppItemIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppItemIdGet`: ItemDto
    fmt.Fprintf(os.Stdout, "Response from `ItemApi.ApiAppItemIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppItemIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ItemDto**](ItemDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppItemIdPut

> ItemDto ApiAppItemIdPut(ctx, id).CreateUpdateItemDto(createUpdateItemDto).Execute()



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
    createUpdateItemDto := *openapiclient.NewCreateUpdateItemDto() // CreateUpdateItemDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ItemApi.ApiAppItemIdPut(context.Background(), id).CreateUpdateItemDto(createUpdateItemDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.ApiAppItemIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppItemIdPut`: ItemDto
    fmt.Fprintf(os.Stdout, "Response from `ItemApi.ApiAppItemIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppItemIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createUpdateItemDto** | [**CreateUpdateItemDto**](CreateUpdateItemDto.md) |  | 

### Return type

[**ItemDto**](ItemDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppItemPost

> ItemDto ApiAppItemPost(ctx).CreateUpdateItemDto(createUpdateItemDto).Execute()



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
    createUpdateItemDto := *openapiclient.NewCreateUpdateItemDto() // CreateUpdateItemDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ItemApi.ApiAppItemPost(context.Background()).CreateUpdateItemDto(createUpdateItemDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.ApiAppItemPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppItemPost`: ItemDto
    fmt.Fprintf(os.Stdout, "Response from `ItemApi.ApiAppItemPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppItemPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUpdateItemDto** | [**CreateUpdateItemDto**](CreateUpdateItemDto.md) |  | 

### Return type

[**ItemDto**](ItemDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppItemSpecialItemsGet

> SpecialItemDto ApiAppItemSpecialItemsGet(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ItemApi.ApiAppItemSpecialItemsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.ApiAppItemSpecialItemsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppItemSpecialItemsGet`: SpecialItemDto
    fmt.Fprintf(os.Stdout, "Response from `ItemApi.ApiAppItemSpecialItemsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppItemSpecialItemsGetRequest struct via the builder pattern


### Return type

[**SpecialItemDto**](SpecialItemDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppItemSyncGet

> ItemDtoPagedResultDto ApiAppItemSyncGet(ctx).AfterVersion(afterVersion).SkipCount(skipCount).MaxResultCount(maxResultCount).Execute()



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
    resp, r, err := api_client.ItemApi.ApiAppItemSyncGet(context.Background()).AfterVersion(afterVersion).SkipCount(skipCount).MaxResultCount(maxResultCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.ApiAppItemSyncGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppItemSyncGet`: ItemDtoPagedResultDto
    fmt.Fprintf(os.Stdout, "Response from `ItemApi.ApiAppItemSyncGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppItemSyncGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **afterVersion** | **int64** |  | 
 **skipCount** | **int32** |  | [default to 0]
 **maxResultCount** | **int32** |  | [default to 100]

### Return type

[**ItemDtoPagedResultDto**](ItemDtoPagedResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


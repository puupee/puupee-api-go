# \MessageSourceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppMessageSourceIdDelete**](MessageSourceApi.md#ApiAppMessageSourceIdDelete) | **Delete** /api/app/message-source/{id} | 
[**ApiAppMessageSourceIdGet**](MessageSourceApi.md#ApiAppMessageSourceIdGet) | **Get** /api/app/message-source/{id} | 
[**ApiAppMessageSourceIdPut**](MessageSourceApi.md#ApiAppMessageSourceIdPut) | **Put** /api/app/message-source/{id} | 
[**ApiAppMessageSourcePost**](MessageSourceApi.md#ApiAppMessageSourcePost) | **Post** /api/app/message-source | 
[**ApiAppMessageSourcePublishedListCategoryIdGet**](MessageSourceApi.md#ApiAppMessageSourcePublishedListCategoryIdGet) | **Get** /api/app/message-source/published-list/{categoryId} | 



## ApiAppMessageSourceIdDelete

> ApiAppMessageSourceIdDelete(ctx, id).Execute()



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
    r, err := apiClient.MessageSourceApi.ApiAppMessageSourceIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSourceApi.ApiAppMessageSourceIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiAppMessageSourceIdDeleteRequest struct via the builder pattern


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


## ApiAppMessageSourceIdGet

> MessageSourceDto ApiAppMessageSourceIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.MessageSourceApi.ApiAppMessageSourceIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSourceApi.ApiAppMessageSourceIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppMessageSourceIdGet`: MessageSourceDto
    fmt.Fprintf(os.Stdout, "Response from `MessageSourceApi.ApiAppMessageSourceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppMessageSourceIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageSourceDto**](MessageSourceDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppMessageSourceIdPut

> CreateUpdateMessageSourceDto ApiAppMessageSourceIdPut(ctx, id).CreateUpdateMessageSourceDto(createUpdateMessageSourceDto).Execute()



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
    createUpdateMessageSourceDto := *openapiclient.NewCreateUpdateMessageSourceDto() // CreateUpdateMessageSourceDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSourceApi.ApiAppMessageSourceIdPut(context.Background(), id).CreateUpdateMessageSourceDto(createUpdateMessageSourceDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSourceApi.ApiAppMessageSourceIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppMessageSourceIdPut`: CreateUpdateMessageSourceDto
    fmt.Fprintf(os.Stdout, "Response from `MessageSourceApi.ApiAppMessageSourceIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppMessageSourceIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createUpdateMessageSourceDto** | [**CreateUpdateMessageSourceDto**](CreateUpdateMessageSourceDto.md) |  | 

### Return type

[**CreateUpdateMessageSourceDto**](CreateUpdateMessageSourceDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppMessageSourcePost

> CreateUpdateMessageSourceDto ApiAppMessageSourcePost(ctx).CreateUpdateMessageSourceDto(createUpdateMessageSourceDto).Execute()



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
    createUpdateMessageSourceDto := *openapiclient.NewCreateUpdateMessageSourceDto() // CreateUpdateMessageSourceDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSourceApi.ApiAppMessageSourcePost(context.Background()).CreateUpdateMessageSourceDto(createUpdateMessageSourceDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSourceApi.ApiAppMessageSourcePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppMessageSourcePost`: CreateUpdateMessageSourceDto
    fmt.Fprintf(os.Stdout, "Response from `MessageSourceApi.ApiAppMessageSourcePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppMessageSourcePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUpdateMessageSourceDto** | [**CreateUpdateMessageSourceDto**](CreateUpdateMessageSourceDto.md) |  | 

### Return type

[**CreateUpdateMessageSourceDto**](CreateUpdateMessageSourceDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppMessageSourcePublishedListCategoryIdGet

> []MessageSourceDto ApiAppMessageSourcePublishedListCategoryIdGet(ctx, categoryId).Execute()



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
    categoryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSourceApi.ApiAppMessageSourcePublishedListCategoryIdGet(context.Background(), categoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSourceApi.ApiAppMessageSourcePublishedListCategoryIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppMessageSourcePublishedListCategoryIdGet`: []MessageSourceDto
    fmt.Fprintf(os.Stdout, "Response from `MessageSourceApi.ApiAppMessageSourcePublishedListCategoryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppMessageSourcePublishedListCategoryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MessageSourceDto**](MessageSourceDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \ApiKeysApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiApiKeysGet**](ApiKeysApi.md#ApiApiKeysGet) | **Get** /api/api-keys | 
[**ApiApiKeysIdDelete**](ApiKeysApi.md#ApiApiKeysIdDelete) | **Delete** /api/api-keys/{id} | 
[**ApiApiKeysIdGet**](ApiKeysApi.md#ApiApiKeysIdGet) | **Get** /api/api-keys/{id} | 
[**ApiApiKeysIdPut**](ApiKeysApi.md#ApiApiKeysIdPut) | **Put** /api/api-keys/{id} | 
[**ApiApiKeysPost**](ApiKeysApi.md#ApiApiKeysPost) | **Post** /api/api-keys | 



## ApiApiKeysGet

> ApiKeyDtoPagedResultDto ApiApiKeysGet(ctx).Sorting(sorting).SkipCount(skipCount).MaxResultCount(maxResultCount).Execute()



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
    sorting := "sorting_example" // string |  (optional)
    skipCount := int32(56) // int32 |  (optional)
    maxResultCount := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiKeysApi.ApiApiKeysGet(context.Background()).Sorting(sorting).SkipCount(skipCount).MaxResultCount(maxResultCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysApi.ApiApiKeysGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiApiKeysGet`: ApiKeyDtoPagedResultDto
    fmt.Fprintf(os.Stdout, "Response from `ApiKeysApi.ApiApiKeysGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiApiKeysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sorting** | **string** |  | 
 **skipCount** | **int32** |  | 
 **maxResultCount** | **int32** |  | 

### Return type

[**ApiKeyDtoPagedResultDto**](ApiKeyDtoPagedResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiApiKeysIdDelete

> ApiApiKeysIdDelete(ctx, id).Execute()



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
    r, err := apiClient.ApiKeysApi.ApiApiKeysIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysApi.ApiApiKeysIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiApiKeysIdDeleteRequest struct via the builder pattern


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


## ApiApiKeysIdGet

> ApiKeyDto ApiApiKeysIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.ApiKeysApi.ApiApiKeysIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysApi.ApiApiKeysIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiApiKeysIdGet`: ApiKeyDto
    fmt.Fprintf(os.Stdout, "Response from `ApiKeysApi.ApiApiKeysIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiApiKeysIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiKeyDto**](ApiKeyDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiApiKeysIdPut

> ApiKeyDto ApiApiKeysIdPut(ctx, id).Body(body).Execute()



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
    body := *openapiclient.NewApiKeyUpdateDto("Name_example") // ApiKeyUpdateDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiKeysApi.ApiApiKeysIdPut(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysApi.ApiApiKeysIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiApiKeysIdPut`: ApiKeyDto
    fmt.Fprintf(os.Stdout, "Response from `ApiKeysApi.ApiApiKeysIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiApiKeysIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiKeyUpdateDto**](ApiKeyUpdateDto.md) |  | 

### Return type

[**ApiKeyDto**](ApiKeyDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiApiKeysPost

> ApiKeyDto ApiApiKeysPost(ctx).Body(body).Execute()



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
    body := *openapiclient.NewApiKeyCreateDto("Name_example", "Key_example") // ApiKeyCreateDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiKeysApi.ApiApiKeysPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysApi.ApiApiKeysPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiApiKeysPost`: ApiKeyDto
    fmt.Fprintf(os.Stdout, "Response from `ApiKeysApi.ApiApiKeysPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiApiKeysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiKeyCreateDto**](ApiKeyCreateDto.md) |  | 

### Return type

[**ApiKeyDto**](ApiKeyDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


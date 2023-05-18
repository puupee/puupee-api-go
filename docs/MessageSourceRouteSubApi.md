# \MessageSourceRouteSubApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppMessageSourceRouteSubGet**](MessageSourceRouteSubApi.md#ApiAppMessageSourceRouteSubGet) | **Get** /api/app/message-source-route-sub | 
[**ApiAppMessageSourceRouteSubIdDelete**](MessageSourceRouteSubApi.md#ApiAppMessageSourceRouteSubIdDelete) | **Delete** /api/app/message-source-route-sub/{id} | 
[**ApiAppMessageSourceRouteSubIdGet**](MessageSourceRouteSubApi.md#ApiAppMessageSourceRouteSubIdGet) | **Get** /api/app/message-source-route-sub/{id} | 
[**ApiAppMessageSourceRouteSubIdPut**](MessageSourceRouteSubApi.md#ApiAppMessageSourceRouteSubIdPut) | **Put** /api/app/message-source-route-sub/{id} | 
[**ApiAppMessageSourceRouteSubPost**](MessageSourceRouteSubApi.md#ApiAppMessageSourceRouteSubPost) | **Post** /api/app/message-source-route-sub | 



## ApiAppMessageSourceRouteSubGet

> []MessageSourceRouteSubDto ApiAppMessageSourceRouteSubGet(ctx).Execute()



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
    resp, r, err := apiClient.MessageSourceRouteSubApi.ApiAppMessageSourceRouteSubGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSourceRouteSubApi.ApiAppMessageSourceRouteSubGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppMessageSourceRouteSubGet`: []MessageSourceRouteSubDto
    fmt.Fprintf(os.Stdout, "Response from `MessageSourceRouteSubApi.ApiAppMessageSourceRouteSubGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppMessageSourceRouteSubGetRequest struct via the builder pattern


### Return type

[**[]MessageSourceRouteSubDto**](MessageSourceRouteSubDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppMessageSourceRouteSubIdDelete

> ApiAppMessageSourceRouteSubIdDelete(ctx, id).Execute()



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
    r, err := apiClient.MessageSourceRouteSubApi.ApiAppMessageSourceRouteSubIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSourceRouteSubApi.ApiAppMessageSourceRouteSubIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiAppMessageSourceRouteSubIdDeleteRequest struct via the builder pattern


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


## ApiAppMessageSourceRouteSubIdGet

> MessageSourceRouteSubDto ApiAppMessageSourceRouteSubIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.MessageSourceRouteSubApi.ApiAppMessageSourceRouteSubIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSourceRouteSubApi.ApiAppMessageSourceRouteSubIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppMessageSourceRouteSubIdGet`: MessageSourceRouteSubDto
    fmt.Fprintf(os.Stdout, "Response from `MessageSourceRouteSubApi.ApiAppMessageSourceRouteSubIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppMessageSourceRouteSubIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageSourceRouteSubDto**](MessageSourceRouteSubDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppMessageSourceRouteSubIdPut

> ApiAppMessageSourceRouteSubIdPut(ctx, id).CreateUpdateMessageSourceRouteSubDto(createUpdateMessageSourceRouteSubDto).Execute()



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
    createUpdateMessageSourceRouteSubDto := *openapiclient.NewCreateUpdateMessageSourceRouteSubDto() // CreateUpdateMessageSourceRouteSubDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MessageSourceRouteSubApi.ApiAppMessageSourceRouteSubIdPut(context.Background(), id).CreateUpdateMessageSourceRouteSubDto(createUpdateMessageSourceRouteSubDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSourceRouteSubApi.ApiAppMessageSourceRouteSubIdPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiAppMessageSourceRouteSubIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createUpdateMessageSourceRouteSubDto** | [**CreateUpdateMessageSourceRouteSubDto**](CreateUpdateMessageSourceRouteSubDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppMessageSourceRouteSubPost

> ApiAppMessageSourceRouteSubPost(ctx).CreateUpdateMessageSourceRouteSubDto(createUpdateMessageSourceRouteSubDto).Execute()



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
    createUpdateMessageSourceRouteSubDto := *openapiclient.NewCreateUpdateMessageSourceRouteSubDto() // CreateUpdateMessageSourceRouteSubDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MessageSourceRouteSubApi.ApiAppMessageSourceRouteSubPost(context.Background()).CreateUpdateMessageSourceRouteSubDto(createUpdateMessageSourceRouteSubDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSourceRouteSubApi.ApiAppMessageSourceRouteSubPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppMessageSourceRouteSubPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUpdateMessageSourceRouteSubDto** | [**CreateUpdateMessageSourceRouteSubDto**](CreateUpdateMessageSourceRouteSubDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


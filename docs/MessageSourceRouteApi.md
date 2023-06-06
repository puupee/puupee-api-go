# \MessageSourceRouteApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppMessageSourceRouteIdDelete**](MessageSourceRouteApi.md#ApiAppMessageSourceRouteIdDelete) | **Delete** /api/app/message-source-route/{id} | 
[**ApiAppMessageSourceRouteIdGet**](MessageSourceRouteApi.md#ApiAppMessageSourceRouteIdGet) | **Get** /api/app/message-source-route/{id} | 
[**ApiAppMessageSourceRouteIdPut**](MessageSourceRouteApi.md#ApiAppMessageSourceRouteIdPut) | **Put** /api/app/message-source-route/{id} | 
[**ApiAppMessageSourceRoutePost**](MessageSourceRouteApi.md#ApiAppMessageSourceRoutePost) | **Post** /api/app/message-source-route | 
[**ApiAppMessageSourceRoutePublishedListSourceIdGet**](MessageSourceRouteApi.md#ApiAppMessageSourceRoutePublishedListSourceIdGet) | **Get** /api/app/message-source-route/published-list/{sourceId} | 



## ApiAppMessageSourceRouteIdDelete

> ApiAppMessageSourceRouteIdDelete(ctx, id).Execute()



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
    r, err := apiClient.MessageSourceRouteApi.ApiAppMessageSourceRouteIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSourceRouteApi.ApiAppMessageSourceRouteIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiAppMessageSourceRouteIdDeleteRequest struct via the builder pattern


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


## ApiAppMessageSourceRouteIdGet

> MessageSourceRouteDto ApiAppMessageSourceRouteIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.MessageSourceRouteApi.ApiAppMessageSourceRouteIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSourceRouteApi.ApiAppMessageSourceRouteIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppMessageSourceRouteIdGet`: MessageSourceRouteDto
    fmt.Fprintf(os.Stdout, "Response from `MessageSourceRouteApi.ApiAppMessageSourceRouteIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppMessageSourceRouteIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageSourceRouteDto**](MessageSourceRouteDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppMessageSourceRouteIdPut

> ApiAppMessageSourceRouteIdPut(ctx, id).Body(body).Execute()



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
    body := *openapiclient.NewCreateUpdateMessageSourceRouteDto() // CreateUpdateMessageSourceRouteDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MessageSourceRouteApi.ApiAppMessageSourceRouteIdPut(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSourceRouteApi.ApiAppMessageSourceRouteIdPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiAppMessageSourceRouteIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateUpdateMessageSourceRouteDto**](CreateUpdateMessageSourceRouteDto.md) |  | 

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


## ApiAppMessageSourceRoutePost

> ApiAppMessageSourceRoutePost(ctx).Body(body).Execute()



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
    body := *openapiclient.NewCreateUpdateMessageSourceRouteDto() // CreateUpdateMessageSourceRouteDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MessageSourceRouteApi.ApiAppMessageSourceRoutePost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSourceRouteApi.ApiAppMessageSourceRoutePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppMessageSourceRoutePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateUpdateMessageSourceRouteDto**](CreateUpdateMessageSourceRouteDto.md) |  | 

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


## ApiAppMessageSourceRoutePublishedListSourceIdGet

> []MessageSourceRouteDto ApiAppMessageSourceRoutePublishedListSourceIdGet(ctx, sourceId).Execute()



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
    sourceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSourceRouteApi.ApiAppMessageSourceRoutePublishedListSourceIdGet(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSourceRouteApi.ApiAppMessageSourceRoutePublishedListSourceIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppMessageSourceRoutePublishedListSourceIdGet`: []MessageSourceRouteDto
    fmt.Fprintf(os.Stdout, "Response from `MessageSourceRouteApi.ApiAppMessageSourceRoutePublishedListSourceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppMessageSourceRoutePublishedListSourceIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MessageSourceRouteDto**](MessageSourceRouteDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


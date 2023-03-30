# \MessageTemplateApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppMessageTemplateGet**](MessageTemplateApi.md#ApiAppMessageTemplateGet) | **Get** /api/app/message-template | 
[**ApiAppMessageTemplateIdDelete**](MessageTemplateApi.md#ApiAppMessageTemplateIdDelete) | **Delete** /api/app/message-template/{id} | 
[**ApiAppMessageTemplateIdGet**](MessageTemplateApi.md#ApiAppMessageTemplateIdGet) | **Get** /api/app/message-template/{id} | 
[**ApiAppMessageTemplateIdPut**](MessageTemplateApi.md#ApiAppMessageTemplateIdPut) | **Put** /api/app/message-template/{id} | 
[**ApiAppMessageTemplatePost**](MessageTemplateApi.md#ApiAppMessageTemplatePost) | **Post** /api/app/message-template | 



## ApiAppMessageTemplateGet

> []MessageTemplateDto ApiAppMessageTemplateGet(ctx).Execute()



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageTemplateApi.ApiAppMessageTemplateGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageTemplateApi.ApiAppMessageTemplateGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppMessageTemplateGet`: []MessageTemplateDto
    fmt.Fprintf(os.Stdout, "Response from `MessageTemplateApi.ApiAppMessageTemplateGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppMessageTemplateGetRequest struct via the builder pattern


### Return type

[**[]MessageTemplateDto**](MessageTemplateDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppMessageTemplateIdDelete

> ApiAppMessageTemplateIdDelete(ctx, id).Execute()



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageTemplateApi.ApiAppMessageTemplateIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageTemplateApi.ApiAppMessageTemplateIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiAppMessageTemplateIdDeleteRequest struct via the builder pattern


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


## ApiAppMessageTemplateIdGet

> MessageTemplateDto ApiAppMessageTemplateIdGet(ctx, id).Execute()



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageTemplateApi.ApiAppMessageTemplateIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageTemplateApi.ApiAppMessageTemplateIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppMessageTemplateIdGet`: MessageTemplateDto
    fmt.Fprintf(os.Stdout, "Response from `MessageTemplateApi.ApiAppMessageTemplateIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppMessageTemplateIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageTemplateDto**](MessageTemplateDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppMessageTemplateIdPut

> MessageTemplateDto ApiAppMessageTemplateIdPut(ctx, id).Body(body).Execute()



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
    body := *openapiclient.NewCreateOrUpdateMessageTemplateDto() // CreateOrUpdateMessageTemplateDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageTemplateApi.ApiAppMessageTemplateIdPut(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageTemplateApi.ApiAppMessageTemplateIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppMessageTemplateIdPut`: MessageTemplateDto
    fmt.Fprintf(os.Stdout, "Response from `MessageTemplateApi.ApiAppMessageTemplateIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppMessageTemplateIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateOrUpdateMessageTemplateDto**](CreateOrUpdateMessageTemplateDto.md) |  | 

### Return type

[**MessageTemplateDto**](MessageTemplateDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppMessageTemplatePost

> MessageTemplateDto ApiAppMessageTemplatePost(ctx).Body(body).Execute()



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
    body := *openapiclient.NewCreateOrUpdateMessageTemplateDto() // CreateOrUpdateMessageTemplateDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageTemplateApi.ApiAppMessageTemplatePost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageTemplateApi.ApiAppMessageTemplatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppMessageTemplatePost`: MessageTemplateDto
    fmt.Fprintf(os.Stdout, "Response from `MessageTemplateApi.ApiAppMessageTemplatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppMessageTemplatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateOrUpdateMessageTemplateDto**](CreateOrUpdateMessageTemplateDto.md) |  | 

### Return type

[**MessageTemplateDto**](MessageTemplateDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


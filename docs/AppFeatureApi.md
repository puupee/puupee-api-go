# \AppFeatureApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppAppFeatureGet**](AppFeatureApi.md#ApiAppAppFeatureGet) | **Get** /api/app/app-feature | 
[**ApiAppAppFeatureIdDelete**](AppFeatureApi.md#ApiAppAppFeatureIdDelete) | **Delete** /api/app/app-feature/{id} | 
[**ApiAppAppFeaturePost**](AppFeatureApi.md#ApiAppAppFeaturePost) | **Post** /api/app/app-feature | 
[**ApiAppAppFeaturePut**](AppFeatureApi.md#ApiAppAppFeaturePut) | **Put** /api/app/app-feature | 



## ApiAppAppFeatureGet

> []AppFeatureDto ApiAppAppFeatureGet(ctx).AppId(appId).Execute()



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
    appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppFeatureApi.ApiAppAppFeatureGet(context.Background()).AppId(appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppFeatureApi.ApiAppAppFeatureGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppAppFeatureGet`: []AppFeatureDto
    fmt.Fprintf(os.Stdout, "Response from `AppFeatureApi.ApiAppAppFeatureGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppAppFeatureGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string** |  | 

### Return type

[**[]AppFeatureDto**](AppFeatureDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppAppFeatureIdDelete

> ApiAppAppFeatureIdDelete(ctx, id).Execute()



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
    resp, r, err := apiClient.AppFeatureApi.ApiAppAppFeatureIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppFeatureApi.ApiAppAppFeatureIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiAppAppFeatureIdDeleteRequest struct via the builder pattern


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


## ApiAppAppFeaturePost

> AppFeatureDto ApiAppAppFeaturePost(ctx).Body(body).Execute()



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
    body := *openapiclient.NewCreateOrUpdateAppFeatureDto() // CreateOrUpdateAppFeatureDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppFeatureApi.ApiAppAppFeaturePost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppFeatureApi.ApiAppAppFeaturePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppAppFeaturePost`: AppFeatureDto
    fmt.Fprintf(os.Stdout, "Response from `AppFeatureApi.ApiAppAppFeaturePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppAppFeaturePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateOrUpdateAppFeatureDto**](CreateOrUpdateAppFeatureDto.md) |  | 

### Return type

[**AppFeatureDto**](AppFeatureDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppAppFeaturePut

> AppFeatureDto ApiAppAppFeaturePut(ctx).Body(body).Execute()



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
    body := *openapiclient.NewCreateOrUpdateAppFeatureDto() // CreateOrUpdateAppFeatureDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppFeatureApi.ApiAppAppFeaturePut(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppFeatureApi.ApiAppAppFeaturePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppAppFeaturePut`: AppFeatureDto
    fmt.Fprintf(os.Stdout, "Response from `AppFeatureApi.ApiAppAppFeaturePut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppAppFeaturePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateOrUpdateAppFeatureDto**](CreateOrUpdateAppFeatureDto.md) |  | 

### Return type

[**AppFeatureDto**](AppFeatureDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


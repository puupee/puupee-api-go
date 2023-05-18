# \AppSdkApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppAppSdkGet**](AppSdkApi.md#ApiAppAppSdkGet) | **Get** /api/app/app-sdk | 
[**ApiAppAppSdkIdDelete**](AppSdkApi.md#ApiAppAppSdkIdDelete) | **Delete** /api/app/app-sdk/{id} | 
[**ApiAppAppSdkIdPut**](AppSdkApi.md#ApiAppAppSdkIdPut) | **Put** /api/app/app-sdk/{id} | 
[**ApiAppAppSdkPost**](AppSdkApi.md#ApiAppAppSdkPost) | **Post** /api/app/app-sdk | 



## ApiAppAppSdkGet

> []AppSdkDto ApiAppAppSdkGet(ctx).Execute()



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
    resp, r, err := apiClient.AppSdkApi.ApiAppAppSdkGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppSdkApi.ApiAppAppSdkGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppAppSdkGet`: []AppSdkDto
    fmt.Fprintf(os.Stdout, "Response from `AppSdkApi.ApiAppAppSdkGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppAppSdkGetRequest struct via the builder pattern


### Return type

[**[]AppSdkDto**](AppSdkDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppAppSdkIdDelete

> ApiAppAppSdkIdDelete(ctx, id).Execute()



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
    r, err := apiClient.AppSdkApi.ApiAppAppSdkIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppSdkApi.ApiAppAppSdkIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiAppAppSdkIdDeleteRequest struct via the builder pattern


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


## ApiAppAppSdkIdPut

> AppSdkDto ApiAppAppSdkIdPut(ctx, id).CreateOrUpdateAppSdkDto(createOrUpdateAppSdkDto).Execute()



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
    createOrUpdateAppSdkDto := *openapiclient.NewCreateOrUpdateAppSdkDto() // CreateOrUpdateAppSdkDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppSdkApi.ApiAppAppSdkIdPut(context.Background(), id).CreateOrUpdateAppSdkDto(createOrUpdateAppSdkDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppSdkApi.ApiAppAppSdkIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppAppSdkIdPut`: AppSdkDto
    fmt.Fprintf(os.Stdout, "Response from `AppSdkApi.ApiAppAppSdkIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppAppSdkIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrUpdateAppSdkDto** | [**CreateOrUpdateAppSdkDto**](CreateOrUpdateAppSdkDto.md) |  | 

### Return type

[**AppSdkDto**](AppSdkDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppAppSdkPost

> AppSdkDto ApiAppAppSdkPost(ctx).CreateOrUpdateAppSdkDto(createOrUpdateAppSdkDto).Execute()



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
    createOrUpdateAppSdkDto := *openapiclient.NewCreateOrUpdateAppSdkDto() // CreateOrUpdateAppSdkDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppSdkApi.ApiAppAppSdkPost(context.Background()).CreateOrUpdateAppSdkDto(createOrUpdateAppSdkDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppSdkApi.ApiAppAppSdkPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppAppSdkPost`: AppSdkDto
    fmt.Fprintf(os.Stdout, "Response from `AppSdkApi.ApiAppAppSdkPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppAppSdkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrUpdateAppSdkDto** | [**CreateOrUpdateAppSdkDto**](CreateOrUpdateAppSdkDto.md) |  | 

### Return type

[**AppSdkDto**](AppSdkDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


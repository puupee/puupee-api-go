# \AppReleaseApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppAppReleaseGet**](AppReleaseApi.md#ApiAppAppReleaseGet) | **Get** /api/app/app-release | 
[**ApiAppAppReleaseIdDelete**](AppReleaseApi.md#ApiAppAppReleaseIdDelete) | **Delete** /api/app/app-release/{id} | 
[**ApiAppAppReleaseIdGet**](AppReleaseApi.md#ApiAppAppReleaseIdGet) | **Get** /api/app/app-release/{id} | 
[**ApiAppAppReleaseIdPut**](AppReleaseApi.md#ApiAppAppReleaseIdPut) | **Put** /api/app/app-release/{id} | 
[**ApiAppAppReleaseLatestGet**](AppReleaseApi.md#ApiAppAppReleaseLatestGet) | **Get** /api/app/app-release/latest | 
[**ApiAppAppReleasePost**](AppReleaseApi.md#ApiAppAppReleasePost) | **Post** /api/app/app-release | 



## ApiAppAppReleaseGet

> AppReleaseDtoPagedResultDto ApiAppAppReleaseGet(ctx).AppId(appId).Environment(environment).PlatformName(platformName).PlatformValue(platformValue).Sorting(sorting).SkipCount(skipCount).MaxResultCount(maxResultCount).Execute()



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
    appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    environment := "environment_example" // string |  (optional)
    platformName := "platformName_example" // string |  (optional)
    platformValue := "platformValue_example" // string |  (optional)
    sorting := "sorting_example" // string |  (optional)
    skipCount := int32(56) // int32 |  (optional)
    maxResultCount := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppReleaseApi.ApiAppAppReleaseGet(context.Background()).AppId(appId).Environment(environment).PlatformName(platformName).PlatformValue(platformValue).Sorting(sorting).SkipCount(skipCount).MaxResultCount(maxResultCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppReleaseApi.ApiAppAppReleaseGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppAppReleaseGet`: AppReleaseDtoPagedResultDto
    fmt.Fprintf(os.Stdout, "Response from `AppReleaseApi.ApiAppAppReleaseGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppAppReleaseGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string** |  | 
 **environment** | **string** |  | 
 **platformName** | **string** |  | 
 **platformValue** | **string** |  | 
 **sorting** | **string** |  | 
 **skipCount** | **int32** |  | 
 **maxResultCount** | **int32** |  | 

### Return type

[**AppReleaseDtoPagedResultDto**](AppReleaseDtoPagedResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppAppReleaseIdDelete

> ApiAppAppReleaseIdDelete(ctx, id).Execute()



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
    r, err := apiClient.AppReleaseApi.ApiAppAppReleaseIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppReleaseApi.ApiAppAppReleaseIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiAppAppReleaseIdDeleteRequest struct via the builder pattern


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


## ApiAppAppReleaseIdGet

> AppReleaseDto ApiAppAppReleaseIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.AppReleaseApi.ApiAppAppReleaseIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppReleaseApi.ApiAppAppReleaseIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppAppReleaseIdGet`: AppReleaseDto
    fmt.Fprintf(os.Stdout, "Response from `AppReleaseApi.ApiAppAppReleaseIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppAppReleaseIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppReleaseDto**](AppReleaseDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppAppReleaseIdPut

> AppReleaseDto ApiAppAppReleaseIdPut(ctx, id).CreateOrUpdateAppReleaseDto(createOrUpdateAppReleaseDto).Execute()



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
    createOrUpdateAppReleaseDto := *openapiclient.NewCreateOrUpdateAppReleaseDto() // CreateOrUpdateAppReleaseDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppReleaseApi.ApiAppAppReleaseIdPut(context.Background(), id).CreateOrUpdateAppReleaseDto(createOrUpdateAppReleaseDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppReleaseApi.ApiAppAppReleaseIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppAppReleaseIdPut`: AppReleaseDto
    fmt.Fprintf(os.Stdout, "Response from `AppReleaseApi.ApiAppAppReleaseIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppAppReleaseIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrUpdateAppReleaseDto** | [**CreateOrUpdateAppReleaseDto**](CreateOrUpdateAppReleaseDto.md) |  | 

### Return type

[**AppReleaseDto**](AppReleaseDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppAppReleaseLatestGet

> AppReleaseDto ApiAppAppReleaseLatestGet(ctx).AppName(appName).Platform(platform).ProductType(productType).Environment(environment).Execute()



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
    appName := "appName_example" // string |  (optional)
    platform := "platform_example" // string |  (optional)
    productType := "productType_example" // string |  (optional)
    environment := "environment_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppReleaseApi.ApiAppAppReleaseLatestGet(context.Background()).AppName(appName).Platform(platform).ProductType(productType).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppReleaseApi.ApiAppAppReleaseLatestGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppAppReleaseLatestGet`: AppReleaseDto
    fmt.Fprintf(os.Stdout, "Response from `AppReleaseApi.ApiAppAppReleaseLatestGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppAppReleaseLatestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appName** | **string** |  | 
 **platform** | **string** |  | 
 **productType** | **string** |  | 
 **environment** | **string** |  | 

### Return type

[**AppReleaseDto**](AppReleaseDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppAppReleasePost

> AppReleaseDto ApiAppAppReleasePost(ctx).CreateOrUpdateAppReleaseDto(createOrUpdateAppReleaseDto).Execute()



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
    createOrUpdateAppReleaseDto := *openapiclient.NewCreateOrUpdateAppReleaseDto() // CreateOrUpdateAppReleaseDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppReleaseApi.ApiAppAppReleasePost(context.Background()).CreateOrUpdateAppReleaseDto(createOrUpdateAppReleaseDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppReleaseApi.ApiAppAppReleasePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppAppReleasePost`: AppReleaseDto
    fmt.Fprintf(os.Stdout, "Response from `AppReleaseApi.ApiAppAppReleasePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppAppReleasePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrUpdateAppReleaseDto** | [**CreateOrUpdateAppReleaseDto**](CreateOrUpdateAppReleaseDto.md) |  | 

### Return type

[**AppReleaseDto**](AppReleaseDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


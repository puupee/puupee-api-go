# \TenantApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiMultiTenancyTenantsGet**](TenantApi.md#ApiMultiTenancyTenantsGet) | **Get** /api/multi-tenancy/tenants | 
[**ApiMultiTenancyTenantsIdDefaultConnectionStringDelete**](TenantApi.md#ApiMultiTenancyTenantsIdDefaultConnectionStringDelete) | **Delete** /api/multi-tenancy/tenants/{id}/default-connection-string | 
[**ApiMultiTenancyTenantsIdDefaultConnectionStringGet**](TenantApi.md#ApiMultiTenancyTenantsIdDefaultConnectionStringGet) | **Get** /api/multi-tenancy/tenants/{id}/default-connection-string | 
[**ApiMultiTenancyTenantsIdDefaultConnectionStringPut**](TenantApi.md#ApiMultiTenancyTenantsIdDefaultConnectionStringPut) | **Put** /api/multi-tenancy/tenants/{id}/default-connection-string | 
[**ApiMultiTenancyTenantsIdDelete**](TenantApi.md#ApiMultiTenancyTenantsIdDelete) | **Delete** /api/multi-tenancy/tenants/{id} | 
[**ApiMultiTenancyTenantsIdGet**](TenantApi.md#ApiMultiTenancyTenantsIdGet) | **Get** /api/multi-tenancy/tenants/{id} | 
[**ApiMultiTenancyTenantsIdPut**](TenantApi.md#ApiMultiTenancyTenantsIdPut) | **Put** /api/multi-tenancy/tenants/{id} | 
[**ApiMultiTenancyTenantsPost**](TenantApi.md#ApiMultiTenancyTenantsPost) | **Post** /api/multi-tenancy/tenants | 



## ApiMultiTenancyTenantsGet

> TenantDtoPagedResultDto ApiMultiTenancyTenantsGet(ctx).Filter(filter).Sorting(sorting).SkipCount(skipCount).MaxResultCount(maxResultCount).Execute()



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
    filter := "filter_example" // string |  (optional)
    sorting := "sorting_example" // string |  (optional)
    skipCount := int32(56) // int32 |  (optional)
    maxResultCount := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantApi.ApiMultiTenancyTenantsGet(context.Background()).Filter(filter).Sorting(sorting).SkipCount(skipCount).MaxResultCount(maxResultCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.ApiMultiTenancyTenantsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiMultiTenancyTenantsGet`: TenantDtoPagedResultDto
    fmt.Fprintf(os.Stdout, "Response from `TenantApi.ApiMultiTenancyTenantsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiMultiTenancyTenantsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** |  | 
 **sorting** | **string** |  | 
 **skipCount** | **int32** |  | 
 **maxResultCount** | **int32** |  | 

### Return type

[**TenantDtoPagedResultDto**](TenantDtoPagedResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiMultiTenancyTenantsIdDefaultConnectionStringDelete

> ApiMultiTenancyTenantsIdDefaultConnectionStringDelete(ctx, id).Execute()



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
    r, err := apiClient.TenantApi.ApiMultiTenancyTenantsIdDefaultConnectionStringDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.ApiMultiTenancyTenantsIdDefaultConnectionStringDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiMultiTenancyTenantsIdDefaultConnectionStringDeleteRequest struct via the builder pattern


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


## ApiMultiTenancyTenantsIdDefaultConnectionStringGet

> string ApiMultiTenancyTenantsIdDefaultConnectionStringGet(ctx, id).Execute()



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
    resp, r, err := apiClient.TenantApi.ApiMultiTenancyTenantsIdDefaultConnectionStringGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.ApiMultiTenancyTenantsIdDefaultConnectionStringGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiMultiTenancyTenantsIdDefaultConnectionStringGet`: string
    fmt.Fprintf(os.Stdout, "Response from `TenantApi.ApiMultiTenancyTenantsIdDefaultConnectionStringGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiMultiTenancyTenantsIdDefaultConnectionStringGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiMultiTenancyTenantsIdDefaultConnectionStringPut

> ApiMultiTenancyTenantsIdDefaultConnectionStringPut(ctx, id).DefaultConnectionString(defaultConnectionString).Execute()



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
    defaultConnectionString := "defaultConnectionString_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TenantApi.ApiMultiTenancyTenantsIdDefaultConnectionStringPut(context.Background(), id).DefaultConnectionString(defaultConnectionString).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.ApiMultiTenancyTenantsIdDefaultConnectionStringPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiMultiTenancyTenantsIdDefaultConnectionStringPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **defaultConnectionString** | **string** |  | 

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


## ApiMultiTenancyTenantsIdDelete

> ApiMultiTenancyTenantsIdDelete(ctx, id).Execute()



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
    r, err := apiClient.TenantApi.ApiMultiTenancyTenantsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.ApiMultiTenancyTenantsIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiMultiTenancyTenantsIdDeleteRequest struct via the builder pattern


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


## ApiMultiTenancyTenantsIdGet

> TenantDto ApiMultiTenancyTenantsIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.TenantApi.ApiMultiTenancyTenantsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.ApiMultiTenancyTenantsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiMultiTenancyTenantsIdGet`: TenantDto
    fmt.Fprintf(os.Stdout, "Response from `TenantApi.ApiMultiTenancyTenantsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiMultiTenancyTenantsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TenantDto**](TenantDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiMultiTenancyTenantsIdPut

> TenantDto ApiMultiTenancyTenantsIdPut(ctx, id).Body(body).Execute()



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
    body := *openapiclient.NewTenantUpdateDto("Name_example") // TenantUpdateDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantApi.ApiMultiTenancyTenantsIdPut(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.ApiMultiTenancyTenantsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiMultiTenancyTenantsIdPut`: TenantDto
    fmt.Fprintf(os.Stdout, "Response from `TenantApi.ApiMultiTenancyTenantsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiMultiTenancyTenantsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**TenantUpdateDto**](TenantUpdateDto.md) |  | 

### Return type

[**TenantDto**](TenantDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiMultiTenancyTenantsPost

> TenantDto ApiMultiTenancyTenantsPost(ctx).Body(body).Execute()



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
    body := *openapiclient.NewTenantCreateDto("Name_example", "AdminEmailAddress_example", "AdminPassword_example") // TenantCreateDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantApi.ApiMultiTenancyTenantsPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.ApiMultiTenancyTenantsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiMultiTenancyTenantsPost`: TenantDto
    fmt.Fprintf(os.Stdout, "Response from `TenantApi.ApiMultiTenancyTenantsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiMultiTenancyTenantsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TenantCreateDto**](TenantCreateDto.md) |  | 

### Return type

[**TenantDto**](TenantDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


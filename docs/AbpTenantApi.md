# \AbpTenantApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAbpMultiTenancyTenantsByIdIdGet**](AbpTenantApi.md#ApiAbpMultiTenancyTenantsByIdIdGet) | **Get** /api/abp/multi-tenancy/tenants/by-id/{id} | 
[**ApiAbpMultiTenancyTenantsByNameNameGet**](AbpTenantApi.md#ApiAbpMultiTenancyTenantsByNameNameGet) | **Get** /api/abp/multi-tenancy/tenants/by-name/{name} | 



## ApiAbpMultiTenancyTenantsByIdIdGet

> FindTenantResultDto ApiAbpMultiTenancyTenantsByIdIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.AbpTenantApi.ApiAbpMultiTenancyTenantsByIdIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AbpTenantApi.ApiAbpMultiTenancyTenantsByIdIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAbpMultiTenancyTenantsByIdIdGet`: FindTenantResultDto
    fmt.Fprintf(os.Stdout, "Response from `AbpTenantApi.ApiAbpMultiTenancyTenantsByIdIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAbpMultiTenancyTenantsByIdIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FindTenantResultDto**](FindTenantResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAbpMultiTenancyTenantsByNameNameGet

> FindTenantResultDto ApiAbpMultiTenancyTenantsByNameNameGet(ctx, name).Execute()



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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AbpTenantApi.ApiAbpMultiTenancyTenantsByNameNameGet(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AbpTenantApi.ApiAbpMultiTenancyTenantsByNameNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAbpMultiTenancyTenantsByNameNameGet`: FindTenantResultDto
    fmt.Fprintf(os.Stdout, "Response from `AbpTenantApi.ApiAbpMultiTenancyTenantsByNameNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAbpMultiTenancyTenantsByNameNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FindTenantResultDto**](FindTenantResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


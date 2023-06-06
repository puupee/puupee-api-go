# \PermissionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiPermissionManagementPermissionsGet**](PermissionsApi.md#ApiPermissionManagementPermissionsGet) | **Get** /api/permission-management/permissions | 
[**ApiPermissionManagementPermissionsPut**](PermissionsApi.md#ApiPermissionManagementPermissionsPut) | **Put** /api/permission-management/permissions | 



## ApiPermissionManagementPermissionsGet

> GetPermissionListResultDto ApiPermissionManagementPermissionsGet(ctx).ProviderName(providerName).ProviderKey(providerKey).Execute()



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
    providerName := "providerName_example" // string |  (optional)
    providerKey := "providerKey_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.ApiPermissionManagementPermissionsGet(context.Background()).ProviderName(providerName).ProviderKey(providerKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.ApiPermissionManagementPermissionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiPermissionManagementPermissionsGet`: GetPermissionListResultDto
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.ApiPermissionManagementPermissionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPermissionManagementPermissionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerName** | **string** |  | 
 **providerKey** | **string** |  | 

### Return type

[**GetPermissionListResultDto**](GetPermissionListResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPermissionManagementPermissionsPut

> ApiPermissionManagementPermissionsPut(ctx).ProviderName(providerName).ProviderKey(providerKey).Body(body).Execute()



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
    providerName := "providerName_example" // string |  (optional)
    providerKey := "providerKey_example" // string |  (optional)
    body := *openapiclient.NewUpdatePermissionsDto() // UpdatePermissionsDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PermissionsApi.ApiPermissionManagementPermissionsPut(context.Background()).ProviderName(providerName).ProviderKey(providerKey).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.ApiPermissionManagementPermissionsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPermissionManagementPermissionsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerName** | **string** |  | 
 **providerKey** | **string** |  | 
 **body** | [**UpdatePermissionsDto**](UpdatePermissionsDto.md) |  | 

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


# \FeaturesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiFeatureManagementFeaturesDelete**](FeaturesApi.md#ApiFeatureManagementFeaturesDelete) | **Delete** /api/feature-management/features | 
[**ApiFeatureManagementFeaturesGet**](FeaturesApi.md#ApiFeatureManagementFeaturesGet) | **Get** /api/feature-management/features | 
[**ApiFeatureManagementFeaturesPut**](FeaturesApi.md#ApiFeatureManagementFeaturesPut) | **Put** /api/feature-management/features | 



## ApiFeatureManagementFeaturesDelete

> ApiFeatureManagementFeaturesDelete(ctx).ProviderName(providerName).ProviderKey(providerKey).Execute()



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
    r, err := apiClient.FeaturesApi.ApiFeatureManagementFeaturesDelete(context.Background()).ProviderName(providerName).ProviderKey(providerKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.ApiFeatureManagementFeaturesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiFeatureManagementFeaturesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerName** | **string** |  | 
 **providerKey** | **string** |  | 

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


## ApiFeatureManagementFeaturesGet

> GetFeatureListResultDto ApiFeatureManagementFeaturesGet(ctx).ProviderName(providerName).ProviderKey(providerKey).Execute()



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
    resp, r, err := apiClient.FeaturesApi.ApiFeatureManagementFeaturesGet(context.Background()).ProviderName(providerName).ProviderKey(providerKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.ApiFeatureManagementFeaturesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiFeatureManagementFeaturesGet`: GetFeatureListResultDto
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.ApiFeatureManagementFeaturesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiFeatureManagementFeaturesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerName** | **string** |  | 
 **providerKey** | **string** |  | 

### Return type

[**GetFeatureListResultDto**](GetFeatureListResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiFeatureManagementFeaturesPut

> ApiFeatureManagementFeaturesPut(ctx).ProviderName(providerName).ProviderKey(providerKey).UpdateFeaturesDto(updateFeaturesDto).Execute()



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
    updateFeaturesDto := *openapiclient.NewUpdateFeaturesDto() // UpdateFeaturesDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FeaturesApi.ApiFeatureManagementFeaturesPut(context.Background()).ProviderName(providerName).ProviderKey(providerKey).UpdateFeaturesDto(updateFeaturesDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.ApiFeatureManagementFeaturesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiFeatureManagementFeaturesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerName** | **string** |  | 
 **providerKey** | **string** |  | 
 **updateFeaturesDto** | [**UpdateFeaturesDto**](UpdateFeaturesDto.md) |  | 

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


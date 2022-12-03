# \AbpApplicationConfigurationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAbpApplicationConfigurationGet**](AbpApplicationConfigurationApi.md#ApiAbpApplicationConfigurationGet) | **Get** /api/abp/application-configuration | 



## ApiAbpApplicationConfigurationGet

> ApplicationConfigurationDto ApiAbpApplicationConfigurationGet(ctx).IncludeLocalizationResources(includeLocalizationResources).Execute()



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
    includeLocalizationResources := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AbpApplicationConfigurationApi.ApiAbpApplicationConfigurationGet(context.Background()).IncludeLocalizationResources(includeLocalizationResources).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AbpApplicationConfigurationApi.ApiAbpApplicationConfigurationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAbpApplicationConfigurationGet`: ApplicationConfigurationDto
    fmt.Fprintf(os.Stdout, "Response from `AbpApplicationConfigurationApi.ApiAbpApplicationConfigurationGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAbpApplicationConfigurationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeLocalizationResources** | **bool** |  | 

### Return type

[**ApplicationConfigurationDto**](ApplicationConfigurationDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


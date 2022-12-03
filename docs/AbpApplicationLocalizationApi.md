# \AbpApplicationLocalizationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAbpApplicationLocalizationGet**](AbpApplicationLocalizationApi.md#ApiAbpApplicationLocalizationGet) | **Get** /api/abp/application-localization | 



## ApiAbpApplicationLocalizationGet

> ApplicationLocalizationDto ApiAbpApplicationLocalizationGet(ctx).CultureName(cultureName).OnlyDynamics(onlyDynamics).Execute()



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
    cultureName := "cultureName_example" // string | 
    onlyDynamics := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AbpApplicationLocalizationApi.ApiAbpApplicationLocalizationGet(context.Background()).CultureName(cultureName).OnlyDynamics(onlyDynamics).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AbpApplicationLocalizationApi.ApiAbpApplicationLocalizationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAbpApplicationLocalizationGet`: ApplicationLocalizationDto
    fmt.Fprintf(os.Stdout, "Response from `AbpApplicationLocalizationApi.ApiAbpApplicationLocalizationGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAbpApplicationLocalizationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cultureName** | **string** |  | 
 **onlyDynamics** | **bool** |  | 

### Return type

[**ApplicationLocalizationDto**](ApplicationLocalizationDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


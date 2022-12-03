# \AbpApiDefinitionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAbpApiDefinitionGet**](AbpApiDefinitionApi.md#ApiAbpApiDefinitionGet) | **Get** /api/abp/api-definition | 



## ApiAbpApiDefinitionGet

> ApplicationApiDescriptionModel ApiAbpApiDefinitionGet(ctx).IncludeTypes(includeTypes).Execute()



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
    includeTypes := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AbpApiDefinitionApi.ApiAbpApiDefinitionGet(context.Background()).IncludeTypes(includeTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AbpApiDefinitionApi.ApiAbpApiDefinitionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAbpApiDefinitionGet`: ApplicationApiDescriptionModel
    fmt.Fprintf(os.Stdout, "Response from `AbpApiDefinitionApi.ApiAbpApiDefinitionGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAbpApiDefinitionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeTypes** | **bool** |  | 

### Return type

[**ApplicationApiDescriptionModel**](ApplicationApiDescriptionModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


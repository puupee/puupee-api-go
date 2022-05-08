# \FileApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppFileFileOrCredentialsCreatorIdGet**](FileApi.md#ApiAppFileFileOrCredentialsCreatorIdGet) | **Get** /api/app/file/file-or-credentials/{creatorId} | 
[**ApiAppFilePreSignUrlPost**](FileApi.md#ApiAppFilePreSignUrlPost) | **Post** /api/app/file/pre-sign-url | 



## ApiAppFileFileOrCredentialsCreatorIdGet

> FileOrCredentialsDto ApiAppFileFileOrCredentialsCreatorIdGet(ctx, creatorId).RapidCode(rapidCode).Execute()



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
    creatorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    rapidCode := "rapidCode_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FileApi.ApiAppFileFileOrCredentialsCreatorIdGet(context.Background(), creatorId).RapidCode(rapidCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileApi.ApiAppFileFileOrCredentialsCreatorIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppFileFileOrCredentialsCreatorIdGet`: FileOrCredentialsDto
    fmt.Fprintf(os.Stdout, "Response from `FileApi.ApiAppFileFileOrCredentialsCreatorIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creatorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppFileFileOrCredentialsCreatorIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rapidCode** | **string** |  | 

### Return type

[**FileOrCredentialsDto**](FileOrCredentialsDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppFilePreSignUrlPost

> string ApiAppFilePreSignUrlPost(ctx).Key(key).Execute()



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
    key := "key_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FileApi.ApiAppFilePreSignUrlPost(context.Background()).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileApi.ApiAppFilePreSignUrlPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppFilePreSignUrlPost`: string
    fmt.Fprintf(os.Stdout, "Response from `FileApi.ApiAppFilePreSignUrlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppFilePreSignUrlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** |  | 

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


# \FileApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppFilePreSignUrlPost**](FileApi.md#ApiAppFilePreSignUrlPost) | **Post** /api/app/file/pre-sign-url | 
[**ApiAppFileUploadCredentialsGet**](FileApi.md#ApiAppFileUploadCredentialsGet) | **Get** /api/app/file/upload-credentials | 



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


## ApiAppFileUploadCredentialsGet

> UploadCredentials ApiAppFileUploadCredentialsGet(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FileApi.ApiAppFileUploadCredentialsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileApi.ApiAppFileUploadCredentialsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppFileUploadCredentialsGet`: UploadCredentials
    fmt.Fprintf(os.Stdout, "Response from `FileApi.ApiAppFileUploadCredentialsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppFileUploadCredentialsGetRequest struct via the builder pattern


### Return type

[**UploadCredentials**](UploadCredentials.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


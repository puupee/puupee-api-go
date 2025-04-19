# \StorageObjectApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFileOrCredentials**](StorageObjectApi.md#GetFileOrCredentials) | **Get** /api/app/storage-object/file-or-credentials | 
[**PreSignUrl**](StorageObjectApi.md#PreSignUrl) | **Post** /api/app/storage-object/pre-sign-url | 



## GetFileOrCredentials

> StorageObjectOrCredentialsDto GetFileOrCredentials(ctx).RapidCode(rapidCode).Bucket(bucket).Execute()



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
    rapidCode := "rapidCode_example" // string |  (optional)
    bucket := "bucket_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageObjectApi.GetFileOrCredentials(context.Background()).RapidCode(rapidCode).Bucket(bucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageObjectApi.GetFileOrCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFileOrCredentials`: StorageObjectOrCredentialsDto
    fmt.Fprintf(os.Stdout, "Response from `StorageObjectApi.GetFileOrCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFileOrCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rapidCode** | **string** |  | 
 **bucket** | **string** |  | 

### Return type

[**StorageObjectOrCredentialsDto**](StorageObjectOrCredentialsDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreSignUrl

> string PreSignUrl(ctx).Key(key).Bucket(bucket).Execute()



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
    key := "key_example" // string |  (optional)
    bucket := "bucket_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageObjectApi.PreSignUrl(context.Background()).Key(key).Bucket(bucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageObjectApi.PreSignUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreSignUrl`: string
    fmt.Fprintf(os.Stdout, "Response from `StorageObjectApi.PreSignUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreSignUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** |  | 
 **bucket** | **string** |  | 

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


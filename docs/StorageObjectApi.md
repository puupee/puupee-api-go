# \StorageObjectApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppStorageObjectFileGet**](StorageObjectApi.md#ApiAppStorageObjectFileGet) | **Get** /api/app/storage-object/file | 
[**ApiAppStorageObjectFileOrCredentialsGet**](StorageObjectApi.md#ApiAppStorageObjectFileOrCredentialsGet) | **Get** /api/app/storage-object/file-or-credentials | 
[**ApiAppStorageObjectPreSignUrlPost**](StorageObjectApi.md#ApiAppStorageObjectPreSignUrlPost) | **Post** /api/app/storage-object/pre-sign-url | 
[**ApiAppStorageObjectThumbGet**](StorageObjectApi.md#ApiAppStorageObjectThumbGet) | **Get** /api/app/storage-object/thumb | 



## ApiAppStorageObjectFileGet

> ApiAppStorageObjectFileGet(ctx).Key(key).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StorageObjectApi.ApiAppStorageObjectFileGet(context.Background()).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageObjectApi.ApiAppStorageObjectFileGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppStorageObjectFileGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppStorageObjectFileOrCredentialsGet

> StorageObjectOrCredentialsDto ApiAppStorageObjectFileOrCredentialsGet(ctx).RapidCode(rapidCode).Key(key).Execute()



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
    key := "key_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageObjectApi.ApiAppStorageObjectFileOrCredentialsGet(context.Background()).RapidCode(rapidCode).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageObjectApi.ApiAppStorageObjectFileOrCredentialsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppStorageObjectFileOrCredentialsGet`: StorageObjectOrCredentialsDto
    fmt.Fprintf(os.Stdout, "Response from `StorageObjectApi.ApiAppStorageObjectFileOrCredentialsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppStorageObjectFileOrCredentialsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rapidCode** | **string** |  | 
 **key** | **string** |  | 

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


## ApiAppStorageObjectPreSignUrlPost

> string ApiAppStorageObjectPreSignUrlPost(ctx).Key(key).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageObjectApi.ApiAppStorageObjectPreSignUrlPost(context.Background()).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageObjectApi.ApiAppStorageObjectPreSignUrlPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppStorageObjectPreSignUrlPost`: string
    fmt.Fprintf(os.Stdout, "Response from `StorageObjectApi.ApiAppStorageObjectPreSignUrlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppStorageObjectPreSignUrlPostRequest struct via the builder pattern


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


## ApiAppStorageObjectThumbGet

> ApiAppStorageObjectThumbGet(ctx).Key(key).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StorageObjectApi.ApiAppStorageObjectThumbGet(context.Background()).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageObjectApi.ApiAppStorageObjectThumbGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppStorageObjectThumbGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


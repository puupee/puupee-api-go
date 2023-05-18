# \AvatarApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppAvatarCredentialsGet**](AvatarApi.md#ApiAppAvatarCredentialsGet) | **Get** /api/app/avatar/credentials | 
[**ApiAppAvatarPost**](AvatarApi.md#ApiAppAvatarPost) | **Post** /api/app/avatar | 



## ApiAppAvatarCredentialsGet

> StorageObjectCredentials ApiAppAvatarCredentialsGet(ctx).Key(key).Execute()



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
    resp, r, err := apiClient.AvatarApi.ApiAppAvatarCredentialsGet(context.Background()).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvatarApi.ApiAppAvatarCredentialsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppAvatarCredentialsGet`: StorageObjectCredentials
    fmt.Fprintf(os.Stdout, "Response from `AvatarApi.ApiAppAvatarCredentialsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppAvatarCredentialsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** |  | 

### Return type

[**StorageObjectCredentials**](StorageObjectCredentials.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppAvatarPost

> AvatarDto ApiAppAvatarPost(ctx).CreateAvatarDto(createAvatarDto).Execute()



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
    createAvatarDto := *openapiclient.NewCreateAvatarDto() // CreateAvatarDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvatarApi.ApiAppAvatarPost(context.Background()).CreateAvatarDto(createAvatarDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvatarApi.ApiAppAvatarPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppAvatarPost`: AvatarDto
    fmt.Fprintf(os.Stdout, "Response from `AvatarApi.ApiAppAvatarPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppAvatarPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAvatarDto** | [**CreateAvatarDto**](CreateAvatarDto.md) |  | 

### Return type

[**AvatarDto**](AvatarDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


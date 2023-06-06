# \EmailSettingsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiSettingManagementEmailingGet**](EmailSettingsApi.md#ApiSettingManagementEmailingGet) | **Get** /api/setting-management/emailing | 
[**ApiSettingManagementEmailingPost**](EmailSettingsApi.md#ApiSettingManagementEmailingPost) | **Post** /api/setting-management/emailing | 
[**ApiSettingManagementEmailingSendTestEmailPost**](EmailSettingsApi.md#ApiSettingManagementEmailingSendTestEmailPost) | **Post** /api/setting-management/emailing/send-test-email | 



## ApiSettingManagementEmailingGet

> EmailSettingsDto ApiSettingManagementEmailingGet(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailSettingsApi.ApiSettingManagementEmailingGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailSettingsApi.ApiSettingManagementEmailingGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSettingManagementEmailingGet`: EmailSettingsDto
    fmt.Fprintf(os.Stdout, "Response from `EmailSettingsApi.ApiSettingManagementEmailingGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiSettingManagementEmailingGetRequest struct via the builder pattern


### Return type

[**EmailSettingsDto**](EmailSettingsDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSettingManagementEmailingPost

> ApiSettingManagementEmailingPost(ctx).Body(body).Execute()



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
    body := *openapiclient.NewUpdateEmailSettingsDto("DefaultFromAddress_example", "DefaultFromDisplayName_example") // UpdateEmailSettingsDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EmailSettingsApi.ApiSettingManagementEmailingPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailSettingsApi.ApiSettingManagementEmailingPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiSettingManagementEmailingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateEmailSettingsDto**](UpdateEmailSettingsDto.md) |  | 

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


## ApiSettingManagementEmailingSendTestEmailPost

> ApiSettingManagementEmailingSendTestEmailPost(ctx).Body(body).Execute()



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
    body := *openapiclient.NewSendTestEmailInput("SenderEmailAddress_example", "TargetEmailAddress_example", "Subject_example") // SendTestEmailInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EmailSettingsApi.ApiSettingManagementEmailingSendTestEmailPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailSettingsApi.ApiSettingManagementEmailingSendTestEmailPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiSettingManagementEmailingSendTestEmailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SendTestEmailInput**](SendTestEmailInput.md) |  | 

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


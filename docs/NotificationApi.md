# \NotificationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppNotificationBarkApiKeyMessageGet**](NotificationApi.md#ApiAppNotificationBarkApiKeyMessageGet) | **Get** /api/app/notification/bark/{apiKey}/{message} | 
[**ApiAppNotificationGet**](NotificationApi.md#ApiAppNotificationGet) | **Get** /api/app/notification | 
[**ApiAppNotificationPushPost**](NotificationApi.md#ApiAppNotificationPushPost) | **Post** /api/app/notification/push | 



## ApiAppNotificationBarkApiKeyMessageGet

> ApiAppNotificationBarkApiKeyMessageGet(ctx, apiKey, message).AutomaticallyCopy(automaticallyCopy).Copy(copy).Url(url).IsArchive(isArchive).Group(group).Icon(icon).Level(level).Execute()



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
    apiKey := "apiKey_example" // string | 
    message := "message_example" // string | 
    automaticallyCopy := int32(56) // int32 |  (optional) (default to 0)
    copy := "copy_example" // string |  (optional)
    url := "url_example" // string |  (optional)
    isArchive := "isArchive_example" // string |  (optional)
    group := "group_example" // string |  (optional)
    icon := "icon_example" // string |  (optional)
    level := "level_example" // string |  (optional) (default to "active")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationApi.ApiAppNotificationBarkApiKeyMessageGet(context.Background(), apiKey, message).AutomaticallyCopy(automaticallyCopy).Copy(copy).Url(url).IsArchive(isArchive).Group(group).Icon(icon).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationApi.ApiAppNotificationBarkApiKeyMessageGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKey** | **string** |  | 
**message** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppNotificationBarkApiKeyMessageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **automaticallyCopy** | **int32** |  | [default to 0]
 **copy** | **string** |  | 
 **url** | **string** |  | 
 **isArchive** | **string** |  | 
 **group** | **string** |  | 
 **icon** | **string** |  | 
 **level** | **string** |  | [default to &quot;active&quot;]

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


## ApiAppNotificationGet

> NotificationInfoDtoPagedResultDto ApiAppNotificationGet(ctx).Sorting(sorting).SkipCount(skipCount).MaxResultCount(maxResultCount).Execute()



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
    sorting := "sorting_example" // string |  (optional)
    skipCount := int32(56) // int32 |  (optional)
    maxResultCount := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationApi.ApiAppNotificationGet(context.Background()).Sorting(sorting).SkipCount(skipCount).MaxResultCount(maxResultCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationApi.ApiAppNotificationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppNotificationGet`: NotificationInfoDtoPagedResultDto
    fmt.Fprintf(os.Stdout, "Response from `NotificationApi.ApiAppNotificationGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppNotificationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sorting** | **string** |  | 
 **skipCount** | **int32** |  | 
 **maxResultCount** | **int32** |  | 

### Return type

[**NotificationInfoDtoPagedResultDto**](NotificationInfoDtoPagedResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppNotificationPushPost

> ApiAppNotificationPushPost(ctx).CreatePushNotificationDto(createPushNotificationDto).Execute()



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
    createPushNotificationDto := *openapiclient.NewCreatePushNotificationDto() // CreatePushNotificationDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationApi.ApiAppNotificationPushPost(context.Background()).CreatePushNotificationDto(createPushNotificationDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationApi.ApiAppNotificationPushPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppNotificationPushPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPushNotificationDto** | [**CreatePushNotificationDto**](CreatePushNotificationDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


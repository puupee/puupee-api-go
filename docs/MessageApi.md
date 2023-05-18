# \MessageApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppMessagePublishPost**](MessageApi.md#ApiAppMessagePublishPost) | **Post** /api/app/message/publish | 
[**ApiAppMessageRecallPost**](MessageApi.md#ApiAppMessageRecallPost) | **Post** /api/app/message/recall | 
[**ApiAppMessageSubscribePost**](MessageApi.md#ApiAppMessageSubscribePost) | **Post** /api/app/message/subscribe | 
[**ApiAppMessageUnsubscribePost**](MessageApi.md#ApiAppMessageUnsubscribePost) | **Post** /api/app/message/unsubscribe | 



## ApiAppMessagePublishPost

> ApiAppMessagePublishPost(ctx).MessagePublishDto(messagePublishDto).Execute()



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
    messagePublishDto := *openapiclient.NewMessagePublishDto() // MessagePublishDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MessageApi.ApiAppMessagePublishPost(context.Background()).MessagePublishDto(messagePublishDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageApi.ApiAppMessagePublishPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppMessagePublishPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messagePublishDto** | [**MessagePublishDto**](MessagePublishDto.md) |  | 

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


## ApiAppMessageRecallPost

> ApiAppMessageRecallPost(ctx).MessageRecallDto(messageRecallDto).Execute()



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
    messageRecallDto := *openapiclient.NewMessageRecallDto() // MessageRecallDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MessageApi.ApiAppMessageRecallPost(context.Background()).MessageRecallDto(messageRecallDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageApi.ApiAppMessageRecallPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppMessageRecallPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messageRecallDto** | [**MessageRecallDto**](MessageRecallDto.md) |  | 

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


## ApiAppMessageSubscribePost

> ApiAppMessageSubscribePost(ctx).MessageSubscribeDto(messageSubscribeDto).Execute()



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
    messageSubscribeDto := *openapiclient.NewMessageSubscribeDto() // MessageSubscribeDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MessageApi.ApiAppMessageSubscribePost(context.Background()).MessageSubscribeDto(messageSubscribeDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageApi.ApiAppMessageSubscribePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppMessageSubscribePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messageSubscribeDto** | [**MessageSubscribeDto**](MessageSubscribeDto.md) |  | 

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


## ApiAppMessageUnsubscribePost

> ApiAppMessageUnsubscribePost(ctx).MessageUnsubscribeDto(messageUnsubscribeDto).Execute()



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
    messageUnsubscribeDto := *openapiclient.NewMessageUnsubscribeDto() // MessageUnsubscribeDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MessageApi.ApiAppMessageUnsubscribePost(context.Background()).MessageUnsubscribeDto(messageUnsubscribeDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageApi.ApiAppMessageUnsubscribePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppMessageUnsubscribePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messageUnsubscribeDto** | [**MessageUnsubscribeDto**](MessageUnsubscribeDto.md) |  | 

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


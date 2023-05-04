# \SubscriptionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppSubscriptionAppleNotificationsPost**](SubscriptionApi.md#ApiAppSubscriptionAppleNotificationsPost) | **Post** /api/app/subscription/apple-notifications | 
[**ApiAppSubscriptionGet**](SubscriptionApi.md#ApiAppSubscriptionGet) | **Get** /api/app/subscription | 
[**ApiAppSubscriptionOrderPost**](SubscriptionApi.md#ApiAppSubscriptionOrderPost) | **Post** /api/app/subscription/order | 
[**ApiAppSubscriptionVerifyReceiptPost**](SubscriptionApi.md#ApiAppSubscriptionVerifyReceiptPost) | **Post** /api/app/subscription/verify-receipt | 



## ApiAppSubscriptionAppleNotificationsPost

> ApiAppSubscriptionAppleNotificationsPost(ctx).Body(body).Execute()



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
    body := *openapiclient.NewAnv2() // Anv2 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.ApiAppSubscriptionAppleNotificationsPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.ApiAppSubscriptionAppleNotificationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppSubscriptionAppleNotificationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Anv2**](Anv2.md) |  | 

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


## ApiAppSubscriptionGet

> SubscriptionDto ApiAppSubscriptionGet(ctx).AppId(appId).Execute()



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
    appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.ApiAppSubscriptionGet(context.Background()).AppId(appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.ApiAppSubscriptionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppSubscriptionGet`: SubscriptionDto
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.ApiAppSubscriptionGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppSubscriptionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string** |  | 

### Return type

[**SubscriptionDto**](SubscriptionDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppSubscriptionOrderPost

> SubscriptionOrderDto ApiAppSubscriptionOrderPost(ctx).Body(body).Execute()



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
    body := *openapiclient.NewCreateOrGetSubscriptionOrderDto() // CreateOrGetSubscriptionOrderDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.ApiAppSubscriptionOrderPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.ApiAppSubscriptionOrderPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppSubscriptionOrderPost`: SubscriptionOrderDto
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.ApiAppSubscriptionOrderPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppSubscriptionOrderPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateOrGetSubscriptionOrderDto**](CreateOrGetSubscriptionOrderDto.md) |  | 

### Return type

[**SubscriptionOrderDto**](SubscriptionOrderDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppSubscriptionVerifyReceiptPost

> VerifyReceiptResult ApiAppSubscriptionVerifyReceiptPost(ctx).Body(body).Execute()



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
    body := *openapiclient.NewVerifyReceiptDto() // VerifyReceiptDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.ApiAppSubscriptionVerifyReceiptPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.ApiAppSubscriptionVerifyReceiptPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppSubscriptionVerifyReceiptPost`: VerifyReceiptResult
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.ApiAppSubscriptionVerifyReceiptPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppSubscriptionVerifyReceiptPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VerifyReceiptDto**](VerifyReceiptDto.md) |  | 

### Return type

[**VerifyReceiptResult**](VerifyReceiptResult.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


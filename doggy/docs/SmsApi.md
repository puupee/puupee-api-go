# \SmsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppSmsSendChangePhoneCodePost**](SmsApi.md#ApiAppSmsSendChangePhoneCodePost) | **Post** /api/app/sms/send-change-phone-code | 
[**ApiAppSmsSendLoginCodePost**](SmsApi.md#ApiAppSmsSendLoginCodePost) | **Post** /api/app/sms/send-login-code | 



## ApiAppSmsSendChangePhoneCodePost

> ApiAppSmsSendChangePhoneCodePost(ctx).SendSmsCodeDto(sendSmsCodeDto).Execute()



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
    sendSmsCodeDto := *openapiclient.NewSendSmsCodeDto() // SendSmsCodeDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SmsApi.ApiAppSmsSendChangePhoneCodePost(context.Background()).SendSmsCodeDto(sendSmsCodeDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmsApi.ApiAppSmsSendChangePhoneCodePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppSmsSendChangePhoneCodePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendSmsCodeDto** | [**SendSmsCodeDto**](SendSmsCodeDto.md) |  | 

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


## ApiAppSmsSendLoginCodePost

> ApiAppSmsSendLoginCodePost(ctx).SendSmsCodeDto(sendSmsCodeDto).Execute()



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
    sendSmsCodeDto := *openapiclient.NewSendSmsCodeDto() // SendSmsCodeDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SmsApi.ApiAppSmsSendLoginCodePost(context.Background()).SendSmsCodeDto(sendSmsCodeDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmsApi.ApiAppSmsSendLoginCodePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppSmsSendLoginCodePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendSmsCodeDto** | [**SendSmsCodeDto**](SendSmsCodeDto.md) |  | 

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


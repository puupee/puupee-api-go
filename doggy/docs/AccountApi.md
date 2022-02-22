# \AccountApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAccountRegisterPost**](AccountApi.md#ApiAccountRegisterPost) | **Post** /api/account/register | 
[**ApiAccountResetPasswordPost**](AccountApi.md#ApiAccountResetPasswordPost) | **Post** /api/account/reset-password | 
[**ApiAccountSendPasswordResetCodePost**](AccountApi.md#ApiAccountSendPasswordResetCodePost) | **Post** /api/account/send-password-reset-code | 



## ApiAccountRegisterPost

> IdentityUserDto ApiAccountRegisterPost(ctx).RegisterDto(registerDto).Execute()



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
    registerDto := *openapiclient.NewRegisterDto("UserName_example", "EmailAddress_example", "Password_example", "AppName_example") // RegisterDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.ApiAccountRegisterPost(context.Background()).RegisterDto(registerDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.ApiAccountRegisterPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountRegisterPost`: IdentityUserDto
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.ApiAccountRegisterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountRegisterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerDto** | [**RegisterDto**](RegisterDto.md) |  | 

### Return type

[**IdentityUserDto**](IdentityUserDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountResetPasswordPost

> ApiAccountResetPasswordPost(ctx).ResetPasswordDto(resetPasswordDto).Execute()



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
    resetPasswordDto := *openapiclient.NewResetPasswordDto("ResetToken_example", "Password_example") // ResetPasswordDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.ApiAccountResetPasswordPost(context.Background()).ResetPasswordDto(resetPasswordDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.ApiAccountResetPasswordPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountResetPasswordPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordDto** | [**ResetPasswordDto**](ResetPasswordDto.md) |  | 

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


## ApiAccountSendPasswordResetCodePost

> ApiAccountSendPasswordResetCodePost(ctx).SendPasswordResetCodeDto(sendPasswordResetCodeDto).Execute()



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
    sendPasswordResetCodeDto := *openapiclient.NewSendPasswordResetCodeDto("Email_example", "AppName_example") // SendPasswordResetCodeDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.ApiAccountSendPasswordResetCodePost(context.Background()).SendPasswordResetCodeDto(sendPasswordResetCodeDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.ApiAccountSendPasswordResetCodePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountSendPasswordResetCodePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendPasswordResetCodeDto** | [**SendPasswordResetCodeDto**](SendPasswordResetCodeDto.md) |  | 

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


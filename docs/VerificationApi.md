# \VerificationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppVerificationSendCodePost**](VerificationApi.md#ApiAppVerificationSendCodePost) | **Post** /api/app/verification/send-code | 



## ApiAppVerificationSendCodePost

> ApiAppVerificationSendCodePost(ctx).SendVerificationCodeDto(sendVerificationCodeDto).Execute()



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
    sendVerificationCodeDto := *openapiclient.NewSendVerificationCodeDto() // SendVerificationCodeDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VerificationApi.ApiAppVerificationSendCodePost(context.Background()).SendVerificationCodeDto(sendVerificationCodeDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VerificationApi.ApiAppVerificationSendCodePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppVerificationSendCodePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendVerificationCodeDto** | [**SendVerificationCodeDto**](SendVerificationCodeDto.md) |  | 

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


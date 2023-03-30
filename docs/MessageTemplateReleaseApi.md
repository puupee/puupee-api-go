# \MessageTemplateReleaseApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppMessageTemplateReleaseGet**](MessageTemplateReleaseApi.md#ApiAppMessageTemplateReleaseGet) | **Get** /api/app/message-template-release | 
[**ApiAppMessageTemplateReleaseIdGet**](MessageTemplateReleaseApi.md#ApiAppMessageTemplateReleaseIdGet) | **Get** /api/app/message-template-release/{id} | 
[**ApiAppMessageTemplateReleasePost**](MessageTemplateReleaseApi.md#ApiAppMessageTemplateReleasePost) | **Post** /api/app/message-template-release | 



## ApiAppMessageTemplateReleaseGet

> []map[string]interface{} ApiAppMessageTemplateReleaseGet(ctx).TemplateId(templateId).Execute()



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
    templateId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageTemplateReleaseApi.ApiAppMessageTemplateReleaseGet(context.Background()).TemplateId(templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageTemplateReleaseApi.ApiAppMessageTemplateReleaseGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppMessageTemplateReleaseGet`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MessageTemplateReleaseApi.ApiAppMessageTemplateReleaseGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppMessageTemplateReleaseGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateId** | **string** |  | 

### Return type

**[]map[string]interface{}**

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppMessageTemplateReleaseIdGet

> map[string]interface{} ApiAppMessageTemplateReleaseIdGet(ctx, id).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageTemplateReleaseApi.ApiAppMessageTemplateReleaseIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageTemplateReleaseApi.ApiAppMessageTemplateReleaseIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppMessageTemplateReleaseIdGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MessageTemplateReleaseApi.ApiAppMessageTemplateReleaseIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppMessageTemplateReleaseIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppMessageTemplateReleasePost

> map[string]interface{} ApiAppMessageTemplateReleasePost(ctx).Body(body).Execute()



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
    body := *openapiclient.NewCreateMessageTemplateReleaseDto() // CreateMessageTemplateReleaseDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageTemplateReleaseApi.ApiAppMessageTemplateReleasePost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageTemplateReleaseApi.ApiAppMessageTemplateReleasePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppMessageTemplateReleasePost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MessageTemplateReleaseApi.ApiAppMessageTemplateReleasePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppMessageTemplateReleasePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateMessageTemplateReleaseDto**](CreateMessageTemplateReleaseDto.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


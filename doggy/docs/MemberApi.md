# \MemberApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppMemberGet**](MemberApi.md#ApiAppMemberGet) | **Get** /api/app/member | 



## ApiAppMemberGet

> MemberDto ApiAppMemberGet(ctx).CreatorId(creatorId).Execute()



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
    creatorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MemberApi.ApiAppMemberGet(context.Background()).CreatorId(creatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MemberApi.ApiAppMemberGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppMemberGet`: MemberDto
    fmt.Fprintf(os.Stdout, "Response from `MemberApi.ApiAppMemberGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppMemberGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **creatorId** | **string** |  | 

### Return type

[**MemberDto**](MemberDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


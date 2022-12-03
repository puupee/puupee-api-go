# \SyncStateApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppSyncStateGet**](SyncStateApi.md#ApiAppSyncStateGet) | **Get** /api/app/sync-state | 
[**ApiAppSyncStatePuupeeChangedEtoPost**](SyncStateApi.md#ApiAppSyncStatePuupeeChangedEtoPost) | **Post** /api/app/sync-state/puupee-changed-eto | 



## ApiAppSyncStateGet

> SyncStateDto ApiAppSyncStateGet(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyncStateApi.ApiAppSyncStateGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyncStateApi.ApiAppSyncStateGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppSyncStateGet`: SyncStateDto
    fmt.Fprintf(os.Stdout, "Response from `SyncStateApi.ApiAppSyncStateGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppSyncStateGetRequest struct via the builder pattern


### Return type

[**SyncStateDto**](SyncStateDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppSyncStatePuupeeChangedEtoPost

> PuupeeChangedEto ApiAppSyncStatePuupeeChangedEtoPost(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyncStateApi.ApiAppSyncStatePuupeeChangedEtoPost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyncStateApi.ApiAppSyncStatePuupeeChangedEtoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppSyncStatePuupeeChangedEtoPost`: PuupeeChangedEto
    fmt.Fprintf(os.Stdout, "Response from `SyncStateApi.ApiAppSyncStatePuupeeChangedEtoPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppSyncStatePuupeeChangedEtoPostRequest struct via the builder pattern


### Return type

[**PuupeeChangedEto**](PuupeeChangedEto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


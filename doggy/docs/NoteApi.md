# \NoteApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppNoteNoteSpecsGet**](NoteApi.md#ApiAppNoteNoteSpecsGet) | **Get** /api/app/note/note-specs | 



## ApiAppNoteNoteSpecsGet

> NoteSpecDto ApiAppNoteNoteSpecsGet(ctx).Execute()



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NoteApi.ApiAppNoteNoteSpecsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NoteApi.ApiAppNoteNoteSpecsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppNoteNoteSpecsGet`: NoteSpecDto
    fmt.Fprintf(os.Stdout, "Response from `NoteApi.ApiAppNoteNoteSpecsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAppNoteNoteSpecsGetRequest struct via the builder pattern


### Return type

[**NoteSpecDto**](NoteSpecDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


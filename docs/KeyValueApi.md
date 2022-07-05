# \KeyValueApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAppKeyValueBoolGet**](KeyValueApi.md#ApiAppKeyValueBoolGet) | **Get** /api/app/key-value/bool | 
[**ApiAppKeyValueDateTimeGet**](KeyValueApi.md#ApiAppKeyValueDateTimeGet) | **Get** /api/app/key-value/date-time | 
[**ApiAppKeyValueDecimalGet**](KeyValueApi.md#ApiAppKeyValueDecimalGet) | **Get** /api/app/key-value/decimal | 
[**ApiAppKeyValueDoubleGet**](KeyValueApi.md#ApiAppKeyValueDoubleGet) | **Get** /api/app/key-value/double | 
[**ApiAppKeyValueIntGet**](KeyValueApi.md#ApiAppKeyValueIntGet) | **Get** /api/app/key-value/int | 
[**ApiAppKeyValueSetBoolPost**](KeyValueApi.md#ApiAppKeyValueSetBoolPost) | **Post** /api/app/key-value/set-bool | 
[**ApiAppKeyValueSetDateTimePost**](KeyValueApi.md#ApiAppKeyValueSetDateTimePost) | **Post** /api/app/key-value/set-date-time | 
[**ApiAppKeyValueSetDecimalPost**](KeyValueApi.md#ApiAppKeyValueSetDecimalPost) | **Post** /api/app/key-value/set-decimal | 
[**ApiAppKeyValueSetDoublePost**](KeyValueApi.md#ApiAppKeyValueSetDoublePost) | **Post** /api/app/key-value/set-double | 
[**ApiAppKeyValueSetIntPost**](KeyValueApi.md#ApiAppKeyValueSetIntPost) | **Post** /api/app/key-value/set-int | 
[**ApiAppKeyValueSetStringPost**](KeyValueApi.md#ApiAppKeyValueSetStringPost) | **Post** /api/app/key-value/set-string | 
[**ApiAppKeyValueStringGet**](KeyValueApi.md#ApiAppKeyValueStringGet) | **Get** /api/app/key-value/string | 



## ApiAppKeyValueBoolGet

> BooleanKeyValue ApiAppKeyValueBoolGet(ctx).Key(key).Execute()



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
    key := "key_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeyValueApi.ApiAppKeyValueBoolGet(context.Background()).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyValueApi.ApiAppKeyValueBoolGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppKeyValueBoolGet`: BooleanKeyValue
    fmt.Fprintf(os.Stdout, "Response from `KeyValueApi.ApiAppKeyValueBoolGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppKeyValueBoolGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** |  | 

### Return type

[**BooleanKeyValue**](BooleanKeyValue.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppKeyValueDateTimeGet

> DateTimeKeyValue ApiAppKeyValueDateTimeGet(ctx).Key(key).Execute()



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
    key := "key_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeyValueApi.ApiAppKeyValueDateTimeGet(context.Background()).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyValueApi.ApiAppKeyValueDateTimeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppKeyValueDateTimeGet`: DateTimeKeyValue
    fmt.Fprintf(os.Stdout, "Response from `KeyValueApi.ApiAppKeyValueDateTimeGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppKeyValueDateTimeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** |  | 

### Return type

[**DateTimeKeyValue**](DateTimeKeyValue.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppKeyValueDecimalGet

> DecimalKeyValue ApiAppKeyValueDecimalGet(ctx).Key(key).Execute()



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
    key := "key_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeyValueApi.ApiAppKeyValueDecimalGet(context.Background()).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyValueApi.ApiAppKeyValueDecimalGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppKeyValueDecimalGet`: DecimalKeyValue
    fmt.Fprintf(os.Stdout, "Response from `KeyValueApi.ApiAppKeyValueDecimalGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppKeyValueDecimalGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** |  | 

### Return type

[**DecimalKeyValue**](DecimalKeyValue.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppKeyValueDoubleGet

> DoubleKeyValue ApiAppKeyValueDoubleGet(ctx).Key(key).Execute()



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
    key := "key_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeyValueApi.ApiAppKeyValueDoubleGet(context.Background()).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyValueApi.ApiAppKeyValueDoubleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppKeyValueDoubleGet`: DoubleKeyValue
    fmt.Fprintf(os.Stdout, "Response from `KeyValueApi.ApiAppKeyValueDoubleGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppKeyValueDoubleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** |  | 

### Return type

[**DoubleKeyValue**](DoubleKeyValue.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppKeyValueIntGet

> Int32KeyValue ApiAppKeyValueIntGet(ctx).Key(key).Execute()



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
    key := "key_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeyValueApi.ApiAppKeyValueIntGet(context.Background()).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyValueApi.ApiAppKeyValueIntGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppKeyValueIntGet`: Int32KeyValue
    fmt.Fprintf(os.Stdout, "Response from `KeyValueApi.ApiAppKeyValueIntGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppKeyValueIntGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** |  | 

### Return type

[**Int32KeyValue**](Int32KeyValue.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAppKeyValueSetBoolPost

> ApiAppKeyValueSetBoolPost(ctx).Key(key).BooleanSetKeyValueDto(booleanSetKeyValueDto).Execute()



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
    key := "key_example" // string |  (optional)
    booleanSetKeyValueDto := *openapiclient.NewBooleanSetKeyValueDto() // BooleanSetKeyValueDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeyValueApi.ApiAppKeyValueSetBoolPost(context.Background()).Key(key).BooleanSetKeyValueDto(booleanSetKeyValueDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyValueApi.ApiAppKeyValueSetBoolPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppKeyValueSetBoolPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** |  | 
 **booleanSetKeyValueDto** | [**BooleanSetKeyValueDto**](BooleanSetKeyValueDto.md) |  | 

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


## ApiAppKeyValueSetDateTimePost

> ApiAppKeyValueSetDateTimePost(ctx).Key(key).DateTimeSetKeyValueDto(dateTimeSetKeyValueDto).Execute()



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
    key := "key_example" // string |  (optional)
    dateTimeSetKeyValueDto := *openapiclient.NewDateTimeSetKeyValueDto() // DateTimeSetKeyValueDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeyValueApi.ApiAppKeyValueSetDateTimePost(context.Background()).Key(key).DateTimeSetKeyValueDto(dateTimeSetKeyValueDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyValueApi.ApiAppKeyValueSetDateTimePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppKeyValueSetDateTimePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** |  | 
 **dateTimeSetKeyValueDto** | [**DateTimeSetKeyValueDto**](DateTimeSetKeyValueDto.md) |  | 

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


## ApiAppKeyValueSetDecimalPost

> ApiAppKeyValueSetDecimalPost(ctx).Key(key).DecimalSetKeyValueDto(decimalSetKeyValueDto).Execute()



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
    key := "key_example" // string |  (optional)
    decimalSetKeyValueDto := *openapiclient.NewDecimalSetKeyValueDto() // DecimalSetKeyValueDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeyValueApi.ApiAppKeyValueSetDecimalPost(context.Background()).Key(key).DecimalSetKeyValueDto(decimalSetKeyValueDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyValueApi.ApiAppKeyValueSetDecimalPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppKeyValueSetDecimalPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** |  | 
 **decimalSetKeyValueDto** | [**DecimalSetKeyValueDto**](DecimalSetKeyValueDto.md) |  | 

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


## ApiAppKeyValueSetDoublePost

> ApiAppKeyValueSetDoublePost(ctx).Key(key).DoubleSetKeyValueDto(doubleSetKeyValueDto).Execute()



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
    key := "key_example" // string |  (optional)
    doubleSetKeyValueDto := *openapiclient.NewDoubleSetKeyValueDto() // DoubleSetKeyValueDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeyValueApi.ApiAppKeyValueSetDoublePost(context.Background()).Key(key).DoubleSetKeyValueDto(doubleSetKeyValueDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyValueApi.ApiAppKeyValueSetDoublePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppKeyValueSetDoublePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** |  | 
 **doubleSetKeyValueDto** | [**DoubleSetKeyValueDto**](DoubleSetKeyValueDto.md) |  | 

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


## ApiAppKeyValueSetIntPost

> ApiAppKeyValueSetIntPost(ctx).Key(key).Int32SetKeyValueDto(int32SetKeyValueDto).Execute()



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
    key := "key_example" // string |  (optional)
    int32SetKeyValueDto := *openapiclient.NewInt32SetKeyValueDto() // Int32SetKeyValueDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeyValueApi.ApiAppKeyValueSetIntPost(context.Background()).Key(key).Int32SetKeyValueDto(int32SetKeyValueDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyValueApi.ApiAppKeyValueSetIntPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppKeyValueSetIntPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** |  | 
 **int32SetKeyValueDto** | [**Int32SetKeyValueDto**](Int32SetKeyValueDto.md) |  | 

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


## ApiAppKeyValueSetStringPost

> ApiAppKeyValueSetStringPost(ctx).Key(key).StringSetKeyValueDto(stringSetKeyValueDto).Execute()



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
    key := "key_example" // string |  (optional)
    stringSetKeyValueDto := *openapiclient.NewStringSetKeyValueDto() // StringSetKeyValueDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeyValueApi.ApiAppKeyValueSetStringPost(context.Background()).Key(key).StringSetKeyValueDto(stringSetKeyValueDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyValueApi.ApiAppKeyValueSetStringPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppKeyValueSetStringPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** |  | 
 **stringSetKeyValueDto** | [**StringSetKeyValueDto**](StringSetKeyValueDto.md) |  | 

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


## ApiAppKeyValueStringGet

> StringKeyValue ApiAppKeyValueStringGet(ctx).Key(key).Execute()



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
    key := "key_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeyValueApi.ApiAppKeyValueStringGet(context.Background()).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyValueApi.ApiAppKeyValueStringGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAppKeyValueStringGet`: StringKeyValue
    fmt.Fprintf(os.Stdout, "Response from `KeyValueApi.ApiAppKeyValueStringGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAppKeyValueStringGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** |  | 

### Return type

[**StringKeyValue**](StringKeyValue.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


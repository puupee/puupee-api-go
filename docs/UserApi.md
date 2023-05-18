# \UserApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiIdentityUsersAssignableRolesGet**](UserApi.md#ApiIdentityUsersAssignableRolesGet) | **Get** /api/identity/users/assignable-roles | 
[**ApiIdentityUsersByEmailEmailGet**](UserApi.md#ApiIdentityUsersByEmailEmailGet) | **Get** /api/identity/users/by-email/{email} | 
[**ApiIdentityUsersByUsernameUserNameGet**](UserApi.md#ApiIdentityUsersByUsernameUserNameGet) | **Get** /api/identity/users/by-username/{userName} | 
[**ApiIdentityUsersGet**](UserApi.md#ApiIdentityUsersGet) | **Get** /api/identity/users | 
[**ApiIdentityUsersIdDelete**](UserApi.md#ApiIdentityUsersIdDelete) | **Delete** /api/identity/users/{id} | 
[**ApiIdentityUsersIdGet**](UserApi.md#ApiIdentityUsersIdGet) | **Get** /api/identity/users/{id} | 
[**ApiIdentityUsersIdPut**](UserApi.md#ApiIdentityUsersIdPut) | **Put** /api/identity/users/{id} | 
[**ApiIdentityUsersIdRolesGet**](UserApi.md#ApiIdentityUsersIdRolesGet) | **Get** /api/identity/users/{id}/roles | 
[**ApiIdentityUsersIdRolesPut**](UserApi.md#ApiIdentityUsersIdRolesPut) | **Put** /api/identity/users/{id}/roles | 
[**ApiIdentityUsersPost**](UserApi.md#ApiIdentityUsersPost) | **Post** /api/identity/users | 



## ApiIdentityUsersAssignableRolesGet

> IdentityRoleDtoListResultDto ApiIdentityUsersAssignableRolesGet(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ApiIdentityUsersAssignableRolesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiIdentityUsersAssignableRolesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiIdentityUsersAssignableRolesGet`: IdentityRoleDtoListResultDto
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ApiIdentityUsersAssignableRolesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiIdentityUsersAssignableRolesGetRequest struct via the builder pattern


### Return type

[**IdentityRoleDtoListResultDto**](IdentityRoleDtoListResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiIdentityUsersByEmailEmailGet

> IdentityUserDto ApiIdentityUsersByEmailEmailGet(ctx, email).Execute()



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
    email := "email_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ApiIdentityUsersByEmailEmailGet(context.Background(), email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiIdentityUsersByEmailEmailGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiIdentityUsersByEmailEmailGet`: IdentityUserDto
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ApiIdentityUsersByEmailEmailGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiIdentityUsersByEmailEmailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityUserDto**](IdentityUserDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiIdentityUsersByUsernameUserNameGet

> IdentityUserDto ApiIdentityUsersByUsernameUserNameGet(ctx, userName).Execute()



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
    userName := "userName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ApiIdentityUsersByUsernameUserNameGet(context.Background(), userName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiIdentityUsersByUsernameUserNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiIdentityUsersByUsernameUserNameGet`: IdentityUserDto
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ApiIdentityUsersByUsernameUserNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiIdentityUsersByUsernameUserNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityUserDto**](IdentityUserDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiIdentityUsersGet

> IdentityUserDtoPagedResultDto ApiIdentityUsersGet(ctx).Filter(filter).Sorting(sorting).SkipCount(skipCount).MaxResultCount(maxResultCount).Execute()



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
    filter := "filter_example" // string |  (optional)
    sorting := "sorting_example" // string |  (optional)
    skipCount := int32(56) // int32 |  (optional)
    maxResultCount := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ApiIdentityUsersGet(context.Background()).Filter(filter).Sorting(sorting).SkipCount(skipCount).MaxResultCount(maxResultCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiIdentityUsersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiIdentityUsersGet`: IdentityUserDtoPagedResultDto
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ApiIdentityUsersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiIdentityUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** |  | 
 **sorting** | **string** |  | 
 **skipCount** | **int32** |  | 
 **maxResultCount** | **int32** |  | 

### Return type

[**IdentityUserDtoPagedResultDto**](IdentityUserDtoPagedResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiIdentityUsersIdDelete

> ApiIdentityUsersIdDelete(ctx, id).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserApi.ApiIdentityUsersIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiIdentityUsersIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiIdentityUsersIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiIdentityUsersIdGet

> IdentityUserDto ApiIdentityUsersIdGet(ctx, id).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ApiIdentityUsersIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiIdentityUsersIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiIdentityUsersIdGet`: IdentityUserDto
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ApiIdentityUsersIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiIdentityUsersIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityUserDto**](IdentityUserDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiIdentityUsersIdPut

> IdentityUserDto ApiIdentityUsersIdPut(ctx, id).IdentityUserUpdateDto(identityUserUpdateDto).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    identityUserUpdateDto := *openapiclient.NewIdentityUserUpdateDto("UserName_example", "Email_example") // IdentityUserUpdateDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ApiIdentityUsersIdPut(context.Background(), id).IdentityUserUpdateDto(identityUserUpdateDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiIdentityUsersIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiIdentityUsersIdPut`: IdentityUserDto
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ApiIdentityUsersIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiIdentityUsersIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityUserUpdateDto** | [**IdentityUserUpdateDto**](IdentityUserUpdateDto.md) |  | 

### Return type

[**IdentityUserDto**](IdentityUserDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiIdentityUsersIdRolesGet

> IdentityRoleDtoListResultDto ApiIdentityUsersIdRolesGet(ctx, id).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ApiIdentityUsersIdRolesGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiIdentityUsersIdRolesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiIdentityUsersIdRolesGet`: IdentityRoleDtoListResultDto
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ApiIdentityUsersIdRolesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiIdentityUsersIdRolesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityRoleDtoListResultDto**](IdentityRoleDtoListResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiIdentityUsersIdRolesPut

> ApiIdentityUsersIdRolesPut(ctx, id).IdentityUserUpdateRolesDto(identityUserUpdateRolesDto).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    identityUserUpdateRolesDto := *openapiclient.NewIdentityUserUpdateRolesDto([]string{"RoleNames_example"}) // IdentityUserUpdateRolesDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserApi.ApiIdentityUsersIdRolesPut(context.Background(), id).IdentityUserUpdateRolesDto(identityUserUpdateRolesDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiIdentityUsersIdRolesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiIdentityUsersIdRolesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityUserUpdateRolesDto** | [**IdentityUserUpdateRolesDto**](IdentityUserUpdateRolesDto.md) |  | 

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


## ApiIdentityUsersPost

> IdentityUserDto ApiIdentityUsersPost(ctx).IdentityUserCreateDto(identityUserCreateDto).Execute()



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
    identityUserCreateDto := *openapiclient.NewIdentityUserCreateDto("UserName_example", "Email_example", "Password_example") // IdentityUserCreateDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ApiIdentityUsersPost(context.Background()).IdentityUserCreateDto(identityUserCreateDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiIdentityUsersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiIdentityUsersPost`: IdentityUserDto
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ApiIdentityUsersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiIdentityUsersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityUserCreateDto** | [**IdentityUserCreateDto**](IdentityUserCreateDto.md) |  | 

### Return type

[**IdentityUserDto**](IdentityUserDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


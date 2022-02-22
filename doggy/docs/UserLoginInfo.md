# UserLoginInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserNameOrEmailAddress** | **string** |  | 
**Password** | **string** |  | 
**RememberMe** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserLoginInfo

`func NewUserLoginInfo(userNameOrEmailAddress string, password string, ) *UserLoginInfo`

NewUserLoginInfo instantiates a new UserLoginInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLoginInfoWithDefaults

`func NewUserLoginInfoWithDefaults() *UserLoginInfo`

NewUserLoginInfoWithDefaults instantiates a new UserLoginInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserNameOrEmailAddress

`func (o *UserLoginInfo) GetUserNameOrEmailAddress() string`

GetUserNameOrEmailAddress returns the UserNameOrEmailAddress field if non-nil, zero value otherwise.

### GetUserNameOrEmailAddressOk

`func (o *UserLoginInfo) GetUserNameOrEmailAddressOk() (*string, bool)`

GetUserNameOrEmailAddressOk returns a tuple with the UserNameOrEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameOrEmailAddress

`func (o *UserLoginInfo) SetUserNameOrEmailAddress(v string)`

SetUserNameOrEmailAddress sets UserNameOrEmailAddress field to given value.


### GetPassword

`func (o *UserLoginInfo) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserLoginInfo) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserLoginInfo) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetRememberMe

`func (o *UserLoginInfo) GetRememberMe() bool`

GetRememberMe returns the RememberMe field if non-nil, zero value otherwise.

### GetRememberMeOk

`func (o *UserLoginInfo) GetRememberMeOk() (*bool, bool)`

GetRememberMeOk returns a tuple with the RememberMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberMe

`func (o *UserLoginInfo) SetRememberMe(v bool)`

SetRememberMe sets RememberMe field to given value.

### HasRememberMe

`func (o *UserLoginInfo) HasRememberMe() bool`

HasRememberMe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



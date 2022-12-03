# SendPasswordResetCodeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**AppName** | **string** |  | 
**ReturnUrl** | Pointer to **string** |  | [optional] 
**ReturnUrlHash** | Pointer to **string** |  | [optional] 

## Methods

### NewSendPasswordResetCodeDto

`func NewSendPasswordResetCodeDto(email string, appName string, ) *SendPasswordResetCodeDto`

NewSendPasswordResetCodeDto instantiates a new SendPasswordResetCodeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendPasswordResetCodeDtoWithDefaults

`func NewSendPasswordResetCodeDtoWithDefaults() *SendPasswordResetCodeDto`

NewSendPasswordResetCodeDtoWithDefaults instantiates a new SendPasswordResetCodeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SendPasswordResetCodeDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SendPasswordResetCodeDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SendPasswordResetCodeDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAppName

`func (o *SendPasswordResetCodeDto) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *SendPasswordResetCodeDto) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *SendPasswordResetCodeDto) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetReturnUrl

`func (o *SendPasswordResetCodeDto) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *SendPasswordResetCodeDto) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *SendPasswordResetCodeDto) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *SendPasswordResetCodeDto) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetReturnUrlHash

`func (o *SendPasswordResetCodeDto) GetReturnUrlHash() string`

GetReturnUrlHash returns the ReturnUrlHash field if non-nil, zero value otherwise.

### GetReturnUrlHashOk

`func (o *SendPasswordResetCodeDto) GetReturnUrlHashOk() (*string, bool)`

GetReturnUrlHashOk returns a tuple with the ReturnUrlHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrlHash

`func (o *SendPasswordResetCodeDto) SetReturnUrlHash(v string)`

SetReturnUrlHash sets ReturnUrlHash field to given value.

### HasReturnUrlHash

`func (o *SendPasswordResetCodeDto) HasReturnUrlHash() bool`

HasReturnUrlHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



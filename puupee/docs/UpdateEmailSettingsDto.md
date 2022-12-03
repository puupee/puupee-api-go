# UpdateEmailSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SmtpHost** | Pointer to **string** |  | [optional] 
**SmtpPort** | Pointer to **int32** |  | [optional] 
**SmtpUserName** | Pointer to **string** |  | [optional] 
**SmtpPassword** | Pointer to **string** |  | [optional] 
**SmtpDomain** | Pointer to **string** |  | [optional] 
**SmtpEnableSsl** | Pointer to **bool** |  | [optional] 
**SmtpUseDefaultCredentials** | Pointer to **bool** |  | [optional] 
**DefaultFromAddress** | **string** |  | 
**DefaultFromDisplayName** | **string** |  | 

## Methods

### NewUpdateEmailSettingsDto

`func NewUpdateEmailSettingsDto(defaultFromAddress string, defaultFromDisplayName string, ) *UpdateEmailSettingsDto`

NewUpdateEmailSettingsDto instantiates a new UpdateEmailSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEmailSettingsDtoWithDefaults

`func NewUpdateEmailSettingsDtoWithDefaults() *UpdateEmailSettingsDto`

NewUpdateEmailSettingsDtoWithDefaults instantiates a new UpdateEmailSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmtpHost

`func (o *UpdateEmailSettingsDto) GetSmtpHost() string`

GetSmtpHost returns the SmtpHost field if non-nil, zero value otherwise.

### GetSmtpHostOk

`func (o *UpdateEmailSettingsDto) GetSmtpHostOk() (*string, bool)`

GetSmtpHostOk returns a tuple with the SmtpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpHost

`func (o *UpdateEmailSettingsDto) SetSmtpHost(v string)`

SetSmtpHost sets SmtpHost field to given value.

### HasSmtpHost

`func (o *UpdateEmailSettingsDto) HasSmtpHost() bool`

HasSmtpHost returns a boolean if a field has been set.

### GetSmtpPort

`func (o *UpdateEmailSettingsDto) GetSmtpPort() int32`

GetSmtpPort returns the SmtpPort field if non-nil, zero value otherwise.

### GetSmtpPortOk

`func (o *UpdateEmailSettingsDto) GetSmtpPortOk() (*int32, bool)`

GetSmtpPortOk returns a tuple with the SmtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPort

`func (o *UpdateEmailSettingsDto) SetSmtpPort(v int32)`

SetSmtpPort sets SmtpPort field to given value.

### HasSmtpPort

`func (o *UpdateEmailSettingsDto) HasSmtpPort() bool`

HasSmtpPort returns a boolean if a field has been set.

### GetSmtpUserName

`func (o *UpdateEmailSettingsDto) GetSmtpUserName() string`

GetSmtpUserName returns the SmtpUserName field if non-nil, zero value otherwise.

### GetSmtpUserNameOk

`func (o *UpdateEmailSettingsDto) GetSmtpUserNameOk() (*string, bool)`

GetSmtpUserNameOk returns a tuple with the SmtpUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpUserName

`func (o *UpdateEmailSettingsDto) SetSmtpUserName(v string)`

SetSmtpUserName sets SmtpUserName field to given value.

### HasSmtpUserName

`func (o *UpdateEmailSettingsDto) HasSmtpUserName() bool`

HasSmtpUserName returns a boolean if a field has been set.

### GetSmtpPassword

`func (o *UpdateEmailSettingsDto) GetSmtpPassword() string`

GetSmtpPassword returns the SmtpPassword field if non-nil, zero value otherwise.

### GetSmtpPasswordOk

`func (o *UpdateEmailSettingsDto) GetSmtpPasswordOk() (*string, bool)`

GetSmtpPasswordOk returns a tuple with the SmtpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPassword

`func (o *UpdateEmailSettingsDto) SetSmtpPassword(v string)`

SetSmtpPassword sets SmtpPassword field to given value.

### HasSmtpPassword

`func (o *UpdateEmailSettingsDto) HasSmtpPassword() bool`

HasSmtpPassword returns a boolean if a field has been set.

### GetSmtpDomain

`func (o *UpdateEmailSettingsDto) GetSmtpDomain() string`

GetSmtpDomain returns the SmtpDomain field if non-nil, zero value otherwise.

### GetSmtpDomainOk

`func (o *UpdateEmailSettingsDto) GetSmtpDomainOk() (*string, bool)`

GetSmtpDomainOk returns a tuple with the SmtpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpDomain

`func (o *UpdateEmailSettingsDto) SetSmtpDomain(v string)`

SetSmtpDomain sets SmtpDomain field to given value.

### HasSmtpDomain

`func (o *UpdateEmailSettingsDto) HasSmtpDomain() bool`

HasSmtpDomain returns a boolean if a field has been set.

### GetSmtpEnableSsl

`func (o *UpdateEmailSettingsDto) GetSmtpEnableSsl() bool`

GetSmtpEnableSsl returns the SmtpEnableSsl field if non-nil, zero value otherwise.

### GetSmtpEnableSslOk

`func (o *UpdateEmailSettingsDto) GetSmtpEnableSslOk() (*bool, bool)`

GetSmtpEnableSslOk returns a tuple with the SmtpEnableSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpEnableSsl

`func (o *UpdateEmailSettingsDto) SetSmtpEnableSsl(v bool)`

SetSmtpEnableSsl sets SmtpEnableSsl field to given value.

### HasSmtpEnableSsl

`func (o *UpdateEmailSettingsDto) HasSmtpEnableSsl() bool`

HasSmtpEnableSsl returns a boolean if a field has been set.

### GetSmtpUseDefaultCredentials

`func (o *UpdateEmailSettingsDto) GetSmtpUseDefaultCredentials() bool`

GetSmtpUseDefaultCredentials returns the SmtpUseDefaultCredentials field if non-nil, zero value otherwise.

### GetSmtpUseDefaultCredentialsOk

`func (o *UpdateEmailSettingsDto) GetSmtpUseDefaultCredentialsOk() (*bool, bool)`

GetSmtpUseDefaultCredentialsOk returns a tuple with the SmtpUseDefaultCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpUseDefaultCredentials

`func (o *UpdateEmailSettingsDto) SetSmtpUseDefaultCredentials(v bool)`

SetSmtpUseDefaultCredentials sets SmtpUseDefaultCredentials field to given value.

### HasSmtpUseDefaultCredentials

`func (o *UpdateEmailSettingsDto) HasSmtpUseDefaultCredentials() bool`

HasSmtpUseDefaultCredentials returns a boolean if a field has been set.

### GetDefaultFromAddress

`func (o *UpdateEmailSettingsDto) GetDefaultFromAddress() string`

GetDefaultFromAddress returns the DefaultFromAddress field if non-nil, zero value otherwise.

### GetDefaultFromAddressOk

`func (o *UpdateEmailSettingsDto) GetDefaultFromAddressOk() (*string, bool)`

GetDefaultFromAddressOk returns a tuple with the DefaultFromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFromAddress

`func (o *UpdateEmailSettingsDto) SetDefaultFromAddress(v string)`

SetDefaultFromAddress sets DefaultFromAddress field to given value.


### GetDefaultFromDisplayName

`func (o *UpdateEmailSettingsDto) GetDefaultFromDisplayName() string`

GetDefaultFromDisplayName returns the DefaultFromDisplayName field if non-nil, zero value otherwise.

### GetDefaultFromDisplayNameOk

`func (o *UpdateEmailSettingsDto) GetDefaultFromDisplayNameOk() (*string, bool)`

GetDefaultFromDisplayNameOk returns a tuple with the DefaultFromDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFromDisplayName

`func (o *UpdateEmailSettingsDto) SetDefaultFromDisplayName(v string)`

SetDefaultFromDisplayName sets DefaultFromDisplayName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



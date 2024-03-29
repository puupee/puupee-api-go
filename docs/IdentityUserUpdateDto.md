# IdentityUserUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtraProperties** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**UserName** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Surname** | Pointer to **string** |  | [optional] 
**Email** | **string** |  | 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**LockoutEnabled** | Pointer to **bool** |  | [optional] 
**RoleNames** | Pointer to **[]string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**ConcurrencyStamp** | Pointer to **string** |  | [optional] 

## Methods

### NewIdentityUserUpdateDto

`func NewIdentityUserUpdateDto(userName string, email string, ) *IdentityUserUpdateDto`

NewIdentityUserUpdateDto instantiates a new IdentityUserUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityUserUpdateDtoWithDefaults

`func NewIdentityUserUpdateDtoWithDefaults() *IdentityUserUpdateDto`

NewIdentityUserUpdateDtoWithDefaults instantiates a new IdentityUserUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtraProperties

`func (o *IdentityUserUpdateDto) GetExtraProperties() map[string]map[string]interface{}`

GetExtraProperties returns the ExtraProperties field if non-nil, zero value otherwise.

### GetExtraPropertiesOk

`func (o *IdentityUserUpdateDto) GetExtraPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetExtraPropertiesOk returns a tuple with the ExtraProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraProperties

`func (o *IdentityUserUpdateDto) SetExtraProperties(v map[string]map[string]interface{})`

SetExtraProperties sets ExtraProperties field to given value.

### HasExtraProperties

`func (o *IdentityUserUpdateDto) HasExtraProperties() bool`

HasExtraProperties returns a boolean if a field has been set.

### GetUserName

`func (o *IdentityUserUpdateDto) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *IdentityUserUpdateDto) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *IdentityUserUpdateDto) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetName

`func (o *IdentityUserUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityUserUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityUserUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityUserUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSurname

`func (o *IdentityUserUpdateDto) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *IdentityUserUpdateDto) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *IdentityUserUpdateDto) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *IdentityUserUpdateDto) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetEmail

`func (o *IdentityUserUpdateDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IdentityUserUpdateDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IdentityUserUpdateDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPhoneNumber

`func (o *IdentityUserUpdateDto) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *IdentityUserUpdateDto) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *IdentityUserUpdateDto) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *IdentityUserUpdateDto) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetIsActive

`func (o *IdentityUserUpdateDto) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *IdentityUserUpdateDto) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *IdentityUserUpdateDto) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *IdentityUserUpdateDto) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLockoutEnabled

`func (o *IdentityUserUpdateDto) GetLockoutEnabled() bool`

GetLockoutEnabled returns the LockoutEnabled field if non-nil, zero value otherwise.

### GetLockoutEnabledOk

`func (o *IdentityUserUpdateDto) GetLockoutEnabledOk() (*bool, bool)`

GetLockoutEnabledOk returns a tuple with the LockoutEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutEnabled

`func (o *IdentityUserUpdateDto) SetLockoutEnabled(v bool)`

SetLockoutEnabled sets LockoutEnabled field to given value.

### HasLockoutEnabled

`func (o *IdentityUserUpdateDto) HasLockoutEnabled() bool`

HasLockoutEnabled returns a boolean if a field has been set.

### GetRoleNames

`func (o *IdentityUserUpdateDto) GetRoleNames() []string`

GetRoleNames returns the RoleNames field if non-nil, zero value otherwise.

### GetRoleNamesOk

`func (o *IdentityUserUpdateDto) GetRoleNamesOk() (*[]string, bool)`

GetRoleNamesOk returns a tuple with the RoleNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleNames

`func (o *IdentityUserUpdateDto) SetRoleNames(v []string)`

SetRoleNames sets RoleNames field to given value.

### HasRoleNames

`func (o *IdentityUserUpdateDto) HasRoleNames() bool`

HasRoleNames returns a boolean if a field has been set.

### GetPassword

`func (o *IdentityUserUpdateDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IdentityUserUpdateDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IdentityUserUpdateDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IdentityUserUpdateDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetConcurrencyStamp

`func (o *IdentityUserUpdateDto) GetConcurrencyStamp() string`

GetConcurrencyStamp returns the ConcurrencyStamp field if non-nil, zero value otherwise.

### GetConcurrencyStampOk

`func (o *IdentityUserUpdateDto) GetConcurrencyStampOk() (*string, bool)`

GetConcurrencyStampOk returns a tuple with the ConcurrencyStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyStamp

`func (o *IdentityUserUpdateDto) SetConcurrencyStamp(v string)`

SetConcurrencyStamp sets ConcurrencyStamp field to given value.

### HasConcurrencyStamp

`func (o *IdentityUserUpdateDto) HasConcurrencyStamp() bool`

HasConcurrencyStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



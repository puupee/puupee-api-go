# UserData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Surname** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EmailConfirmed** | Pointer to **bool** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**PhoneNumberConfirmed** | Pointer to **bool** |  | [optional] 
**ExtraProperties** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewUserData

`func NewUserData() *UserData`

NewUserData instantiates a new UserData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDataWithDefaults

`func NewUserDataWithDefaults() *UserData`

NewUserDataWithDefaults instantiates a new UserData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTenantId

`func (o *UserData) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UserData) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UserData) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *UserData) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUserName

`func (o *UserData) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UserData) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UserData) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UserData) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetName

`func (o *UserData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSurname

`func (o *UserData) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *UserData) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *UserData) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *UserData) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetIsActive

`func (o *UserData) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UserData) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UserData) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UserData) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetEmail

`func (o *UserData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserData) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserData) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmailConfirmed

`func (o *UserData) GetEmailConfirmed() bool`

GetEmailConfirmed returns the EmailConfirmed field if non-nil, zero value otherwise.

### GetEmailConfirmedOk

`func (o *UserData) GetEmailConfirmedOk() (*bool, bool)`

GetEmailConfirmedOk returns a tuple with the EmailConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConfirmed

`func (o *UserData) SetEmailConfirmed(v bool)`

SetEmailConfirmed sets EmailConfirmed field to given value.

### HasEmailConfirmed

`func (o *UserData) HasEmailConfirmed() bool`

HasEmailConfirmed returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *UserData) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UserData) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UserData) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *UserData) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPhoneNumberConfirmed

`func (o *UserData) GetPhoneNumberConfirmed() bool`

GetPhoneNumberConfirmed returns the PhoneNumberConfirmed field if non-nil, zero value otherwise.

### GetPhoneNumberConfirmedOk

`func (o *UserData) GetPhoneNumberConfirmedOk() (*bool, bool)`

GetPhoneNumberConfirmedOk returns a tuple with the PhoneNumberConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberConfirmed

`func (o *UserData) SetPhoneNumberConfirmed(v bool)`

SetPhoneNumberConfirmed sets PhoneNumberConfirmed field to given value.

### HasPhoneNumberConfirmed

`func (o *UserData) HasPhoneNumberConfirmed() bool`

HasPhoneNumberConfirmed returns a boolean if a field has been set.

### GetExtraProperties

`func (o *UserData) GetExtraProperties() map[string]map[string]interface{}`

GetExtraProperties returns the ExtraProperties field if non-nil, zero value otherwise.

### GetExtraPropertiesOk

`func (o *UserData) GetExtraPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetExtraPropertiesOk returns a tuple with the ExtraProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraProperties

`func (o *UserData) SetExtraProperties(v map[string]map[string]interface{})`

SetExtraProperties sets ExtraProperties field to given value.

### HasExtraProperties

`func (o *UserData) HasExtraProperties() bool`

HasExtraProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



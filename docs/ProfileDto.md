# ProfileDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtraProperties** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**UserName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Surname** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**IsExternal** | Pointer to **bool** |  | [optional] 
**HasPassword** | Pointer to **bool** |  | [optional] 
**ConcurrencyStamp** | Pointer to **string** |  | [optional] 

## Methods

### NewProfileDto

`func NewProfileDto() *ProfileDto`

NewProfileDto instantiates a new ProfileDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileDtoWithDefaults

`func NewProfileDtoWithDefaults() *ProfileDto`

NewProfileDtoWithDefaults instantiates a new ProfileDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtraProperties

`func (o *ProfileDto) GetExtraProperties() map[string]map[string]interface{}`

GetExtraProperties returns the ExtraProperties field if non-nil, zero value otherwise.

### GetExtraPropertiesOk

`func (o *ProfileDto) GetExtraPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetExtraPropertiesOk returns a tuple with the ExtraProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraProperties

`func (o *ProfileDto) SetExtraProperties(v map[string]map[string]interface{})`

SetExtraProperties sets ExtraProperties field to given value.

### HasExtraProperties

`func (o *ProfileDto) HasExtraProperties() bool`

HasExtraProperties returns a boolean if a field has been set.

### GetUserName

`func (o *ProfileDto) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ProfileDto) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ProfileDto) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ProfileDto) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetEmail

`func (o *ProfileDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ProfileDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ProfileDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ProfileDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *ProfileDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProfileDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProfileDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProfileDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSurname

`func (o *ProfileDto) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *ProfileDto) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *ProfileDto) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *ProfileDto) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *ProfileDto) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ProfileDto) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ProfileDto) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ProfileDto) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetIsExternal

`func (o *ProfileDto) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *ProfileDto) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *ProfileDto) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *ProfileDto) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetHasPassword

`func (o *ProfileDto) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *ProfileDto) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *ProfileDto) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *ProfileDto) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.

### GetConcurrencyStamp

`func (o *ProfileDto) GetConcurrencyStamp() string`

GetConcurrencyStamp returns the ConcurrencyStamp field if non-nil, zero value otherwise.

### GetConcurrencyStampOk

`func (o *ProfileDto) GetConcurrencyStampOk() (*string, bool)`

GetConcurrencyStampOk returns a tuple with the ConcurrencyStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyStamp

`func (o *ProfileDto) SetConcurrencyStamp(v string)`

SetConcurrencyStamp sets ConcurrencyStamp field to given value.

### HasConcurrencyStamp

`func (o *ProfileDto) HasConcurrencyStamp() bool`

HasConcurrencyStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



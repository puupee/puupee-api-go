# IdentityRoleUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtraProperties** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Name** | **string** |  | 
**IsDefault** | Pointer to **bool** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**ConcurrencyStamp** | Pointer to **string** |  | [optional] 

## Methods

### NewIdentityRoleUpdateDto

`func NewIdentityRoleUpdateDto(name string, ) *IdentityRoleUpdateDto`

NewIdentityRoleUpdateDto instantiates a new IdentityRoleUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityRoleUpdateDtoWithDefaults

`func NewIdentityRoleUpdateDtoWithDefaults() *IdentityRoleUpdateDto`

NewIdentityRoleUpdateDtoWithDefaults instantiates a new IdentityRoleUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtraProperties

`func (o *IdentityRoleUpdateDto) GetExtraProperties() map[string]map[string]interface{}`

GetExtraProperties returns the ExtraProperties field if non-nil, zero value otherwise.

### GetExtraPropertiesOk

`func (o *IdentityRoleUpdateDto) GetExtraPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetExtraPropertiesOk returns a tuple with the ExtraProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraProperties

`func (o *IdentityRoleUpdateDto) SetExtraProperties(v map[string]map[string]interface{})`

SetExtraProperties sets ExtraProperties field to given value.

### HasExtraProperties

`func (o *IdentityRoleUpdateDto) HasExtraProperties() bool`

HasExtraProperties returns a boolean if a field has been set.

### GetName

`func (o *IdentityRoleUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityRoleUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityRoleUpdateDto) SetName(v string)`

SetName sets Name field to given value.


### GetIsDefault

`func (o *IdentityRoleUpdateDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *IdentityRoleUpdateDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *IdentityRoleUpdateDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *IdentityRoleUpdateDto) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsPublic

`func (o *IdentityRoleUpdateDto) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *IdentityRoleUpdateDto) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *IdentityRoleUpdateDto) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *IdentityRoleUpdateDto) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetConcurrencyStamp

`func (o *IdentityRoleUpdateDto) GetConcurrencyStamp() string`

GetConcurrencyStamp returns the ConcurrencyStamp field if non-nil, zero value otherwise.

### GetConcurrencyStampOk

`func (o *IdentityRoleUpdateDto) GetConcurrencyStampOk() (*string, bool)`

GetConcurrencyStampOk returns a tuple with the ConcurrencyStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyStamp

`func (o *IdentityRoleUpdateDto) SetConcurrencyStamp(v string)`

SetConcurrencyStamp sets ConcurrencyStamp field to given value.

### HasConcurrencyStamp

`func (o *IdentityRoleUpdateDto) HasConcurrencyStamp() bool`

HasConcurrencyStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



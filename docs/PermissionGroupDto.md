# PermissionGroupDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**DisplayNameKey** | Pointer to **string** |  | [optional] 
**DisplayNameResource** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**[]PermissionGrantInfoDto**](PermissionGrantInfoDto.md) |  | [optional] 

## Methods

### NewPermissionGroupDto

`func NewPermissionGroupDto() *PermissionGroupDto`

NewPermissionGroupDto instantiates a new PermissionGroupDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionGroupDtoWithDefaults

`func NewPermissionGroupDtoWithDefaults() *PermissionGroupDto`

NewPermissionGroupDtoWithDefaults instantiates a new PermissionGroupDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PermissionGroupDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PermissionGroupDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PermissionGroupDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PermissionGroupDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *PermissionGroupDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PermissionGroupDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PermissionGroupDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PermissionGroupDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDisplayNameKey

`func (o *PermissionGroupDto) GetDisplayNameKey() string`

GetDisplayNameKey returns the DisplayNameKey field if non-nil, zero value otherwise.

### GetDisplayNameKeyOk

`func (o *PermissionGroupDto) GetDisplayNameKeyOk() (*string, bool)`

GetDisplayNameKeyOk returns a tuple with the DisplayNameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNameKey

`func (o *PermissionGroupDto) SetDisplayNameKey(v string)`

SetDisplayNameKey sets DisplayNameKey field to given value.

### HasDisplayNameKey

`func (o *PermissionGroupDto) HasDisplayNameKey() bool`

HasDisplayNameKey returns a boolean if a field has been set.

### GetDisplayNameResource

`func (o *PermissionGroupDto) GetDisplayNameResource() string`

GetDisplayNameResource returns the DisplayNameResource field if non-nil, zero value otherwise.

### GetDisplayNameResourceOk

`func (o *PermissionGroupDto) GetDisplayNameResourceOk() (*string, bool)`

GetDisplayNameResourceOk returns a tuple with the DisplayNameResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNameResource

`func (o *PermissionGroupDto) SetDisplayNameResource(v string)`

SetDisplayNameResource sets DisplayNameResource field to given value.

### HasDisplayNameResource

`func (o *PermissionGroupDto) HasDisplayNameResource() bool`

HasDisplayNameResource returns a boolean if a field has been set.

### GetPermissions

`func (o *PermissionGroupDto) GetPermissions() []PermissionGrantInfoDto`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PermissionGroupDto) GetPermissionsOk() (*[]PermissionGrantInfoDto, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PermissionGroupDto) SetPermissions(v []PermissionGrantInfoDto)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *PermissionGroupDto) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



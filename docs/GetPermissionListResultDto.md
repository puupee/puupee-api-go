# GetPermissionListResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityDisplayName** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to [**[]PermissionGroupDto**](PermissionGroupDto.md) |  | [optional] 

## Methods

### NewGetPermissionListResultDto

`func NewGetPermissionListResultDto() *GetPermissionListResultDto`

NewGetPermissionListResultDto instantiates a new GetPermissionListResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPermissionListResultDtoWithDefaults

`func NewGetPermissionListResultDtoWithDefaults() *GetPermissionListResultDto`

NewGetPermissionListResultDtoWithDefaults instantiates a new GetPermissionListResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityDisplayName

`func (o *GetPermissionListResultDto) GetEntityDisplayName() string`

GetEntityDisplayName returns the EntityDisplayName field if non-nil, zero value otherwise.

### GetEntityDisplayNameOk

`func (o *GetPermissionListResultDto) GetEntityDisplayNameOk() (*string, bool)`

GetEntityDisplayNameOk returns a tuple with the EntityDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityDisplayName

`func (o *GetPermissionListResultDto) SetEntityDisplayName(v string)`

SetEntityDisplayName sets EntityDisplayName field to given value.

### HasEntityDisplayName

`func (o *GetPermissionListResultDto) HasEntityDisplayName() bool`

HasEntityDisplayName returns a boolean if a field has been set.

### GetGroups

`func (o *GetPermissionListResultDto) GetGroups() []PermissionGroupDto`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GetPermissionListResultDto) GetGroupsOk() (*[]PermissionGroupDto, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GetPermissionListResultDto) SetGroups(v []PermissionGroupDto)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *GetPermissionListResultDto) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



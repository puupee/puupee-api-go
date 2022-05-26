# CreateUpdateTagDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**ParentTagId** | Pointer to **NullableString** |  | [optional] 
**SyncVersion** | Pointer to **int64** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**DeletionTime** | Pointer to **NullableTime** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**LastModificationTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreateUpdateTagDto

`func NewCreateUpdateTagDto() *CreateUpdateTagDto`

NewCreateUpdateTagDto instantiates a new CreateUpdateTagDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdateTagDtoWithDefaults

`func NewCreateUpdateTagDtoWithDefaults() *CreateUpdateTagDto`

NewCreateUpdateTagDtoWithDefaults instantiates a new CreateUpdateTagDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateUpdateTagDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUpdateTagDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUpdateTagDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateUpdateTagDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateUpdateTagDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateUpdateTagDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParentTagId

`func (o *CreateUpdateTagDto) GetParentTagId() string`

GetParentTagId returns the ParentTagId field if non-nil, zero value otherwise.

### GetParentTagIdOk

`func (o *CreateUpdateTagDto) GetParentTagIdOk() (*string, bool)`

GetParentTagIdOk returns a tuple with the ParentTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTagId

`func (o *CreateUpdateTagDto) SetParentTagId(v string)`

SetParentTagId sets ParentTagId field to given value.

### HasParentTagId

`func (o *CreateUpdateTagDto) HasParentTagId() bool`

HasParentTagId returns a boolean if a field has been set.

### SetParentTagIdNil

`func (o *CreateUpdateTagDto) SetParentTagIdNil(b bool)`

 SetParentTagIdNil sets the value for ParentTagId to be an explicit nil

### UnsetParentTagId
`func (o *CreateUpdateTagDto) UnsetParentTagId()`

UnsetParentTagId ensures that no value is present for ParentTagId, not even an explicit nil
### GetSyncVersion

`func (o *CreateUpdateTagDto) GetSyncVersion() int64`

GetSyncVersion returns the SyncVersion field if non-nil, zero value otherwise.

### GetSyncVersionOk

`func (o *CreateUpdateTagDto) GetSyncVersionOk() (*int64, bool)`

GetSyncVersionOk returns a tuple with the SyncVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncVersion

`func (o *CreateUpdateTagDto) SetSyncVersion(v int64)`

SetSyncVersion sets SyncVersion field to given value.

### HasSyncVersion

`func (o *CreateUpdateTagDto) HasSyncVersion() bool`

HasSyncVersion returns a boolean if a field has been set.

### GetIsDeleted

`func (o *CreateUpdateTagDto) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CreateUpdateTagDto) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CreateUpdateTagDto) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *CreateUpdateTagDto) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeletionTime

`func (o *CreateUpdateTagDto) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *CreateUpdateTagDto) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *CreateUpdateTagDto) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *CreateUpdateTagDto) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### SetDeletionTimeNil

`func (o *CreateUpdateTagDto) SetDeletionTimeNil(b bool)`

 SetDeletionTimeNil sets the value for DeletionTime to be an explicit nil

### UnsetDeletionTime
`func (o *CreateUpdateTagDto) UnsetDeletionTime()`

UnsetDeletionTime ensures that no value is present for DeletionTime, not even an explicit nil
### GetCreationTime

`func (o *CreateUpdateTagDto) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CreateUpdateTagDto) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CreateUpdateTagDto) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *CreateUpdateTagDto) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetLastModificationTime

`func (o *CreateUpdateTagDto) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *CreateUpdateTagDto) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *CreateUpdateTagDto) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *CreateUpdateTagDto) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



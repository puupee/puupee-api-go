# CreateUpdateTodoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Priority** | Pointer to [**Priority**](Priority.md) |  | [optional] 
**TagIds** | Pointer to **[]string** |  | [optional] 
**IsDone** | Pointer to **bool** |  | [optional] 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**EndAt** | Pointer to **NullableTime** |  | [optional] 
**SyncVersion** | Pointer to **int64** |  | [optional] 
**DoneAt** | Pointer to **NullableTime** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**DeletionTime** | Pointer to **NullableTime** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**LastModificationTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreateUpdateTodoDto

`func NewCreateUpdateTodoDto() *CreateUpdateTodoDto`

NewCreateUpdateTodoDto instantiates a new CreateUpdateTodoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdateTodoDtoWithDefaults

`func NewCreateUpdateTodoDtoWithDefaults() *CreateUpdateTodoDto`

NewCreateUpdateTodoDtoWithDefaults instantiates a new CreateUpdateTodoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateUpdateTodoDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateUpdateTodoDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateUpdateTodoDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateUpdateTodoDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CreateUpdateTodoDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CreateUpdateTodoDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPriority

`func (o *CreateUpdateTodoDto) GetPriority() Priority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateUpdateTodoDto) GetPriorityOk() (*Priority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateUpdateTodoDto) SetPriority(v Priority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CreateUpdateTodoDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTagIds

`func (o *CreateUpdateTodoDto) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *CreateUpdateTodoDto) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *CreateUpdateTodoDto) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *CreateUpdateTodoDto) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### SetTagIdsNil

`func (o *CreateUpdateTodoDto) SetTagIdsNil(b bool)`

 SetTagIdsNil sets the value for TagIds to be an explicit nil

### UnsetTagIds
`func (o *CreateUpdateTodoDto) UnsetTagIds()`

UnsetTagIds ensures that no value is present for TagIds, not even an explicit nil
### GetIsDone

`func (o *CreateUpdateTodoDto) GetIsDone() bool`

GetIsDone returns the IsDone field if non-nil, zero value otherwise.

### GetIsDoneOk

`func (o *CreateUpdateTodoDto) GetIsDoneOk() (*bool, bool)`

GetIsDoneOk returns a tuple with the IsDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDone

`func (o *CreateUpdateTodoDto) SetIsDone(v bool)`

SetIsDone sets IsDone field to given value.

### HasIsDone

`func (o *CreateUpdateTodoDto) HasIsDone() bool`

HasIsDone returns a boolean if a field has been set.

### GetParentId

`func (o *CreateUpdateTodoDto) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateUpdateTodoDto) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateUpdateTodoDto) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateUpdateTodoDto) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *CreateUpdateTodoDto) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *CreateUpdateTodoDto) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetEndAt

`func (o *CreateUpdateTodoDto) GetEndAt() time.Time`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *CreateUpdateTodoDto) GetEndAtOk() (*time.Time, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *CreateUpdateTodoDto) SetEndAt(v time.Time)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *CreateUpdateTodoDto) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### SetEndAtNil

`func (o *CreateUpdateTodoDto) SetEndAtNil(b bool)`

 SetEndAtNil sets the value for EndAt to be an explicit nil

### UnsetEndAt
`func (o *CreateUpdateTodoDto) UnsetEndAt()`

UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil
### GetSyncVersion

`func (o *CreateUpdateTodoDto) GetSyncVersion() int64`

GetSyncVersion returns the SyncVersion field if non-nil, zero value otherwise.

### GetSyncVersionOk

`func (o *CreateUpdateTodoDto) GetSyncVersionOk() (*int64, bool)`

GetSyncVersionOk returns a tuple with the SyncVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncVersion

`func (o *CreateUpdateTodoDto) SetSyncVersion(v int64)`

SetSyncVersion sets SyncVersion field to given value.

### HasSyncVersion

`func (o *CreateUpdateTodoDto) HasSyncVersion() bool`

HasSyncVersion returns a boolean if a field has been set.

### GetDoneAt

`func (o *CreateUpdateTodoDto) GetDoneAt() time.Time`

GetDoneAt returns the DoneAt field if non-nil, zero value otherwise.

### GetDoneAtOk

`func (o *CreateUpdateTodoDto) GetDoneAtOk() (*time.Time, bool)`

GetDoneAtOk returns a tuple with the DoneAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoneAt

`func (o *CreateUpdateTodoDto) SetDoneAt(v time.Time)`

SetDoneAt sets DoneAt field to given value.

### HasDoneAt

`func (o *CreateUpdateTodoDto) HasDoneAt() bool`

HasDoneAt returns a boolean if a field has been set.

### SetDoneAtNil

`func (o *CreateUpdateTodoDto) SetDoneAtNil(b bool)`

 SetDoneAtNil sets the value for DoneAt to be an explicit nil

### UnsetDoneAt
`func (o *CreateUpdateTodoDto) UnsetDoneAt()`

UnsetDoneAt ensures that no value is present for DoneAt, not even an explicit nil
### GetIsDeleted

`func (o *CreateUpdateTodoDto) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CreateUpdateTodoDto) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CreateUpdateTodoDto) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *CreateUpdateTodoDto) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeletionTime

`func (o *CreateUpdateTodoDto) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *CreateUpdateTodoDto) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *CreateUpdateTodoDto) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *CreateUpdateTodoDto) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### SetDeletionTimeNil

`func (o *CreateUpdateTodoDto) SetDeletionTimeNil(b bool)`

 SetDeletionTimeNil sets the value for DeletionTime to be an explicit nil

### UnsetDeletionTime
`func (o *CreateUpdateTodoDto) UnsetDeletionTime()`

UnsetDeletionTime ensures that no value is present for DeletionTime, not even an explicit nil
### GetCreationTime

`func (o *CreateUpdateTodoDto) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CreateUpdateTodoDto) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CreateUpdateTodoDto) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *CreateUpdateTodoDto) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetLastModificationTime

`func (o *CreateUpdateTodoDto) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *CreateUpdateTodoDto) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *CreateUpdateTodoDto) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *CreateUpdateTodoDto) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



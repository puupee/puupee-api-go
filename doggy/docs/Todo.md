# Todo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ExtraProperties** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**ConcurrencyStamp** | Pointer to **NullableString** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**CreatorId** | Pointer to **NullableString** |  | [optional] 
**LastModificationTime** | Pointer to **NullableTime** |  | [optional] 
**LastModifierId** | Pointer to **NullableString** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**DeleterId** | Pointer to **NullableString** |  | [optional] 
**DeletionTime** | Pointer to **NullableTime** |  | [optional] 
**Deleter** | Pointer to [**IdentityUser**](IdentityUser.md) |  | [optional] 
**Creator** | Pointer to [**IdentityUser**](IdentityUser.md) |  | [optional] 
**LastModifier** | Pointer to [**IdentityUser**](IdentityUser.md) |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Priority** | Pointer to [**Priority**](Priority.md) |  | [optional] 
**DoneAt** | Pointer to **NullableTime** |  | [optional] 
**IsDone** | Pointer to **bool** |  | [optional] 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**Parent** | Pointer to [**Todo**](Todo.md) |  | [optional] 
**Children** | Pointer to [**[]Todo**](Todo.md) |  | [optional] 
**EndAt** | Pointer to **NullableTime** |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**TodoTags** | Pointer to [**[]TodoTag**](TodoTag.md) |  | [optional] 

## Methods

### NewTodo

`func NewTodo() *Todo`

NewTodo instantiates a new Todo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTodoWithDefaults

`func NewTodoWithDefaults() *Todo`

NewTodoWithDefaults instantiates a new Todo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Todo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Todo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Todo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Todo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExtraProperties

`func (o *Todo) GetExtraProperties() map[string]interface{}`

GetExtraProperties returns the ExtraProperties field if non-nil, zero value otherwise.

### GetExtraPropertiesOk

`func (o *Todo) GetExtraPropertiesOk() (*map[string]interface{}, bool)`

GetExtraPropertiesOk returns a tuple with the ExtraProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraProperties

`func (o *Todo) SetExtraProperties(v map[string]interface{})`

SetExtraProperties sets ExtraProperties field to given value.

### HasExtraProperties

`func (o *Todo) HasExtraProperties() bool`

HasExtraProperties returns a boolean if a field has been set.

### SetExtraPropertiesNil

`func (o *Todo) SetExtraPropertiesNil(b bool)`

 SetExtraPropertiesNil sets the value for ExtraProperties to be an explicit nil

### UnsetExtraProperties
`func (o *Todo) UnsetExtraProperties()`

UnsetExtraProperties ensures that no value is present for ExtraProperties, not even an explicit nil
### GetConcurrencyStamp

`func (o *Todo) GetConcurrencyStamp() string`

GetConcurrencyStamp returns the ConcurrencyStamp field if non-nil, zero value otherwise.

### GetConcurrencyStampOk

`func (o *Todo) GetConcurrencyStampOk() (*string, bool)`

GetConcurrencyStampOk returns a tuple with the ConcurrencyStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyStamp

`func (o *Todo) SetConcurrencyStamp(v string)`

SetConcurrencyStamp sets ConcurrencyStamp field to given value.

### HasConcurrencyStamp

`func (o *Todo) HasConcurrencyStamp() bool`

HasConcurrencyStamp returns a boolean if a field has been set.

### SetConcurrencyStampNil

`func (o *Todo) SetConcurrencyStampNil(b bool)`

 SetConcurrencyStampNil sets the value for ConcurrencyStamp to be an explicit nil

### UnsetConcurrencyStamp
`func (o *Todo) UnsetConcurrencyStamp()`

UnsetConcurrencyStamp ensures that no value is present for ConcurrencyStamp, not even an explicit nil
### GetCreationTime

`func (o *Todo) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Todo) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Todo) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *Todo) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *Todo) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *Todo) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *Todo) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *Todo) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### SetCreatorIdNil

`func (o *Todo) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *Todo) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetLastModificationTime

`func (o *Todo) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *Todo) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *Todo) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *Todo) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### SetLastModificationTimeNil

`func (o *Todo) SetLastModificationTimeNil(b bool)`

 SetLastModificationTimeNil sets the value for LastModificationTime to be an explicit nil

### UnsetLastModificationTime
`func (o *Todo) UnsetLastModificationTime()`

UnsetLastModificationTime ensures that no value is present for LastModificationTime, not even an explicit nil
### GetLastModifierId

`func (o *Todo) GetLastModifierId() string`

GetLastModifierId returns the LastModifierId field if non-nil, zero value otherwise.

### GetLastModifierIdOk

`func (o *Todo) GetLastModifierIdOk() (*string, bool)`

GetLastModifierIdOk returns a tuple with the LastModifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierId

`func (o *Todo) SetLastModifierId(v string)`

SetLastModifierId sets LastModifierId field to given value.

### HasLastModifierId

`func (o *Todo) HasLastModifierId() bool`

HasLastModifierId returns a boolean if a field has been set.

### SetLastModifierIdNil

`func (o *Todo) SetLastModifierIdNil(b bool)`

 SetLastModifierIdNil sets the value for LastModifierId to be an explicit nil

### UnsetLastModifierId
`func (o *Todo) UnsetLastModifierId()`

UnsetLastModifierId ensures that no value is present for LastModifierId, not even an explicit nil
### GetIsDeleted

`func (o *Todo) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *Todo) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *Todo) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *Todo) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeleterId

`func (o *Todo) GetDeleterId() string`

GetDeleterId returns the DeleterId field if non-nil, zero value otherwise.

### GetDeleterIdOk

`func (o *Todo) GetDeleterIdOk() (*string, bool)`

GetDeleterIdOk returns a tuple with the DeleterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleterId

`func (o *Todo) SetDeleterId(v string)`

SetDeleterId sets DeleterId field to given value.

### HasDeleterId

`func (o *Todo) HasDeleterId() bool`

HasDeleterId returns a boolean if a field has been set.

### SetDeleterIdNil

`func (o *Todo) SetDeleterIdNil(b bool)`

 SetDeleterIdNil sets the value for DeleterId to be an explicit nil

### UnsetDeleterId
`func (o *Todo) UnsetDeleterId()`

UnsetDeleterId ensures that no value is present for DeleterId, not even an explicit nil
### GetDeletionTime

`func (o *Todo) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *Todo) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *Todo) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *Todo) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### SetDeletionTimeNil

`func (o *Todo) SetDeletionTimeNil(b bool)`

 SetDeletionTimeNil sets the value for DeletionTime to be an explicit nil

### UnsetDeletionTime
`func (o *Todo) UnsetDeletionTime()`

UnsetDeletionTime ensures that no value is present for DeletionTime, not even an explicit nil
### GetDeleter

`func (o *Todo) GetDeleter() IdentityUser`

GetDeleter returns the Deleter field if non-nil, zero value otherwise.

### GetDeleterOk

`func (o *Todo) GetDeleterOk() (*IdentityUser, bool)`

GetDeleterOk returns a tuple with the Deleter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleter

`func (o *Todo) SetDeleter(v IdentityUser)`

SetDeleter sets Deleter field to given value.

### HasDeleter

`func (o *Todo) HasDeleter() bool`

HasDeleter returns a boolean if a field has been set.

### GetCreator

`func (o *Todo) GetCreator() IdentityUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Todo) GetCreatorOk() (*IdentityUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Todo) SetCreator(v IdentityUser)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *Todo) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetLastModifier

`func (o *Todo) GetLastModifier() IdentityUser`

GetLastModifier returns the LastModifier field if non-nil, zero value otherwise.

### GetLastModifierOk

`func (o *Todo) GetLastModifierOk() (*IdentityUser, bool)`

GetLastModifierOk returns a tuple with the LastModifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifier

`func (o *Todo) SetLastModifier(v IdentityUser)`

SetLastModifier sets LastModifier field to given value.

### HasLastModifier

`func (o *Todo) HasLastModifier() bool`

HasLastModifier returns a boolean if a field has been set.

### GetTitle

`func (o *Todo) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Todo) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Todo) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Todo) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *Todo) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Todo) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPriority

`func (o *Todo) GetPriority() Priority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Todo) GetPriorityOk() (*Priority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Todo) SetPriority(v Priority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Todo) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetDoneAt

`func (o *Todo) GetDoneAt() time.Time`

GetDoneAt returns the DoneAt field if non-nil, zero value otherwise.

### GetDoneAtOk

`func (o *Todo) GetDoneAtOk() (*time.Time, bool)`

GetDoneAtOk returns a tuple with the DoneAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoneAt

`func (o *Todo) SetDoneAt(v time.Time)`

SetDoneAt sets DoneAt field to given value.

### HasDoneAt

`func (o *Todo) HasDoneAt() bool`

HasDoneAt returns a boolean if a field has been set.

### SetDoneAtNil

`func (o *Todo) SetDoneAtNil(b bool)`

 SetDoneAtNil sets the value for DoneAt to be an explicit nil

### UnsetDoneAt
`func (o *Todo) UnsetDoneAt()`

UnsetDoneAt ensures that no value is present for DoneAt, not even an explicit nil
### GetIsDone

`func (o *Todo) GetIsDone() bool`

GetIsDone returns the IsDone field if non-nil, zero value otherwise.

### GetIsDoneOk

`func (o *Todo) GetIsDoneOk() (*bool, bool)`

GetIsDoneOk returns a tuple with the IsDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDone

`func (o *Todo) SetIsDone(v bool)`

SetIsDone sets IsDone field to given value.

### HasIsDone

`func (o *Todo) HasIsDone() bool`

HasIsDone returns a boolean if a field has been set.

### GetParentId

`func (o *Todo) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Todo) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Todo) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Todo) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *Todo) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *Todo) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetParent

`func (o *Todo) GetParent() Todo`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Todo) GetParentOk() (*Todo, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Todo) SetParent(v Todo)`

SetParent sets Parent field to given value.

### HasParent

`func (o *Todo) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetChildren

`func (o *Todo) GetChildren() []Todo`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Todo) GetChildrenOk() (*[]Todo, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Todo) SetChildren(v []Todo)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *Todo) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### SetChildrenNil

`func (o *Todo) SetChildrenNil(b bool)`

 SetChildrenNil sets the value for Children to be an explicit nil

### UnsetChildren
`func (o *Todo) UnsetChildren()`

UnsetChildren ensures that no value is present for Children, not even an explicit nil
### GetEndAt

`func (o *Todo) GetEndAt() time.Time`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *Todo) GetEndAtOk() (*time.Time, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *Todo) SetEndAt(v time.Time)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *Todo) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### SetEndAtNil

`func (o *Todo) SetEndAtNil(b bool)`

 SetEndAtNil sets the value for EndAt to be an explicit nil

### UnsetEndAt
`func (o *Todo) UnsetEndAt()`

UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil
### GetTags

`func (o *Todo) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Todo) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Todo) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Todo) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *Todo) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *Todo) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTodoTags

`func (o *Todo) GetTodoTags() []TodoTag`

GetTodoTags returns the TodoTags field if non-nil, zero value otherwise.

### GetTodoTagsOk

`func (o *Todo) GetTodoTagsOk() (*[]TodoTag, bool)`

GetTodoTagsOk returns a tuple with the TodoTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodoTags

`func (o *Todo) SetTodoTags(v []TodoTag)`

SetTodoTags sets TodoTags field to given value.

### HasTodoTags

`func (o *Todo) HasTodoTags() bool`

HasTodoTags returns a boolean if a field has been set.

### SetTodoTagsNil

`func (o *Todo) SetTodoTagsNil(b bool)`

 SetTodoTagsNil sets the value for TodoTags to be an explicit nil

### UnsetTodoTags
`func (o *Todo) UnsetTodoTags()`

UnsetTodoTags ensures that no value is present for TodoTags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



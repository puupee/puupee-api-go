# Tag

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
**Name** | Pointer to **NullableString** |  | [optional] 
**TagCount** | Pointer to **int32** |  | [optional] 
**ParentTagId** | Pointer to **NullableString** |  | [optional] 
**ParentTag** | Pointer to [**Tag**](Tag.md) |  | [optional] 
**Children** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**Todos** | Pointer to [**[]Todo**](Todo.md) |  | [optional] 
**TodoTags** | Pointer to [**[]TodoTag**](TodoTag.md) |  | [optional] 
**Items** | Pointer to [**[]Item**](Item.md) |  | [optional] 
**ItemTags** | Pointer to [**[]ItemTag**](ItemTag.md) |  | [optional] 
**FullPath** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewTag

`func NewTag() *Tag`

NewTag instantiates a new Tag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagWithDefaults

`func NewTagWithDefaults() *Tag`

NewTagWithDefaults instantiates a new Tag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Tag) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tag) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tag) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Tag) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExtraProperties

`func (o *Tag) GetExtraProperties() map[string]interface{}`

GetExtraProperties returns the ExtraProperties field if non-nil, zero value otherwise.

### GetExtraPropertiesOk

`func (o *Tag) GetExtraPropertiesOk() (*map[string]interface{}, bool)`

GetExtraPropertiesOk returns a tuple with the ExtraProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraProperties

`func (o *Tag) SetExtraProperties(v map[string]interface{})`

SetExtraProperties sets ExtraProperties field to given value.

### HasExtraProperties

`func (o *Tag) HasExtraProperties() bool`

HasExtraProperties returns a boolean if a field has been set.

### SetExtraPropertiesNil

`func (o *Tag) SetExtraPropertiesNil(b bool)`

 SetExtraPropertiesNil sets the value for ExtraProperties to be an explicit nil

### UnsetExtraProperties
`func (o *Tag) UnsetExtraProperties()`

UnsetExtraProperties ensures that no value is present for ExtraProperties, not even an explicit nil
### GetConcurrencyStamp

`func (o *Tag) GetConcurrencyStamp() string`

GetConcurrencyStamp returns the ConcurrencyStamp field if non-nil, zero value otherwise.

### GetConcurrencyStampOk

`func (o *Tag) GetConcurrencyStampOk() (*string, bool)`

GetConcurrencyStampOk returns a tuple with the ConcurrencyStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyStamp

`func (o *Tag) SetConcurrencyStamp(v string)`

SetConcurrencyStamp sets ConcurrencyStamp field to given value.

### HasConcurrencyStamp

`func (o *Tag) HasConcurrencyStamp() bool`

HasConcurrencyStamp returns a boolean if a field has been set.

### SetConcurrencyStampNil

`func (o *Tag) SetConcurrencyStampNil(b bool)`

 SetConcurrencyStampNil sets the value for ConcurrencyStamp to be an explicit nil

### UnsetConcurrencyStamp
`func (o *Tag) UnsetConcurrencyStamp()`

UnsetConcurrencyStamp ensures that no value is present for ConcurrencyStamp, not even an explicit nil
### GetCreationTime

`func (o *Tag) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Tag) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Tag) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *Tag) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *Tag) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *Tag) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *Tag) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *Tag) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### SetCreatorIdNil

`func (o *Tag) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *Tag) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetLastModificationTime

`func (o *Tag) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *Tag) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *Tag) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *Tag) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### SetLastModificationTimeNil

`func (o *Tag) SetLastModificationTimeNil(b bool)`

 SetLastModificationTimeNil sets the value for LastModificationTime to be an explicit nil

### UnsetLastModificationTime
`func (o *Tag) UnsetLastModificationTime()`

UnsetLastModificationTime ensures that no value is present for LastModificationTime, not even an explicit nil
### GetLastModifierId

`func (o *Tag) GetLastModifierId() string`

GetLastModifierId returns the LastModifierId field if non-nil, zero value otherwise.

### GetLastModifierIdOk

`func (o *Tag) GetLastModifierIdOk() (*string, bool)`

GetLastModifierIdOk returns a tuple with the LastModifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierId

`func (o *Tag) SetLastModifierId(v string)`

SetLastModifierId sets LastModifierId field to given value.

### HasLastModifierId

`func (o *Tag) HasLastModifierId() bool`

HasLastModifierId returns a boolean if a field has been set.

### SetLastModifierIdNil

`func (o *Tag) SetLastModifierIdNil(b bool)`

 SetLastModifierIdNil sets the value for LastModifierId to be an explicit nil

### UnsetLastModifierId
`func (o *Tag) UnsetLastModifierId()`

UnsetLastModifierId ensures that no value is present for LastModifierId, not even an explicit nil
### GetIsDeleted

`func (o *Tag) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *Tag) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *Tag) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *Tag) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeleterId

`func (o *Tag) GetDeleterId() string`

GetDeleterId returns the DeleterId field if non-nil, zero value otherwise.

### GetDeleterIdOk

`func (o *Tag) GetDeleterIdOk() (*string, bool)`

GetDeleterIdOk returns a tuple with the DeleterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleterId

`func (o *Tag) SetDeleterId(v string)`

SetDeleterId sets DeleterId field to given value.

### HasDeleterId

`func (o *Tag) HasDeleterId() bool`

HasDeleterId returns a boolean if a field has been set.

### SetDeleterIdNil

`func (o *Tag) SetDeleterIdNil(b bool)`

 SetDeleterIdNil sets the value for DeleterId to be an explicit nil

### UnsetDeleterId
`func (o *Tag) UnsetDeleterId()`

UnsetDeleterId ensures that no value is present for DeleterId, not even an explicit nil
### GetDeletionTime

`func (o *Tag) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *Tag) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *Tag) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *Tag) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### SetDeletionTimeNil

`func (o *Tag) SetDeletionTimeNil(b bool)`

 SetDeletionTimeNil sets the value for DeletionTime to be an explicit nil

### UnsetDeletionTime
`func (o *Tag) UnsetDeletionTime()`

UnsetDeletionTime ensures that no value is present for DeletionTime, not even an explicit nil
### GetDeleter

`func (o *Tag) GetDeleter() IdentityUser`

GetDeleter returns the Deleter field if non-nil, zero value otherwise.

### GetDeleterOk

`func (o *Tag) GetDeleterOk() (*IdentityUser, bool)`

GetDeleterOk returns a tuple with the Deleter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleter

`func (o *Tag) SetDeleter(v IdentityUser)`

SetDeleter sets Deleter field to given value.

### HasDeleter

`func (o *Tag) HasDeleter() bool`

HasDeleter returns a boolean if a field has been set.

### GetCreator

`func (o *Tag) GetCreator() IdentityUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Tag) GetCreatorOk() (*IdentityUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Tag) SetCreator(v IdentityUser)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *Tag) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetLastModifier

`func (o *Tag) GetLastModifier() IdentityUser`

GetLastModifier returns the LastModifier field if non-nil, zero value otherwise.

### GetLastModifierOk

`func (o *Tag) GetLastModifierOk() (*IdentityUser, bool)`

GetLastModifierOk returns a tuple with the LastModifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifier

`func (o *Tag) SetLastModifier(v IdentityUser)`

SetLastModifier sets LastModifier field to given value.

### HasLastModifier

`func (o *Tag) HasLastModifier() bool`

HasLastModifier returns a boolean if a field has been set.

### GetName

`func (o *Tag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tag) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Tag) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Tag) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Tag) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTagCount

`func (o *Tag) GetTagCount() int32`

GetTagCount returns the TagCount field if non-nil, zero value otherwise.

### GetTagCountOk

`func (o *Tag) GetTagCountOk() (*int32, bool)`

GetTagCountOk returns a tuple with the TagCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCount

`func (o *Tag) SetTagCount(v int32)`

SetTagCount sets TagCount field to given value.

### HasTagCount

`func (o *Tag) HasTagCount() bool`

HasTagCount returns a boolean if a field has been set.

### GetParentTagId

`func (o *Tag) GetParentTagId() string`

GetParentTagId returns the ParentTagId field if non-nil, zero value otherwise.

### GetParentTagIdOk

`func (o *Tag) GetParentTagIdOk() (*string, bool)`

GetParentTagIdOk returns a tuple with the ParentTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTagId

`func (o *Tag) SetParentTagId(v string)`

SetParentTagId sets ParentTagId field to given value.

### HasParentTagId

`func (o *Tag) HasParentTagId() bool`

HasParentTagId returns a boolean if a field has been set.

### SetParentTagIdNil

`func (o *Tag) SetParentTagIdNil(b bool)`

 SetParentTagIdNil sets the value for ParentTagId to be an explicit nil

### UnsetParentTagId
`func (o *Tag) UnsetParentTagId()`

UnsetParentTagId ensures that no value is present for ParentTagId, not even an explicit nil
### GetParentTag

`func (o *Tag) GetParentTag() Tag`

GetParentTag returns the ParentTag field if non-nil, zero value otherwise.

### GetParentTagOk

`func (o *Tag) GetParentTagOk() (*Tag, bool)`

GetParentTagOk returns a tuple with the ParentTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTag

`func (o *Tag) SetParentTag(v Tag)`

SetParentTag sets ParentTag field to given value.

### HasParentTag

`func (o *Tag) HasParentTag() bool`

HasParentTag returns a boolean if a field has been set.

### GetChildren

`func (o *Tag) GetChildren() []Tag`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Tag) GetChildrenOk() (*[]Tag, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Tag) SetChildren(v []Tag)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *Tag) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### SetChildrenNil

`func (o *Tag) SetChildrenNil(b bool)`

 SetChildrenNil sets the value for Children to be an explicit nil

### UnsetChildren
`func (o *Tag) UnsetChildren()`

UnsetChildren ensures that no value is present for Children, not even an explicit nil
### GetTodos

`func (o *Tag) GetTodos() []Todo`

GetTodos returns the Todos field if non-nil, zero value otherwise.

### GetTodosOk

`func (o *Tag) GetTodosOk() (*[]Todo, bool)`

GetTodosOk returns a tuple with the Todos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodos

`func (o *Tag) SetTodos(v []Todo)`

SetTodos sets Todos field to given value.

### HasTodos

`func (o *Tag) HasTodos() bool`

HasTodos returns a boolean if a field has been set.

### SetTodosNil

`func (o *Tag) SetTodosNil(b bool)`

 SetTodosNil sets the value for Todos to be an explicit nil

### UnsetTodos
`func (o *Tag) UnsetTodos()`

UnsetTodos ensures that no value is present for Todos, not even an explicit nil
### GetTodoTags

`func (o *Tag) GetTodoTags() []TodoTag`

GetTodoTags returns the TodoTags field if non-nil, zero value otherwise.

### GetTodoTagsOk

`func (o *Tag) GetTodoTagsOk() (*[]TodoTag, bool)`

GetTodoTagsOk returns a tuple with the TodoTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodoTags

`func (o *Tag) SetTodoTags(v []TodoTag)`

SetTodoTags sets TodoTags field to given value.

### HasTodoTags

`func (o *Tag) HasTodoTags() bool`

HasTodoTags returns a boolean if a field has been set.

### SetTodoTagsNil

`func (o *Tag) SetTodoTagsNil(b bool)`

 SetTodoTagsNil sets the value for TodoTags to be an explicit nil

### UnsetTodoTags
`func (o *Tag) UnsetTodoTags()`

UnsetTodoTags ensures that no value is present for TodoTags, not even an explicit nil
### GetItems

`func (o *Tag) GetItems() []Item`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Tag) GetItemsOk() (*[]Item, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Tag) SetItems(v []Item)`

SetItems sets Items field to given value.

### HasItems

`func (o *Tag) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *Tag) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *Tag) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetItemTags

`func (o *Tag) GetItemTags() []ItemTag`

GetItemTags returns the ItemTags field if non-nil, zero value otherwise.

### GetItemTagsOk

`func (o *Tag) GetItemTagsOk() (*[]ItemTag, bool)`

GetItemTagsOk returns a tuple with the ItemTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTags

`func (o *Tag) SetItemTags(v []ItemTag)`

SetItemTags sets ItemTags field to given value.

### HasItemTags

`func (o *Tag) HasItemTags() bool`

HasItemTags returns a boolean if a field has been set.

### SetItemTagsNil

`func (o *Tag) SetItemTagsNil(b bool)`

 SetItemTagsNil sets the value for ItemTags to be an explicit nil

### UnsetItemTags
`func (o *Tag) UnsetItemTags()`

UnsetItemTags ensures that no value is present for ItemTags, not even an explicit nil
### GetFullPath

`func (o *Tag) GetFullPath() string`

GetFullPath returns the FullPath field if non-nil, zero value otherwise.

### GetFullPathOk

`func (o *Tag) GetFullPathOk() (*string, bool)`

GetFullPathOk returns a tuple with the FullPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPath

`func (o *Tag) SetFullPath(v string)`

SetFullPath sets FullPath field to given value.

### HasFullPath

`func (o *Tag) HasFullPath() bool`

HasFullPath returns a boolean if a field has been set.

### SetFullPathNil

`func (o *Tag) SetFullPathNil(b bool)`

 SetFullPathNil sets the value for FullPath to be an explicit nil

### UnsetFullPath
`func (o *Tag) UnsetFullPath()`

UnsetFullPath ensures that no value is present for FullPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



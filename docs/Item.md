# Item

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
**Description** | Pointer to **NullableString** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**ParentItemId** | Pointer to **NullableString** |  | [optional] 
**ParentItem** | Pointer to [**Item**](Item.md) |  | [optional] 
**Children** | Pointer to [**[]Item**](Item.md) |  | [optional] 
**ThumbId** | Pointer to **NullableString** |  | [optional] 
**Thumb** | Pointer to [**ItemThumb**](ItemThumb.md) |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Md5** | Pointer to **NullableString** |  | [optional] 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**ItemType**](ItemType.md) |  | [optional] 
**Extension** | Pointer to **NullableString** |  | [optional] 
**StorageClass** | Pointer to **NullableString** |  | [optional] 
**FileCreatedAt** | Pointer to **NullableTime** |  | [optional] 
**FileUpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**FileId** | Pointer to **NullableString** |  | [optional] 
**File** | Pointer to [**File**](File.md) |  | [optional] 
**DisplayStyle** | Pointer to [**DisplayStyle**](DisplayStyle.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**ItemTags** | Pointer to [**[]ItemTag**](ItemTag.md) |  | [optional] 

## Methods

### NewItem

`func NewItem() *Item`

NewItem instantiates a new Item object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemWithDefaults

`func NewItemWithDefaults() *Item`

NewItemWithDefaults instantiates a new Item object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Item) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Item) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Item) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Item) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExtraProperties

`func (o *Item) GetExtraProperties() map[string]interface{}`

GetExtraProperties returns the ExtraProperties field if non-nil, zero value otherwise.

### GetExtraPropertiesOk

`func (o *Item) GetExtraPropertiesOk() (*map[string]interface{}, bool)`

GetExtraPropertiesOk returns a tuple with the ExtraProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraProperties

`func (o *Item) SetExtraProperties(v map[string]interface{})`

SetExtraProperties sets ExtraProperties field to given value.

### HasExtraProperties

`func (o *Item) HasExtraProperties() bool`

HasExtraProperties returns a boolean if a field has been set.

### SetExtraPropertiesNil

`func (o *Item) SetExtraPropertiesNil(b bool)`

 SetExtraPropertiesNil sets the value for ExtraProperties to be an explicit nil

### UnsetExtraProperties
`func (o *Item) UnsetExtraProperties()`

UnsetExtraProperties ensures that no value is present for ExtraProperties, not even an explicit nil
### GetConcurrencyStamp

`func (o *Item) GetConcurrencyStamp() string`

GetConcurrencyStamp returns the ConcurrencyStamp field if non-nil, zero value otherwise.

### GetConcurrencyStampOk

`func (o *Item) GetConcurrencyStampOk() (*string, bool)`

GetConcurrencyStampOk returns a tuple with the ConcurrencyStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyStamp

`func (o *Item) SetConcurrencyStamp(v string)`

SetConcurrencyStamp sets ConcurrencyStamp field to given value.

### HasConcurrencyStamp

`func (o *Item) HasConcurrencyStamp() bool`

HasConcurrencyStamp returns a boolean if a field has been set.

### SetConcurrencyStampNil

`func (o *Item) SetConcurrencyStampNil(b bool)`

 SetConcurrencyStampNil sets the value for ConcurrencyStamp to be an explicit nil

### UnsetConcurrencyStamp
`func (o *Item) UnsetConcurrencyStamp()`

UnsetConcurrencyStamp ensures that no value is present for ConcurrencyStamp, not even an explicit nil
### GetCreationTime

`func (o *Item) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Item) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Item) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *Item) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *Item) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *Item) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *Item) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *Item) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### SetCreatorIdNil

`func (o *Item) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *Item) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetLastModificationTime

`func (o *Item) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *Item) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *Item) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *Item) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### SetLastModificationTimeNil

`func (o *Item) SetLastModificationTimeNil(b bool)`

 SetLastModificationTimeNil sets the value for LastModificationTime to be an explicit nil

### UnsetLastModificationTime
`func (o *Item) UnsetLastModificationTime()`

UnsetLastModificationTime ensures that no value is present for LastModificationTime, not even an explicit nil
### GetLastModifierId

`func (o *Item) GetLastModifierId() string`

GetLastModifierId returns the LastModifierId field if non-nil, zero value otherwise.

### GetLastModifierIdOk

`func (o *Item) GetLastModifierIdOk() (*string, bool)`

GetLastModifierIdOk returns a tuple with the LastModifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierId

`func (o *Item) SetLastModifierId(v string)`

SetLastModifierId sets LastModifierId field to given value.

### HasLastModifierId

`func (o *Item) HasLastModifierId() bool`

HasLastModifierId returns a boolean if a field has been set.

### SetLastModifierIdNil

`func (o *Item) SetLastModifierIdNil(b bool)`

 SetLastModifierIdNil sets the value for LastModifierId to be an explicit nil

### UnsetLastModifierId
`func (o *Item) UnsetLastModifierId()`

UnsetLastModifierId ensures that no value is present for LastModifierId, not even an explicit nil
### GetIsDeleted

`func (o *Item) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *Item) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *Item) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *Item) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeleterId

`func (o *Item) GetDeleterId() string`

GetDeleterId returns the DeleterId field if non-nil, zero value otherwise.

### GetDeleterIdOk

`func (o *Item) GetDeleterIdOk() (*string, bool)`

GetDeleterIdOk returns a tuple with the DeleterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleterId

`func (o *Item) SetDeleterId(v string)`

SetDeleterId sets DeleterId field to given value.

### HasDeleterId

`func (o *Item) HasDeleterId() bool`

HasDeleterId returns a boolean if a field has been set.

### SetDeleterIdNil

`func (o *Item) SetDeleterIdNil(b bool)`

 SetDeleterIdNil sets the value for DeleterId to be an explicit nil

### UnsetDeleterId
`func (o *Item) UnsetDeleterId()`

UnsetDeleterId ensures that no value is present for DeleterId, not even an explicit nil
### GetDeletionTime

`func (o *Item) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *Item) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *Item) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *Item) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### SetDeletionTimeNil

`func (o *Item) SetDeletionTimeNil(b bool)`

 SetDeletionTimeNil sets the value for DeletionTime to be an explicit nil

### UnsetDeletionTime
`func (o *Item) UnsetDeletionTime()`

UnsetDeletionTime ensures that no value is present for DeletionTime, not even an explicit nil
### GetDeleter

`func (o *Item) GetDeleter() IdentityUser`

GetDeleter returns the Deleter field if non-nil, zero value otherwise.

### GetDeleterOk

`func (o *Item) GetDeleterOk() (*IdentityUser, bool)`

GetDeleterOk returns a tuple with the Deleter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleter

`func (o *Item) SetDeleter(v IdentityUser)`

SetDeleter sets Deleter field to given value.

### HasDeleter

`func (o *Item) HasDeleter() bool`

HasDeleter returns a boolean if a field has been set.

### GetCreator

`func (o *Item) GetCreator() IdentityUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Item) GetCreatorOk() (*IdentityUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Item) SetCreator(v IdentityUser)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *Item) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetLastModifier

`func (o *Item) GetLastModifier() IdentityUser`

GetLastModifier returns the LastModifier field if non-nil, zero value otherwise.

### GetLastModifierOk

`func (o *Item) GetLastModifierOk() (*IdentityUser, bool)`

GetLastModifierOk returns a tuple with the LastModifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifier

`func (o *Item) SetLastModifier(v IdentityUser)`

SetLastModifier sets LastModifier field to given value.

### HasLastModifier

`func (o *Item) HasLastModifier() bool`

HasLastModifier returns a boolean if a field has been set.

### GetName

`func (o *Item) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Item) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Item) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Item) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Item) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Item) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *Item) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Item) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Item) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Item) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Item) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Item) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTotalCount

`func (o *Item) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *Item) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *Item) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *Item) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetPassword

`func (o *Item) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Item) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Item) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Item) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *Item) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *Item) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetParentItemId

`func (o *Item) GetParentItemId() string`

GetParentItemId returns the ParentItemId field if non-nil, zero value otherwise.

### GetParentItemIdOk

`func (o *Item) GetParentItemIdOk() (*string, bool)`

GetParentItemIdOk returns a tuple with the ParentItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentItemId

`func (o *Item) SetParentItemId(v string)`

SetParentItemId sets ParentItemId field to given value.

### HasParentItemId

`func (o *Item) HasParentItemId() bool`

HasParentItemId returns a boolean if a field has been set.

### SetParentItemIdNil

`func (o *Item) SetParentItemIdNil(b bool)`

 SetParentItemIdNil sets the value for ParentItemId to be an explicit nil

### UnsetParentItemId
`func (o *Item) UnsetParentItemId()`

UnsetParentItemId ensures that no value is present for ParentItemId, not even an explicit nil
### GetParentItem

`func (o *Item) GetParentItem() Item`

GetParentItem returns the ParentItem field if non-nil, zero value otherwise.

### GetParentItemOk

`func (o *Item) GetParentItemOk() (*Item, bool)`

GetParentItemOk returns a tuple with the ParentItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentItem

`func (o *Item) SetParentItem(v Item)`

SetParentItem sets ParentItem field to given value.

### HasParentItem

`func (o *Item) HasParentItem() bool`

HasParentItem returns a boolean if a field has been set.

### GetChildren

`func (o *Item) GetChildren() []Item`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Item) GetChildrenOk() (*[]Item, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Item) SetChildren(v []Item)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *Item) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### SetChildrenNil

`func (o *Item) SetChildrenNil(b bool)`

 SetChildrenNil sets the value for Children to be an explicit nil

### UnsetChildren
`func (o *Item) UnsetChildren()`

UnsetChildren ensures that no value is present for Children, not even an explicit nil
### GetThumbId

`func (o *Item) GetThumbId() string`

GetThumbId returns the ThumbId field if non-nil, zero value otherwise.

### GetThumbIdOk

`func (o *Item) GetThumbIdOk() (*string, bool)`

GetThumbIdOk returns a tuple with the ThumbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbId

`func (o *Item) SetThumbId(v string)`

SetThumbId sets ThumbId field to given value.

### HasThumbId

`func (o *Item) HasThumbId() bool`

HasThumbId returns a boolean if a field has been set.

### SetThumbIdNil

`func (o *Item) SetThumbIdNil(b bool)`

 SetThumbIdNil sets the value for ThumbId to be an explicit nil

### UnsetThumbId
`func (o *Item) UnsetThumbId()`

UnsetThumbId ensures that no value is present for ThumbId, not even an explicit nil
### GetThumb

`func (o *Item) GetThumb() ItemThumb`

GetThumb returns the Thumb field if non-nil, zero value otherwise.

### GetThumbOk

`func (o *Item) GetThumbOk() (*ItemThumb, bool)`

GetThumbOk returns a tuple with the Thumb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumb

`func (o *Item) SetThumb(v ItemThumb)`

SetThumb sets Thumb field to given value.

### HasThumb

`func (o *Item) HasThumb() bool`

HasThumb returns a boolean if a field has been set.

### GetKey

`func (o *Item) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Item) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Item) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Item) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *Item) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *Item) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetSize

`func (o *Item) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Item) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Item) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *Item) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMd5

`func (o *Item) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *Item) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *Item) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *Item) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### SetMd5Nil

`func (o *Item) SetMd5Nil(b bool)`

 SetMd5Nil sets the value for Md5 to be an explicit nil

### UnsetMd5
`func (o *Item) UnsetMd5()`

UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
### GetContentType

`func (o *Item) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *Item) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *Item) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *Item) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *Item) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *Item) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetType

`func (o *Item) GetType() ItemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Item) GetTypeOk() (*ItemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Item) SetType(v ItemType)`

SetType sets Type field to given value.

### HasType

`func (o *Item) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExtension

`func (o *Item) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *Item) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *Item) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *Item) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### SetExtensionNil

`func (o *Item) SetExtensionNil(b bool)`

 SetExtensionNil sets the value for Extension to be an explicit nil

### UnsetExtension
`func (o *Item) UnsetExtension()`

UnsetExtension ensures that no value is present for Extension, not even an explicit nil
### GetStorageClass

`func (o *Item) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *Item) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *Item) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *Item) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### SetStorageClassNil

`func (o *Item) SetStorageClassNil(b bool)`

 SetStorageClassNil sets the value for StorageClass to be an explicit nil

### UnsetStorageClass
`func (o *Item) UnsetStorageClass()`

UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil
### GetFileCreatedAt

`func (o *Item) GetFileCreatedAt() time.Time`

GetFileCreatedAt returns the FileCreatedAt field if non-nil, zero value otherwise.

### GetFileCreatedAtOk

`func (o *Item) GetFileCreatedAtOk() (*time.Time, bool)`

GetFileCreatedAtOk returns a tuple with the FileCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCreatedAt

`func (o *Item) SetFileCreatedAt(v time.Time)`

SetFileCreatedAt sets FileCreatedAt field to given value.

### HasFileCreatedAt

`func (o *Item) HasFileCreatedAt() bool`

HasFileCreatedAt returns a boolean if a field has been set.

### SetFileCreatedAtNil

`func (o *Item) SetFileCreatedAtNil(b bool)`

 SetFileCreatedAtNil sets the value for FileCreatedAt to be an explicit nil

### UnsetFileCreatedAt
`func (o *Item) UnsetFileCreatedAt()`

UnsetFileCreatedAt ensures that no value is present for FileCreatedAt, not even an explicit nil
### GetFileUpdatedAt

`func (o *Item) GetFileUpdatedAt() time.Time`

GetFileUpdatedAt returns the FileUpdatedAt field if non-nil, zero value otherwise.

### GetFileUpdatedAtOk

`func (o *Item) GetFileUpdatedAtOk() (*time.Time, bool)`

GetFileUpdatedAtOk returns a tuple with the FileUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUpdatedAt

`func (o *Item) SetFileUpdatedAt(v time.Time)`

SetFileUpdatedAt sets FileUpdatedAt field to given value.

### HasFileUpdatedAt

`func (o *Item) HasFileUpdatedAt() bool`

HasFileUpdatedAt returns a boolean if a field has been set.

### SetFileUpdatedAtNil

`func (o *Item) SetFileUpdatedAtNil(b bool)`

 SetFileUpdatedAtNil sets the value for FileUpdatedAt to be an explicit nil

### UnsetFileUpdatedAt
`func (o *Item) UnsetFileUpdatedAt()`

UnsetFileUpdatedAt ensures that no value is present for FileUpdatedAt, not even an explicit nil
### GetFileId

`func (o *Item) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *Item) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *Item) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *Item) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### SetFileIdNil

`func (o *Item) SetFileIdNil(b bool)`

 SetFileIdNil sets the value for FileId to be an explicit nil

### UnsetFileId
`func (o *Item) UnsetFileId()`

UnsetFileId ensures that no value is present for FileId, not even an explicit nil
### GetFile

`func (o *Item) GetFile() File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *Item) GetFileOk() (*File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *Item) SetFile(v File)`

SetFile sets File field to given value.

### HasFile

`func (o *Item) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetDisplayStyle

`func (o *Item) GetDisplayStyle() DisplayStyle`

GetDisplayStyle returns the DisplayStyle field if non-nil, zero value otherwise.

### GetDisplayStyleOk

`func (o *Item) GetDisplayStyleOk() (*DisplayStyle, bool)`

GetDisplayStyleOk returns a tuple with the DisplayStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayStyle

`func (o *Item) SetDisplayStyle(v DisplayStyle)`

SetDisplayStyle sets DisplayStyle field to given value.

### HasDisplayStyle

`func (o *Item) HasDisplayStyle() bool`

HasDisplayStyle returns a boolean if a field has been set.

### GetTags

`func (o *Item) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Item) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Item) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Item) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *Item) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *Item) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetItemTags

`func (o *Item) GetItemTags() []ItemTag`

GetItemTags returns the ItemTags field if non-nil, zero value otherwise.

### GetItemTagsOk

`func (o *Item) GetItemTagsOk() (*[]ItemTag, bool)`

GetItemTagsOk returns a tuple with the ItemTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTags

`func (o *Item) SetItemTags(v []ItemTag)`

SetItemTags sets ItemTags field to given value.

### HasItemTags

`func (o *Item) HasItemTags() bool`

HasItemTags returns a boolean if a field has been set.

### SetItemTagsNil

`func (o *Item) SetItemTagsNil(b bool)`

 SetItemTagsNil sets the value for ItemTags to be an explicit nil

### UnsetItemTags
`func (o *Item) UnsetItemTags()`

UnsetItemTags ensures that no value is present for ItemTags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateUpdateItemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**ParentItemId** | Pointer to **NullableString** |  | [optional] 
**TagIds** | Pointer to **[]string** |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] 
**Md5** | Pointer to **NullableString** |  | [optional] 
**SliceMd5** | Pointer to **NullableString** |  | [optional] 
**RapidCode** | Pointer to **NullableString** |  | [optional] 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**ItemType**](ItemType.md) |  | [optional] 
**DisplayStyle** | Pointer to [**DisplayStyle**](DisplayStyle.md) |  | [optional] 
**Extension** | Pointer to **NullableString** |  | [optional] 
**StorageClass** | Pointer to **NullableString** |  | [optional] 
**FileCreatedAt** | Pointer to **NullableTime** |  | [optional] 
**FileUpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**SyncVersion** | Pointer to **int64** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**DeletionTime** | Pointer to **NullableTime** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**LastModificationTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreateUpdateItemDto

`func NewCreateUpdateItemDto() *CreateUpdateItemDto`

NewCreateUpdateItemDto instantiates a new CreateUpdateItemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdateItemDtoWithDefaults

`func NewCreateUpdateItemDtoWithDefaults() *CreateUpdateItemDto`

NewCreateUpdateItemDtoWithDefaults instantiates a new CreateUpdateItemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateUpdateItemDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUpdateItemDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUpdateItemDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateUpdateItemDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateUpdateItemDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateUpdateItemDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CreateUpdateItemDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateUpdateItemDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateUpdateItemDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateUpdateItemDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateUpdateItemDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateUpdateItemDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPassword

`func (o *CreateUpdateItemDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateUpdateItemDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateUpdateItemDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateUpdateItemDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *CreateUpdateItemDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CreateUpdateItemDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetParentItemId

`func (o *CreateUpdateItemDto) GetParentItemId() string`

GetParentItemId returns the ParentItemId field if non-nil, zero value otherwise.

### GetParentItemIdOk

`func (o *CreateUpdateItemDto) GetParentItemIdOk() (*string, bool)`

GetParentItemIdOk returns a tuple with the ParentItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentItemId

`func (o *CreateUpdateItemDto) SetParentItemId(v string)`

SetParentItemId sets ParentItemId field to given value.

### HasParentItemId

`func (o *CreateUpdateItemDto) HasParentItemId() bool`

HasParentItemId returns a boolean if a field has been set.

### SetParentItemIdNil

`func (o *CreateUpdateItemDto) SetParentItemIdNil(b bool)`

 SetParentItemIdNil sets the value for ParentItemId to be an explicit nil

### UnsetParentItemId
`func (o *CreateUpdateItemDto) UnsetParentItemId()`

UnsetParentItemId ensures that no value is present for ParentItemId, not even an explicit nil
### GetTagIds

`func (o *CreateUpdateItemDto) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *CreateUpdateItemDto) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *CreateUpdateItemDto) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *CreateUpdateItemDto) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### SetTagIdsNil

`func (o *CreateUpdateItemDto) SetTagIdsNil(b bool)`

 SetTagIdsNil sets the value for TagIds to be an explicit nil

### UnsetTagIds
`func (o *CreateUpdateItemDto) UnsetTagIds()`

UnsetTagIds ensures that no value is present for TagIds, not even an explicit nil
### GetKey

`func (o *CreateUpdateItemDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateUpdateItemDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateUpdateItemDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateUpdateItemDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *CreateUpdateItemDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *CreateUpdateItemDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetMd5

`func (o *CreateUpdateItemDto) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *CreateUpdateItemDto) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *CreateUpdateItemDto) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *CreateUpdateItemDto) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### SetMd5Nil

`func (o *CreateUpdateItemDto) SetMd5Nil(b bool)`

 SetMd5Nil sets the value for Md5 to be an explicit nil

### UnsetMd5
`func (o *CreateUpdateItemDto) UnsetMd5()`

UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
### GetSliceMd5

`func (o *CreateUpdateItemDto) GetSliceMd5() string`

GetSliceMd5 returns the SliceMd5 field if non-nil, zero value otherwise.

### GetSliceMd5Ok

`func (o *CreateUpdateItemDto) GetSliceMd5Ok() (*string, bool)`

GetSliceMd5Ok returns a tuple with the SliceMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceMd5

`func (o *CreateUpdateItemDto) SetSliceMd5(v string)`

SetSliceMd5 sets SliceMd5 field to given value.

### HasSliceMd5

`func (o *CreateUpdateItemDto) HasSliceMd5() bool`

HasSliceMd5 returns a boolean if a field has been set.

### SetSliceMd5Nil

`func (o *CreateUpdateItemDto) SetSliceMd5Nil(b bool)`

 SetSliceMd5Nil sets the value for SliceMd5 to be an explicit nil

### UnsetSliceMd5
`func (o *CreateUpdateItemDto) UnsetSliceMd5()`

UnsetSliceMd5 ensures that no value is present for SliceMd5, not even an explicit nil
### GetRapidCode

`func (o *CreateUpdateItemDto) GetRapidCode() string`

GetRapidCode returns the RapidCode field if non-nil, zero value otherwise.

### GetRapidCodeOk

`func (o *CreateUpdateItemDto) GetRapidCodeOk() (*string, bool)`

GetRapidCodeOk returns a tuple with the RapidCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRapidCode

`func (o *CreateUpdateItemDto) SetRapidCode(v string)`

SetRapidCode sets RapidCode field to given value.

### HasRapidCode

`func (o *CreateUpdateItemDto) HasRapidCode() bool`

HasRapidCode returns a boolean if a field has been set.

### SetRapidCodeNil

`func (o *CreateUpdateItemDto) SetRapidCodeNil(b bool)`

 SetRapidCodeNil sets the value for RapidCode to be an explicit nil

### UnsetRapidCode
`func (o *CreateUpdateItemDto) UnsetRapidCode()`

UnsetRapidCode ensures that no value is present for RapidCode, not even an explicit nil
### GetContentType

`func (o *CreateUpdateItemDto) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CreateUpdateItemDto) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CreateUpdateItemDto) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *CreateUpdateItemDto) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *CreateUpdateItemDto) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *CreateUpdateItemDto) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetType

`func (o *CreateUpdateItemDto) GetType() ItemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateUpdateItemDto) GetTypeOk() (*ItemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateUpdateItemDto) SetType(v ItemType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateUpdateItemDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisplayStyle

`func (o *CreateUpdateItemDto) GetDisplayStyle() DisplayStyle`

GetDisplayStyle returns the DisplayStyle field if non-nil, zero value otherwise.

### GetDisplayStyleOk

`func (o *CreateUpdateItemDto) GetDisplayStyleOk() (*DisplayStyle, bool)`

GetDisplayStyleOk returns a tuple with the DisplayStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayStyle

`func (o *CreateUpdateItemDto) SetDisplayStyle(v DisplayStyle)`

SetDisplayStyle sets DisplayStyle field to given value.

### HasDisplayStyle

`func (o *CreateUpdateItemDto) HasDisplayStyle() bool`

HasDisplayStyle returns a boolean if a field has been set.

### GetExtension

`func (o *CreateUpdateItemDto) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *CreateUpdateItemDto) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *CreateUpdateItemDto) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *CreateUpdateItemDto) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### SetExtensionNil

`func (o *CreateUpdateItemDto) SetExtensionNil(b bool)`

 SetExtensionNil sets the value for Extension to be an explicit nil

### UnsetExtension
`func (o *CreateUpdateItemDto) UnsetExtension()`

UnsetExtension ensures that no value is present for Extension, not even an explicit nil
### GetStorageClass

`func (o *CreateUpdateItemDto) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *CreateUpdateItemDto) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *CreateUpdateItemDto) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *CreateUpdateItemDto) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### SetStorageClassNil

`func (o *CreateUpdateItemDto) SetStorageClassNil(b bool)`

 SetStorageClassNil sets the value for StorageClass to be an explicit nil

### UnsetStorageClass
`func (o *CreateUpdateItemDto) UnsetStorageClass()`

UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil
### GetFileCreatedAt

`func (o *CreateUpdateItemDto) GetFileCreatedAt() time.Time`

GetFileCreatedAt returns the FileCreatedAt field if non-nil, zero value otherwise.

### GetFileCreatedAtOk

`func (o *CreateUpdateItemDto) GetFileCreatedAtOk() (*time.Time, bool)`

GetFileCreatedAtOk returns a tuple with the FileCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCreatedAt

`func (o *CreateUpdateItemDto) SetFileCreatedAt(v time.Time)`

SetFileCreatedAt sets FileCreatedAt field to given value.

### HasFileCreatedAt

`func (o *CreateUpdateItemDto) HasFileCreatedAt() bool`

HasFileCreatedAt returns a boolean if a field has been set.

### SetFileCreatedAtNil

`func (o *CreateUpdateItemDto) SetFileCreatedAtNil(b bool)`

 SetFileCreatedAtNil sets the value for FileCreatedAt to be an explicit nil

### UnsetFileCreatedAt
`func (o *CreateUpdateItemDto) UnsetFileCreatedAt()`

UnsetFileCreatedAt ensures that no value is present for FileCreatedAt, not even an explicit nil
### GetFileUpdatedAt

`func (o *CreateUpdateItemDto) GetFileUpdatedAt() time.Time`

GetFileUpdatedAt returns the FileUpdatedAt field if non-nil, zero value otherwise.

### GetFileUpdatedAtOk

`func (o *CreateUpdateItemDto) GetFileUpdatedAtOk() (*time.Time, bool)`

GetFileUpdatedAtOk returns a tuple with the FileUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUpdatedAt

`func (o *CreateUpdateItemDto) SetFileUpdatedAt(v time.Time)`

SetFileUpdatedAt sets FileUpdatedAt field to given value.

### HasFileUpdatedAt

`func (o *CreateUpdateItemDto) HasFileUpdatedAt() bool`

HasFileUpdatedAt returns a boolean if a field has been set.

### SetFileUpdatedAtNil

`func (o *CreateUpdateItemDto) SetFileUpdatedAtNil(b bool)`

 SetFileUpdatedAtNil sets the value for FileUpdatedAt to be an explicit nil

### UnsetFileUpdatedAt
`func (o *CreateUpdateItemDto) UnsetFileUpdatedAt()`

UnsetFileUpdatedAt ensures that no value is present for FileUpdatedAt, not even an explicit nil
### GetSyncVersion

`func (o *CreateUpdateItemDto) GetSyncVersion() int64`

GetSyncVersion returns the SyncVersion field if non-nil, zero value otherwise.

### GetSyncVersionOk

`func (o *CreateUpdateItemDto) GetSyncVersionOk() (*int64, bool)`

GetSyncVersionOk returns a tuple with the SyncVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncVersion

`func (o *CreateUpdateItemDto) SetSyncVersion(v int64)`

SetSyncVersion sets SyncVersion field to given value.

### HasSyncVersion

`func (o *CreateUpdateItemDto) HasSyncVersion() bool`

HasSyncVersion returns a boolean if a field has been set.

### GetIsDeleted

`func (o *CreateUpdateItemDto) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CreateUpdateItemDto) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CreateUpdateItemDto) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *CreateUpdateItemDto) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeletionTime

`func (o *CreateUpdateItemDto) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *CreateUpdateItemDto) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *CreateUpdateItemDto) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *CreateUpdateItemDto) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### SetDeletionTimeNil

`func (o *CreateUpdateItemDto) SetDeletionTimeNil(b bool)`

 SetDeletionTimeNil sets the value for DeletionTime to be an explicit nil

### UnsetDeletionTime
`func (o *CreateUpdateItemDto) UnsetDeletionTime()`

UnsetDeletionTime ensures that no value is present for DeletionTime, not even an explicit nil
### GetCreationTime

`func (o *CreateUpdateItemDto) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CreateUpdateItemDto) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CreateUpdateItemDto) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *CreateUpdateItemDto) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetLastModificationTime

`func (o *CreateUpdateItemDto) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *CreateUpdateItemDto) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *CreateUpdateItemDto) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *CreateUpdateItemDto) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



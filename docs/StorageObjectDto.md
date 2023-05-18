# StorageObjectDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**CreatorId** | Pointer to **NullableString** |  | [optional] 
**LastModificationTime** | Pointer to **NullableTime** |  | [optional] 
**LastModifierId** | Pointer to **NullableString** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**DeleterId** | Pointer to **NullableString** |  | [optional] 
**DeletionTime** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Md5** | Pointer to **NullableString** |  | [optional] 
**SliceMd5** | Pointer to **NullableString** |  | [optional] 
**RapidCode** | Pointer to **NullableString** |  | [optional] 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**Extension** | Pointer to **NullableString** |  | [optional] 
**StorageClass** | Pointer to **NullableString** |  | [optional] 
**StorageObjectCreatedAt** | Pointer to **NullableTime** |  | [optional] 
**StorageObjectUpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**SyncVersion** | Pointer to **int64** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewStorageObjectDto

`func NewStorageObjectDto() *StorageObjectDto`

NewStorageObjectDto instantiates a new StorageObjectDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageObjectDtoWithDefaults

`func NewStorageObjectDtoWithDefaults() *StorageObjectDto`

NewStorageObjectDtoWithDefaults instantiates a new StorageObjectDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageObjectDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageObjectDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageObjectDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StorageObjectDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreationTime

`func (o *StorageObjectDto) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *StorageObjectDto) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *StorageObjectDto) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *StorageObjectDto) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *StorageObjectDto) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *StorageObjectDto) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *StorageObjectDto) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *StorageObjectDto) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### SetCreatorIdNil

`func (o *StorageObjectDto) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *StorageObjectDto) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetLastModificationTime

`func (o *StorageObjectDto) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *StorageObjectDto) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *StorageObjectDto) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *StorageObjectDto) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### SetLastModificationTimeNil

`func (o *StorageObjectDto) SetLastModificationTimeNil(b bool)`

 SetLastModificationTimeNil sets the value for LastModificationTime to be an explicit nil

### UnsetLastModificationTime
`func (o *StorageObjectDto) UnsetLastModificationTime()`

UnsetLastModificationTime ensures that no value is present for LastModificationTime, not even an explicit nil
### GetLastModifierId

`func (o *StorageObjectDto) GetLastModifierId() string`

GetLastModifierId returns the LastModifierId field if non-nil, zero value otherwise.

### GetLastModifierIdOk

`func (o *StorageObjectDto) GetLastModifierIdOk() (*string, bool)`

GetLastModifierIdOk returns a tuple with the LastModifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierId

`func (o *StorageObjectDto) SetLastModifierId(v string)`

SetLastModifierId sets LastModifierId field to given value.

### HasLastModifierId

`func (o *StorageObjectDto) HasLastModifierId() bool`

HasLastModifierId returns a boolean if a field has been set.

### SetLastModifierIdNil

`func (o *StorageObjectDto) SetLastModifierIdNil(b bool)`

 SetLastModifierIdNil sets the value for LastModifierId to be an explicit nil

### UnsetLastModifierId
`func (o *StorageObjectDto) UnsetLastModifierId()`

UnsetLastModifierId ensures that no value is present for LastModifierId, not even an explicit nil
### GetIsDeleted

`func (o *StorageObjectDto) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *StorageObjectDto) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *StorageObjectDto) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *StorageObjectDto) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeleterId

`func (o *StorageObjectDto) GetDeleterId() string`

GetDeleterId returns the DeleterId field if non-nil, zero value otherwise.

### GetDeleterIdOk

`func (o *StorageObjectDto) GetDeleterIdOk() (*string, bool)`

GetDeleterIdOk returns a tuple with the DeleterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleterId

`func (o *StorageObjectDto) SetDeleterId(v string)`

SetDeleterId sets DeleterId field to given value.

### HasDeleterId

`func (o *StorageObjectDto) HasDeleterId() bool`

HasDeleterId returns a boolean if a field has been set.

### SetDeleterIdNil

`func (o *StorageObjectDto) SetDeleterIdNil(b bool)`

 SetDeleterIdNil sets the value for DeleterId to be an explicit nil

### UnsetDeleterId
`func (o *StorageObjectDto) UnsetDeleterId()`

UnsetDeleterId ensures that no value is present for DeleterId, not even an explicit nil
### GetDeletionTime

`func (o *StorageObjectDto) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *StorageObjectDto) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *StorageObjectDto) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *StorageObjectDto) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### SetDeletionTimeNil

`func (o *StorageObjectDto) SetDeletionTimeNil(b bool)`

 SetDeletionTimeNil sets the value for DeletionTime to be an explicit nil

### UnsetDeletionTime
`func (o *StorageObjectDto) UnsetDeletionTime()`

UnsetDeletionTime ensures that no value is present for DeletionTime, not even an explicit nil
### GetName

`func (o *StorageObjectDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageObjectDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageObjectDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageObjectDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *StorageObjectDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *StorageObjectDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetKey

`func (o *StorageObjectDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *StorageObjectDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *StorageObjectDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *StorageObjectDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *StorageObjectDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *StorageObjectDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetUrl

`func (o *StorageObjectDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StorageObjectDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StorageObjectDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StorageObjectDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *StorageObjectDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *StorageObjectDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetSize

`func (o *StorageObjectDto) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageObjectDto) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageObjectDto) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageObjectDto) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMd5

`func (o *StorageObjectDto) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *StorageObjectDto) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *StorageObjectDto) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *StorageObjectDto) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### SetMd5Nil

`func (o *StorageObjectDto) SetMd5Nil(b bool)`

 SetMd5Nil sets the value for Md5 to be an explicit nil

### UnsetMd5
`func (o *StorageObjectDto) UnsetMd5()`

UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
### GetSliceMd5

`func (o *StorageObjectDto) GetSliceMd5() string`

GetSliceMd5 returns the SliceMd5 field if non-nil, zero value otherwise.

### GetSliceMd5Ok

`func (o *StorageObjectDto) GetSliceMd5Ok() (*string, bool)`

GetSliceMd5Ok returns a tuple with the SliceMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceMd5

`func (o *StorageObjectDto) SetSliceMd5(v string)`

SetSliceMd5 sets SliceMd5 field to given value.

### HasSliceMd5

`func (o *StorageObjectDto) HasSliceMd5() bool`

HasSliceMd5 returns a boolean if a field has been set.

### SetSliceMd5Nil

`func (o *StorageObjectDto) SetSliceMd5Nil(b bool)`

 SetSliceMd5Nil sets the value for SliceMd5 to be an explicit nil

### UnsetSliceMd5
`func (o *StorageObjectDto) UnsetSliceMd5()`

UnsetSliceMd5 ensures that no value is present for SliceMd5, not even an explicit nil
### GetRapidCode

`func (o *StorageObjectDto) GetRapidCode() string`

GetRapidCode returns the RapidCode field if non-nil, zero value otherwise.

### GetRapidCodeOk

`func (o *StorageObjectDto) GetRapidCodeOk() (*string, bool)`

GetRapidCodeOk returns a tuple with the RapidCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRapidCode

`func (o *StorageObjectDto) SetRapidCode(v string)`

SetRapidCode sets RapidCode field to given value.

### HasRapidCode

`func (o *StorageObjectDto) HasRapidCode() bool`

HasRapidCode returns a boolean if a field has been set.

### SetRapidCodeNil

`func (o *StorageObjectDto) SetRapidCodeNil(b bool)`

 SetRapidCodeNil sets the value for RapidCode to be an explicit nil

### UnsetRapidCode
`func (o *StorageObjectDto) UnsetRapidCode()`

UnsetRapidCode ensures that no value is present for RapidCode, not even an explicit nil
### GetContentType

`func (o *StorageObjectDto) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *StorageObjectDto) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *StorageObjectDto) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *StorageObjectDto) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *StorageObjectDto) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *StorageObjectDto) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetExtension

`func (o *StorageObjectDto) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *StorageObjectDto) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *StorageObjectDto) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *StorageObjectDto) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### SetExtensionNil

`func (o *StorageObjectDto) SetExtensionNil(b bool)`

 SetExtensionNil sets the value for Extension to be an explicit nil

### UnsetExtension
`func (o *StorageObjectDto) UnsetExtension()`

UnsetExtension ensures that no value is present for Extension, not even an explicit nil
### GetStorageClass

`func (o *StorageObjectDto) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *StorageObjectDto) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *StorageObjectDto) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *StorageObjectDto) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### SetStorageClassNil

`func (o *StorageObjectDto) SetStorageClassNil(b bool)`

 SetStorageClassNil sets the value for StorageClass to be an explicit nil

### UnsetStorageClass
`func (o *StorageObjectDto) UnsetStorageClass()`

UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil
### GetStorageObjectCreatedAt

`func (o *StorageObjectDto) GetStorageObjectCreatedAt() time.Time`

GetStorageObjectCreatedAt returns the StorageObjectCreatedAt field if non-nil, zero value otherwise.

### GetStorageObjectCreatedAtOk

`func (o *StorageObjectDto) GetStorageObjectCreatedAtOk() (*time.Time, bool)`

GetStorageObjectCreatedAtOk returns a tuple with the StorageObjectCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageObjectCreatedAt

`func (o *StorageObjectDto) SetStorageObjectCreatedAt(v time.Time)`

SetStorageObjectCreatedAt sets StorageObjectCreatedAt field to given value.

### HasStorageObjectCreatedAt

`func (o *StorageObjectDto) HasStorageObjectCreatedAt() bool`

HasStorageObjectCreatedAt returns a boolean if a field has been set.

### SetStorageObjectCreatedAtNil

`func (o *StorageObjectDto) SetStorageObjectCreatedAtNil(b bool)`

 SetStorageObjectCreatedAtNil sets the value for StorageObjectCreatedAt to be an explicit nil

### UnsetStorageObjectCreatedAt
`func (o *StorageObjectDto) UnsetStorageObjectCreatedAt()`

UnsetStorageObjectCreatedAt ensures that no value is present for StorageObjectCreatedAt, not even an explicit nil
### GetStorageObjectUpdatedAt

`func (o *StorageObjectDto) GetStorageObjectUpdatedAt() time.Time`

GetStorageObjectUpdatedAt returns the StorageObjectUpdatedAt field if non-nil, zero value otherwise.

### GetStorageObjectUpdatedAtOk

`func (o *StorageObjectDto) GetStorageObjectUpdatedAtOk() (*time.Time, bool)`

GetStorageObjectUpdatedAtOk returns a tuple with the StorageObjectUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageObjectUpdatedAt

`func (o *StorageObjectDto) SetStorageObjectUpdatedAt(v time.Time)`

SetStorageObjectUpdatedAt sets StorageObjectUpdatedAt field to given value.

### HasStorageObjectUpdatedAt

`func (o *StorageObjectDto) HasStorageObjectUpdatedAt() bool`

HasStorageObjectUpdatedAt returns a boolean if a field has been set.

### SetStorageObjectUpdatedAtNil

`func (o *StorageObjectDto) SetStorageObjectUpdatedAtNil(b bool)`

 SetStorageObjectUpdatedAtNil sets the value for StorageObjectUpdatedAt to be an explicit nil

### UnsetStorageObjectUpdatedAt
`func (o *StorageObjectDto) UnsetStorageObjectUpdatedAt()`

UnsetStorageObjectUpdatedAt ensures that no value is present for StorageObjectUpdatedAt, not even an explicit nil
### GetSyncVersion

`func (o *StorageObjectDto) GetSyncVersion() int64`

GetSyncVersion returns the SyncVersion field if non-nil, zero value otherwise.

### GetSyncVersionOk

`func (o *StorageObjectDto) GetSyncVersionOk() (*int64, bool)`

GetSyncVersionOk returns a tuple with the SyncVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncVersion

`func (o *StorageObjectDto) SetSyncVersion(v int64)`

SetSyncVersion sets SyncVersion field to given value.

### HasSyncVersion

`func (o *StorageObjectDto) HasSyncVersion() bool`

HasSyncVersion returns a boolean if a field has been set.

### GetPassword

`func (o *StorageObjectDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *StorageObjectDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *StorageObjectDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *StorageObjectDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *StorageObjectDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *StorageObjectDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



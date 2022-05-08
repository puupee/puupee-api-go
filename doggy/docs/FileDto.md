# FileDto

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
**Size** | Pointer to **int32** |  | [optional] 
**Md5** | Pointer to **NullableString** |  | [optional] 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**Extension** | Pointer to **NullableString** |  | [optional] 
**StorageClass** | Pointer to **NullableString** |  | [optional] 
**FileCreatedAt** | Pointer to **NullableTime** |  | [optional] 
**FileUpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**FileId** | Pointer to **NullableString** |  | [optional] 
**File** | Pointer to [**FileDto**](FileDto.md) |  | [optional] 
**Folders** | Pointer to [**[]ItemDto**](ItemDto.md) |  | [optional] 
**Tags** | Pointer to [**[]TagDto**](TagDto.md) |  | [optional] 

## Methods

### NewFileDto

`func NewFileDto() *FileDto`

NewFileDto instantiates a new FileDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileDtoWithDefaults

`func NewFileDtoWithDefaults() *FileDto`

NewFileDtoWithDefaults instantiates a new FileDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FileDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreationTime

`func (o *FileDto) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *FileDto) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *FileDto) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *FileDto) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *FileDto) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *FileDto) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *FileDto) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *FileDto) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### SetCreatorIdNil

`func (o *FileDto) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *FileDto) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetLastModificationTime

`func (o *FileDto) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *FileDto) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *FileDto) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *FileDto) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### SetLastModificationTimeNil

`func (o *FileDto) SetLastModificationTimeNil(b bool)`

 SetLastModificationTimeNil sets the value for LastModificationTime to be an explicit nil

### UnsetLastModificationTime
`func (o *FileDto) UnsetLastModificationTime()`

UnsetLastModificationTime ensures that no value is present for LastModificationTime, not even an explicit nil
### GetLastModifierId

`func (o *FileDto) GetLastModifierId() string`

GetLastModifierId returns the LastModifierId field if non-nil, zero value otherwise.

### GetLastModifierIdOk

`func (o *FileDto) GetLastModifierIdOk() (*string, bool)`

GetLastModifierIdOk returns a tuple with the LastModifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierId

`func (o *FileDto) SetLastModifierId(v string)`

SetLastModifierId sets LastModifierId field to given value.

### HasLastModifierId

`func (o *FileDto) HasLastModifierId() bool`

HasLastModifierId returns a boolean if a field has been set.

### SetLastModifierIdNil

`func (o *FileDto) SetLastModifierIdNil(b bool)`

 SetLastModifierIdNil sets the value for LastModifierId to be an explicit nil

### UnsetLastModifierId
`func (o *FileDto) UnsetLastModifierId()`

UnsetLastModifierId ensures that no value is present for LastModifierId, not even an explicit nil
### GetIsDeleted

`func (o *FileDto) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *FileDto) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *FileDto) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *FileDto) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeleterId

`func (o *FileDto) GetDeleterId() string`

GetDeleterId returns the DeleterId field if non-nil, zero value otherwise.

### GetDeleterIdOk

`func (o *FileDto) GetDeleterIdOk() (*string, bool)`

GetDeleterIdOk returns a tuple with the DeleterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleterId

`func (o *FileDto) SetDeleterId(v string)`

SetDeleterId sets DeleterId field to given value.

### HasDeleterId

`func (o *FileDto) HasDeleterId() bool`

HasDeleterId returns a boolean if a field has been set.

### SetDeleterIdNil

`func (o *FileDto) SetDeleterIdNil(b bool)`

 SetDeleterIdNil sets the value for DeleterId to be an explicit nil

### UnsetDeleterId
`func (o *FileDto) UnsetDeleterId()`

UnsetDeleterId ensures that no value is present for DeleterId, not even an explicit nil
### GetDeletionTime

`func (o *FileDto) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *FileDto) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *FileDto) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *FileDto) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### SetDeletionTimeNil

`func (o *FileDto) SetDeletionTimeNil(b bool)`

 SetDeletionTimeNil sets the value for DeletionTime to be an explicit nil

### UnsetDeletionTime
`func (o *FileDto) UnsetDeletionTime()`

UnsetDeletionTime ensures that no value is present for DeletionTime, not even an explicit nil
### GetName

`func (o *FileDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FileDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FileDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FileDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetKey

`func (o *FileDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FileDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FileDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FileDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *FileDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *FileDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetUrl

`func (o *FileDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FileDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FileDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FileDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *FileDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *FileDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetSize

`func (o *FileDto) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FileDto) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FileDto) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *FileDto) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMd5

`func (o *FileDto) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *FileDto) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *FileDto) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *FileDto) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### SetMd5Nil

`func (o *FileDto) SetMd5Nil(b bool)`

 SetMd5Nil sets the value for Md5 to be an explicit nil

### UnsetMd5
`func (o *FileDto) UnsetMd5()`

UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
### GetContentType

`func (o *FileDto) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *FileDto) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *FileDto) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *FileDto) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *FileDto) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *FileDto) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetExtension

`func (o *FileDto) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *FileDto) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *FileDto) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *FileDto) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### SetExtensionNil

`func (o *FileDto) SetExtensionNil(b bool)`

 SetExtensionNil sets the value for Extension to be an explicit nil

### UnsetExtension
`func (o *FileDto) UnsetExtension()`

UnsetExtension ensures that no value is present for Extension, not even an explicit nil
### GetStorageClass

`func (o *FileDto) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *FileDto) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *FileDto) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *FileDto) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### SetStorageClassNil

`func (o *FileDto) SetStorageClassNil(b bool)`

 SetStorageClassNil sets the value for StorageClass to be an explicit nil

### UnsetStorageClass
`func (o *FileDto) UnsetStorageClass()`

UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil
### GetFileCreatedAt

`func (o *FileDto) GetFileCreatedAt() time.Time`

GetFileCreatedAt returns the FileCreatedAt field if non-nil, zero value otherwise.

### GetFileCreatedAtOk

`func (o *FileDto) GetFileCreatedAtOk() (*time.Time, bool)`

GetFileCreatedAtOk returns a tuple with the FileCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCreatedAt

`func (o *FileDto) SetFileCreatedAt(v time.Time)`

SetFileCreatedAt sets FileCreatedAt field to given value.

### HasFileCreatedAt

`func (o *FileDto) HasFileCreatedAt() bool`

HasFileCreatedAt returns a boolean if a field has been set.

### SetFileCreatedAtNil

`func (o *FileDto) SetFileCreatedAtNil(b bool)`

 SetFileCreatedAtNil sets the value for FileCreatedAt to be an explicit nil

### UnsetFileCreatedAt
`func (o *FileDto) UnsetFileCreatedAt()`

UnsetFileCreatedAt ensures that no value is present for FileCreatedAt, not even an explicit nil
### GetFileUpdatedAt

`func (o *FileDto) GetFileUpdatedAt() time.Time`

GetFileUpdatedAt returns the FileUpdatedAt field if non-nil, zero value otherwise.

### GetFileUpdatedAtOk

`func (o *FileDto) GetFileUpdatedAtOk() (*time.Time, bool)`

GetFileUpdatedAtOk returns a tuple with the FileUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUpdatedAt

`func (o *FileDto) SetFileUpdatedAt(v time.Time)`

SetFileUpdatedAt sets FileUpdatedAt field to given value.

### HasFileUpdatedAt

`func (o *FileDto) HasFileUpdatedAt() bool`

HasFileUpdatedAt returns a boolean if a field has been set.

### SetFileUpdatedAtNil

`func (o *FileDto) SetFileUpdatedAtNil(b bool)`

 SetFileUpdatedAtNil sets the value for FileUpdatedAt to be an explicit nil

### UnsetFileUpdatedAt
`func (o *FileDto) UnsetFileUpdatedAt()`

UnsetFileUpdatedAt ensures that no value is present for FileUpdatedAt, not even an explicit nil
### GetFileId

`func (o *FileDto) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *FileDto) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *FileDto) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *FileDto) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### SetFileIdNil

`func (o *FileDto) SetFileIdNil(b bool)`

 SetFileIdNil sets the value for FileId to be an explicit nil

### UnsetFileId
`func (o *FileDto) UnsetFileId()`

UnsetFileId ensures that no value is present for FileId, not even an explicit nil
### GetFile

`func (o *FileDto) GetFile() FileDto`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *FileDto) GetFileOk() (*FileDto, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *FileDto) SetFile(v FileDto)`

SetFile sets File field to given value.

### HasFile

`func (o *FileDto) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetFolders

`func (o *FileDto) GetFolders() []ItemDto`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *FileDto) GetFoldersOk() (*[]ItemDto, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *FileDto) SetFolders(v []ItemDto)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *FileDto) HasFolders() bool`

HasFolders returns a boolean if a field has been set.

### SetFoldersNil

`func (o *FileDto) SetFoldersNil(b bool)`

 SetFoldersNil sets the value for Folders to be an explicit nil

### UnsetFolders
`func (o *FileDto) UnsetFolders()`

UnsetFolders ensures that no value is present for Folders, not even an explicit nil
### GetTags

`func (o *FileDto) GetTags() []TagDto`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FileDto) GetTagsOk() (*[]TagDto, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FileDto) SetTags(v []TagDto)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FileDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *FileDto) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *FileDto) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



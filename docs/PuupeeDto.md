# PuupeeDto

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
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Text** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 
**Format** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Md5** | Pointer to **NullableString** |  | [optional] 
**SliceMd5** | Pointer to **NullableString** |  | [optional] 
**RapidCode** | Pointer to **NullableString** |  | [optional] 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**DisplayStyle** | Pointer to **NullableString** |  | [optional] 
**Extension** | Pointer to **NullableString** |  | [optional] 
**StorageClass** | Pointer to **NullableString** |  | [optional] 
**StorageObjectCreatedAt** | Pointer to **NullableTime** |  | [optional] 
**StorageObjectUpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**StorageObjectId** | Pointer to **NullableString** |  | [optional] 
**StorageObject** | Pointer to [**StorageObjectDto**](StorageObjectDto.md) |  | [optional] 
**Thumb** | Pointer to [**StorageObjectDto**](StorageObjectDto.md) |  | [optional] 
**Priority** | Pointer to **NullableInt32** |  | [optional] 
**DoneAt** | Pointer to **NullableTime** |  | [optional] 
**IsDone** | Pointer to **bool** |  | [optional] 
**StartAt** | Pointer to **NullableTime** |  | [optional] 
**EndAt** | Pointer to **NullableTime** |  | [optional] 
**NotifyAt** | Pointer to **NullableTime** |  | [optional] 
**NotifyTimingType** | Pointer to **NullableString** |  | [optional] 
**NotifyTimingUnit** | Pointer to **NullableString** |  | [optional] 
**NotifyTimingValue** | Pointer to **NullableInt32** |  | [optional] 
**Repeat** | Pointer to **NullableString** |  | [optional] 
**RepeatOffAt** | Pointer to **NullableTime** |  | [optional] 
**RepeatOffTimes** | Pointer to **NullableInt32** |  | [optional] 
**Repetitions** | Pointer to **NullableInt32** |  | [optional] 
**SyncVersion** | Pointer to **int64** |  | [optional] 
**IsHidden** | Pointer to **bool** |  | [optional] 
**Tagging** | Pointer to **NullableString** |  | [optional] 
**LastModifierDeviceToken** | Pointer to **NullableString** |  | [optional] 
**LastModifierDevice** | Pointer to **NullableString** |  | [optional] 
**AppName** | Pointer to **NullableString** |  | [optional] 
**SortIndex** | Pointer to **int32** |  | [optional] 

## Methods

### NewPuupeeDto

`func NewPuupeeDto() *PuupeeDto`

NewPuupeeDto instantiates a new PuupeeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPuupeeDtoWithDefaults

`func NewPuupeeDtoWithDefaults() *PuupeeDto`

NewPuupeeDtoWithDefaults instantiates a new PuupeeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PuupeeDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PuupeeDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PuupeeDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PuupeeDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreationTime

`func (o *PuupeeDto) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *PuupeeDto) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *PuupeeDto) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *PuupeeDto) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *PuupeeDto) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *PuupeeDto) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *PuupeeDto) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *PuupeeDto) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### SetCreatorIdNil

`func (o *PuupeeDto) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *PuupeeDto) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetLastModificationTime

`func (o *PuupeeDto) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *PuupeeDto) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *PuupeeDto) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *PuupeeDto) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### SetLastModificationTimeNil

`func (o *PuupeeDto) SetLastModificationTimeNil(b bool)`

 SetLastModificationTimeNil sets the value for LastModificationTime to be an explicit nil

### UnsetLastModificationTime
`func (o *PuupeeDto) UnsetLastModificationTime()`

UnsetLastModificationTime ensures that no value is present for LastModificationTime, not even an explicit nil
### GetLastModifierId

`func (o *PuupeeDto) GetLastModifierId() string`

GetLastModifierId returns the LastModifierId field if non-nil, zero value otherwise.

### GetLastModifierIdOk

`func (o *PuupeeDto) GetLastModifierIdOk() (*string, bool)`

GetLastModifierIdOk returns a tuple with the LastModifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierId

`func (o *PuupeeDto) SetLastModifierId(v string)`

SetLastModifierId sets LastModifierId field to given value.

### HasLastModifierId

`func (o *PuupeeDto) HasLastModifierId() bool`

HasLastModifierId returns a boolean if a field has been set.

### SetLastModifierIdNil

`func (o *PuupeeDto) SetLastModifierIdNil(b bool)`

 SetLastModifierIdNil sets the value for LastModifierId to be an explicit nil

### UnsetLastModifierId
`func (o *PuupeeDto) UnsetLastModifierId()`

UnsetLastModifierId ensures that no value is present for LastModifierId, not even an explicit nil
### GetIsDeleted

`func (o *PuupeeDto) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *PuupeeDto) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *PuupeeDto) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *PuupeeDto) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeleterId

`func (o *PuupeeDto) GetDeleterId() string`

GetDeleterId returns the DeleterId field if non-nil, zero value otherwise.

### GetDeleterIdOk

`func (o *PuupeeDto) GetDeleterIdOk() (*string, bool)`

GetDeleterIdOk returns a tuple with the DeleterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleterId

`func (o *PuupeeDto) SetDeleterId(v string)`

SetDeleterId sets DeleterId field to given value.

### HasDeleterId

`func (o *PuupeeDto) HasDeleterId() bool`

HasDeleterId returns a boolean if a field has been set.

### SetDeleterIdNil

`func (o *PuupeeDto) SetDeleterIdNil(b bool)`

 SetDeleterIdNil sets the value for DeleterId to be an explicit nil

### UnsetDeleterId
`func (o *PuupeeDto) UnsetDeleterId()`

UnsetDeleterId ensures that no value is present for DeleterId, not even an explicit nil
### GetDeletionTime

`func (o *PuupeeDto) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *PuupeeDto) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *PuupeeDto) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *PuupeeDto) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### SetDeletionTimeNil

`func (o *PuupeeDto) SetDeletionTimeNil(b bool)`

 SetDeletionTimeNil sets the value for DeletionTime to be an explicit nil

### UnsetDeletionTime
`func (o *PuupeeDto) UnsetDeletionTime()`

UnsetDeletionTime ensures that no value is present for DeletionTime, not even an explicit nil
### GetName

`func (o *PuupeeDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PuupeeDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PuupeeDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PuupeeDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PuupeeDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PuupeeDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTitle

`func (o *PuupeeDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PuupeeDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PuupeeDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PuupeeDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *PuupeeDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *PuupeeDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *PuupeeDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PuupeeDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PuupeeDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PuupeeDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PuupeeDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PuupeeDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetText

`func (o *PuupeeDto) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PuupeeDto) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PuupeeDto) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *PuupeeDto) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *PuupeeDto) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *PuupeeDto) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetContent

`func (o *PuupeeDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PuupeeDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PuupeeDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *PuupeeDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *PuupeeDto) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *PuupeeDto) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetFormat

`func (o *PuupeeDto) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PuupeeDto) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PuupeeDto) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *PuupeeDto) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *PuupeeDto) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *PuupeeDto) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetPassword

`func (o *PuupeeDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PuupeeDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PuupeeDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PuupeeDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *PuupeeDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *PuupeeDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetParentId

`func (o *PuupeeDto) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *PuupeeDto) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *PuupeeDto) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *PuupeeDto) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *PuupeeDto) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *PuupeeDto) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetKey

`func (o *PuupeeDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PuupeeDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PuupeeDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PuupeeDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *PuupeeDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *PuupeeDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetUrl

`func (o *PuupeeDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PuupeeDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PuupeeDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PuupeeDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *PuupeeDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *PuupeeDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetSize

`func (o *PuupeeDto) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PuupeeDto) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PuupeeDto) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *PuupeeDto) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMd5

`func (o *PuupeeDto) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *PuupeeDto) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *PuupeeDto) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *PuupeeDto) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### SetMd5Nil

`func (o *PuupeeDto) SetMd5Nil(b bool)`

 SetMd5Nil sets the value for Md5 to be an explicit nil

### UnsetMd5
`func (o *PuupeeDto) UnsetMd5()`

UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
### GetSliceMd5

`func (o *PuupeeDto) GetSliceMd5() string`

GetSliceMd5 returns the SliceMd5 field if non-nil, zero value otherwise.

### GetSliceMd5Ok

`func (o *PuupeeDto) GetSliceMd5Ok() (*string, bool)`

GetSliceMd5Ok returns a tuple with the SliceMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceMd5

`func (o *PuupeeDto) SetSliceMd5(v string)`

SetSliceMd5 sets SliceMd5 field to given value.

### HasSliceMd5

`func (o *PuupeeDto) HasSliceMd5() bool`

HasSliceMd5 returns a boolean if a field has been set.

### SetSliceMd5Nil

`func (o *PuupeeDto) SetSliceMd5Nil(b bool)`

 SetSliceMd5Nil sets the value for SliceMd5 to be an explicit nil

### UnsetSliceMd5
`func (o *PuupeeDto) UnsetSliceMd5()`

UnsetSliceMd5 ensures that no value is present for SliceMd5, not even an explicit nil
### GetRapidCode

`func (o *PuupeeDto) GetRapidCode() string`

GetRapidCode returns the RapidCode field if non-nil, zero value otherwise.

### GetRapidCodeOk

`func (o *PuupeeDto) GetRapidCodeOk() (*string, bool)`

GetRapidCodeOk returns a tuple with the RapidCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRapidCode

`func (o *PuupeeDto) SetRapidCode(v string)`

SetRapidCode sets RapidCode field to given value.

### HasRapidCode

`func (o *PuupeeDto) HasRapidCode() bool`

HasRapidCode returns a boolean if a field has been set.

### SetRapidCodeNil

`func (o *PuupeeDto) SetRapidCodeNil(b bool)`

 SetRapidCodeNil sets the value for RapidCode to be an explicit nil

### UnsetRapidCode
`func (o *PuupeeDto) UnsetRapidCode()`

UnsetRapidCode ensures that no value is present for RapidCode, not even an explicit nil
### GetContentType

`func (o *PuupeeDto) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PuupeeDto) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PuupeeDto) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *PuupeeDto) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *PuupeeDto) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *PuupeeDto) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetType

`func (o *PuupeeDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PuupeeDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PuupeeDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PuupeeDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *PuupeeDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PuupeeDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetDisplayStyle

`func (o *PuupeeDto) GetDisplayStyle() string`

GetDisplayStyle returns the DisplayStyle field if non-nil, zero value otherwise.

### GetDisplayStyleOk

`func (o *PuupeeDto) GetDisplayStyleOk() (*string, bool)`

GetDisplayStyleOk returns a tuple with the DisplayStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayStyle

`func (o *PuupeeDto) SetDisplayStyle(v string)`

SetDisplayStyle sets DisplayStyle field to given value.

### HasDisplayStyle

`func (o *PuupeeDto) HasDisplayStyle() bool`

HasDisplayStyle returns a boolean if a field has been set.

### SetDisplayStyleNil

`func (o *PuupeeDto) SetDisplayStyleNil(b bool)`

 SetDisplayStyleNil sets the value for DisplayStyle to be an explicit nil

### UnsetDisplayStyle
`func (o *PuupeeDto) UnsetDisplayStyle()`

UnsetDisplayStyle ensures that no value is present for DisplayStyle, not even an explicit nil
### GetExtension

`func (o *PuupeeDto) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *PuupeeDto) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *PuupeeDto) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *PuupeeDto) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### SetExtensionNil

`func (o *PuupeeDto) SetExtensionNil(b bool)`

 SetExtensionNil sets the value for Extension to be an explicit nil

### UnsetExtension
`func (o *PuupeeDto) UnsetExtension()`

UnsetExtension ensures that no value is present for Extension, not even an explicit nil
### GetStorageClass

`func (o *PuupeeDto) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *PuupeeDto) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *PuupeeDto) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *PuupeeDto) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### SetStorageClassNil

`func (o *PuupeeDto) SetStorageClassNil(b bool)`

 SetStorageClassNil sets the value for StorageClass to be an explicit nil

### UnsetStorageClass
`func (o *PuupeeDto) UnsetStorageClass()`

UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil
### GetStorageObjectCreatedAt

`func (o *PuupeeDto) GetStorageObjectCreatedAt() time.Time`

GetStorageObjectCreatedAt returns the StorageObjectCreatedAt field if non-nil, zero value otherwise.

### GetStorageObjectCreatedAtOk

`func (o *PuupeeDto) GetStorageObjectCreatedAtOk() (*time.Time, bool)`

GetStorageObjectCreatedAtOk returns a tuple with the StorageObjectCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageObjectCreatedAt

`func (o *PuupeeDto) SetStorageObjectCreatedAt(v time.Time)`

SetStorageObjectCreatedAt sets StorageObjectCreatedAt field to given value.

### HasStorageObjectCreatedAt

`func (o *PuupeeDto) HasStorageObjectCreatedAt() bool`

HasStorageObjectCreatedAt returns a boolean if a field has been set.

### SetStorageObjectCreatedAtNil

`func (o *PuupeeDto) SetStorageObjectCreatedAtNil(b bool)`

 SetStorageObjectCreatedAtNil sets the value for StorageObjectCreatedAt to be an explicit nil

### UnsetStorageObjectCreatedAt
`func (o *PuupeeDto) UnsetStorageObjectCreatedAt()`

UnsetStorageObjectCreatedAt ensures that no value is present for StorageObjectCreatedAt, not even an explicit nil
### GetStorageObjectUpdatedAt

`func (o *PuupeeDto) GetStorageObjectUpdatedAt() time.Time`

GetStorageObjectUpdatedAt returns the StorageObjectUpdatedAt field if non-nil, zero value otherwise.

### GetStorageObjectUpdatedAtOk

`func (o *PuupeeDto) GetStorageObjectUpdatedAtOk() (*time.Time, bool)`

GetStorageObjectUpdatedAtOk returns a tuple with the StorageObjectUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageObjectUpdatedAt

`func (o *PuupeeDto) SetStorageObjectUpdatedAt(v time.Time)`

SetStorageObjectUpdatedAt sets StorageObjectUpdatedAt field to given value.

### HasStorageObjectUpdatedAt

`func (o *PuupeeDto) HasStorageObjectUpdatedAt() bool`

HasStorageObjectUpdatedAt returns a boolean if a field has been set.

### SetStorageObjectUpdatedAtNil

`func (o *PuupeeDto) SetStorageObjectUpdatedAtNil(b bool)`

 SetStorageObjectUpdatedAtNil sets the value for StorageObjectUpdatedAt to be an explicit nil

### UnsetStorageObjectUpdatedAt
`func (o *PuupeeDto) UnsetStorageObjectUpdatedAt()`

UnsetStorageObjectUpdatedAt ensures that no value is present for StorageObjectUpdatedAt, not even an explicit nil
### GetStorageObjectId

`func (o *PuupeeDto) GetStorageObjectId() string`

GetStorageObjectId returns the StorageObjectId field if non-nil, zero value otherwise.

### GetStorageObjectIdOk

`func (o *PuupeeDto) GetStorageObjectIdOk() (*string, bool)`

GetStorageObjectIdOk returns a tuple with the StorageObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageObjectId

`func (o *PuupeeDto) SetStorageObjectId(v string)`

SetStorageObjectId sets StorageObjectId field to given value.

### HasStorageObjectId

`func (o *PuupeeDto) HasStorageObjectId() bool`

HasStorageObjectId returns a boolean if a field has been set.

### SetStorageObjectIdNil

`func (o *PuupeeDto) SetStorageObjectIdNil(b bool)`

 SetStorageObjectIdNil sets the value for StorageObjectId to be an explicit nil

### UnsetStorageObjectId
`func (o *PuupeeDto) UnsetStorageObjectId()`

UnsetStorageObjectId ensures that no value is present for StorageObjectId, not even an explicit nil
### GetStorageObject

`func (o *PuupeeDto) GetStorageObject() StorageObjectDto`

GetStorageObject returns the StorageObject field if non-nil, zero value otherwise.

### GetStorageObjectOk

`func (o *PuupeeDto) GetStorageObjectOk() (*StorageObjectDto, bool)`

GetStorageObjectOk returns a tuple with the StorageObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageObject

`func (o *PuupeeDto) SetStorageObject(v StorageObjectDto)`

SetStorageObject sets StorageObject field to given value.

### HasStorageObject

`func (o *PuupeeDto) HasStorageObject() bool`

HasStorageObject returns a boolean if a field has been set.

### GetThumb

`func (o *PuupeeDto) GetThumb() StorageObjectDto`

GetThumb returns the Thumb field if non-nil, zero value otherwise.

### GetThumbOk

`func (o *PuupeeDto) GetThumbOk() (*StorageObjectDto, bool)`

GetThumbOk returns a tuple with the Thumb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumb

`func (o *PuupeeDto) SetThumb(v StorageObjectDto)`

SetThumb sets Thumb field to given value.

### HasThumb

`func (o *PuupeeDto) HasThumb() bool`

HasThumb returns a boolean if a field has been set.

### GetPriority

`func (o *PuupeeDto) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PuupeeDto) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PuupeeDto) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PuupeeDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *PuupeeDto) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *PuupeeDto) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetDoneAt

`func (o *PuupeeDto) GetDoneAt() time.Time`

GetDoneAt returns the DoneAt field if non-nil, zero value otherwise.

### GetDoneAtOk

`func (o *PuupeeDto) GetDoneAtOk() (*time.Time, bool)`

GetDoneAtOk returns a tuple with the DoneAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoneAt

`func (o *PuupeeDto) SetDoneAt(v time.Time)`

SetDoneAt sets DoneAt field to given value.

### HasDoneAt

`func (o *PuupeeDto) HasDoneAt() bool`

HasDoneAt returns a boolean if a field has been set.

### SetDoneAtNil

`func (o *PuupeeDto) SetDoneAtNil(b bool)`

 SetDoneAtNil sets the value for DoneAt to be an explicit nil

### UnsetDoneAt
`func (o *PuupeeDto) UnsetDoneAt()`

UnsetDoneAt ensures that no value is present for DoneAt, not even an explicit nil
### GetIsDone

`func (o *PuupeeDto) GetIsDone() bool`

GetIsDone returns the IsDone field if non-nil, zero value otherwise.

### GetIsDoneOk

`func (o *PuupeeDto) GetIsDoneOk() (*bool, bool)`

GetIsDoneOk returns a tuple with the IsDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDone

`func (o *PuupeeDto) SetIsDone(v bool)`

SetIsDone sets IsDone field to given value.

### HasIsDone

`func (o *PuupeeDto) HasIsDone() bool`

HasIsDone returns a boolean if a field has been set.

### GetStartAt

`func (o *PuupeeDto) GetStartAt() time.Time`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *PuupeeDto) GetStartAtOk() (*time.Time, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *PuupeeDto) SetStartAt(v time.Time)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *PuupeeDto) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### SetStartAtNil

`func (o *PuupeeDto) SetStartAtNil(b bool)`

 SetStartAtNil sets the value for StartAt to be an explicit nil

### UnsetStartAt
`func (o *PuupeeDto) UnsetStartAt()`

UnsetStartAt ensures that no value is present for StartAt, not even an explicit nil
### GetEndAt

`func (o *PuupeeDto) GetEndAt() time.Time`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *PuupeeDto) GetEndAtOk() (*time.Time, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *PuupeeDto) SetEndAt(v time.Time)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *PuupeeDto) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### SetEndAtNil

`func (o *PuupeeDto) SetEndAtNil(b bool)`

 SetEndAtNil sets the value for EndAt to be an explicit nil

### UnsetEndAt
`func (o *PuupeeDto) UnsetEndAt()`

UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil
### GetNotifyAt

`func (o *PuupeeDto) GetNotifyAt() time.Time`

GetNotifyAt returns the NotifyAt field if non-nil, zero value otherwise.

### GetNotifyAtOk

`func (o *PuupeeDto) GetNotifyAtOk() (*time.Time, bool)`

GetNotifyAtOk returns a tuple with the NotifyAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyAt

`func (o *PuupeeDto) SetNotifyAt(v time.Time)`

SetNotifyAt sets NotifyAt field to given value.

### HasNotifyAt

`func (o *PuupeeDto) HasNotifyAt() bool`

HasNotifyAt returns a boolean if a field has been set.

### SetNotifyAtNil

`func (o *PuupeeDto) SetNotifyAtNil(b bool)`

 SetNotifyAtNil sets the value for NotifyAt to be an explicit nil

### UnsetNotifyAt
`func (o *PuupeeDto) UnsetNotifyAt()`

UnsetNotifyAt ensures that no value is present for NotifyAt, not even an explicit nil
### GetNotifyTimingType

`func (o *PuupeeDto) GetNotifyTimingType() string`

GetNotifyTimingType returns the NotifyTimingType field if non-nil, zero value otherwise.

### GetNotifyTimingTypeOk

`func (o *PuupeeDto) GetNotifyTimingTypeOk() (*string, bool)`

GetNotifyTimingTypeOk returns a tuple with the NotifyTimingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTimingType

`func (o *PuupeeDto) SetNotifyTimingType(v string)`

SetNotifyTimingType sets NotifyTimingType field to given value.

### HasNotifyTimingType

`func (o *PuupeeDto) HasNotifyTimingType() bool`

HasNotifyTimingType returns a boolean if a field has been set.

### SetNotifyTimingTypeNil

`func (o *PuupeeDto) SetNotifyTimingTypeNil(b bool)`

 SetNotifyTimingTypeNil sets the value for NotifyTimingType to be an explicit nil

### UnsetNotifyTimingType
`func (o *PuupeeDto) UnsetNotifyTimingType()`

UnsetNotifyTimingType ensures that no value is present for NotifyTimingType, not even an explicit nil
### GetNotifyTimingUnit

`func (o *PuupeeDto) GetNotifyTimingUnit() string`

GetNotifyTimingUnit returns the NotifyTimingUnit field if non-nil, zero value otherwise.

### GetNotifyTimingUnitOk

`func (o *PuupeeDto) GetNotifyTimingUnitOk() (*string, bool)`

GetNotifyTimingUnitOk returns a tuple with the NotifyTimingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTimingUnit

`func (o *PuupeeDto) SetNotifyTimingUnit(v string)`

SetNotifyTimingUnit sets NotifyTimingUnit field to given value.

### HasNotifyTimingUnit

`func (o *PuupeeDto) HasNotifyTimingUnit() bool`

HasNotifyTimingUnit returns a boolean if a field has been set.

### SetNotifyTimingUnitNil

`func (o *PuupeeDto) SetNotifyTimingUnitNil(b bool)`

 SetNotifyTimingUnitNil sets the value for NotifyTimingUnit to be an explicit nil

### UnsetNotifyTimingUnit
`func (o *PuupeeDto) UnsetNotifyTimingUnit()`

UnsetNotifyTimingUnit ensures that no value is present for NotifyTimingUnit, not even an explicit nil
### GetNotifyTimingValue

`func (o *PuupeeDto) GetNotifyTimingValue() int32`

GetNotifyTimingValue returns the NotifyTimingValue field if non-nil, zero value otherwise.

### GetNotifyTimingValueOk

`func (o *PuupeeDto) GetNotifyTimingValueOk() (*int32, bool)`

GetNotifyTimingValueOk returns a tuple with the NotifyTimingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTimingValue

`func (o *PuupeeDto) SetNotifyTimingValue(v int32)`

SetNotifyTimingValue sets NotifyTimingValue field to given value.

### HasNotifyTimingValue

`func (o *PuupeeDto) HasNotifyTimingValue() bool`

HasNotifyTimingValue returns a boolean if a field has been set.

### SetNotifyTimingValueNil

`func (o *PuupeeDto) SetNotifyTimingValueNil(b bool)`

 SetNotifyTimingValueNil sets the value for NotifyTimingValue to be an explicit nil

### UnsetNotifyTimingValue
`func (o *PuupeeDto) UnsetNotifyTimingValue()`

UnsetNotifyTimingValue ensures that no value is present for NotifyTimingValue, not even an explicit nil
### GetRepeat

`func (o *PuupeeDto) GetRepeat() string`

GetRepeat returns the Repeat field if non-nil, zero value otherwise.

### GetRepeatOk

`func (o *PuupeeDto) GetRepeatOk() (*string, bool)`

GetRepeatOk returns a tuple with the Repeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeat

`func (o *PuupeeDto) SetRepeat(v string)`

SetRepeat sets Repeat field to given value.

### HasRepeat

`func (o *PuupeeDto) HasRepeat() bool`

HasRepeat returns a boolean if a field has been set.

### SetRepeatNil

`func (o *PuupeeDto) SetRepeatNil(b bool)`

 SetRepeatNil sets the value for Repeat to be an explicit nil

### UnsetRepeat
`func (o *PuupeeDto) UnsetRepeat()`

UnsetRepeat ensures that no value is present for Repeat, not even an explicit nil
### GetRepeatOffAt

`func (o *PuupeeDto) GetRepeatOffAt() time.Time`

GetRepeatOffAt returns the RepeatOffAt field if non-nil, zero value otherwise.

### GetRepeatOffAtOk

`func (o *PuupeeDto) GetRepeatOffAtOk() (*time.Time, bool)`

GetRepeatOffAtOk returns a tuple with the RepeatOffAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatOffAt

`func (o *PuupeeDto) SetRepeatOffAt(v time.Time)`

SetRepeatOffAt sets RepeatOffAt field to given value.

### HasRepeatOffAt

`func (o *PuupeeDto) HasRepeatOffAt() bool`

HasRepeatOffAt returns a boolean if a field has been set.

### SetRepeatOffAtNil

`func (o *PuupeeDto) SetRepeatOffAtNil(b bool)`

 SetRepeatOffAtNil sets the value for RepeatOffAt to be an explicit nil

### UnsetRepeatOffAt
`func (o *PuupeeDto) UnsetRepeatOffAt()`

UnsetRepeatOffAt ensures that no value is present for RepeatOffAt, not even an explicit nil
### GetRepeatOffTimes

`func (o *PuupeeDto) GetRepeatOffTimes() int32`

GetRepeatOffTimes returns the RepeatOffTimes field if non-nil, zero value otherwise.

### GetRepeatOffTimesOk

`func (o *PuupeeDto) GetRepeatOffTimesOk() (*int32, bool)`

GetRepeatOffTimesOk returns a tuple with the RepeatOffTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatOffTimes

`func (o *PuupeeDto) SetRepeatOffTimes(v int32)`

SetRepeatOffTimes sets RepeatOffTimes field to given value.

### HasRepeatOffTimes

`func (o *PuupeeDto) HasRepeatOffTimes() bool`

HasRepeatOffTimes returns a boolean if a field has been set.

### SetRepeatOffTimesNil

`func (o *PuupeeDto) SetRepeatOffTimesNil(b bool)`

 SetRepeatOffTimesNil sets the value for RepeatOffTimes to be an explicit nil

### UnsetRepeatOffTimes
`func (o *PuupeeDto) UnsetRepeatOffTimes()`

UnsetRepeatOffTimes ensures that no value is present for RepeatOffTimes, not even an explicit nil
### GetRepetitions

`func (o *PuupeeDto) GetRepetitions() int32`

GetRepetitions returns the Repetitions field if non-nil, zero value otherwise.

### GetRepetitionsOk

`func (o *PuupeeDto) GetRepetitionsOk() (*int32, bool)`

GetRepetitionsOk returns a tuple with the Repetitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitions

`func (o *PuupeeDto) SetRepetitions(v int32)`

SetRepetitions sets Repetitions field to given value.

### HasRepetitions

`func (o *PuupeeDto) HasRepetitions() bool`

HasRepetitions returns a boolean if a field has been set.

### SetRepetitionsNil

`func (o *PuupeeDto) SetRepetitionsNil(b bool)`

 SetRepetitionsNil sets the value for Repetitions to be an explicit nil

### UnsetRepetitions
`func (o *PuupeeDto) UnsetRepetitions()`

UnsetRepetitions ensures that no value is present for Repetitions, not even an explicit nil
### GetSyncVersion

`func (o *PuupeeDto) GetSyncVersion() int64`

GetSyncVersion returns the SyncVersion field if non-nil, zero value otherwise.

### GetSyncVersionOk

`func (o *PuupeeDto) GetSyncVersionOk() (*int64, bool)`

GetSyncVersionOk returns a tuple with the SyncVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncVersion

`func (o *PuupeeDto) SetSyncVersion(v int64)`

SetSyncVersion sets SyncVersion field to given value.

### HasSyncVersion

`func (o *PuupeeDto) HasSyncVersion() bool`

HasSyncVersion returns a boolean if a field has been set.

### GetIsHidden

`func (o *PuupeeDto) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *PuupeeDto) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *PuupeeDto) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *PuupeeDto) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetTagging

`func (o *PuupeeDto) GetTagging() string`

GetTagging returns the Tagging field if non-nil, zero value otherwise.

### GetTaggingOk

`func (o *PuupeeDto) GetTaggingOk() (*string, bool)`

GetTaggingOk returns a tuple with the Tagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagging

`func (o *PuupeeDto) SetTagging(v string)`

SetTagging sets Tagging field to given value.

### HasTagging

`func (o *PuupeeDto) HasTagging() bool`

HasTagging returns a boolean if a field has been set.

### SetTaggingNil

`func (o *PuupeeDto) SetTaggingNil(b bool)`

 SetTaggingNil sets the value for Tagging to be an explicit nil

### UnsetTagging
`func (o *PuupeeDto) UnsetTagging()`

UnsetTagging ensures that no value is present for Tagging, not even an explicit nil
### GetLastModifierDeviceToken

`func (o *PuupeeDto) GetLastModifierDeviceToken() string`

GetLastModifierDeviceToken returns the LastModifierDeviceToken field if non-nil, zero value otherwise.

### GetLastModifierDeviceTokenOk

`func (o *PuupeeDto) GetLastModifierDeviceTokenOk() (*string, bool)`

GetLastModifierDeviceTokenOk returns a tuple with the LastModifierDeviceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierDeviceToken

`func (o *PuupeeDto) SetLastModifierDeviceToken(v string)`

SetLastModifierDeviceToken sets LastModifierDeviceToken field to given value.

### HasLastModifierDeviceToken

`func (o *PuupeeDto) HasLastModifierDeviceToken() bool`

HasLastModifierDeviceToken returns a boolean if a field has been set.

### SetLastModifierDeviceTokenNil

`func (o *PuupeeDto) SetLastModifierDeviceTokenNil(b bool)`

 SetLastModifierDeviceTokenNil sets the value for LastModifierDeviceToken to be an explicit nil

### UnsetLastModifierDeviceToken
`func (o *PuupeeDto) UnsetLastModifierDeviceToken()`

UnsetLastModifierDeviceToken ensures that no value is present for LastModifierDeviceToken, not even an explicit nil
### GetLastModifierDevice

`func (o *PuupeeDto) GetLastModifierDevice() string`

GetLastModifierDevice returns the LastModifierDevice field if non-nil, zero value otherwise.

### GetLastModifierDeviceOk

`func (o *PuupeeDto) GetLastModifierDeviceOk() (*string, bool)`

GetLastModifierDeviceOk returns a tuple with the LastModifierDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierDevice

`func (o *PuupeeDto) SetLastModifierDevice(v string)`

SetLastModifierDevice sets LastModifierDevice field to given value.

### HasLastModifierDevice

`func (o *PuupeeDto) HasLastModifierDevice() bool`

HasLastModifierDevice returns a boolean if a field has been set.

### SetLastModifierDeviceNil

`func (o *PuupeeDto) SetLastModifierDeviceNil(b bool)`

 SetLastModifierDeviceNil sets the value for LastModifierDevice to be an explicit nil

### UnsetLastModifierDevice
`func (o *PuupeeDto) UnsetLastModifierDevice()`

UnsetLastModifierDevice ensures that no value is present for LastModifierDevice, not even an explicit nil
### GetAppName

`func (o *PuupeeDto) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *PuupeeDto) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *PuupeeDto) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *PuupeeDto) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### SetAppNameNil

`func (o *PuupeeDto) SetAppNameNil(b bool)`

 SetAppNameNil sets the value for AppName to be an explicit nil

### UnsetAppName
`func (o *PuupeeDto) UnsetAppName()`

UnsetAppName ensures that no value is present for AppName, not even an explicit nil
### GetSortIndex

`func (o *PuupeeDto) GetSortIndex() int32`

GetSortIndex returns the SortIndex field if non-nil, zero value otherwise.

### GetSortIndexOk

`func (o *PuupeeDto) GetSortIndexOk() (*int32, bool)`

GetSortIndexOk returns a tuple with the SortIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortIndex

`func (o *PuupeeDto) SetSortIndex(v int32)`

SetSortIndex sets SortIndex field to given value.

### HasSortIndex

`func (o *PuupeeDto) HasSortIndex() bool`

HasSortIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateOrUpdatePuupeeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Title** | Pointer to **string** |  | [optional] 
**IsHidden** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Md5** | Pointer to **string** |  | [optional] 
**SliceMd5** | Pointer to **string** |  | [optional] 
**RapidCode** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DisplayStyle** | Pointer to **string** |  | [optional] 
**Extension** | Pointer to **string** |  | [optional] 
**StorageClass** | Pointer to **string** |  | [optional] 
**StorageObjectCreatedAt** | Pointer to **time.Time** |  | [optional] 
**StorageObjectUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**SyncVersion** | Pointer to **int64** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**DeletionTime** | Pointer to **time.Time** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**LastModificationTime** | Pointer to **time.Time** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**StartAt** | Pointer to **time.Time** |  | [optional] 
**EndAt** | Pointer to **time.Time** |  | [optional] 
**NotifyAt** | Pointer to **time.Time** |  | [optional] 
**NotifyTimingType** | Pointer to **string** |  | [optional] 
**NotifyTimingUnit** | Pointer to **string** |  | [optional] 
**NotifyTimingValue** | Pointer to **int32** |  | [optional] 
**Repeat** | Pointer to **string** |  | [optional] 
**RepeatOffAt** | Pointer to **time.Time** |  | [optional] 
**RepeatOffTimes** | Pointer to **int32** |  | [optional] 
**Repetitions** | Pointer to **int32** |  | [optional] 
**IsDone** | Pointer to **bool** |  | [optional] 
**DoneAt** | Pointer to **time.Time** |  | [optional] 
**CreatorId** | Pointer to **string** |  | [optional] 
**LastModifierId** | Pointer to **string** |  | [optional] 
**DeleterId** | Pointer to **string** |  | [optional] 
**Tagging** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**LastModifierDeviceToken** | Pointer to **string** |  | [optional] 
**LastModifierDevice** | Pointer to **string** |  | [optional] 
**App** | Pointer to **string** |  | [optional] 
**PushToUser** | Pointer to **bool** |  | [optional] 
**SortIndex** | Pointer to **int32** |  | [optional] 

## Methods

### NewCreateOrUpdatePuupeeDto

`func NewCreateOrUpdatePuupeeDto(id string, name string, ) *CreateOrUpdatePuupeeDto`

NewCreateOrUpdatePuupeeDto instantiates a new CreateOrUpdatePuupeeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdatePuupeeDtoWithDefaults

`func NewCreateOrUpdatePuupeeDtoWithDefaults() *CreateOrUpdatePuupeeDto`

NewCreateOrUpdatePuupeeDtoWithDefaults instantiates a new CreateOrUpdatePuupeeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateOrUpdatePuupeeDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateOrUpdatePuupeeDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateOrUpdatePuupeeDto) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CreateOrUpdatePuupeeDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrUpdatePuupeeDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrUpdatePuupeeDto) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *CreateOrUpdatePuupeeDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateOrUpdatePuupeeDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateOrUpdatePuupeeDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateOrUpdatePuupeeDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetIsHidden

`func (o *CreateOrUpdatePuupeeDto) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *CreateOrUpdatePuupeeDto) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *CreateOrUpdatePuupeeDto) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *CreateOrUpdatePuupeeDto) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetDescription

`func (o *CreateOrUpdatePuupeeDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrUpdatePuupeeDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrUpdatePuupeeDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOrUpdatePuupeeDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetText

`func (o *CreateOrUpdatePuupeeDto) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CreateOrUpdatePuupeeDto) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CreateOrUpdatePuupeeDto) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CreateOrUpdatePuupeeDto) HasText() bool`

HasText returns a boolean if a field has been set.

### GetContent

`func (o *CreateOrUpdatePuupeeDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateOrUpdatePuupeeDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateOrUpdatePuupeeDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CreateOrUpdatePuupeeDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetFormat

`func (o *CreateOrUpdatePuupeeDto) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateOrUpdatePuupeeDto) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateOrUpdatePuupeeDto) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CreateOrUpdatePuupeeDto) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetPassword

`func (o *CreateOrUpdatePuupeeDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateOrUpdatePuupeeDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateOrUpdatePuupeeDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateOrUpdatePuupeeDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetParentId

`func (o *CreateOrUpdatePuupeeDto) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateOrUpdatePuupeeDto) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateOrUpdatePuupeeDto) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateOrUpdatePuupeeDto) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetKey

`func (o *CreateOrUpdatePuupeeDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateOrUpdatePuupeeDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateOrUpdatePuupeeDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateOrUpdatePuupeeDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMd5

`func (o *CreateOrUpdatePuupeeDto) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *CreateOrUpdatePuupeeDto) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *CreateOrUpdatePuupeeDto) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *CreateOrUpdatePuupeeDto) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetSliceMd5

`func (o *CreateOrUpdatePuupeeDto) GetSliceMd5() string`

GetSliceMd5 returns the SliceMd5 field if non-nil, zero value otherwise.

### GetSliceMd5Ok

`func (o *CreateOrUpdatePuupeeDto) GetSliceMd5Ok() (*string, bool)`

GetSliceMd5Ok returns a tuple with the SliceMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceMd5

`func (o *CreateOrUpdatePuupeeDto) SetSliceMd5(v string)`

SetSliceMd5 sets SliceMd5 field to given value.

### HasSliceMd5

`func (o *CreateOrUpdatePuupeeDto) HasSliceMd5() bool`

HasSliceMd5 returns a boolean if a field has been set.

### GetRapidCode

`func (o *CreateOrUpdatePuupeeDto) GetRapidCode() string`

GetRapidCode returns the RapidCode field if non-nil, zero value otherwise.

### GetRapidCodeOk

`func (o *CreateOrUpdatePuupeeDto) GetRapidCodeOk() (*string, bool)`

GetRapidCodeOk returns a tuple with the RapidCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRapidCode

`func (o *CreateOrUpdatePuupeeDto) SetRapidCode(v string)`

SetRapidCode sets RapidCode field to given value.

### HasRapidCode

`func (o *CreateOrUpdatePuupeeDto) HasRapidCode() bool`

HasRapidCode returns a boolean if a field has been set.

### GetContentType

`func (o *CreateOrUpdatePuupeeDto) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CreateOrUpdatePuupeeDto) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CreateOrUpdatePuupeeDto) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *CreateOrUpdatePuupeeDto) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetType

`func (o *CreateOrUpdatePuupeeDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateOrUpdatePuupeeDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateOrUpdatePuupeeDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateOrUpdatePuupeeDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisplayStyle

`func (o *CreateOrUpdatePuupeeDto) GetDisplayStyle() string`

GetDisplayStyle returns the DisplayStyle field if non-nil, zero value otherwise.

### GetDisplayStyleOk

`func (o *CreateOrUpdatePuupeeDto) GetDisplayStyleOk() (*string, bool)`

GetDisplayStyleOk returns a tuple with the DisplayStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayStyle

`func (o *CreateOrUpdatePuupeeDto) SetDisplayStyle(v string)`

SetDisplayStyle sets DisplayStyle field to given value.

### HasDisplayStyle

`func (o *CreateOrUpdatePuupeeDto) HasDisplayStyle() bool`

HasDisplayStyle returns a boolean if a field has been set.

### GetExtension

`func (o *CreateOrUpdatePuupeeDto) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *CreateOrUpdatePuupeeDto) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *CreateOrUpdatePuupeeDto) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *CreateOrUpdatePuupeeDto) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetStorageClass

`func (o *CreateOrUpdatePuupeeDto) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *CreateOrUpdatePuupeeDto) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *CreateOrUpdatePuupeeDto) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *CreateOrUpdatePuupeeDto) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### GetStorageObjectCreatedAt

`func (o *CreateOrUpdatePuupeeDto) GetStorageObjectCreatedAt() time.Time`

GetStorageObjectCreatedAt returns the StorageObjectCreatedAt field if non-nil, zero value otherwise.

### GetStorageObjectCreatedAtOk

`func (o *CreateOrUpdatePuupeeDto) GetStorageObjectCreatedAtOk() (*time.Time, bool)`

GetStorageObjectCreatedAtOk returns a tuple with the StorageObjectCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageObjectCreatedAt

`func (o *CreateOrUpdatePuupeeDto) SetStorageObjectCreatedAt(v time.Time)`

SetStorageObjectCreatedAt sets StorageObjectCreatedAt field to given value.

### HasStorageObjectCreatedAt

`func (o *CreateOrUpdatePuupeeDto) HasStorageObjectCreatedAt() bool`

HasStorageObjectCreatedAt returns a boolean if a field has been set.

### GetStorageObjectUpdatedAt

`func (o *CreateOrUpdatePuupeeDto) GetStorageObjectUpdatedAt() time.Time`

GetStorageObjectUpdatedAt returns the StorageObjectUpdatedAt field if non-nil, zero value otherwise.

### GetStorageObjectUpdatedAtOk

`func (o *CreateOrUpdatePuupeeDto) GetStorageObjectUpdatedAtOk() (*time.Time, bool)`

GetStorageObjectUpdatedAtOk returns a tuple with the StorageObjectUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageObjectUpdatedAt

`func (o *CreateOrUpdatePuupeeDto) SetStorageObjectUpdatedAt(v time.Time)`

SetStorageObjectUpdatedAt sets StorageObjectUpdatedAt field to given value.

### HasStorageObjectUpdatedAt

`func (o *CreateOrUpdatePuupeeDto) HasStorageObjectUpdatedAt() bool`

HasStorageObjectUpdatedAt returns a boolean if a field has been set.

### GetSyncVersion

`func (o *CreateOrUpdatePuupeeDto) GetSyncVersion() int64`

GetSyncVersion returns the SyncVersion field if non-nil, zero value otherwise.

### GetSyncVersionOk

`func (o *CreateOrUpdatePuupeeDto) GetSyncVersionOk() (*int64, bool)`

GetSyncVersionOk returns a tuple with the SyncVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncVersion

`func (o *CreateOrUpdatePuupeeDto) SetSyncVersion(v int64)`

SetSyncVersion sets SyncVersion field to given value.

### HasSyncVersion

`func (o *CreateOrUpdatePuupeeDto) HasSyncVersion() bool`

HasSyncVersion returns a boolean if a field has been set.

### GetIsDeleted

`func (o *CreateOrUpdatePuupeeDto) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CreateOrUpdatePuupeeDto) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CreateOrUpdatePuupeeDto) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *CreateOrUpdatePuupeeDto) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeletionTime

`func (o *CreateOrUpdatePuupeeDto) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *CreateOrUpdatePuupeeDto) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *CreateOrUpdatePuupeeDto) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *CreateOrUpdatePuupeeDto) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### GetCreationTime

`func (o *CreateOrUpdatePuupeeDto) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CreateOrUpdatePuupeeDto) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CreateOrUpdatePuupeeDto) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *CreateOrUpdatePuupeeDto) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetLastModificationTime

`func (o *CreateOrUpdatePuupeeDto) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *CreateOrUpdatePuupeeDto) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *CreateOrUpdatePuupeeDto) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *CreateOrUpdatePuupeeDto) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### GetPriority

`func (o *CreateOrUpdatePuupeeDto) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateOrUpdatePuupeeDto) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateOrUpdatePuupeeDto) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CreateOrUpdatePuupeeDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStartAt

`func (o *CreateOrUpdatePuupeeDto) GetStartAt() time.Time`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *CreateOrUpdatePuupeeDto) GetStartAtOk() (*time.Time, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *CreateOrUpdatePuupeeDto) SetStartAt(v time.Time)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *CreateOrUpdatePuupeeDto) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetEndAt

`func (o *CreateOrUpdatePuupeeDto) GetEndAt() time.Time`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *CreateOrUpdatePuupeeDto) GetEndAtOk() (*time.Time, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *CreateOrUpdatePuupeeDto) SetEndAt(v time.Time)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *CreateOrUpdatePuupeeDto) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### GetNotifyAt

`func (o *CreateOrUpdatePuupeeDto) GetNotifyAt() time.Time`

GetNotifyAt returns the NotifyAt field if non-nil, zero value otherwise.

### GetNotifyAtOk

`func (o *CreateOrUpdatePuupeeDto) GetNotifyAtOk() (*time.Time, bool)`

GetNotifyAtOk returns a tuple with the NotifyAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyAt

`func (o *CreateOrUpdatePuupeeDto) SetNotifyAt(v time.Time)`

SetNotifyAt sets NotifyAt field to given value.

### HasNotifyAt

`func (o *CreateOrUpdatePuupeeDto) HasNotifyAt() bool`

HasNotifyAt returns a boolean if a field has been set.

### GetNotifyTimingType

`func (o *CreateOrUpdatePuupeeDto) GetNotifyTimingType() string`

GetNotifyTimingType returns the NotifyTimingType field if non-nil, zero value otherwise.

### GetNotifyTimingTypeOk

`func (o *CreateOrUpdatePuupeeDto) GetNotifyTimingTypeOk() (*string, bool)`

GetNotifyTimingTypeOk returns a tuple with the NotifyTimingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTimingType

`func (o *CreateOrUpdatePuupeeDto) SetNotifyTimingType(v string)`

SetNotifyTimingType sets NotifyTimingType field to given value.

### HasNotifyTimingType

`func (o *CreateOrUpdatePuupeeDto) HasNotifyTimingType() bool`

HasNotifyTimingType returns a boolean if a field has been set.

### GetNotifyTimingUnit

`func (o *CreateOrUpdatePuupeeDto) GetNotifyTimingUnit() string`

GetNotifyTimingUnit returns the NotifyTimingUnit field if non-nil, zero value otherwise.

### GetNotifyTimingUnitOk

`func (o *CreateOrUpdatePuupeeDto) GetNotifyTimingUnitOk() (*string, bool)`

GetNotifyTimingUnitOk returns a tuple with the NotifyTimingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTimingUnit

`func (o *CreateOrUpdatePuupeeDto) SetNotifyTimingUnit(v string)`

SetNotifyTimingUnit sets NotifyTimingUnit field to given value.

### HasNotifyTimingUnit

`func (o *CreateOrUpdatePuupeeDto) HasNotifyTimingUnit() bool`

HasNotifyTimingUnit returns a boolean if a field has been set.

### GetNotifyTimingValue

`func (o *CreateOrUpdatePuupeeDto) GetNotifyTimingValue() int32`

GetNotifyTimingValue returns the NotifyTimingValue field if non-nil, zero value otherwise.

### GetNotifyTimingValueOk

`func (o *CreateOrUpdatePuupeeDto) GetNotifyTimingValueOk() (*int32, bool)`

GetNotifyTimingValueOk returns a tuple with the NotifyTimingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTimingValue

`func (o *CreateOrUpdatePuupeeDto) SetNotifyTimingValue(v int32)`

SetNotifyTimingValue sets NotifyTimingValue field to given value.

### HasNotifyTimingValue

`func (o *CreateOrUpdatePuupeeDto) HasNotifyTimingValue() bool`

HasNotifyTimingValue returns a boolean if a field has been set.

### GetRepeat

`func (o *CreateOrUpdatePuupeeDto) GetRepeat() string`

GetRepeat returns the Repeat field if non-nil, zero value otherwise.

### GetRepeatOk

`func (o *CreateOrUpdatePuupeeDto) GetRepeatOk() (*string, bool)`

GetRepeatOk returns a tuple with the Repeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeat

`func (o *CreateOrUpdatePuupeeDto) SetRepeat(v string)`

SetRepeat sets Repeat field to given value.

### HasRepeat

`func (o *CreateOrUpdatePuupeeDto) HasRepeat() bool`

HasRepeat returns a boolean if a field has been set.

### GetRepeatOffAt

`func (o *CreateOrUpdatePuupeeDto) GetRepeatOffAt() time.Time`

GetRepeatOffAt returns the RepeatOffAt field if non-nil, zero value otherwise.

### GetRepeatOffAtOk

`func (o *CreateOrUpdatePuupeeDto) GetRepeatOffAtOk() (*time.Time, bool)`

GetRepeatOffAtOk returns a tuple with the RepeatOffAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatOffAt

`func (o *CreateOrUpdatePuupeeDto) SetRepeatOffAt(v time.Time)`

SetRepeatOffAt sets RepeatOffAt field to given value.

### HasRepeatOffAt

`func (o *CreateOrUpdatePuupeeDto) HasRepeatOffAt() bool`

HasRepeatOffAt returns a boolean if a field has been set.

### GetRepeatOffTimes

`func (o *CreateOrUpdatePuupeeDto) GetRepeatOffTimes() int32`

GetRepeatOffTimes returns the RepeatOffTimes field if non-nil, zero value otherwise.

### GetRepeatOffTimesOk

`func (o *CreateOrUpdatePuupeeDto) GetRepeatOffTimesOk() (*int32, bool)`

GetRepeatOffTimesOk returns a tuple with the RepeatOffTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatOffTimes

`func (o *CreateOrUpdatePuupeeDto) SetRepeatOffTimes(v int32)`

SetRepeatOffTimes sets RepeatOffTimes field to given value.

### HasRepeatOffTimes

`func (o *CreateOrUpdatePuupeeDto) HasRepeatOffTimes() bool`

HasRepeatOffTimes returns a boolean if a field has been set.

### GetRepetitions

`func (o *CreateOrUpdatePuupeeDto) GetRepetitions() int32`

GetRepetitions returns the Repetitions field if non-nil, zero value otherwise.

### GetRepetitionsOk

`func (o *CreateOrUpdatePuupeeDto) GetRepetitionsOk() (*int32, bool)`

GetRepetitionsOk returns a tuple with the Repetitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitions

`func (o *CreateOrUpdatePuupeeDto) SetRepetitions(v int32)`

SetRepetitions sets Repetitions field to given value.

### HasRepetitions

`func (o *CreateOrUpdatePuupeeDto) HasRepetitions() bool`

HasRepetitions returns a boolean if a field has been set.

### GetIsDone

`func (o *CreateOrUpdatePuupeeDto) GetIsDone() bool`

GetIsDone returns the IsDone field if non-nil, zero value otherwise.

### GetIsDoneOk

`func (o *CreateOrUpdatePuupeeDto) GetIsDoneOk() (*bool, bool)`

GetIsDoneOk returns a tuple with the IsDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDone

`func (o *CreateOrUpdatePuupeeDto) SetIsDone(v bool)`

SetIsDone sets IsDone field to given value.

### HasIsDone

`func (o *CreateOrUpdatePuupeeDto) HasIsDone() bool`

HasIsDone returns a boolean if a field has been set.

### GetDoneAt

`func (o *CreateOrUpdatePuupeeDto) GetDoneAt() time.Time`

GetDoneAt returns the DoneAt field if non-nil, zero value otherwise.

### GetDoneAtOk

`func (o *CreateOrUpdatePuupeeDto) GetDoneAtOk() (*time.Time, bool)`

GetDoneAtOk returns a tuple with the DoneAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoneAt

`func (o *CreateOrUpdatePuupeeDto) SetDoneAt(v time.Time)`

SetDoneAt sets DoneAt field to given value.

### HasDoneAt

`func (o *CreateOrUpdatePuupeeDto) HasDoneAt() bool`

HasDoneAt returns a boolean if a field has been set.

### GetCreatorId

`func (o *CreateOrUpdatePuupeeDto) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *CreateOrUpdatePuupeeDto) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *CreateOrUpdatePuupeeDto) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *CreateOrUpdatePuupeeDto) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetLastModifierId

`func (o *CreateOrUpdatePuupeeDto) GetLastModifierId() string`

GetLastModifierId returns the LastModifierId field if non-nil, zero value otherwise.

### GetLastModifierIdOk

`func (o *CreateOrUpdatePuupeeDto) GetLastModifierIdOk() (*string, bool)`

GetLastModifierIdOk returns a tuple with the LastModifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierId

`func (o *CreateOrUpdatePuupeeDto) SetLastModifierId(v string)`

SetLastModifierId sets LastModifierId field to given value.

### HasLastModifierId

`func (o *CreateOrUpdatePuupeeDto) HasLastModifierId() bool`

HasLastModifierId returns a boolean if a field has been set.

### GetDeleterId

`func (o *CreateOrUpdatePuupeeDto) GetDeleterId() string`

GetDeleterId returns the DeleterId field if non-nil, zero value otherwise.

### GetDeleterIdOk

`func (o *CreateOrUpdatePuupeeDto) GetDeleterIdOk() (*string, bool)`

GetDeleterIdOk returns a tuple with the DeleterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleterId

`func (o *CreateOrUpdatePuupeeDto) SetDeleterId(v string)`

SetDeleterId sets DeleterId field to given value.

### HasDeleterId

`func (o *CreateOrUpdatePuupeeDto) HasDeleterId() bool`

HasDeleterId returns a boolean if a field has been set.

### GetTagging

`func (o *CreateOrUpdatePuupeeDto) GetTagging() string`

GetTagging returns the Tagging field if non-nil, zero value otherwise.

### GetTaggingOk

`func (o *CreateOrUpdatePuupeeDto) GetTaggingOk() (*string, bool)`

GetTaggingOk returns a tuple with the Tagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagging

`func (o *CreateOrUpdatePuupeeDto) SetTagging(v string)`

SetTagging sets Tagging field to given value.

### HasTagging

`func (o *CreateOrUpdatePuupeeDto) HasTagging() bool`

HasTagging returns a boolean if a field has been set.

### GetUrl

`func (o *CreateOrUpdatePuupeeDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateOrUpdatePuupeeDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateOrUpdatePuupeeDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateOrUpdatePuupeeDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSize

`func (o *CreateOrUpdatePuupeeDto) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateOrUpdatePuupeeDto) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateOrUpdatePuupeeDto) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateOrUpdatePuupeeDto) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetLastModifierDeviceToken

`func (o *CreateOrUpdatePuupeeDto) GetLastModifierDeviceToken() string`

GetLastModifierDeviceToken returns the LastModifierDeviceToken field if non-nil, zero value otherwise.

### GetLastModifierDeviceTokenOk

`func (o *CreateOrUpdatePuupeeDto) GetLastModifierDeviceTokenOk() (*string, bool)`

GetLastModifierDeviceTokenOk returns a tuple with the LastModifierDeviceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierDeviceToken

`func (o *CreateOrUpdatePuupeeDto) SetLastModifierDeviceToken(v string)`

SetLastModifierDeviceToken sets LastModifierDeviceToken field to given value.

### HasLastModifierDeviceToken

`func (o *CreateOrUpdatePuupeeDto) HasLastModifierDeviceToken() bool`

HasLastModifierDeviceToken returns a boolean if a field has been set.

### GetLastModifierDevice

`func (o *CreateOrUpdatePuupeeDto) GetLastModifierDevice() string`

GetLastModifierDevice returns the LastModifierDevice field if non-nil, zero value otherwise.

### GetLastModifierDeviceOk

`func (o *CreateOrUpdatePuupeeDto) GetLastModifierDeviceOk() (*string, bool)`

GetLastModifierDeviceOk returns a tuple with the LastModifierDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierDevice

`func (o *CreateOrUpdatePuupeeDto) SetLastModifierDevice(v string)`

SetLastModifierDevice sets LastModifierDevice field to given value.

### HasLastModifierDevice

`func (o *CreateOrUpdatePuupeeDto) HasLastModifierDevice() bool`

HasLastModifierDevice returns a boolean if a field has been set.

### GetApp

`func (o *CreateOrUpdatePuupeeDto) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *CreateOrUpdatePuupeeDto) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *CreateOrUpdatePuupeeDto) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *CreateOrUpdatePuupeeDto) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetPushToUser

`func (o *CreateOrUpdatePuupeeDto) GetPushToUser() bool`

GetPushToUser returns the PushToUser field if non-nil, zero value otherwise.

### GetPushToUserOk

`func (o *CreateOrUpdatePuupeeDto) GetPushToUserOk() (*bool, bool)`

GetPushToUserOk returns a tuple with the PushToUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushToUser

`func (o *CreateOrUpdatePuupeeDto) SetPushToUser(v bool)`

SetPushToUser sets PushToUser field to given value.

### HasPushToUser

`func (o *CreateOrUpdatePuupeeDto) HasPushToUser() bool`

HasPushToUser returns a boolean if a field has been set.

### GetSortIndex

`func (o *CreateOrUpdatePuupeeDto) GetSortIndex() int32`

GetSortIndex returns the SortIndex field if non-nil, zero value otherwise.

### GetSortIndexOk

`func (o *CreateOrUpdatePuupeeDto) GetSortIndexOk() (*int32, bool)`

GetSortIndexOk returns a tuple with the SortIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortIndex

`func (o *CreateOrUpdatePuupeeDto) SetSortIndex(v int32)`

SetSortIndex sets SortIndex field to given value.

### HasSortIndex

`func (o *CreateOrUpdatePuupeeDto) HasSortIndex() bool`

HasSortIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



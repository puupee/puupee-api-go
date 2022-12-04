# CreateUpdatePuupeeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Title** | Pointer to **string** |  | [optional] 
**IsHidden** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
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

## Methods

### NewCreateUpdatePuupeeDto

`func NewCreateUpdatePuupeeDto(id string, name string, ) *CreateUpdatePuupeeDto`

NewCreateUpdatePuupeeDto instantiates a new CreateUpdatePuupeeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdatePuupeeDtoWithDefaults

`func NewCreateUpdatePuupeeDtoWithDefaults() *CreateUpdatePuupeeDto`

NewCreateUpdatePuupeeDtoWithDefaults instantiates a new CreateUpdatePuupeeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateUpdatePuupeeDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateUpdatePuupeeDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateUpdatePuupeeDto) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CreateUpdatePuupeeDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUpdatePuupeeDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUpdatePuupeeDto) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *CreateUpdatePuupeeDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateUpdatePuupeeDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateUpdatePuupeeDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateUpdatePuupeeDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetIsHidden

`func (o *CreateUpdatePuupeeDto) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *CreateUpdatePuupeeDto) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *CreateUpdatePuupeeDto) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *CreateUpdatePuupeeDto) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetDescription

`func (o *CreateUpdatePuupeeDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateUpdatePuupeeDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateUpdatePuupeeDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateUpdatePuupeeDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetText

`func (o *CreateUpdatePuupeeDto) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CreateUpdatePuupeeDto) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CreateUpdatePuupeeDto) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CreateUpdatePuupeeDto) HasText() bool`

HasText returns a boolean if a field has been set.

### GetFormat

`func (o *CreateUpdatePuupeeDto) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateUpdatePuupeeDto) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateUpdatePuupeeDto) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CreateUpdatePuupeeDto) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetPassword

`func (o *CreateUpdatePuupeeDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateUpdatePuupeeDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateUpdatePuupeeDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateUpdatePuupeeDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetParentId

`func (o *CreateUpdatePuupeeDto) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateUpdatePuupeeDto) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateUpdatePuupeeDto) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateUpdatePuupeeDto) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetKey

`func (o *CreateUpdatePuupeeDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateUpdatePuupeeDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateUpdatePuupeeDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateUpdatePuupeeDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMd5

`func (o *CreateUpdatePuupeeDto) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *CreateUpdatePuupeeDto) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *CreateUpdatePuupeeDto) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *CreateUpdatePuupeeDto) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetSliceMd5

`func (o *CreateUpdatePuupeeDto) GetSliceMd5() string`

GetSliceMd5 returns the SliceMd5 field if non-nil, zero value otherwise.

### GetSliceMd5Ok

`func (o *CreateUpdatePuupeeDto) GetSliceMd5Ok() (*string, bool)`

GetSliceMd5Ok returns a tuple with the SliceMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceMd5

`func (o *CreateUpdatePuupeeDto) SetSliceMd5(v string)`

SetSliceMd5 sets SliceMd5 field to given value.

### HasSliceMd5

`func (o *CreateUpdatePuupeeDto) HasSliceMd5() bool`

HasSliceMd5 returns a boolean if a field has been set.

### GetRapidCode

`func (o *CreateUpdatePuupeeDto) GetRapidCode() string`

GetRapidCode returns the RapidCode field if non-nil, zero value otherwise.

### GetRapidCodeOk

`func (o *CreateUpdatePuupeeDto) GetRapidCodeOk() (*string, bool)`

GetRapidCodeOk returns a tuple with the RapidCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRapidCode

`func (o *CreateUpdatePuupeeDto) SetRapidCode(v string)`

SetRapidCode sets RapidCode field to given value.

### HasRapidCode

`func (o *CreateUpdatePuupeeDto) HasRapidCode() bool`

HasRapidCode returns a boolean if a field has been set.

### GetContentType

`func (o *CreateUpdatePuupeeDto) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CreateUpdatePuupeeDto) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CreateUpdatePuupeeDto) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *CreateUpdatePuupeeDto) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetType

`func (o *CreateUpdatePuupeeDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateUpdatePuupeeDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateUpdatePuupeeDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateUpdatePuupeeDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisplayStyle

`func (o *CreateUpdatePuupeeDto) GetDisplayStyle() string`

GetDisplayStyle returns the DisplayStyle field if non-nil, zero value otherwise.

### GetDisplayStyleOk

`func (o *CreateUpdatePuupeeDto) GetDisplayStyleOk() (*string, bool)`

GetDisplayStyleOk returns a tuple with the DisplayStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayStyle

`func (o *CreateUpdatePuupeeDto) SetDisplayStyle(v string)`

SetDisplayStyle sets DisplayStyle field to given value.

### HasDisplayStyle

`func (o *CreateUpdatePuupeeDto) HasDisplayStyle() bool`

HasDisplayStyle returns a boolean if a field has been set.

### GetExtension

`func (o *CreateUpdatePuupeeDto) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *CreateUpdatePuupeeDto) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *CreateUpdatePuupeeDto) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *CreateUpdatePuupeeDto) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetStorageClass

`func (o *CreateUpdatePuupeeDto) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *CreateUpdatePuupeeDto) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *CreateUpdatePuupeeDto) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *CreateUpdatePuupeeDto) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### GetStorageObjectCreatedAt

`func (o *CreateUpdatePuupeeDto) GetStorageObjectCreatedAt() time.Time`

GetStorageObjectCreatedAt returns the StorageObjectCreatedAt field if non-nil, zero value otherwise.

### GetStorageObjectCreatedAtOk

`func (o *CreateUpdatePuupeeDto) GetStorageObjectCreatedAtOk() (*time.Time, bool)`

GetStorageObjectCreatedAtOk returns a tuple with the StorageObjectCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageObjectCreatedAt

`func (o *CreateUpdatePuupeeDto) SetStorageObjectCreatedAt(v time.Time)`

SetStorageObjectCreatedAt sets StorageObjectCreatedAt field to given value.

### HasStorageObjectCreatedAt

`func (o *CreateUpdatePuupeeDto) HasStorageObjectCreatedAt() bool`

HasStorageObjectCreatedAt returns a boolean if a field has been set.

### GetStorageObjectUpdatedAt

`func (o *CreateUpdatePuupeeDto) GetStorageObjectUpdatedAt() time.Time`

GetStorageObjectUpdatedAt returns the StorageObjectUpdatedAt field if non-nil, zero value otherwise.

### GetStorageObjectUpdatedAtOk

`func (o *CreateUpdatePuupeeDto) GetStorageObjectUpdatedAtOk() (*time.Time, bool)`

GetStorageObjectUpdatedAtOk returns a tuple with the StorageObjectUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageObjectUpdatedAt

`func (o *CreateUpdatePuupeeDto) SetStorageObjectUpdatedAt(v time.Time)`

SetStorageObjectUpdatedAt sets StorageObjectUpdatedAt field to given value.

### HasStorageObjectUpdatedAt

`func (o *CreateUpdatePuupeeDto) HasStorageObjectUpdatedAt() bool`

HasStorageObjectUpdatedAt returns a boolean if a field has been set.

### GetSyncVersion

`func (o *CreateUpdatePuupeeDto) GetSyncVersion() int64`

GetSyncVersion returns the SyncVersion field if non-nil, zero value otherwise.

### GetSyncVersionOk

`func (o *CreateUpdatePuupeeDto) GetSyncVersionOk() (*int64, bool)`

GetSyncVersionOk returns a tuple with the SyncVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncVersion

`func (o *CreateUpdatePuupeeDto) SetSyncVersion(v int64)`

SetSyncVersion sets SyncVersion field to given value.

### HasSyncVersion

`func (o *CreateUpdatePuupeeDto) HasSyncVersion() bool`

HasSyncVersion returns a boolean if a field has been set.

### GetIsDeleted

`func (o *CreateUpdatePuupeeDto) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CreateUpdatePuupeeDto) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CreateUpdatePuupeeDto) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *CreateUpdatePuupeeDto) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeletionTime

`func (o *CreateUpdatePuupeeDto) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *CreateUpdatePuupeeDto) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *CreateUpdatePuupeeDto) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *CreateUpdatePuupeeDto) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### GetCreationTime

`func (o *CreateUpdatePuupeeDto) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CreateUpdatePuupeeDto) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CreateUpdatePuupeeDto) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *CreateUpdatePuupeeDto) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetLastModificationTime

`func (o *CreateUpdatePuupeeDto) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *CreateUpdatePuupeeDto) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *CreateUpdatePuupeeDto) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *CreateUpdatePuupeeDto) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### GetPriority

`func (o *CreateUpdatePuupeeDto) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateUpdatePuupeeDto) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateUpdatePuupeeDto) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CreateUpdatePuupeeDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStartAt

`func (o *CreateUpdatePuupeeDto) GetStartAt() time.Time`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *CreateUpdatePuupeeDto) GetStartAtOk() (*time.Time, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *CreateUpdatePuupeeDto) SetStartAt(v time.Time)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *CreateUpdatePuupeeDto) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetEndAt

`func (o *CreateUpdatePuupeeDto) GetEndAt() time.Time`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *CreateUpdatePuupeeDto) GetEndAtOk() (*time.Time, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *CreateUpdatePuupeeDto) SetEndAt(v time.Time)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *CreateUpdatePuupeeDto) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### GetNotifyAt

`func (o *CreateUpdatePuupeeDto) GetNotifyAt() time.Time`

GetNotifyAt returns the NotifyAt field if non-nil, zero value otherwise.

### GetNotifyAtOk

`func (o *CreateUpdatePuupeeDto) GetNotifyAtOk() (*time.Time, bool)`

GetNotifyAtOk returns a tuple with the NotifyAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyAt

`func (o *CreateUpdatePuupeeDto) SetNotifyAt(v time.Time)`

SetNotifyAt sets NotifyAt field to given value.

### HasNotifyAt

`func (o *CreateUpdatePuupeeDto) HasNotifyAt() bool`

HasNotifyAt returns a boolean if a field has been set.

### GetNotifyTimingType

`func (o *CreateUpdatePuupeeDto) GetNotifyTimingType() string`

GetNotifyTimingType returns the NotifyTimingType field if non-nil, zero value otherwise.

### GetNotifyTimingTypeOk

`func (o *CreateUpdatePuupeeDto) GetNotifyTimingTypeOk() (*string, bool)`

GetNotifyTimingTypeOk returns a tuple with the NotifyTimingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTimingType

`func (o *CreateUpdatePuupeeDto) SetNotifyTimingType(v string)`

SetNotifyTimingType sets NotifyTimingType field to given value.

### HasNotifyTimingType

`func (o *CreateUpdatePuupeeDto) HasNotifyTimingType() bool`

HasNotifyTimingType returns a boolean if a field has been set.

### GetNotifyTimingUnit

`func (o *CreateUpdatePuupeeDto) GetNotifyTimingUnit() string`

GetNotifyTimingUnit returns the NotifyTimingUnit field if non-nil, zero value otherwise.

### GetNotifyTimingUnitOk

`func (o *CreateUpdatePuupeeDto) GetNotifyTimingUnitOk() (*string, bool)`

GetNotifyTimingUnitOk returns a tuple with the NotifyTimingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTimingUnit

`func (o *CreateUpdatePuupeeDto) SetNotifyTimingUnit(v string)`

SetNotifyTimingUnit sets NotifyTimingUnit field to given value.

### HasNotifyTimingUnit

`func (o *CreateUpdatePuupeeDto) HasNotifyTimingUnit() bool`

HasNotifyTimingUnit returns a boolean if a field has been set.

### GetNotifyTimingValue

`func (o *CreateUpdatePuupeeDto) GetNotifyTimingValue() int32`

GetNotifyTimingValue returns the NotifyTimingValue field if non-nil, zero value otherwise.

### GetNotifyTimingValueOk

`func (o *CreateUpdatePuupeeDto) GetNotifyTimingValueOk() (*int32, bool)`

GetNotifyTimingValueOk returns a tuple with the NotifyTimingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTimingValue

`func (o *CreateUpdatePuupeeDto) SetNotifyTimingValue(v int32)`

SetNotifyTimingValue sets NotifyTimingValue field to given value.

### HasNotifyTimingValue

`func (o *CreateUpdatePuupeeDto) HasNotifyTimingValue() bool`

HasNotifyTimingValue returns a boolean if a field has been set.

### GetRepeat

`func (o *CreateUpdatePuupeeDto) GetRepeat() string`

GetRepeat returns the Repeat field if non-nil, zero value otherwise.

### GetRepeatOk

`func (o *CreateUpdatePuupeeDto) GetRepeatOk() (*string, bool)`

GetRepeatOk returns a tuple with the Repeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeat

`func (o *CreateUpdatePuupeeDto) SetRepeat(v string)`

SetRepeat sets Repeat field to given value.

### HasRepeat

`func (o *CreateUpdatePuupeeDto) HasRepeat() bool`

HasRepeat returns a boolean if a field has been set.

### GetRepeatOffAt

`func (o *CreateUpdatePuupeeDto) GetRepeatOffAt() time.Time`

GetRepeatOffAt returns the RepeatOffAt field if non-nil, zero value otherwise.

### GetRepeatOffAtOk

`func (o *CreateUpdatePuupeeDto) GetRepeatOffAtOk() (*time.Time, bool)`

GetRepeatOffAtOk returns a tuple with the RepeatOffAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatOffAt

`func (o *CreateUpdatePuupeeDto) SetRepeatOffAt(v time.Time)`

SetRepeatOffAt sets RepeatOffAt field to given value.

### HasRepeatOffAt

`func (o *CreateUpdatePuupeeDto) HasRepeatOffAt() bool`

HasRepeatOffAt returns a boolean if a field has been set.

### GetRepeatOffTimes

`func (o *CreateUpdatePuupeeDto) GetRepeatOffTimes() int32`

GetRepeatOffTimes returns the RepeatOffTimes field if non-nil, zero value otherwise.

### GetRepeatOffTimesOk

`func (o *CreateUpdatePuupeeDto) GetRepeatOffTimesOk() (*int32, bool)`

GetRepeatOffTimesOk returns a tuple with the RepeatOffTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatOffTimes

`func (o *CreateUpdatePuupeeDto) SetRepeatOffTimes(v int32)`

SetRepeatOffTimes sets RepeatOffTimes field to given value.

### HasRepeatOffTimes

`func (o *CreateUpdatePuupeeDto) HasRepeatOffTimes() bool`

HasRepeatOffTimes returns a boolean if a field has been set.

### GetIsDone

`func (o *CreateUpdatePuupeeDto) GetIsDone() bool`

GetIsDone returns the IsDone field if non-nil, zero value otherwise.

### GetIsDoneOk

`func (o *CreateUpdatePuupeeDto) GetIsDoneOk() (*bool, bool)`

GetIsDoneOk returns a tuple with the IsDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDone

`func (o *CreateUpdatePuupeeDto) SetIsDone(v bool)`

SetIsDone sets IsDone field to given value.

### HasIsDone

`func (o *CreateUpdatePuupeeDto) HasIsDone() bool`

HasIsDone returns a boolean if a field has been set.

### GetDoneAt

`func (o *CreateUpdatePuupeeDto) GetDoneAt() time.Time`

GetDoneAt returns the DoneAt field if non-nil, zero value otherwise.

### GetDoneAtOk

`func (o *CreateUpdatePuupeeDto) GetDoneAtOk() (*time.Time, bool)`

GetDoneAtOk returns a tuple with the DoneAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoneAt

`func (o *CreateUpdatePuupeeDto) SetDoneAt(v time.Time)`

SetDoneAt sets DoneAt field to given value.

### HasDoneAt

`func (o *CreateUpdatePuupeeDto) HasDoneAt() bool`

HasDoneAt returns a boolean if a field has been set.

### GetCreatorId

`func (o *CreateUpdatePuupeeDto) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *CreateUpdatePuupeeDto) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *CreateUpdatePuupeeDto) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *CreateUpdatePuupeeDto) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetLastModifierId

`func (o *CreateUpdatePuupeeDto) GetLastModifierId() string`

GetLastModifierId returns the LastModifierId field if non-nil, zero value otherwise.

### GetLastModifierIdOk

`func (o *CreateUpdatePuupeeDto) GetLastModifierIdOk() (*string, bool)`

GetLastModifierIdOk returns a tuple with the LastModifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierId

`func (o *CreateUpdatePuupeeDto) SetLastModifierId(v string)`

SetLastModifierId sets LastModifierId field to given value.

### HasLastModifierId

`func (o *CreateUpdatePuupeeDto) HasLastModifierId() bool`

HasLastModifierId returns a boolean if a field has been set.

### GetDeleterId

`func (o *CreateUpdatePuupeeDto) GetDeleterId() string`

GetDeleterId returns the DeleterId field if non-nil, zero value otherwise.

### GetDeleterIdOk

`func (o *CreateUpdatePuupeeDto) GetDeleterIdOk() (*string, bool)`

GetDeleterIdOk returns a tuple with the DeleterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleterId

`func (o *CreateUpdatePuupeeDto) SetDeleterId(v string)`

SetDeleterId sets DeleterId field to given value.

### HasDeleterId

`func (o *CreateUpdatePuupeeDto) HasDeleterId() bool`

HasDeleterId returns a boolean if a field has been set.

### GetTagging

`func (o *CreateUpdatePuupeeDto) GetTagging() string`

GetTagging returns the Tagging field if non-nil, zero value otherwise.

### GetTaggingOk

`func (o *CreateUpdatePuupeeDto) GetTaggingOk() (*string, bool)`

GetTaggingOk returns a tuple with the Tagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagging

`func (o *CreateUpdatePuupeeDto) SetTagging(v string)`

SetTagging sets Tagging field to given value.

### HasTagging

`func (o *CreateUpdatePuupeeDto) HasTagging() bool`

HasTagging returns a boolean if a field has been set.

### GetUrl

`func (o *CreateUpdatePuupeeDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateUpdatePuupeeDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateUpdatePuupeeDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateUpdatePuupeeDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSize

`func (o *CreateUpdatePuupeeDto) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateUpdatePuupeeDto) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateUpdatePuupeeDto) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateUpdatePuupeeDto) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetLastModifierDeviceToken

`func (o *CreateUpdatePuupeeDto) GetLastModifierDeviceToken() string`

GetLastModifierDeviceToken returns the LastModifierDeviceToken field if non-nil, zero value otherwise.

### GetLastModifierDeviceTokenOk

`func (o *CreateUpdatePuupeeDto) GetLastModifierDeviceTokenOk() (*string, bool)`

GetLastModifierDeviceTokenOk returns a tuple with the LastModifierDeviceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierDeviceToken

`func (o *CreateUpdatePuupeeDto) SetLastModifierDeviceToken(v string)`

SetLastModifierDeviceToken sets LastModifierDeviceToken field to given value.

### HasLastModifierDeviceToken

`func (o *CreateUpdatePuupeeDto) HasLastModifierDeviceToken() bool`

HasLastModifierDeviceToken returns a boolean if a field has been set.

### GetLastModifierDevice

`func (o *CreateUpdatePuupeeDto) GetLastModifierDevice() string`

GetLastModifierDevice returns the LastModifierDevice field if non-nil, zero value otherwise.

### GetLastModifierDeviceOk

`func (o *CreateUpdatePuupeeDto) GetLastModifierDeviceOk() (*string, bool)`

GetLastModifierDeviceOk returns a tuple with the LastModifierDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierDevice

`func (o *CreateUpdatePuupeeDto) SetLastModifierDevice(v string)`

SetLastModifierDevice sets LastModifierDevice field to given value.

### HasLastModifierDevice

`func (o *CreateUpdatePuupeeDto) HasLastModifierDevice() bool`

HasLastModifierDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



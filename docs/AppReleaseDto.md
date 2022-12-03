# AppReleaseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**CreatorId** | Pointer to **string** |  | [optional] 
**LastModificationTime** | Pointer to **time.Time** |  | [optional] 
**LastModifierId** | Pointer to **string** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**DeleterId** | Pointer to **string** |  | [optional] 
**DeletionTime** | Pointer to **time.Time** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**RapidCode** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Md5** | Pointer to **string** |  | [optional] 
**SliceMd5** | Pointer to **string** |  | [optional] 
**DownloadUrl** | Pointer to **string** |  | [optional] 
**ProductType** | Pointer to **string** |  | [optional] 
**IsForceUpdate** | Pointer to **bool** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**Channel** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 

## Methods

### NewAppReleaseDto

`func NewAppReleaseDto() *AppReleaseDto`

NewAppReleaseDto instantiates a new AppReleaseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppReleaseDtoWithDefaults

`func NewAppReleaseDtoWithDefaults() *AppReleaseDto`

NewAppReleaseDtoWithDefaults instantiates a new AppReleaseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppReleaseDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppReleaseDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppReleaseDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppReleaseDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreationTime

`func (o *AppReleaseDto) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AppReleaseDto) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AppReleaseDto) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *AppReleaseDto) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *AppReleaseDto) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *AppReleaseDto) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *AppReleaseDto) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *AppReleaseDto) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetLastModificationTime

`func (o *AppReleaseDto) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *AppReleaseDto) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *AppReleaseDto) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *AppReleaseDto) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### GetLastModifierId

`func (o *AppReleaseDto) GetLastModifierId() string`

GetLastModifierId returns the LastModifierId field if non-nil, zero value otherwise.

### GetLastModifierIdOk

`func (o *AppReleaseDto) GetLastModifierIdOk() (*string, bool)`

GetLastModifierIdOk returns a tuple with the LastModifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierId

`func (o *AppReleaseDto) SetLastModifierId(v string)`

SetLastModifierId sets LastModifierId field to given value.

### HasLastModifierId

`func (o *AppReleaseDto) HasLastModifierId() bool`

HasLastModifierId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *AppReleaseDto) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *AppReleaseDto) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *AppReleaseDto) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *AppReleaseDto) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeleterId

`func (o *AppReleaseDto) GetDeleterId() string`

GetDeleterId returns the DeleterId field if non-nil, zero value otherwise.

### GetDeleterIdOk

`func (o *AppReleaseDto) GetDeleterIdOk() (*string, bool)`

GetDeleterIdOk returns a tuple with the DeleterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleterId

`func (o *AppReleaseDto) SetDeleterId(v string)`

SetDeleterId sets DeleterId field to given value.

### HasDeleterId

`func (o *AppReleaseDto) HasDeleterId() bool`

HasDeleterId returns a boolean if a field has been set.

### GetDeletionTime

`func (o *AppReleaseDto) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *AppReleaseDto) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *AppReleaseDto) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *AppReleaseDto) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### GetVersion

`func (o *AppReleaseDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AppReleaseDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AppReleaseDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AppReleaseDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetNotes

`func (o *AppReleaseDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AppReleaseDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AppReleaseDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AppReleaseDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPlatform

`func (o *AppReleaseDto) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AppReleaseDto) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AppReleaseDto) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AppReleaseDto) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetKey

`func (o *AppReleaseDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AppReleaseDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AppReleaseDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AppReleaseDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetRapidCode

`func (o *AppReleaseDto) GetRapidCode() string`

GetRapidCode returns the RapidCode field if non-nil, zero value otherwise.

### GetRapidCodeOk

`func (o *AppReleaseDto) GetRapidCodeOk() (*string, bool)`

GetRapidCodeOk returns a tuple with the RapidCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRapidCode

`func (o *AppReleaseDto) SetRapidCode(v string)`

SetRapidCode sets RapidCode field to given value.

### HasRapidCode

`func (o *AppReleaseDto) HasRapidCode() bool`

HasRapidCode returns a boolean if a field has been set.

### GetSize

`func (o *AppReleaseDto) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AppReleaseDto) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AppReleaseDto) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *AppReleaseDto) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMd5

`func (o *AppReleaseDto) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *AppReleaseDto) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *AppReleaseDto) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *AppReleaseDto) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetSliceMd5

`func (o *AppReleaseDto) GetSliceMd5() string`

GetSliceMd5 returns the SliceMd5 field if non-nil, zero value otherwise.

### GetSliceMd5Ok

`func (o *AppReleaseDto) GetSliceMd5Ok() (*string, bool)`

GetSliceMd5Ok returns a tuple with the SliceMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceMd5

`func (o *AppReleaseDto) SetSliceMd5(v string)`

SetSliceMd5 sets SliceMd5 field to given value.

### HasSliceMd5

`func (o *AppReleaseDto) HasSliceMd5() bool`

HasSliceMd5 returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *AppReleaseDto) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *AppReleaseDto) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *AppReleaseDto) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *AppReleaseDto) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetProductType

`func (o *AppReleaseDto) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *AppReleaseDto) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *AppReleaseDto) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *AppReleaseDto) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetIsForceUpdate

`func (o *AppReleaseDto) GetIsForceUpdate() bool`

GetIsForceUpdate returns the IsForceUpdate field if non-nil, zero value otherwise.

### GetIsForceUpdateOk

`func (o *AppReleaseDto) GetIsForceUpdateOk() (*bool, bool)`

GetIsForceUpdateOk returns a tuple with the IsForceUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForceUpdate

`func (o *AppReleaseDto) SetIsForceUpdate(v bool)`

SetIsForceUpdate sets IsForceUpdate field to given value.

### HasIsForceUpdate

`func (o *AppReleaseDto) HasIsForceUpdate() bool`

HasIsForceUpdate returns a boolean if a field has been set.

### GetAppId

`func (o *AppReleaseDto) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AppReleaseDto) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AppReleaseDto) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *AppReleaseDto) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetIsEnabled

`func (o *AppReleaseDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AppReleaseDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AppReleaseDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *AppReleaseDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetChannel

`func (o *AppReleaseDto) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *AppReleaseDto) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *AppReleaseDto) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *AppReleaseDto) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetEnvironment

`func (o *AppReleaseDto) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AppReleaseDto) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AppReleaseDto) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AppReleaseDto) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



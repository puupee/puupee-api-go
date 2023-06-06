# CreateOrUpdateAppReleaseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** |  | [optional] 
**VersionName** | Pointer to **string** |  | [optional] 
**VersionCode** | Pointer to **int64** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**RapidCode** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Md5** | Pointer to **string** |  | [optional] 
**SliceMd5** | Pointer to **string** |  | [optional] 
**ProductType** | Pointer to **string** |  | [optional] 
**IsForceUpdate** | Pointer to **bool** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**Channel** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateOrUpdateAppReleaseDto

`func NewCreateOrUpdateAppReleaseDto() *CreateOrUpdateAppReleaseDto`

NewCreateOrUpdateAppReleaseDto instantiates a new CreateOrUpdateAppReleaseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateAppReleaseDtoWithDefaults

`func NewCreateOrUpdateAppReleaseDtoWithDefaults() *CreateOrUpdateAppReleaseDto`

NewCreateOrUpdateAppReleaseDtoWithDefaults instantiates a new CreateOrUpdateAppReleaseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *CreateOrUpdateAppReleaseDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateOrUpdateAppReleaseDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateOrUpdateAppReleaseDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CreateOrUpdateAppReleaseDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionName

`func (o *CreateOrUpdateAppReleaseDto) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *CreateOrUpdateAppReleaseDto) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *CreateOrUpdateAppReleaseDto) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *CreateOrUpdateAppReleaseDto) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.

### GetVersionCode

`func (o *CreateOrUpdateAppReleaseDto) GetVersionCode() int64`

GetVersionCode returns the VersionCode field if non-nil, zero value otherwise.

### GetVersionCodeOk

`func (o *CreateOrUpdateAppReleaseDto) GetVersionCodeOk() (*int64, bool)`

GetVersionCodeOk returns a tuple with the VersionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionCode

`func (o *CreateOrUpdateAppReleaseDto) SetVersionCode(v int64)`

SetVersionCode sets VersionCode field to given value.

### HasVersionCode

`func (o *CreateOrUpdateAppReleaseDto) HasVersionCode() bool`

HasVersionCode returns a boolean if a field has been set.

### GetNotes

`func (o *CreateOrUpdateAppReleaseDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreateOrUpdateAppReleaseDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreateOrUpdateAppReleaseDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreateOrUpdateAppReleaseDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPlatform

`func (o *CreateOrUpdateAppReleaseDto) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CreateOrUpdateAppReleaseDto) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CreateOrUpdateAppReleaseDto) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CreateOrUpdateAppReleaseDto) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetKey

`func (o *CreateOrUpdateAppReleaseDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateOrUpdateAppReleaseDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateOrUpdateAppReleaseDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateOrUpdateAppReleaseDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetRapidCode

`func (o *CreateOrUpdateAppReleaseDto) GetRapidCode() string`

GetRapidCode returns the RapidCode field if non-nil, zero value otherwise.

### GetRapidCodeOk

`func (o *CreateOrUpdateAppReleaseDto) GetRapidCodeOk() (*string, bool)`

GetRapidCodeOk returns a tuple with the RapidCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRapidCode

`func (o *CreateOrUpdateAppReleaseDto) SetRapidCode(v string)`

SetRapidCode sets RapidCode field to given value.

### HasRapidCode

`func (o *CreateOrUpdateAppReleaseDto) HasRapidCode() bool`

HasRapidCode returns a boolean if a field has been set.

### GetSize

`func (o *CreateOrUpdateAppReleaseDto) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateOrUpdateAppReleaseDto) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateOrUpdateAppReleaseDto) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateOrUpdateAppReleaseDto) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMd5

`func (o *CreateOrUpdateAppReleaseDto) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *CreateOrUpdateAppReleaseDto) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *CreateOrUpdateAppReleaseDto) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *CreateOrUpdateAppReleaseDto) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetSliceMd5

`func (o *CreateOrUpdateAppReleaseDto) GetSliceMd5() string`

GetSliceMd5 returns the SliceMd5 field if non-nil, zero value otherwise.

### GetSliceMd5Ok

`func (o *CreateOrUpdateAppReleaseDto) GetSliceMd5Ok() (*string, bool)`

GetSliceMd5Ok returns a tuple with the SliceMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceMd5

`func (o *CreateOrUpdateAppReleaseDto) SetSliceMd5(v string)`

SetSliceMd5 sets SliceMd5 field to given value.

### HasSliceMd5

`func (o *CreateOrUpdateAppReleaseDto) HasSliceMd5() bool`

HasSliceMd5 returns a boolean if a field has been set.

### GetProductType

`func (o *CreateOrUpdateAppReleaseDto) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *CreateOrUpdateAppReleaseDto) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *CreateOrUpdateAppReleaseDto) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *CreateOrUpdateAppReleaseDto) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetIsForceUpdate

`func (o *CreateOrUpdateAppReleaseDto) GetIsForceUpdate() bool`

GetIsForceUpdate returns the IsForceUpdate field if non-nil, zero value otherwise.

### GetIsForceUpdateOk

`func (o *CreateOrUpdateAppReleaseDto) GetIsForceUpdateOk() (*bool, bool)`

GetIsForceUpdateOk returns a tuple with the IsForceUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForceUpdate

`func (o *CreateOrUpdateAppReleaseDto) SetIsForceUpdate(v bool)`

SetIsForceUpdate sets IsForceUpdate field to given value.

### HasIsForceUpdate

`func (o *CreateOrUpdateAppReleaseDto) HasIsForceUpdate() bool`

HasIsForceUpdate returns a boolean if a field has been set.

### GetAppId

`func (o *CreateOrUpdateAppReleaseDto) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *CreateOrUpdateAppReleaseDto) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *CreateOrUpdateAppReleaseDto) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *CreateOrUpdateAppReleaseDto) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetIsEnabled

`func (o *CreateOrUpdateAppReleaseDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *CreateOrUpdateAppReleaseDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *CreateOrUpdateAppReleaseDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *CreateOrUpdateAppReleaseDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetChannel

`func (o *CreateOrUpdateAppReleaseDto) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CreateOrUpdateAppReleaseDto) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CreateOrUpdateAppReleaseDto) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *CreateOrUpdateAppReleaseDto) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetEnvironment

`func (o *CreateOrUpdateAppReleaseDto) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CreateOrUpdateAppReleaseDto) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CreateOrUpdateAppReleaseDto) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CreateOrUpdateAppReleaseDto) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



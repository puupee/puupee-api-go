# AppDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Fromework** | Pointer to [**Framework**](Framework.md) |  | [optional] 
**AppType** | Pointer to [**AppType**](AppType.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**GitRepository** | Pointer to **NullableString** |  | [optional] 
**GitRepositoryType** | Pointer to [**GitRepositoryType**](GitRepositoryType.md) |  | [optional] 
**LatestRelease** | Pointer to [**AppReleaseDto**](AppReleaseDto.md) |  | [optional] 

## Methods

### NewAppDto

`func NewAppDto() *AppDto`

NewAppDto instantiates a new AppDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDtoWithDefaults

`func NewAppDtoWithDefaults() *AppDto`

NewAppDtoWithDefaults instantiates a new AppDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AppDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AppDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AppDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDisplayName

`func (o *AppDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AppDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AppDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AppDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *AppDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AppDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetFromework

`func (o *AppDto) GetFromework() Framework`

GetFromework returns the Fromework field if non-nil, zero value otherwise.

### GetFromeworkOk

`func (o *AppDto) GetFromeworkOk() (*Framework, bool)`

GetFromeworkOk returns a tuple with the Fromework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromework

`func (o *AppDto) SetFromework(v Framework)`

SetFromework sets Fromework field to given value.

### HasFromework

`func (o *AppDto) HasFromework() bool`

HasFromework returns a boolean if a field has been set.

### GetAppType

`func (o *AppDto) GetAppType() AppType`

GetAppType returns the AppType field if non-nil, zero value otherwise.

### GetAppTypeOk

`func (o *AppDto) GetAppTypeOk() (*AppType, bool)`

GetAppTypeOk returns a tuple with the AppType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppType

`func (o *AppDto) SetAppType(v AppType)`

SetAppType sets AppType field to given value.

### HasAppType

`func (o *AppDto) HasAppType() bool`

HasAppType returns a boolean if a field has been set.

### GetDescription

`func (o *AppDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AppDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AppDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIcon

`func (o *AppDto) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *AppDto) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *AppDto) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *AppDto) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *AppDto) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *AppDto) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetGitRepository

`func (o *AppDto) GetGitRepository() string`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *AppDto) GetGitRepositoryOk() (*string, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *AppDto) SetGitRepository(v string)`

SetGitRepository sets GitRepository field to given value.

### HasGitRepository

`func (o *AppDto) HasGitRepository() bool`

HasGitRepository returns a boolean if a field has been set.

### SetGitRepositoryNil

`func (o *AppDto) SetGitRepositoryNil(b bool)`

 SetGitRepositoryNil sets the value for GitRepository to be an explicit nil

### UnsetGitRepository
`func (o *AppDto) UnsetGitRepository()`

UnsetGitRepository ensures that no value is present for GitRepository, not even an explicit nil
### GetGitRepositoryType

`func (o *AppDto) GetGitRepositoryType() GitRepositoryType`

GetGitRepositoryType returns the GitRepositoryType field if non-nil, zero value otherwise.

### GetGitRepositoryTypeOk

`func (o *AppDto) GetGitRepositoryTypeOk() (*GitRepositoryType, bool)`

GetGitRepositoryTypeOk returns a tuple with the GitRepositoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepositoryType

`func (o *AppDto) SetGitRepositoryType(v GitRepositoryType)`

SetGitRepositoryType sets GitRepositoryType field to given value.

### HasGitRepositoryType

`func (o *AppDto) HasGitRepositoryType() bool`

HasGitRepositoryType returns a boolean if a field has been set.

### GetLatestRelease

`func (o *AppDto) GetLatestRelease() AppReleaseDto`

GetLatestRelease returns the LatestRelease field if non-nil, zero value otherwise.

### GetLatestReleaseOk

`func (o *AppDto) GetLatestReleaseOk() (*AppReleaseDto, bool)`

GetLatestReleaseOk returns a tuple with the LatestRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRelease

`func (o *AppDto) SetLatestRelease(v AppReleaseDto)`

SetLatestRelease sets LatestRelease field to given value.

### HasLatestRelease

`func (o *AppDto) HasLatestRelease() bool`

HasLatestRelease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



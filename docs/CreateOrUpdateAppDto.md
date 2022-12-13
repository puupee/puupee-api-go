# CreateOrUpdateAppDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Framework** | Pointer to **string** |  | [optional] 
**AppType** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**GitRepository** | Pointer to **string** |  | [optional] 
**GitRepositoryType** | Pointer to **string** |  | [optional] 
**Features** | Pointer to [**[]AppFeatureDto**](AppFeatureDto.md) |  | [optional] 
**Sdks** | Pointer to [**[]AppSdkDto**](AppSdkDto.md) |  | [optional] 

## Methods

### NewCreateOrUpdateAppDto

`func NewCreateOrUpdateAppDto() *CreateOrUpdateAppDto`

NewCreateOrUpdateAppDto instantiates a new CreateOrUpdateAppDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateAppDtoWithDefaults

`func NewCreateOrUpdateAppDtoWithDefaults() *CreateOrUpdateAppDto`

NewCreateOrUpdateAppDtoWithDefaults instantiates a new CreateOrUpdateAppDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOrUpdateAppDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrUpdateAppDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrUpdateAppDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateOrUpdateAppDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateOrUpdateAppDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateOrUpdateAppDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateOrUpdateAppDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateOrUpdateAppDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFramework

`func (o *CreateOrUpdateAppDto) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *CreateOrUpdateAppDto) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *CreateOrUpdateAppDto) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *CreateOrUpdateAppDto) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetAppType

`func (o *CreateOrUpdateAppDto) GetAppType() string`

GetAppType returns the AppType field if non-nil, zero value otherwise.

### GetAppTypeOk

`func (o *CreateOrUpdateAppDto) GetAppTypeOk() (*string, bool)`

GetAppTypeOk returns a tuple with the AppType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppType

`func (o *CreateOrUpdateAppDto) SetAppType(v string)`

SetAppType sets AppType field to given value.

### HasAppType

`func (o *CreateOrUpdateAppDto) HasAppType() bool`

HasAppType returns a boolean if a field has been set.

### GetDescription

`func (o *CreateOrUpdateAppDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrUpdateAppDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrUpdateAppDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOrUpdateAppDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *CreateOrUpdateAppDto) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CreateOrUpdateAppDto) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CreateOrUpdateAppDto) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CreateOrUpdateAppDto) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetGitRepository

`func (o *CreateOrUpdateAppDto) GetGitRepository() string`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *CreateOrUpdateAppDto) GetGitRepositoryOk() (*string, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *CreateOrUpdateAppDto) SetGitRepository(v string)`

SetGitRepository sets GitRepository field to given value.

### HasGitRepository

`func (o *CreateOrUpdateAppDto) HasGitRepository() bool`

HasGitRepository returns a boolean if a field has been set.

### GetGitRepositoryType

`func (o *CreateOrUpdateAppDto) GetGitRepositoryType() string`

GetGitRepositoryType returns the GitRepositoryType field if non-nil, zero value otherwise.

### GetGitRepositoryTypeOk

`func (o *CreateOrUpdateAppDto) GetGitRepositoryTypeOk() (*string, bool)`

GetGitRepositoryTypeOk returns a tuple with the GitRepositoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepositoryType

`func (o *CreateOrUpdateAppDto) SetGitRepositoryType(v string)`

SetGitRepositoryType sets GitRepositoryType field to given value.

### HasGitRepositoryType

`func (o *CreateOrUpdateAppDto) HasGitRepositoryType() bool`

HasGitRepositoryType returns a boolean if a field has been set.

### GetFeatures

`func (o *CreateOrUpdateAppDto) GetFeatures() []AppFeatureDto`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *CreateOrUpdateAppDto) GetFeaturesOk() (*[]AppFeatureDto, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *CreateOrUpdateAppDto) SetFeatures(v []AppFeatureDto)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *CreateOrUpdateAppDto) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetSdks

`func (o *CreateOrUpdateAppDto) GetSdks() []AppSdkDto`

GetSdks returns the Sdks field if non-nil, zero value otherwise.

### GetSdksOk

`func (o *CreateOrUpdateAppDto) GetSdksOk() (*[]AppSdkDto, bool)`

GetSdksOk returns a tuple with the Sdks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdks

`func (o *CreateOrUpdateAppDto) SetSdks(v []AppSdkDto)`

SetSdks sets Sdks field to given value.

### HasSdks

`func (o *CreateOrUpdateAppDto) HasSdks() bool`

HasSdks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



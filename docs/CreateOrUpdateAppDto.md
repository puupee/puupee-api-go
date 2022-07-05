# CreateOrUpdateAppDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Fromework** | Pointer to [**Framework**](Framework.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**GitRepository** | Pointer to **NullableString** |  | [optional] 
**GitRepositoryType** | Pointer to [**GitRepositoryType**](GitRepositoryType.md) |  | [optional] 

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

### SetNameNil

`func (o *CreateOrUpdateAppDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateOrUpdateAppDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
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

### SetDisplayNameNil

`func (o *CreateOrUpdateAppDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CreateOrUpdateAppDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetFromework

`func (o *CreateOrUpdateAppDto) GetFromework() Framework`

GetFromework returns the Fromework field if non-nil, zero value otherwise.

### GetFromeworkOk

`func (o *CreateOrUpdateAppDto) GetFromeworkOk() (*Framework, bool)`

GetFromeworkOk returns a tuple with the Fromework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromework

`func (o *CreateOrUpdateAppDto) SetFromework(v Framework)`

SetFromework sets Fromework field to given value.

### HasFromework

`func (o *CreateOrUpdateAppDto) HasFromework() bool`

HasFromework returns a boolean if a field has been set.

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

### SetDescriptionNil

`func (o *CreateOrUpdateAppDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateOrUpdateAppDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
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

### SetIconNil

`func (o *CreateOrUpdateAppDto) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *CreateOrUpdateAppDto) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
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

### SetGitRepositoryNil

`func (o *CreateOrUpdateAppDto) SetGitRepositoryNil(b bool)`

 SetGitRepositoryNil sets the value for GitRepository to be an explicit nil

### UnsetGitRepository
`func (o *CreateOrUpdateAppDto) UnsetGitRepository()`

UnsetGitRepository ensures that no value is present for GitRepository, not even an explicit nil
### GetGitRepositoryType

`func (o *CreateOrUpdateAppDto) GetGitRepositoryType() GitRepositoryType`

GetGitRepositoryType returns the GitRepositoryType field if non-nil, zero value otherwise.

### GetGitRepositoryTypeOk

`func (o *CreateOrUpdateAppDto) GetGitRepositoryTypeOk() (*GitRepositoryType, bool)`

GetGitRepositoryTypeOk returns a tuple with the GitRepositoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepositoryType

`func (o *CreateOrUpdateAppDto) SetGitRepositoryType(v GitRepositoryType)`

SetGitRepositoryType sets GitRepositoryType field to given value.

### HasGitRepositoryType

`func (o *CreateOrUpdateAppDto) HasGitRepositoryType() bool`

HasGitRepositoryType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


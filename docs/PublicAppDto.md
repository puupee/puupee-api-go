# PublicAppDto

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
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Framework** | Pointer to **string** |  | [optional] 
**AppType** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**HomePage** | Pointer to **string** |  | [optional] 
**SortIndex** | Pointer to **int32** |  | [optional] 
**GitRepository** | Pointer to **string** |  | [optional] 
**GitRepositoryType** | Pointer to **string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**IsPublished** | Pointer to **bool** |  | [optional] 
**LatestReleases** | Pointer to [**[]AppReleaseDto**](AppReleaseDto.md) |  | [optional] 
**Creator** | Pointer to [**IdentityUserDto**](IdentityUserDto.md) |  | [optional] 
**Features** | Pointer to [**[]AppFeatureDto**](AppFeatureDto.md) |  | [optional] 
**Sdks** | Pointer to [**[]AppSdkDto**](AppSdkDto.md) |  | [optional] 
**Subscribed** | Pointer to **bool** |  | [optional] 

## Methods

### NewPublicAppDto

`func NewPublicAppDto() *PublicAppDto`

NewPublicAppDto instantiates a new PublicAppDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAppDtoWithDefaults

`func NewPublicAppDtoWithDefaults() *PublicAppDto`

NewPublicAppDtoWithDefaults instantiates a new PublicAppDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PublicAppDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicAppDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicAppDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PublicAppDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreationTime

`func (o *PublicAppDto) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *PublicAppDto) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *PublicAppDto) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *PublicAppDto) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *PublicAppDto) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *PublicAppDto) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *PublicAppDto) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *PublicAppDto) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetLastModificationTime

`func (o *PublicAppDto) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *PublicAppDto) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *PublicAppDto) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *PublicAppDto) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### GetLastModifierId

`func (o *PublicAppDto) GetLastModifierId() string`

GetLastModifierId returns the LastModifierId field if non-nil, zero value otherwise.

### GetLastModifierIdOk

`func (o *PublicAppDto) GetLastModifierIdOk() (*string, bool)`

GetLastModifierIdOk returns a tuple with the LastModifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierId

`func (o *PublicAppDto) SetLastModifierId(v string)`

SetLastModifierId sets LastModifierId field to given value.

### HasLastModifierId

`func (o *PublicAppDto) HasLastModifierId() bool`

HasLastModifierId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *PublicAppDto) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *PublicAppDto) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *PublicAppDto) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *PublicAppDto) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeleterId

`func (o *PublicAppDto) GetDeleterId() string`

GetDeleterId returns the DeleterId field if non-nil, zero value otherwise.

### GetDeleterIdOk

`func (o *PublicAppDto) GetDeleterIdOk() (*string, bool)`

GetDeleterIdOk returns a tuple with the DeleterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleterId

`func (o *PublicAppDto) SetDeleterId(v string)`

SetDeleterId sets DeleterId field to given value.

### HasDeleterId

`func (o *PublicAppDto) HasDeleterId() bool`

HasDeleterId returns a boolean if a field has been set.

### GetDeletionTime

`func (o *PublicAppDto) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *PublicAppDto) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *PublicAppDto) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *PublicAppDto) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### GetName

`func (o *PublicAppDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicAppDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicAppDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublicAppDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *PublicAppDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PublicAppDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PublicAppDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PublicAppDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFramework

`func (o *PublicAppDto) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *PublicAppDto) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *PublicAppDto) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *PublicAppDto) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetAppType

`func (o *PublicAppDto) GetAppType() string`

GetAppType returns the AppType field if non-nil, zero value otherwise.

### GetAppTypeOk

`func (o *PublicAppDto) GetAppTypeOk() (*string, bool)`

GetAppTypeOk returns a tuple with the AppType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppType

`func (o *PublicAppDto) SetAppType(v string)`

SetAppType sets AppType field to given value.

### HasAppType

`func (o *PublicAppDto) HasAppType() bool`

HasAppType returns a boolean if a field has been set.

### GetDescription

`func (o *PublicAppDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PublicAppDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PublicAppDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PublicAppDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *PublicAppDto) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *PublicAppDto) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *PublicAppDto) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *PublicAppDto) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetHomePage

`func (o *PublicAppDto) GetHomePage() string`

GetHomePage returns the HomePage field if non-nil, zero value otherwise.

### GetHomePageOk

`func (o *PublicAppDto) GetHomePageOk() (*string, bool)`

GetHomePageOk returns a tuple with the HomePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePage

`func (o *PublicAppDto) SetHomePage(v string)`

SetHomePage sets HomePage field to given value.

### HasHomePage

`func (o *PublicAppDto) HasHomePage() bool`

HasHomePage returns a boolean if a field has been set.

### GetSortIndex

`func (o *PublicAppDto) GetSortIndex() int32`

GetSortIndex returns the SortIndex field if non-nil, zero value otherwise.

### GetSortIndexOk

`func (o *PublicAppDto) GetSortIndexOk() (*int32, bool)`

GetSortIndexOk returns a tuple with the SortIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortIndex

`func (o *PublicAppDto) SetSortIndex(v int32)`

SetSortIndex sets SortIndex field to given value.

### HasSortIndex

`func (o *PublicAppDto) HasSortIndex() bool`

HasSortIndex returns a boolean if a field has been set.

### GetGitRepository

`func (o *PublicAppDto) GetGitRepository() string`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *PublicAppDto) GetGitRepositoryOk() (*string, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *PublicAppDto) SetGitRepository(v string)`

SetGitRepository sets GitRepository field to given value.

### HasGitRepository

`func (o *PublicAppDto) HasGitRepository() bool`

HasGitRepository returns a boolean if a field has been set.

### GetGitRepositoryType

`func (o *PublicAppDto) GetGitRepositoryType() string`

GetGitRepositoryType returns the GitRepositoryType field if non-nil, zero value otherwise.

### GetGitRepositoryTypeOk

`func (o *PublicAppDto) GetGitRepositoryTypeOk() (*string, bool)`

GetGitRepositoryTypeOk returns a tuple with the GitRepositoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepositoryType

`func (o *PublicAppDto) SetGitRepositoryType(v string)`

SetGitRepositoryType sets GitRepositoryType field to given value.

### HasGitRepositoryType

`func (o *PublicAppDto) HasGitRepositoryType() bool`

HasGitRepositoryType returns a boolean if a field has been set.

### GetIsEnabled

`func (o *PublicAppDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *PublicAppDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *PublicAppDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *PublicAppDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsPublished

`func (o *PublicAppDto) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *PublicAppDto) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *PublicAppDto) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *PublicAppDto) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetLatestReleases

`func (o *PublicAppDto) GetLatestReleases() []AppReleaseDto`

GetLatestReleases returns the LatestReleases field if non-nil, zero value otherwise.

### GetLatestReleasesOk

`func (o *PublicAppDto) GetLatestReleasesOk() (*[]AppReleaseDto, bool)`

GetLatestReleasesOk returns a tuple with the LatestReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestReleases

`func (o *PublicAppDto) SetLatestReleases(v []AppReleaseDto)`

SetLatestReleases sets LatestReleases field to given value.

### HasLatestReleases

`func (o *PublicAppDto) HasLatestReleases() bool`

HasLatestReleases returns a boolean if a field has been set.

### GetCreator

`func (o *PublicAppDto) GetCreator() IdentityUserDto`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *PublicAppDto) GetCreatorOk() (*IdentityUserDto, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *PublicAppDto) SetCreator(v IdentityUserDto)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *PublicAppDto) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetFeatures

`func (o *PublicAppDto) GetFeatures() []AppFeatureDto`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *PublicAppDto) GetFeaturesOk() (*[]AppFeatureDto, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *PublicAppDto) SetFeatures(v []AppFeatureDto)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *PublicAppDto) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetSdks

`func (o *PublicAppDto) GetSdks() []AppSdkDto`

GetSdks returns the Sdks field if non-nil, zero value otherwise.

### GetSdksOk

`func (o *PublicAppDto) GetSdksOk() (*[]AppSdkDto, bool)`

GetSdksOk returns a tuple with the Sdks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdks

`func (o *PublicAppDto) SetSdks(v []AppSdkDto)`

SetSdks sets Sdks field to given value.

### HasSdks

`func (o *PublicAppDto) HasSdks() bool`

HasSdks returns a boolean if a field has been set.

### GetSubscribed

`func (o *PublicAppDto) GetSubscribed() bool`

GetSubscribed returns the Subscribed field if non-nil, zero value otherwise.

### GetSubscribedOk

`func (o *PublicAppDto) GetSubscribedOk() (*bool, bool)`

GetSubscribedOk returns a tuple with the Subscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribed

`func (o *PublicAppDto) SetSubscribed(v bool)`

SetSubscribed sets Subscribed field to given value.

### HasSubscribed

`func (o *PublicAppDto) HasSubscribed() bool`

HasSubscribed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



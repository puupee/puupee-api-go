# AppWithUserDto

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
**WebhookUrl** | Pointer to **string** |  | [optional] 
**BusinessDomain** | Pointer to **string** |  | [optional] 
**BusinessUrl** | Pointer to **string** |  | [optional] 
**SubscriptionEnabled** | Pointer to **bool** |  | [optional] 
**SubscriptionPlatforms** | Pointer to **string** |  | [optional] 
**LatestReleases** | Pointer to [**[]AppReleaseDto**](AppReleaseDto.md) |  | [optional] 
**Creator** | Pointer to [**IdentityUserDto**](IdentityUserDto.md) |  | [optional] 
**Features** | Pointer to [**[]AppFeatureDto**](AppFeatureDto.md) |  | [optional] 
**Sdks** | Pointer to [**[]AppSdkDto**](AppSdkDto.md) |  | [optional] 
**Subscribed** | Pointer to **bool** |  | [optional] 

## Methods

### NewAppWithUserDto

`func NewAppWithUserDto() *AppWithUserDto`

NewAppWithUserDto instantiates a new AppWithUserDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppWithUserDtoWithDefaults

`func NewAppWithUserDtoWithDefaults() *AppWithUserDto`

NewAppWithUserDtoWithDefaults instantiates a new AppWithUserDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppWithUserDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppWithUserDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppWithUserDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppWithUserDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreationTime

`func (o *AppWithUserDto) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AppWithUserDto) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AppWithUserDto) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *AppWithUserDto) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *AppWithUserDto) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *AppWithUserDto) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *AppWithUserDto) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *AppWithUserDto) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetLastModificationTime

`func (o *AppWithUserDto) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *AppWithUserDto) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *AppWithUserDto) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *AppWithUserDto) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### GetLastModifierId

`func (o *AppWithUserDto) GetLastModifierId() string`

GetLastModifierId returns the LastModifierId field if non-nil, zero value otherwise.

### GetLastModifierIdOk

`func (o *AppWithUserDto) GetLastModifierIdOk() (*string, bool)`

GetLastModifierIdOk returns a tuple with the LastModifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierId

`func (o *AppWithUserDto) SetLastModifierId(v string)`

SetLastModifierId sets LastModifierId field to given value.

### HasLastModifierId

`func (o *AppWithUserDto) HasLastModifierId() bool`

HasLastModifierId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *AppWithUserDto) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *AppWithUserDto) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *AppWithUserDto) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *AppWithUserDto) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeleterId

`func (o *AppWithUserDto) GetDeleterId() string`

GetDeleterId returns the DeleterId field if non-nil, zero value otherwise.

### GetDeleterIdOk

`func (o *AppWithUserDto) GetDeleterIdOk() (*string, bool)`

GetDeleterIdOk returns a tuple with the DeleterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleterId

`func (o *AppWithUserDto) SetDeleterId(v string)`

SetDeleterId sets DeleterId field to given value.

### HasDeleterId

`func (o *AppWithUserDto) HasDeleterId() bool`

HasDeleterId returns a boolean if a field has been set.

### GetDeletionTime

`func (o *AppWithUserDto) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *AppWithUserDto) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *AppWithUserDto) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *AppWithUserDto) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### GetName

`func (o *AppWithUserDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppWithUserDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppWithUserDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppWithUserDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *AppWithUserDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AppWithUserDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AppWithUserDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AppWithUserDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFramework

`func (o *AppWithUserDto) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *AppWithUserDto) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *AppWithUserDto) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *AppWithUserDto) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetAppType

`func (o *AppWithUserDto) GetAppType() string`

GetAppType returns the AppType field if non-nil, zero value otherwise.

### GetAppTypeOk

`func (o *AppWithUserDto) GetAppTypeOk() (*string, bool)`

GetAppTypeOk returns a tuple with the AppType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppType

`func (o *AppWithUserDto) SetAppType(v string)`

SetAppType sets AppType field to given value.

### HasAppType

`func (o *AppWithUserDto) HasAppType() bool`

HasAppType returns a boolean if a field has been set.

### GetDescription

`func (o *AppWithUserDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppWithUserDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppWithUserDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppWithUserDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *AppWithUserDto) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *AppWithUserDto) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *AppWithUserDto) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *AppWithUserDto) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetHomePage

`func (o *AppWithUserDto) GetHomePage() string`

GetHomePage returns the HomePage field if non-nil, zero value otherwise.

### GetHomePageOk

`func (o *AppWithUserDto) GetHomePageOk() (*string, bool)`

GetHomePageOk returns a tuple with the HomePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePage

`func (o *AppWithUserDto) SetHomePage(v string)`

SetHomePage sets HomePage field to given value.

### HasHomePage

`func (o *AppWithUserDto) HasHomePage() bool`

HasHomePage returns a boolean if a field has been set.

### GetSortIndex

`func (o *AppWithUserDto) GetSortIndex() int32`

GetSortIndex returns the SortIndex field if non-nil, zero value otherwise.

### GetSortIndexOk

`func (o *AppWithUserDto) GetSortIndexOk() (*int32, bool)`

GetSortIndexOk returns a tuple with the SortIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortIndex

`func (o *AppWithUserDto) SetSortIndex(v int32)`

SetSortIndex sets SortIndex field to given value.

### HasSortIndex

`func (o *AppWithUserDto) HasSortIndex() bool`

HasSortIndex returns a boolean if a field has been set.

### GetGitRepository

`func (o *AppWithUserDto) GetGitRepository() string`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *AppWithUserDto) GetGitRepositoryOk() (*string, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *AppWithUserDto) SetGitRepository(v string)`

SetGitRepository sets GitRepository field to given value.

### HasGitRepository

`func (o *AppWithUserDto) HasGitRepository() bool`

HasGitRepository returns a boolean if a field has been set.

### GetGitRepositoryType

`func (o *AppWithUserDto) GetGitRepositoryType() string`

GetGitRepositoryType returns the GitRepositoryType field if non-nil, zero value otherwise.

### GetGitRepositoryTypeOk

`func (o *AppWithUserDto) GetGitRepositoryTypeOk() (*string, bool)`

GetGitRepositoryTypeOk returns a tuple with the GitRepositoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepositoryType

`func (o *AppWithUserDto) SetGitRepositoryType(v string)`

SetGitRepositoryType sets GitRepositoryType field to given value.

### HasGitRepositoryType

`func (o *AppWithUserDto) HasGitRepositoryType() bool`

HasGitRepositoryType returns a boolean if a field has been set.

### GetIsEnabled

`func (o *AppWithUserDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AppWithUserDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AppWithUserDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *AppWithUserDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsPublished

`func (o *AppWithUserDto) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *AppWithUserDto) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *AppWithUserDto) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *AppWithUserDto) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *AppWithUserDto) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *AppWithUserDto) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *AppWithUserDto) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *AppWithUserDto) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetBusinessDomain

`func (o *AppWithUserDto) GetBusinessDomain() string`

GetBusinessDomain returns the BusinessDomain field if non-nil, zero value otherwise.

### GetBusinessDomainOk

`func (o *AppWithUserDto) GetBusinessDomainOk() (*string, bool)`

GetBusinessDomainOk returns a tuple with the BusinessDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessDomain

`func (o *AppWithUserDto) SetBusinessDomain(v string)`

SetBusinessDomain sets BusinessDomain field to given value.

### HasBusinessDomain

`func (o *AppWithUserDto) HasBusinessDomain() bool`

HasBusinessDomain returns a boolean if a field has been set.

### GetBusinessUrl

`func (o *AppWithUserDto) GetBusinessUrl() string`

GetBusinessUrl returns the BusinessUrl field if non-nil, zero value otherwise.

### GetBusinessUrlOk

`func (o *AppWithUserDto) GetBusinessUrlOk() (*string, bool)`

GetBusinessUrlOk returns a tuple with the BusinessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUrl

`func (o *AppWithUserDto) SetBusinessUrl(v string)`

SetBusinessUrl sets BusinessUrl field to given value.

### HasBusinessUrl

`func (o *AppWithUserDto) HasBusinessUrl() bool`

HasBusinessUrl returns a boolean if a field has been set.

### GetSubscriptionEnabled

`func (o *AppWithUserDto) GetSubscriptionEnabled() bool`

GetSubscriptionEnabled returns the SubscriptionEnabled field if non-nil, zero value otherwise.

### GetSubscriptionEnabledOk

`func (o *AppWithUserDto) GetSubscriptionEnabledOk() (*bool, bool)`

GetSubscriptionEnabledOk returns a tuple with the SubscriptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionEnabled

`func (o *AppWithUserDto) SetSubscriptionEnabled(v bool)`

SetSubscriptionEnabled sets SubscriptionEnabled field to given value.

### HasSubscriptionEnabled

`func (o *AppWithUserDto) HasSubscriptionEnabled() bool`

HasSubscriptionEnabled returns a boolean if a field has been set.

### GetSubscriptionPlatforms

`func (o *AppWithUserDto) GetSubscriptionPlatforms() string`

GetSubscriptionPlatforms returns the SubscriptionPlatforms field if non-nil, zero value otherwise.

### GetSubscriptionPlatformsOk

`func (o *AppWithUserDto) GetSubscriptionPlatformsOk() (*string, bool)`

GetSubscriptionPlatformsOk returns a tuple with the SubscriptionPlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPlatforms

`func (o *AppWithUserDto) SetSubscriptionPlatforms(v string)`

SetSubscriptionPlatforms sets SubscriptionPlatforms field to given value.

### HasSubscriptionPlatforms

`func (o *AppWithUserDto) HasSubscriptionPlatforms() bool`

HasSubscriptionPlatforms returns a boolean if a field has been set.

### GetLatestReleases

`func (o *AppWithUserDto) GetLatestReleases() []AppReleaseDto`

GetLatestReleases returns the LatestReleases field if non-nil, zero value otherwise.

### GetLatestReleasesOk

`func (o *AppWithUserDto) GetLatestReleasesOk() (*[]AppReleaseDto, bool)`

GetLatestReleasesOk returns a tuple with the LatestReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestReleases

`func (o *AppWithUserDto) SetLatestReleases(v []AppReleaseDto)`

SetLatestReleases sets LatestReleases field to given value.

### HasLatestReleases

`func (o *AppWithUserDto) HasLatestReleases() bool`

HasLatestReleases returns a boolean if a field has been set.

### GetCreator

`func (o *AppWithUserDto) GetCreator() IdentityUserDto`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *AppWithUserDto) GetCreatorOk() (*IdentityUserDto, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *AppWithUserDto) SetCreator(v IdentityUserDto)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *AppWithUserDto) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetFeatures

`func (o *AppWithUserDto) GetFeatures() []AppFeatureDto`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *AppWithUserDto) GetFeaturesOk() (*[]AppFeatureDto, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *AppWithUserDto) SetFeatures(v []AppFeatureDto)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *AppWithUserDto) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetSdks

`func (o *AppWithUserDto) GetSdks() []AppSdkDto`

GetSdks returns the Sdks field if non-nil, zero value otherwise.

### GetSdksOk

`func (o *AppWithUserDto) GetSdksOk() (*[]AppSdkDto, bool)`

GetSdksOk returns a tuple with the Sdks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdks

`func (o *AppWithUserDto) SetSdks(v []AppSdkDto)`

SetSdks sets Sdks field to given value.

### HasSdks

`func (o *AppWithUserDto) HasSdks() bool`

HasSdks returns a boolean if a field has been set.

### GetSubscribed

`func (o *AppWithUserDto) GetSubscribed() bool`

GetSubscribed returns the Subscribed field if non-nil, zero value otherwise.

### GetSubscribedOk

`func (o *AppWithUserDto) GetSubscribedOk() (*bool, bool)`

GetSubscribedOk returns a tuple with the Subscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribed

`func (o *AppWithUserDto) SetSubscribed(v bool)`

SetSubscribed sets Subscribed field to given value.

### HasSubscribed

`func (o *AppWithUserDto) HasSubscribed() bool`

HasSubscribed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AppDto

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
**SubscriptionPlatforms** | Pointer to **string** |  | [optional] 
**FreePlatforms** | Pointer to **string** |  | [optional] 
**SpecJsonSchema** | Pointer to **string** |  | [optional] 
**LatestReleases** | Pointer to [**[]AppReleaseDto**](AppReleaseDto.md) |  | [optional] 
**Creator** | Pointer to [**IdentityUserDto**](IdentityUserDto.md) |  | [optional] 
**Features** | Pointer to [**[]AppFeatureDto**](AppFeatureDto.md) |  | [optional] 
**Sdks** | Pointer to [**[]AppSdkDto**](AppSdkDto.md) |  | [optional] 

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

### GetId

`func (o *AppDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreationTime

`func (o *AppDto) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AppDto) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AppDto) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *AppDto) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *AppDto) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *AppDto) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *AppDto) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *AppDto) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetLastModificationTime

`func (o *AppDto) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *AppDto) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *AppDto) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *AppDto) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### GetLastModifierId

`func (o *AppDto) GetLastModifierId() string`

GetLastModifierId returns the LastModifierId field if non-nil, zero value otherwise.

### GetLastModifierIdOk

`func (o *AppDto) GetLastModifierIdOk() (*string, bool)`

GetLastModifierIdOk returns a tuple with the LastModifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierId

`func (o *AppDto) SetLastModifierId(v string)`

SetLastModifierId sets LastModifierId field to given value.

### HasLastModifierId

`func (o *AppDto) HasLastModifierId() bool`

HasLastModifierId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *AppDto) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *AppDto) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *AppDto) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *AppDto) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeleterId

`func (o *AppDto) GetDeleterId() string`

GetDeleterId returns the DeleterId field if non-nil, zero value otherwise.

### GetDeleterIdOk

`func (o *AppDto) GetDeleterIdOk() (*string, bool)`

GetDeleterIdOk returns a tuple with the DeleterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleterId

`func (o *AppDto) SetDeleterId(v string)`

SetDeleterId sets DeleterId field to given value.

### HasDeleterId

`func (o *AppDto) HasDeleterId() bool`

HasDeleterId returns a boolean if a field has been set.

### GetDeletionTime

`func (o *AppDto) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *AppDto) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *AppDto) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *AppDto) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

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

### GetFramework

`func (o *AppDto) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *AppDto) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *AppDto) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *AppDto) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetAppType

`func (o *AppDto) GetAppType() string`

GetAppType returns the AppType field if non-nil, zero value otherwise.

### GetAppTypeOk

`func (o *AppDto) GetAppTypeOk() (*string, bool)`

GetAppTypeOk returns a tuple with the AppType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppType

`func (o *AppDto) SetAppType(v string)`

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

### GetHomePage

`func (o *AppDto) GetHomePage() string`

GetHomePage returns the HomePage field if non-nil, zero value otherwise.

### GetHomePageOk

`func (o *AppDto) GetHomePageOk() (*string, bool)`

GetHomePageOk returns a tuple with the HomePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePage

`func (o *AppDto) SetHomePage(v string)`

SetHomePage sets HomePage field to given value.

### HasHomePage

`func (o *AppDto) HasHomePage() bool`

HasHomePage returns a boolean if a field has been set.

### GetSortIndex

`func (o *AppDto) GetSortIndex() int32`

GetSortIndex returns the SortIndex field if non-nil, zero value otherwise.

### GetSortIndexOk

`func (o *AppDto) GetSortIndexOk() (*int32, bool)`

GetSortIndexOk returns a tuple with the SortIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortIndex

`func (o *AppDto) SetSortIndex(v int32)`

SetSortIndex sets SortIndex field to given value.

### HasSortIndex

`func (o *AppDto) HasSortIndex() bool`

HasSortIndex returns a boolean if a field has been set.

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

### GetGitRepositoryType

`func (o *AppDto) GetGitRepositoryType() string`

GetGitRepositoryType returns the GitRepositoryType field if non-nil, zero value otherwise.

### GetGitRepositoryTypeOk

`func (o *AppDto) GetGitRepositoryTypeOk() (*string, bool)`

GetGitRepositoryTypeOk returns a tuple with the GitRepositoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepositoryType

`func (o *AppDto) SetGitRepositoryType(v string)`

SetGitRepositoryType sets GitRepositoryType field to given value.

### HasGitRepositoryType

`func (o *AppDto) HasGitRepositoryType() bool`

HasGitRepositoryType returns a boolean if a field has been set.

### GetIsEnabled

`func (o *AppDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AppDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AppDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *AppDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsPublished

`func (o *AppDto) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *AppDto) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *AppDto) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *AppDto) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *AppDto) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *AppDto) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *AppDto) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *AppDto) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetBusinessDomain

`func (o *AppDto) GetBusinessDomain() string`

GetBusinessDomain returns the BusinessDomain field if non-nil, zero value otherwise.

### GetBusinessDomainOk

`func (o *AppDto) GetBusinessDomainOk() (*string, bool)`

GetBusinessDomainOk returns a tuple with the BusinessDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessDomain

`func (o *AppDto) SetBusinessDomain(v string)`

SetBusinessDomain sets BusinessDomain field to given value.

### HasBusinessDomain

`func (o *AppDto) HasBusinessDomain() bool`

HasBusinessDomain returns a boolean if a field has been set.

### GetBusinessUrl

`func (o *AppDto) GetBusinessUrl() string`

GetBusinessUrl returns the BusinessUrl field if non-nil, zero value otherwise.

### GetBusinessUrlOk

`func (o *AppDto) GetBusinessUrlOk() (*string, bool)`

GetBusinessUrlOk returns a tuple with the BusinessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUrl

`func (o *AppDto) SetBusinessUrl(v string)`

SetBusinessUrl sets BusinessUrl field to given value.

### HasBusinessUrl

`func (o *AppDto) HasBusinessUrl() bool`

HasBusinessUrl returns a boolean if a field has been set.

### GetSubscriptionPlatforms

`func (o *AppDto) GetSubscriptionPlatforms() string`

GetSubscriptionPlatforms returns the SubscriptionPlatforms field if non-nil, zero value otherwise.

### GetSubscriptionPlatformsOk

`func (o *AppDto) GetSubscriptionPlatformsOk() (*string, bool)`

GetSubscriptionPlatformsOk returns a tuple with the SubscriptionPlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPlatforms

`func (o *AppDto) SetSubscriptionPlatforms(v string)`

SetSubscriptionPlatforms sets SubscriptionPlatforms field to given value.

### HasSubscriptionPlatforms

`func (o *AppDto) HasSubscriptionPlatforms() bool`

HasSubscriptionPlatforms returns a boolean if a field has been set.

### GetFreePlatforms

`func (o *AppDto) GetFreePlatforms() string`

GetFreePlatforms returns the FreePlatforms field if non-nil, zero value otherwise.

### GetFreePlatformsOk

`func (o *AppDto) GetFreePlatformsOk() (*string, bool)`

GetFreePlatformsOk returns a tuple with the FreePlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreePlatforms

`func (o *AppDto) SetFreePlatforms(v string)`

SetFreePlatforms sets FreePlatforms field to given value.

### HasFreePlatforms

`func (o *AppDto) HasFreePlatforms() bool`

HasFreePlatforms returns a boolean if a field has been set.

### GetSpecJsonSchema

`func (o *AppDto) GetSpecJsonSchema() string`

GetSpecJsonSchema returns the SpecJsonSchema field if non-nil, zero value otherwise.

### GetSpecJsonSchemaOk

`func (o *AppDto) GetSpecJsonSchemaOk() (*string, bool)`

GetSpecJsonSchemaOk returns a tuple with the SpecJsonSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecJsonSchema

`func (o *AppDto) SetSpecJsonSchema(v string)`

SetSpecJsonSchema sets SpecJsonSchema field to given value.

### HasSpecJsonSchema

`func (o *AppDto) HasSpecJsonSchema() bool`

HasSpecJsonSchema returns a boolean if a field has been set.

### GetLatestReleases

`func (o *AppDto) GetLatestReleases() []AppReleaseDto`

GetLatestReleases returns the LatestReleases field if non-nil, zero value otherwise.

### GetLatestReleasesOk

`func (o *AppDto) GetLatestReleasesOk() (*[]AppReleaseDto, bool)`

GetLatestReleasesOk returns a tuple with the LatestReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestReleases

`func (o *AppDto) SetLatestReleases(v []AppReleaseDto)`

SetLatestReleases sets LatestReleases field to given value.

### HasLatestReleases

`func (o *AppDto) HasLatestReleases() bool`

HasLatestReleases returns a boolean if a field has been set.

### GetCreator

`func (o *AppDto) GetCreator() IdentityUserDto`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *AppDto) GetCreatorOk() (*IdentityUserDto, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *AppDto) SetCreator(v IdentityUserDto)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *AppDto) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetFeatures

`func (o *AppDto) GetFeatures() []AppFeatureDto`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *AppDto) GetFeaturesOk() (*[]AppFeatureDto, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *AppDto) SetFeatures(v []AppFeatureDto)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *AppDto) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetSdks

`func (o *AppDto) GetSdks() []AppSdkDto`

GetSdks returns the Sdks field if non-nil, zero value otherwise.

### GetSdksOk

`func (o *AppDto) GetSdksOk() (*[]AppSdkDto, bool)`

GetSdksOk returns a tuple with the Sdks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdks

`func (o *AppDto) SetSdks(v []AppSdkDto)`

SetSdks sets Sdks field to given value.

### HasSdks

`func (o *AppDto) HasSdks() bool`

HasSdks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



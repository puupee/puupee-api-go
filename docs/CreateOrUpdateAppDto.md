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
**HomePage** | Pointer to **string** |  | [optional] 
**SortIndex** | Pointer to **int32** |  | [optional] 
**GitRepository** | Pointer to **string** |  | [optional] 
**GitRepositoryType** | Pointer to **string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**WebhookUrl** | Pointer to **string** |  | [optional] 
**BusinessDomain** | Pointer to **string** |  | [optional] 
**BusinessUrl** | Pointer to **string** |  | [optional] 
**SubscriptionEnabled** | Pointer to **bool** |  | [optional] 
**SubscriptionPlatforms** | Pointer to **string** |  | [optional] 
**SpecJsonSchema** | Pointer to **string** |  | [optional] 
**IsPublished** | Pointer to **bool** |  | [optional] 
**Features** | Pointer to [**[]AppFeatureDto**](AppFeatureDto.md) |  | [optional] 
**Sdks** | Pointer to [**[]AppSdkDto**](AppSdkDto.md) |  | [optional] 
**OpenClient** | Pointer to [**CreateOpenIddictApplicationDto**](CreateOpenIddictApplicationDto.md) |  | [optional] 

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

### GetHomePage

`func (o *CreateOrUpdateAppDto) GetHomePage() string`

GetHomePage returns the HomePage field if non-nil, zero value otherwise.

### GetHomePageOk

`func (o *CreateOrUpdateAppDto) GetHomePageOk() (*string, bool)`

GetHomePageOk returns a tuple with the HomePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePage

`func (o *CreateOrUpdateAppDto) SetHomePage(v string)`

SetHomePage sets HomePage field to given value.

### HasHomePage

`func (o *CreateOrUpdateAppDto) HasHomePage() bool`

HasHomePage returns a boolean if a field has been set.

### GetSortIndex

`func (o *CreateOrUpdateAppDto) GetSortIndex() int32`

GetSortIndex returns the SortIndex field if non-nil, zero value otherwise.

### GetSortIndexOk

`func (o *CreateOrUpdateAppDto) GetSortIndexOk() (*int32, bool)`

GetSortIndexOk returns a tuple with the SortIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortIndex

`func (o *CreateOrUpdateAppDto) SetSortIndex(v int32)`

SetSortIndex sets SortIndex field to given value.

### HasSortIndex

`func (o *CreateOrUpdateAppDto) HasSortIndex() bool`

HasSortIndex returns a boolean if a field has been set.

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

### GetIsEnabled

`func (o *CreateOrUpdateAppDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *CreateOrUpdateAppDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *CreateOrUpdateAppDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *CreateOrUpdateAppDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *CreateOrUpdateAppDto) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *CreateOrUpdateAppDto) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *CreateOrUpdateAppDto) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *CreateOrUpdateAppDto) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetBusinessDomain

`func (o *CreateOrUpdateAppDto) GetBusinessDomain() string`

GetBusinessDomain returns the BusinessDomain field if non-nil, zero value otherwise.

### GetBusinessDomainOk

`func (o *CreateOrUpdateAppDto) GetBusinessDomainOk() (*string, bool)`

GetBusinessDomainOk returns a tuple with the BusinessDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessDomain

`func (o *CreateOrUpdateAppDto) SetBusinessDomain(v string)`

SetBusinessDomain sets BusinessDomain field to given value.

### HasBusinessDomain

`func (o *CreateOrUpdateAppDto) HasBusinessDomain() bool`

HasBusinessDomain returns a boolean if a field has been set.

### GetBusinessUrl

`func (o *CreateOrUpdateAppDto) GetBusinessUrl() string`

GetBusinessUrl returns the BusinessUrl field if non-nil, zero value otherwise.

### GetBusinessUrlOk

`func (o *CreateOrUpdateAppDto) GetBusinessUrlOk() (*string, bool)`

GetBusinessUrlOk returns a tuple with the BusinessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUrl

`func (o *CreateOrUpdateAppDto) SetBusinessUrl(v string)`

SetBusinessUrl sets BusinessUrl field to given value.

### HasBusinessUrl

`func (o *CreateOrUpdateAppDto) HasBusinessUrl() bool`

HasBusinessUrl returns a boolean if a field has been set.

### GetSubscriptionEnabled

`func (o *CreateOrUpdateAppDto) GetSubscriptionEnabled() bool`

GetSubscriptionEnabled returns the SubscriptionEnabled field if non-nil, zero value otherwise.

### GetSubscriptionEnabledOk

`func (o *CreateOrUpdateAppDto) GetSubscriptionEnabledOk() (*bool, bool)`

GetSubscriptionEnabledOk returns a tuple with the SubscriptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionEnabled

`func (o *CreateOrUpdateAppDto) SetSubscriptionEnabled(v bool)`

SetSubscriptionEnabled sets SubscriptionEnabled field to given value.

### HasSubscriptionEnabled

`func (o *CreateOrUpdateAppDto) HasSubscriptionEnabled() bool`

HasSubscriptionEnabled returns a boolean if a field has been set.

### GetSubscriptionPlatforms

`func (o *CreateOrUpdateAppDto) GetSubscriptionPlatforms() string`

GetSubscriptionPlatforms returns the SubscriptionPlatforms field if non-nil, zero value otherwise.

### GetSubscriptionPlatformsOk

`func (o *CreateOrUpdateAppDto) GetSubscriptionPlatformsOk() (*string, bool)`

GetSubscriptionPlatformsOk returns a tuple with the SubscriptionPlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPlatforms

`func (o *CreateOrUpdateAppDto) SetSubscriptionPlatforms(v string)`

SetSubscriptionPlatforms sets SubscriptionPlatforms field to given value.

### HasSubscriptionPlatforms

`func (o *CreateOrUpdateAppDto) HasSubscriptionPlatforms() bool`

HasSubscriptionPlatforms returns a boolean if a field has been set.

### GetSpecJsonSchema

`func (o *CreateOrUpdateAppDto) GetSpecJsonSchema() string`

GetSpecJsonSchema returns the SpecJsonSchema field if non-nil, zero value otherwise.

### GetSpecJsonSchemaOk

`func (o *CreateOrUpdateAppDto) GetSpecJsonSchemaOk() (*string, bool)`

GetSpecJsonSchemaOk returns a tuple with the SpecJsonSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecJsonSchema

`func (o *CreateOrUpdateAppDto) SetSpecJsonSchema(v string)`

SetSpecJsonSchema sets SpecJsonSchema field to given value.

### HasSpecJsonSchema

`func (o *CreateOrUpdateAppDto) HasSpecJsonSchema() bool`

HasSpecJsonSchema returns a boolean if a field has been set.

### GetIsPublished

`func (o *CreateOrUpdateAppDto) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *CreateOrUpdateAppDto) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *CreateOrUpdateAppDto) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *CreateOrUpdateAppDto) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

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

### GetOpenClient

`func (o *CreateOrUpdateAppDto) GetOpenClient() CreateOpenIddictApplicationDto`

GetOpenClient returns the OpenClient field if non-nil, zero value otherwise.

### GetOpenClientOk

`func (o *CreateOrUpdateAppDto) GetOpenClientOk() (*CreateOpenIddictApplicationDto, bool)`

GetOpenClientOk returns a tuple with the OpenClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenClient

`func (o *CreateOrUpdateAppDto) SetOpenClient(v CreateOpenIddictApplicationDto)`

SetOpenClient sets OpenClient field to given value.

### HasOpenClient

`func (o *CreateOrUpdateAppDto) HasOpenClient() bool`

HasOpenClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



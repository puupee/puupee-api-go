/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
	"time"
)

// PublicAppDto struct for PublicAppDto
type PublicAppDto struct {
	Id *string `json:"id,omitempty"`
	CreationTime *time.Time `json:"creationTime,omitempty"`
	CreatorId *string `json:"creatorId,omitempty"`
	LastModificationTime *time.Time `json:"lastModificationTime,omitempty"`
	LastModifierId *string `json:"lastModifierId,omitempty"`
	IsDeleted *bool `json:"isDeleted,omitempty"`
	DeleterId *string `json:"deleterId,omitempty"`
	DeletionTime *time.Time `json:"deletionTime,omitempty"`
	Name *string `json:"name,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Framework *string `json:"framework,omitempty"`
	AppType *string `json:"appType,omitempty"`
	Description *string `json:"description,omitempty"`
	Icon *string `json:"icon,omitempty"`
	HomePage *string `json:"homePage,omitempty"`
	SortIndex *int32 `json:"sortIndex,omitempty"`
	GitRepository *string `json:"gitRepository,omitempty"`
	GitRepositoryType *string `json:"gitRepositoryType,omitempty"`
	IsEnabled *bool `json:"isEnabled,omitempty"`
	IsPublished *bool `json:"isPublished,omitempty"`
	WebhookUrl *string `json:"webhookUrl,omitempty"`
	BusinessDomain *string `json:"businessDomain,omitempty"`
	BusinessUrl *string `json:"businessUrl,omitempty"`
	LatestReleases []AppReleaseDto `json:"latestReleases,omitempty"`
	Creator *IdentityUserDto `json:"creator,omitempty"`
	Features []AppFeatureDto `json:"features,omitempty"`
	Sdks []AppSdkDto `json:"sdks,omitempty"`
	Subscribed *bool `json:"subscribed,omitempty"`
}

// NewPublicAppDto instantiates a new PublicAppDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicAppDto() *PublicAppDto {
	this := PublicAppDto{}
	return &this
}

// NewPublicAppDtoWithDefaults instantiates a new PublicAppDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicAppDtoWithDefaults() *PublicAppDto {
	this := PublicAppDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PublicAppDto) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PublicAppDto) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PublicAppDto) SetId(v string) {
	o.Id = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *PublicAppDto) GetCreationTime() time.Time {
	if o == nil || isNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreationTime) {
    return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *PublicAppDto) HasCreationTime() bool {
	if o != nil && !isNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *PublicAppDto) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *PublicAppDto) GetCreatorId() string {
	if o == nil || isNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetCreatorIdOk() (*string, bool) {
	if o == nil || isNil(o.CreatorId) {
    return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *PublicAppDto) HasCreatorId() bool {
	if o != nil && !isNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *PublicAppDto) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetLastModificationTime returns the LastModificationTime field value if set, zero value otherwise.
func (o *PublicAppDto) GetLastModificationTime() time.Time {
	if o == nil || isNil(o.LastModificationTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModificationTime
}

// GetLastModificationTimeOk returns a tuple with the LastModificationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetLastModificationTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastModificationTime) {
    return nil, false
	}
	return o.LastModificationTime, true
}

// HasLastModificationTime returns a boolean if a field has been set.
func (o *PublicAppDto) HasLastModificationTime() bool {
	if o != nil && !isNil(o.LastModificationTime) {
		return true
	}

	return false
}

// SetLastModificationTime gets a reference to the given time.Time and assigns it to the LastModificationTime field.
func (o *PublicAppDto) SetLastModificationTime(v time.Time) {
	o.LastModificationTime = &v
}

// GetLastModifierId returns the LastModifierId field value if set, zero value otherwise.
func (o *PublicAppDto) GetLastModifierId() string {
	if o == nil || isNil(o.LastModifierId) {
		var ret string
		return ret
	}
	return *o.LastModifierId
}

// GetLastModifierIdOk returns a tuple with the LastModifierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetLastModifierIdOk() (*string, bool) {
	if o == nil || isNil(o.LastModifierId) {
    return nil, false
	}
	return o.LastModifierId, true
}

// HasLastModifierId returns a boolean if a field has been set.
func (o *PublicAppDto) HasLastModifierId() bool {
	if o != nil && !isNil(o.LastModifierId) {
		return true
	}

	return false
}

// SetLastModifierId gets a reference to the given string and assigns it to the LastModifierId field.
func (o *PublicAppDto) SetLastModifierId(v string) {
	o.LastModifierId = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *PublicAppDto) GetIsDeleted() bool {
	if o == nil || isNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetIsDeletedOk() (*bool, bool) {
	if o == nil || isNil(o.IsDeleted) {
    return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *PublicAppDto) HasIsDeleted() bool {
	if o != nil && !isNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *PublicAppDto) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetDeleterId returns the DeleterId field value if set, zero value otherwise.
func (o *PublicAppDto) GetDeleterId() string {
	if o == nil || isNil(o.DeleterId) {
		var ret string
		return ret
	}
	return *o.DeleterId
}

// GetDeleterIdOk returns a tuple with the DeleterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetDeleterIdOk() (*string, bool) {
	if o == nil || isNil(o.DeleterId) {
    return nil, false
	}
	return o.DeleterId, true
}

// HasDeleterId returns a boolean if a field has been set.
func (o *PublicAppDto) HasDeleterId() bool {
	if o != nil && !isNil(o.DeleterId) {
		return true
	}

	return false
}

// SetDeleterId gets a reference to the given string and assigns it to the DeleterId field.
func (o *PublicAppDto) SetDeleterId(v string) {
	o.DeleterId = &v
}

// GetDeletionTime returns the DeletionTime field value if set, zero value otherwise.
func (o *PublicAppDto) GetDeletionTime() time.Time {
	if o == nil || isNil(o.DeletionTime) {
		var ret time.Time
		return ret
	}
	return *o.DeletionTime
}

// GetDeletionTimeOk returns a tuple with the DeletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetDeletionTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.DeletionTime) {
    return nil, false
	}
	return o.DeletionTime, true
}

// HasDeletionTime returns a boolean if a field has been set.
func (o *PublicAppDto) HasDeletionTime() bool {
	if o != nil && !isNil(o.DeletionTime) {
		return true
	}

	return false
}

// SetDeletionTime gets a reference to the given time.Time and assigns it to the DeletionTime field.
func (o *PublicAppDto) SetDeletionTime(v time.Time) {
	o.DeletionTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PublicAppDto) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PublicAppDto) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PublicAppDto) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PublicAppDto) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PublicAppDto) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PublicAppDto) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetFramework returns the Framework field value if set, zero value otherwise.
func (o *PublicAppDto) GetFramework() string {
	if o == nil || isNil(o.Framework) {
		var ret string
		return ret
	}
	return *o.Framework
}

// GetFrameworkOk returns a tuple with the Framework field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetFrameworkOk() (*string, bool) {
	if o == nil || isNil(o.Framework) {
    return nil, false
	}
	return o.Framework, true
}

// HasFramework returns a boolean if a field has been set.
func (o *PublicAppDto) HasFramework() bool {
	if o != nil && !isNil(o.Framework) {
		return true
	}

	return false
}

// SetFramework gets a reference to the given string and assigns it to the Framework field.
func (o *PublicAppDto) SetFramework(v string) {
	o.Framework = &v
}

// GetAppType returns the AppType field value if set, zero value otherwise.
func (o *PublicAppDto) GetAppType() string {
	if o == nil || isNil(o.AppType) {
		var ret string
		return ret
	}
	return *o.AppType
}

// GetAppTypeOk returns a tuple with the AppType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetAppTypeOk() (*string, bool) {
	if o == nil || isNil(o.AppType) {
    return nil, false
	}
	return o.AppType, true
}

// HasAppType returns a boolean if a field has been set.
func (o *PublicAppDto) HasAppType() bool {
	if o != nil && !isNil(o.AppType) {
		return true
	}

	return false
}

// SetAppType gets a reference to the given string and assigns it to the AppType field.
func (o *PublicAppDto) SetAppType(v string) {
	o.AppType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PublicAppDto) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PublicAppDto) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PublicAppDto) SetDescription(v string) {
	o.Description = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *PublicAppDto) GetIcon() string {
	if o == nil || isNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetIconOk() (*string, bool) {
	if o == nil || isNil(o.Icon) {
    return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *PublicAppDto) HasIcon() bool {
	if o != nil && !isNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *PublicAppDto) SetIcon(v string) {
	o.Icon = &v
}

// GetHomePage returns the HomePage field value if set, zero value otherwise.
func (o *PublicAppDto) GetHomePage() string {
	if o == nil || isNil(o.HomePage) {
		var ret string
		return ret
	}
	return *o.HomePage
}

// GetHomePageOk returns a tuple with the HomePage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetHomePageOk() (*string, bool) {
	if o == nil || isNil(o.HomePage) {
    return nil, false
	}
	return o.HomePage, true
}

// HasHomePage returns a boolean if a field has been set.
func (o *PublicAppDto) HasHomePage() bool {
	if o != nil && !isNil(o.HomePage) {
		return true
	}

	return false
}

// SetHomePage gets a reference to the given string and assigns it to the HomePage field.
func (o *PublicAppDto) SetHomePage(v string) {
	o.HomePage = &v
}

// GetSortIndex returns the SortIndex field value if set, zero value otherwise.
func (o *PublicAppDto) GetSortIndex() int32 {
	if o == nil || isNil(o.SortIndex) {
		var ret int32
		return ret
	}
	return *o.SortIndex
}

// GetSortIndexOk returns a tuple with the SortIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetSortIndexOk() (*int32, bool) {
	if o == nil || isNil(o.SortIndex) {
    return nil, false
	}
	return o.SortIndex, true
}

// HasSortIndex returns a boolean if a field has been set.
func (o *PublicAppDto) HasSortIndex() bool {
	if o != nil && !isNil(o.SortIndex) {
		return true
	}

	return false
}

// SetSortIndex gets a reference to the given int32 and assigns it to the SortIndex field.
func (o *PublicAppDto) SetSortIndex(v int32) {
	o.SortIndex = &v
}

// GetGitRepository returns the GitRepository field value if set, zero value otherwise.
func (o *PublicAppDto) GetGitRepository() string {
	if o == nil || isNil(o.GitRepository) {
		var ret string
		return ret
	}
	return *o.GitRepository
}

// GetGitRepositoryOk returns a tuple with the GitRepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetGitRepositoryOk() (*string, bool) {
	if o == nil || isNil(o.GitRepository) {
    return nil, false
	}
	return o.GitRepository, true
}

// HasGitRepository returns a boolean if a field has been set.
func (o *PublicAppDto) HasGitRepository() bool {
	if o != nil && !isNil(o.GitRepository) {
		return true
	}

	return false
}

// SetGitRepository gets a reference to the given string and assigns it to the GitRepository field.
func (o *PublicAppDto) SetGitRepository(v string) {
	o.GitRepository = &v
}

// GetGitRepositoryType returns the GitRepositoryType field value if set, zero value otherwise.
func (o *PublicAppDto) GetGitRepositoryType() string {
	if o == nil || isNil(o.GitRepositoryType) {
		var ret string
		return ret
	}
	return *o.GitRepositoryType
}

// GetGitRepositoryTypeOk returns a tuple with the GitRepositoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetGitRepositoryTypeOk() (*string, bool) {
	if o == nil || isNil(o.GitRepositoryType) {
    return nil, false
	}
	return o.GitRepositoryType, true
}

// HasGitRepositoryType returns a boolean if a field has been set.
func (o *PublicAppDto) HasGitRepositoryType() bool {
	if o != nil && !isNil(o.GitRepositoryType) {
		return true
	}

	return false
}

// SetGitRepositoryType gets a reference to the given string and assigns it to the GitRepositoryType field.
func (o *PublicAppDto) SetGitRepositoryType(v string) {
	o.GitRepositoryType = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *PublicAppDto) GetIsEnabled() bool {
	if o == nil || isNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetIsEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsEnabled) {
    return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *PublicAppDto) HasIsEnabled() bool {
	if o != nil && !isNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *PublicAppDto) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetIsPublished returns the IsPublished field value if set, zero value otherwise.
func (o *PublicAppDto) GetIsPublished() bool {
	if o == nil || isNil(o.IsPublished) {
		var ret bool
		return ret
	}
	return *o.IsPublished
}

// GetIsPublishedOk returns a tuple with the IsPublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetIsPublishedOk() (*bool, bool) {
	if o == nil || isNil(o.IsPublished) {
    return nil, false
	}
	return o.IsPublished, true
}

// HasIsPublished returns a boolean if a field has been set.
func (o *PublicAppDto) HasIsPublished() bool {
	if o != nil && !isNil(o.IsPublished) {
		return true
	}

	return false
}

// SetIsPublished gets a reference to the given bool and assigns it to the IsPublished field.
func (o *PublicAppDto) SetIsPublished(v bool) {
	o.IsPublished = &v
}

// GetWebhookUrl returns the WebhookUrl field value if set, zero value otherwise.
func (o *PublicAppDto) GetWebhookUrl() string {
	if o == nil || isNil(o.WebhookUrl) {
		var ret string
		return ret
	}
	return *o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetWebhookUrlOk() (*string, bool) {
	if o == nil || isNil(o.WebhookUrl) {
    return nil, false
	}
	return o.WebhookUrl, true
}

// HasWebhookUrl returns a boolean if a field has been set.
func (o *PublicAppDto) HasWebhookUrl() bool {
	if o != nil && !isNil(o.WebhookUrl) {
		return true
	}

	return false
}

// SetWebhookUrl gets a reference to the given string and assigns it to the WebhookUrl field.
func (o *PublicAppDto) SetWebhookUrl(v string) {
	o.WebhookUrl = &v
}

// GetBusinessDomain returns the BusinessDomain field value if set, zero value otherwise.
func (o *PublicAppDto) GetBusinessDomain() string {
	if o == nil || isNil(o.BusinessDomain) {
		var ret string
		return ret
	}
	return *o.BusinessDomain
}

// GetBusinessDomainOk returns a tuple with the BusinessDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetBusinessDomainOk() (*string, bool) {
	if o == nil || isNil(o.BusinessDomain) {
    return nil, false
	}
	return o.BusinessDomain, true
}

// HasBusinessDomain returns a boolean if a field has been set.
func (o *PublicAppDto) HasBusinessDomain() bool {
	if o != nil && !isNil(o.BusinessDomain) {
		return true
	}

	return false
}

// SetBusinessDomain gets a reference to the given string and assigns it to the BusinessDomain field.
func (o *PublicAppDto) SetBusinessDomain(v string) {
	o.BusinessDomain = &v
}

// GetBusinessUrl returns the BusinessUrl field value if set, zero value otherwise.
func (o *PublicAppDto) GetBusinessUrl() string {
	if o == nil || isNil(o.BusinessUrl) {
		var ret string
		return ret
	}
	return *o.BusinessUrl
}

// GetBusinessUrlOk returns a tuple with the BusinessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetBusinessUrlOk() (*string, bool) {
	if o == nil || isNil(o.BusinessUrl) {
    return nil, false
	}
	return o.BusinessUrl, true
}

// HasBusinessUrl returns a boolean if a field has been set.
func (o *PublicAppDto) HasBusinessUrl() bool {
	if o != nil && !isNil(o.BusinessUrl) {
		return true
	}

	return false
}

// SetBusinessUrl gets a reference to the given string and assigns it to the BusinessUrl field.
func (o *PublicAppDto) SetBusinessUrl(v string) {
	o.BusinessUrl = &v
}

// GetLatestReleases returns the LatestReleases field value if set, zero value otherwise.
func (o *PublicAppDto) GetLatestReleases() []AppReleaseDto {
	if o == nil || isNil(o.LatestReleases) {
		var ret []AppReleaseDto
		return ret
	}
	return o.LatestReleases
}

// GetLatestReleasesOk returns a tuple with the LatestReleases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetLatestReleasesOk() ([]AppReleaseDto, bool) {
	if o == nil || isNil(o.LatestReleases) {
    return nil, false
	}
	return o.LatestReleases, true
}

// HasLatestReleases returns a boolean if a field has been set.
func (o *PublicAppDto) HasLatestReleases() bool {
	if o != nil && !isNil(o.LatestReleases) {
		return true
	}

	return false
}

// SetLatestReleases gets a reference to the given []AppReleaseDto and assigns it to the LatestReleases field.
func (o *PublicAppDto) SetLatestReleases(v []AppReleaseDto) {
	o.LatestReleases = v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *PublicAppDto) GetCreator() IdentityUserDto {
	if o == nil || isNil(o.Creator) {
		var ret IdentityUserDto
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetCreatorOk() (*IdentityUserDto, bool) {
	if o == nil || isNil(o.Creator) {
    return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *PublicAppDto) HasCreator() bool {
	if o != nil && !isNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given IdentityUserDto and assigns it to the Creator field.
func (o *PublicAppDto) SetCreator(v IdentityUserDto) {
	o.Creator = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *PublicAppDto) GetFeatures() []AppFeatureDto {
	if o == nil || isNil(o.Features) {
		var ret []AppFeatureDto
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetFeaturesOk() ([]AppFeatureDto, bool) {
	if o == nil || isNil(o.Features) {
    return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *PublicAppDto) HasFeatures() bool {
	if o != nil && !isNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []AppFeatureDto and assigns it to the Features field.
func (o *PublicAppDto) SetFeatures(v []AppFeatureDto) {
	o.Features = v
}

// GetSdks returns the Sdks field value if set, zero value otherwise.
func (o *PublicAppDto) GetSdks() []AppSdkDto {
	if o == nil || isNil(o.Sdks) {
		var ret []AppSdkDto
		return ret
	}
	return o.Sdks
}

// GetSdksOk returns a tuple with the Sdks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetSdksOk() ([]AppSdkDto, bool) {
	if o == nil || isNil(o.Sdks) {
    return nil, false
	}
	return o.Sdks, true
}

// HasSdks returns a boolean if a field has been set.
func (o *PublicAppDto) HasSdks() bool {
	if o != nil && !isNil(o.Sdks) {
		return true
	}

	return false
}

// SetSdks gets a reference to the given []AppSdkDto and assigns it to the Sdks field.
func (o *PublicAppDto) SetSdks(v []AppSdkDto) {
	o.Sdks = v
}

// GetSubscribed returns the Subscribed field value if set, zero value otherwise.
func (o *PublicAppDto) GetSubscribed() bool {
	if o == nil || isNil(o.Subscribed) {
		var ret bool
		return ret
	}
	return *o.Subscribed
}

// GetSubscribedOk returns a tuple with the Subscribed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAppDto) GetSubscribedOk() (*bool, bool) {
	if o == nil || isNil(o.Subscribed) {
    return nil, false
	}
	return o.Subscribed, true
}

// HasSubscribed returns a boolean if a field has been set.
func (o *PublicAppDto) HasSubscribed() bool {
	if o != nil && !isNil(o.Subscribed) {
		return true
	}

	return false
}

// SetSubscribed gets a reference to the given bool and assigns it to the Subscribed field.
func (o *PublicAppDto) SetSubscribed(v bool) {
	o.Subscribed = &v
}

func (o PublicAppDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.CreationTime) {
		toSerialize["creationTime"] = o.CreationTime
	}
	if !isNil(o.CreatorId) {
		toSerialize["creatorId"] = o.CreatorId
	}
	if !isNil(o.LastModificationTime) {
		toSerialize["lastModificationTime"] = o.LastModificationTime
	}
	if !isNil(o.LastModifierId) {
		toSerialize["lastModifierId"] = o.LastModifierId
	}
	if !isNil(o.IsDeleted) {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	if !isNil(o.DeleterId) {
		toSerialize["deleterId"] = o.DeleterId
	}
	if !isNil(o.DeletionTime) {
		toSerialize["deletionTime"] = o.DeletionTime
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.Framework) {
		toSerialize["framework"] = o.Framework
	}
	if !isNil(o.AppType) {
		toSerialize["appType"] = o.AppType
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !isNil(o.HomePage) {
		toSerialize["homePage"] = o.HomePage
	}
	if !isNil(o.SortIndex) {
		toSerialize["sortIndex"] = o.SortIndex
	}
	if !isNil(o.GitRepository) {
		toSerialize["gitRepository"] = o.GitRepository
	}
	if !isNil(o.GitRepositoryType) {
		toSerialize["gitRepositoryType"] = o.GitRepositoryType
	}
	if !isNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if !isNil(o.IsPublished) {
		toSerialize["isPublished"] = o.IsPublished
	}
	if !isNil(o.WebhookUrl) {
		toSerialize["webhookUrl"] = o.WebhookUrl
	}
	if !isNil(o.BusinessDomain) {
		toSerialize["businessDomain"] = o.BusinessDomain
	}
	if !isNil(o.BusinessUrl) {
		toSerialize["businessUrl"] = o.BusinessUrl
	}
	if !isNil(o.LatestReleases) {
		toSerialize["latestReleases"] = o.LatestReleases
	}
	if !isNil(o.Creator) {
		toSerialize["creator"] = o.Creator
	}
	if !isNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !isNil(o.Sdks) {
		toSerialize["sdks"] = o.Sdks
	}
	if !isNil(o.Subscribed) {
		toSerialize["subscribed"] = o.Subscribed
	}
	return json.Marshal(toSerialize)
}

type NullablePublicAppDto struct {
	value *PublicAppDto
	isSet bool
}

func (v NullablePublicAppDto) Get() *PublicAppDto {
	return v.value
}

func (v *NullablePublicAppDto) Set(val *PublicAppDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicAppDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicAppDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicAppDto(val *PublicAppDto) *NullablePublicAppDto {
	return &NullablePublicAppDto{value: val, isSet: true}
}

func (v NullablePublicAppDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicAppDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



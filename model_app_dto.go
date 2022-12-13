/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
	"time"
)

// AppDto struct for AppDto
type AppDto struct {
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
	GitRepository *string `json:"gitRepository,omitempty"`
	GitRepositoryType *string `json:"gitRepositoryType,omitempty"`
	LatestReleases []AppReleaseDto `json:"latestReleases,omitempty"`
	Creator *IdentityUserDto `json:"creator,omitempty"`
	Features []AppFeatureDto `json:"features,omitempty"`
	Sdks []AppSdkDto `json:"sdks,omitempty"`
}

// NewAppDto instantiates a new AppDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppDto() *AppDto {
	this := AppDto{}
	return &this
}

// NewAppDtoWithDefaults instantiates a new AppDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppDtoWithDefaults() *AppDto {
	this := AppDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppDto) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDto) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppDto) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AppDto) SetId(v string) {
	o.Id = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *AppDto) GetCreationTime() time.Time {
	if o == nil || isNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDto) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreationTime) {
    return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *AppDto) HasCreationTime() bool {
	if o != nil && !isNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *AppDto) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *AppDto) GetCreatorId() string {
	if o == nil || isNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDto) GetCreatorIdOk() (*string, bool) {
	if o == nil || isNil(o.CreatorId) {
    return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *AppDto) HasCreatorId() bool {
	if o != nil && !isNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *AppDto) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetLastModificationTime returns the LastModificationTime field value if set, zero value otherwise.
func (o *AppDto) GetLastModificationTime() time.Time {
	if o == nil || isNil(o.LastModificationTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModificationTime
}

// GetLastModificationTimeOk returns a tuple with the LastModificationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDto) GetLastModificationTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastModificationTime) {
    return nil, false
	}
	return o.LastModificationTime, true
}

// HasLastModificationTime returns a boolean if a field has been set.
func (o *AppDto) HasLastModificationTime() bool {
	if o != nil && !isNil(o.LastModificationTime) {
		return true
	}

	return false
}

// SetLastModificationTime gets a reference to the given time.Time and assigns it to the LastModificationTime field.
func (o *AppDto) SetLastModificationTime(v time.Time) {
	o.LastModificationTime = &v
}

// GetLastModifierId returns the LastModifierId field value if set, zero value otherwise.
func (o *AppDto) GetLastModifierId() string {
	if o == nil || isNil(o.LastModifierId) {
		var ret string
		return ret
	}
	return *o.LastModifierId
}

// GetLastModifierIdOk returns a tuple with the LastModifierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDto) GetLastModifierIdOk() (*string, bool) {
	if o == nil || isNil(o.LastModifierId) {
    return nil, false
	}
	return o.LastModifierId, true
}

// HasLastModifierId returns a boolean if a field has been set.
func (o *AppDto) HasLastModifierId() bool {
	if o != nil && !isNil(o.LastModifierId) {
		return true
	}

	return false
}

// SetLastModifierId gets a reference to the given string and assigns it to the LastModifierId field.
func (o *AppDto) SetLastModifierId(v string) {
	o.LastModifierId = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *AppDto) GetIsDeleted() bool {
	if o == nil || isNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDto) GetIsDeletedOk() (*bool, bool) {
	if o == nil || isNil(o.IsDeleted) {
    return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *AppDto) HasIsDeleted() bool {
	if o != nil && !isNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *AppDto) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetDeleterId returns the DeleterId field value if set, zero value otherwise.
func (o *AppDto) GetDeleterId() string {
	if o == nil || isNil(o.DeleterId) {
		var ret string
		return ret
	}
	return *o.DeleterId
}

// GetDeleterIdOk returns a tuple with the DeleterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDto) GetDeleterIdOk() (*string, bool) {
	if o == nil || isNil(o.DeleterId) {
    return nil, false
	}
	return o.DeleterId, true
}

// HasDeleterId returns a boolean if a field has been set.
func (o *AppDto) HasDeleterId() bool {
	if o != nil && !isNil(o.DeleterId) {
		return true
	}

	return false
}

// SetDeleterId gets a reference to the given string and assigns it to the DeleterId field.
func (o *AppDto) SetDeleterId(v string) {
	o.DeleterId = &v
}

// GetDeletionTime returns the DeletionTime field value if set, zero value otherwise.
func (o *AppDto) GetDeletionTime() time.Time {
	if o == nil || isNil(o.DeletionTime) {
		var ret time.Time
		return ret
	}
	return *o.DeletionTime
}

// GetDeletionTimeOk returns a tuple with the DeletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDto) GetDeletionTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.DeletionTime) {
    return nil, false
	}
	return o.DeletionTime, true
}

// HasDeletionTime returns a boolean if a field has been set.
func (o *AppDto) HasDeletionTime() bool {
	if o != nil && !isNil(o.DeletionTime) {
		return true
	}

	return false
}

// SetDeletionTime gets a reference to the given time.Time and assigns it to the DeletionTime field.
func (o *AppDto) SetDeletionTime(v time.Time) {
	o.DeletionTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppDto) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDto) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppDto) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppDto) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AppDto) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDto) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AppDto) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AppDto) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetFramework returns the Framework field value if set, zero value otherwise.
func (o *AppDto) GetFramework() string {
	if o == nil || isNil(o.Framework) {
		var ret string
		return ret
	}
	return *o.Framework
}

// GetFrameworkOk returns a tuple with the Framework field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDto) GetFrameworkOk() (*string, bool) {
	if o == nil || isNil(o.Framework) {
    return nil, false
	}
	return o.Framework, true
}

// HasFramework returns a boolean if a field has been set.
func (o *AppDto) HasFramework() bool {
	if o != nil && !isNil(o.Framework) {
		return true
	}

	return false
}

// SetFramework gets a reference to the given string and assigns it to the Framework field.
func (o *AppDto) SetFramework(v string) {
	o.Framework = &v
}

// GetAppType returns the AppType field value if set, zero value otherwise.
func (o *AppDto) GetAppType() string {
	if o == nil || isNil(o.AppType) {
		var ret string
		return ret
	}
	return *o.AppType
}

// GetAppTypeOk returns a tuple with the AppType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDto) GetAppTypeOk() (*string, bool) {
	if o == nil || isNil(o.AppType) {
    return nil, false
	}
	return o.AppType, true
}

// HasAppType returns a boolean if a field has been set.
func (o *AppDto) HasAppType() bool {
	if o != nil && !isNil(o.AppType) {
		return true
	}

	return false
}

// SetAppType gets a reference to the given string and assigns it to the AppType field.
func (o *AppDto) SetAppType(v string) {
	o.AppType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AppDto) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDto) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AppDto) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AppDto) SetDescription(v string) {
	o.Description = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *AppDto) GetIcon() string {
	if o == nil || isNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDto) GetIconOk() (*string, bool) {
	if o == nil || isNil(o.Icon) {
    return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *AppDto) HasIcon() bool {
	if o != nil && !isNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *AppDto) SetIcon(v string) {
	o.Icon = &v
}

// GetGitRepository returns the GitRepository field value if set, zero value otherwise.
func (o *AppDto) GetGitRepository() string {
	if o == nil || isNil(o.GitRepository) {
		var ret string
		return ret
	}
	return *o.GitRepository
}

// GetGitRepositoryOk returns a tuple with the GitRepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDto) GetGitRepositoryOk() (*string, bool) {
	if o == nil || isNil(o.GitRepository) {
    return nil, false
	}
	return o.GitRepository, true
}

// HasGitRepository returns a boolean if a field has been set.
func (o *AppDto) HasGitRepository() bool {
	if o != nil && !isNil(o.GitRepository) {
		return true
	}

	return false
}

// SetGitRepository gets a reference to the given string and assigns it to the GitRepository field.
func (o *AppDto) SetGitRepository(v string) {
	o.GitRepository = &v
}

// GetGitRepositoryType returns the GitRepositoryType field value if set, zero value otherwise.
func (o *AppDto) GetGitRepositoryType() string {
	if o == nil || isNil(o.GitRepositoryType) {
		var ret string
		return ret
	}
	return *o.GitRepositoryType
}

// GetGitRepositoryTypeOk returns a tuple with the GitRepositoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDto) GetGitRepositoryTypeOk() (*string, bool) {
	if o == nil || isNil(o.GitRepositoryType) {
    return nil, false
	}
	return o.GitRepositoryType, true
}

// HasGitRepositoryType returns a boolean if a field has been set.
func (o *AppDto) HasGitRepositoryType() bool {
	if o != nil && !isNil(o.GitRepositoryType) {
		return true
	}

	return false
}

// SetGitRepositoryType gets a reference to the given string and assigns it to the GitRepositoryType field.
func (o *AppDto) SetGitRepositoryType(v string) {
	o.GitRepositoryType = &v
}

// GetLatestReleases returns the LatestReleases field value if set, zero value otherwise.
func (o *AppDto) GetLatestReleases() []AppReleaseDto {
	if o == nil || isNil(o.LatestReleases) {
		var ret []AppReleaseDto
		return ret
	}
	return o.LatestReleases
}

// GetLatestReleasesOk returns a tuple with the LatestReleases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDto) GetLatestReleasesOk() ([]AppReleaseDto, bool) {
	if o == nil || isNil(o.LatestReleases) {
    return nil, false
	}
	return o.LatestReleases, true
}

// HasLatestReleases returns a boolean if a field has been set.
func (o *AppDto) HasLatestReleases() bool {
	if o != nil && !isNil(o.LatestReleases) {
		return true
	}

	return false
}

// SetLatestReleases gets a reference to the given []AppReleaseDto and assigns it to the LatestReleases field.
func (o *AppDto) SetLatestReleases(v []AppReleaseDto) {
	o.LatestReleases = v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *AppDto) GetCreator() IdentityUserDto {
	if o == nil || isNil(o.Creator) {
		var ret IdentityUserDto
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDto) GetCreatorOk() (*IdentityUserDto, bool) {
	if o == nil || isNil(o.Creator) {
    return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *AppDto) HasCreator() bool {
	if o != nil && !isNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given IdentityUserDto and assigns it to the Creator field.
func (o *AppDto) SetCreator(v IdentityUserDto) {
	o.Creator = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *AppDto) GetFeatures() []AppFeatureDto {
	if o == nil || isNil(o.Features) {
		var ret []AppFeatureDto
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDto) GetFeaturesOk() ([]AppFeatureDto, bool) {
	if o == nil || isNil(o.Features) {
    return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *AppDto) HasFeatures() bool {
	if o != nil && !isNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []AppFeatureDto and assigns it to the Features field.
func (o *AppDto) SetFeatures(v []AppFeatureDto) {
	o.Features = v
}

// GetSdks returns the Sdks field value if set, zero value otherwise.
func (o *AppDto) GetSdks() []AppSdkDto {
	if o == nil || isNil(o.Sdks) {
		var ret []AppSdkDto
		return ret
	}
	return o.Sdks
}

// GetSdksOk returns a tuple with the Sdks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDto) GetSdksOk() ([]AppSdkDto, bool) {
	if o == nil || isNil(o.Sdks) {
    return nil, false
	}
	return o.Sdks, true
}

// HasSdks returns a boolean if a field has been set.
func (o *AppDto) HasSdks() bool {
	if o != nil && !isNil(o.Sdks) {
		return true
	}

	return false
}

// SetSdks gets a reference to the given []AppSdkDto and assigns it to the Sdks field.
func (o *AppDto) SetSdks(v []AppSdkDto) {
	o.Sdks = v
}

func (o AppDto) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.GitRepository) {
		toSerialize["gitRepository"] = o.GitRepository
	}
	if !isNil(o.GitRepositoryType) {
		toSerialize["gitRepositoryType"] = o.GitRepositoryType
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
	return json.Marshal(toSerialize)
}

type NullableAppDto struct {
	value *AppDto
	isSet bool
}

func (v NullableAppDto) Get() *AppDto {
	return v.value
}

func (v *NullableAppDto) Set(val *AppDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAppDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAppDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppDto(val *AppDto) *NullableAppDto {
	return &NullableAppDto{value: val, isSet: true}
}

func (v NullableAppDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



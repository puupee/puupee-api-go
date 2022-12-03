/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
)

// CreateOrUpdateAppDto struct for CreateOrUpdateAppDto
type CreateOrUpdateAppDto struct {
	AppId *string `json:"appId,omitempty"`
	Name *string `json:"name,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Framework *string `json:"framework,omitempty"`
	AppType *string `json:"appType,omitempty"`
	Description *string `json:"description,omitempty"`
	Icon *string `json:"icon,omitempty"`
	GitRepository *string `json:"gitRepository,omitempty"`
	GitRepositoryType *string `json:"gitRepositoryType,omitempty"`
}

// NewCreateOrUpdateAppDto instantiates a new CreateOrUpdateAppDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateAppDto() *CreateOrUpdateAppDto {
	this := CreateOrUpdateAppDto{}
	return &this
}

// NewCreateOrUpdateAppDtoWithDefaults instantiates a new CreateOrUpdateAppDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateAppDtoWithDefaults() *CreateOrUpdateAppDto {
	this := CreateOrUpdateAppDto{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
    return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *CreateOrUpdateAppDto) SetAppId(v string) {
	o.AppId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateOrUpdateAppDto) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CreateOrUpdateAppDto) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetFramework returns the Framework field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetFramework() string {
	if o == nil || isNil(o.Framework) {
		var ret string
		return ret
	}
	return *o.Framework
}

// GetFrameworkOk returns a tuple with the Framework field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetFrameworkOk() (*string, bool) {
	if o == nil || isNil(o.Framework) {
    return nil, false
	}
	return o.Framework, true
}

// HasFramework returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasFramework() bool {
	if o != nil && !isNil(o.Framework) {
		return true
	}

	return false
}

// SetFramework gets a reference to the given string and assigns it to the Framework field.
func (o *CreateOrUpdateAppDto) SetFramework(v string) {
	o.Framework = &v
}

// GetAppType returns the AppType field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetAppType() string {
	if o == nil || isNil(o.AppType) {
		var ret string
		return ret
	}
	return *o.AppType
}

// GetAppTypeOk returns a tuple with the AppType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetAppTypeOk() (*string, bool) {
	if o == nil || isNil(o.AppType) {
    return nil, false
	}
	return o.AppType, true
}

// HasAppType returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasAppType() bool {
	if o != nil && !isNil(o.AppType) {
		return true
	}

	return false
}

// SetAppType gets a reference to the given string and assigns it to the AppType field.
func (o *CreateOrUpdateAppDto) SetAppType(v string) {
	o.AppType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateOrUpdateAppDto) SetDescription(v string) {
	o.Description = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetIcon() string {
	if o == nil || isNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetIconOk() (*string, bool) {
	if o == nil || isNil(o.Icon) {
    return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasIcon() bool {
	if o != nil && !isNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *CreateOrUpdateAppDto) SetIcon(v string) {
	o.Icon = &v
}

// GetGitRepository returns the GitRepository field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetGitRepository() string {
	if o == nil || isNil(o.GitRepository) {
		var ret string
		return ret
	}
	return *o.GitRepository
}

// GetGitRepositoryOk returns a tuple with the GitRepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetGitRepositoryOk() (*string, bool) {
	if o == nil || isNil(o.GitRepository) {
    return nil, false
	}
	return o.GitRepository, true
}

// HasGitRepository returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasGitRepository() bool {
	if o != nil && !isNil(o.GitRepository) {
		return true
	}

	return false
}

// SetGitRepository gets a reference to the given string and assigns it to the GitRepository field.
func (o *CreateOrUpdateAppDto) SetGitRepository(v string) {
	o.GitRepository = &v
}

// GetGitRepositoryType returns the GitRepositoryType field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetGitRepositoryType() string {
	if o == nil || isNil(o.GitRepositoryType) {
		var ret string
		return ret
	}
	return *o.GitRepositoryType
}

// GetGitRepositoryTypeOk returns a tuple with the GitRepositoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetGitRepositoryTypeOk() (*string, bool) {
	if o == nil || isNil(o.GitRepositoryType) {
    return nil, false
	}
	return o.GitRepositoryType, true
}

// HasGitRepositoryType returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasGitRepositoryType() bool {
	if o != nil && !isNil(o.GitRepositoryType) {
		return true
	}

	return false
}

// SetGitRepositoryType gets a reference to the given string and assigns it to the GitRepositoryType field.
func (o *CreateOrUpdateAppDto) SetGitRepositoryType(v string) {
	o.GitRepositoryType = &v
}

func (o CreateOrUpdateAppDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AppId) {
		toSerialize["appId"] = o.AppId
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
	return json.Marshal(toSerialize)
}

type NullableCreateOrUpdateAppDto struct {
	value *CreateOrUpdateAppDto
	isSet bool
}

func (v NullableCreateOrUpdateAppDto) Get() *CreateOrUpdateAppDto {
	return v.value
}

func (v *NullableCreateOrUpdateAppDto) Set(val *CreateOrUpdateAppDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateAppDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateAppDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateAppDto(val *CreateOrUpdateAppDto) *NullableCreateOrUpdateAppDto {
	return &NullableCreateOrUpdateAppDto{value: val, isSet: true}
}

func (v NullableCreateOrUpdateAppDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateAppDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


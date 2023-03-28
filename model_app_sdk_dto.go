/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
	"time"
)

// AppSdkDto struct for AppSdkDto
type AppSdkDto struct {
	Id *string `json:"id,omitempty"`
	CreationTime *time.Time `json:"creationTime,omitempty"`
	CreatorId *string `json:"creatorId,omitempty"`
	LastModificationTime *time.Time `json:"lastModificationTime,omitempty"`
	LastModifierId *string `json:"lastModifierId,omitempty"`
	IsDeleted *bool `json:"isDeleted,omitempty"`
	DeleterId *string `json:"deleterId,omitempty"`
	DeletionTime *time.Time `json:"deletionTime,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Privacy *string `json:"privacy,omitempty"`
	PrivacyUrl *string `json:"privacyUrl,omitempty"`
	HomePage *string `json:"homePage,omitempty"`
}

// NewAppSdkDto instantiates a new AppSdkDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppSdkDto() *AppSdkDto {
	this := AppSdkDto{}
	return &this
}

// NewAppSdkDtoWithDefaults instantiates a new AppSdkDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppSdkDtoWithDefaults() *AppSdkDto {
	this := AppSdkDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppSdkDto) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSdkDto) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppSdkDto) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AppSdkDto) SetId(v string) {
	o.Id = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *AppSdkDto) GetCreationTime() time.Time {
	if o == nil || isNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSdkDto) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreationTime) {
    return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *AppSdkDto) HasCreationTime() bool {
	if o != nil && !isNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *AppSdkDto) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *AppSdkDto) GetCreatorId() string {
	if o == nil || isNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSdkDto) GetCreatorIdOk() (*string, bool) {
	if o == nil || isNil(o.CreatorId) {
    return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *AppSdkDto) HasCreatorId() bool {
	if o != nil && !isNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *AppSdkDto) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetLastModificationTime returns the LastModificationTime field value if set, zero value otherwise.
func (o *AppSdkDto) GetLastModificationTime() time.Time {
	if o == nil || isNil(o.LastModificationTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModificationTime
}

// GetLastModificationTimeOk returns a tuple with the LastModificationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSdkDto) GetLastModificationTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastModificationTime) {
    return nil, false
	}
	return o.LastModificationTime, true
}

// HasLastModificationTime returns a boolean if a field has been set.
func (o *AppSdkDto) HasLastModificationTime() bool {
	if o != nil && !isNil(o.LastModificationTime) {
		return true
	}

	return false
}

// SetLastModificationTime gets a reference to the given time.Time and assigns it to the LastModificationTime field.
func (o *AppSdkDto) SetLastModificationTime(v time.Time) {
	o.LastModificationTime = &v
}

// GetLastModifierId returns the LastModifierId field value if set, zero value otherwise.
func (o *AppSdkDto) GetLastModifierId() string {
	if o == nil || isNil(o.LastModifierId) {
		var ret string
		return ret
	}
	return *o.LastModifierId
}

// GetLastModifierIdOk returns a tuple with the LastModifierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSdkDto) GetLastModifierIdOk() (*string, bool) {
	if o == nil || isNil(o.LastModifierId) {
    return nil, false
	}
	return o.LastModifierId, true
}

// HasLastModifierId returns a boolean if a field has been set.
func (o *AppSdkDto) HasLastModifierId() bool {
	if o != nil && !isNil(o.LastModifierId) {
		return true
	}

	return false
}

// SetLastModifierId gets a reference to the given string and assigns it to the LastModifierId field.
func (o *AppSdkDto) SetLastModifierId(v string) {
	o.LastModifierId = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *AppSdkDto) GetIsDeleted() bool {
	if o == nil || isNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSdkDto) GetIsDeletedOk() (*bool, bool) {
	if o == nil || isNil(o.IsDeleted) {
    return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *AppSdkDto) HasIsDeleted() bool {
	if o != nil && !isNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *AppSdkDto) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetDeleterId returns the DeleterId field value if set, zero value otherwise.
func (o *AppSdkDto) GetDeleterId() string {
	if o == nil || isNil(o.DeleterId) {
		var ret string
		return ret
	}
	return *o.DeleterId
}

// GetDeleterIdOk returns a tuple with the DeleterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSdkDto) GetDeleterIdOk() (*string, bool) {
	if o == nil || isNil(o.DeleterId) {
    return nil, false
	}
	return o.DeleterId, true
}

// HasDeleterId returns a boolean if a field has been set.
func (o *AppSdkDto) HasDeleterId() bool {
	if o != nil && !isNil(o.DeleterId) {
		return true
	}

	return false
}

// SetDeleterId gets a reference to the given string and assigns it to the DeleterId field.
func (o *AppSdkDto) SetDeleterId(v string) {
	o.DeleterId = &v
}

// GetDeletionTime returns the DeletionTime field value if set, zero value otherwise.
func (o *AppSdkDto) GetDeletionTime() time.Time {
	if o == nil || isNil(o.DeletionTime) {
		var ret time.Time
		return ret
	}
	return *o.DeletionTime
}

// GetDeletionTimeOk returns a tuple with the DeletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSdkDto) GetDeletionTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.DeletionTime) {
    return nil, false
	}
	return o.DeletionTime, true
}

// HasDeletionTime returns a boolean if a field has been set.
func (o *AppSdkDto) HasDeletionTime() bool {
	if o != nil && !isNil(o.DeletionTime) {
		return true
	}

	return false
}

// SetDeletionTime gets a reference to the given time.Time and assigns it to the DeletionTime field.
func (o *AppSdkDto) SetDeletionTime(v time.Time) {
	o.DeletionTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppSdkDto) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSdkDto) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppSdkDto) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppSdkDto) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AppSdkDto) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSdkDto) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AppSdkDto) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AppSdkDto) SetDescription(v string) {
	o.Description = &v
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *AppSdkDto) GetPrivacy() string {
	if o == nil || isNil(o.Privacy) {
		var ret string
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSdkDto) GetPrivacyOk() (*string, bool) {
	if o == nil || isNil(o.Privacy) {
    return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field has been set.
func (o *AppSdkDto) HasPrivacy() bool {
	if o != nil && !isNil(o.Privacy) {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given string and assigns it to the Privacy field.
func (o *AppSdkDto) SetPrivacy(v string) {
	o.Privacy = &v
}

// GetPrivacyUrl returns the PrivacyUrl field value if set, zero value otherwise.
func (o *AppSdkDto) GetPrivacyUrl() string {
	if o == nil || isNil(o.PrivacyUrl) {
		var ret string
		return ret
	}
	return *o.PrivacyUrl
}

// GetPrivacyUrlOk returns a tuple with the PrivacyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSdkDto) GetPrivacyUrlOk() (*string, bool) {
	if o == nil || isNil(o.PrivacyUrl) {
    return nil, false
	}
	return o.PrivacyUrl, true
}

// HasPrivacyUrl returns a boolean if a field has been set.
func (o *AppSdkDto) HasPrivacyUrl() bool {
	if o != nil && !isNil(o.PrivacyUrl) {
		return true
	}

	return false
}

// SetPrivacyUrl gets a reference to the given string and assigns it to the PrivacyUrl field.
func (o *AppSdkDto) SetPrivacyUrl(v string) {
	o.PrivacyUrl = &v
}

// GetHomePage returns the HomePage field value if set, zero value otherwise.
func (o *AppSdkDto) GetHomePage() string {
	if o == nil || isNil(o.HomePage) {
		var ret string
		return ret
	}
	return *o.HomePage
}

// GetHomePageOk returns a tuple with the HomePage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSdkDto) GetHomePageOk() (*string, bool) {
	if o == nil || isNil(o.HomePage) {
    return nil, false
	}
	return o.HomePage, true
}

// HasHomePage returns a boolean if a field has been set.
func (o *AppSdkDto) HasHomePage() bool {
	if o != nil && !isNil(o.HomePage) {
		return true
	}

	return false
}

// SetHomePage gets a reference to the given string and assigns it to the HomePage field.
func (o *AppSdkDto) SetHomePage(v string) {
	o.HomePage = &v
}

func (o AppSdkDto) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Privacy) {
		toSerialize["privacy"] = o.Privacy
	}
	if !isNil(o.PrivacyUrl) {
		toSerialize["privacyUrl"] = o.PrivacyUrl
	}
	if !isNil(o.HomePage) {
		toSerialize["homePage"] = o.HomePage
	}
	return json.Marshal(toSerialize)
}

type NullableAppSdkDto struct {
	value *AppSdkDto
	isSet bool
}

func (v NullableAppSdkDto) Get() *AppSdkDto {
	return v.value
}

func (v *NullableAppSdkDto) Set(val *AppSdkDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAppSdkDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAppSdkDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppSdkDto(val *AppSdkDto) *NullableAppSdkDto {
	return &NullableAppSdkDto{value: val, isSet: true}
}

func (v NullableAppSdkDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppSdkDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



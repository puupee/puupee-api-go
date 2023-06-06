/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
)

// checks if the PermissionGrantInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionGrantInfoDto{}

// PermissionGrantInfoDto struct for PermissionGrantInfoDto
type PermissionGrantInfoDto struct {
	Name *string `json:"name,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	ParentName *string `json:"parentName,omitempty"`
	IsGranted *bool `json:"isGranted,omitempty"`
	AllowedProviders []string `json:"allowedProviders,omitempty"`
	GrantedProviders []ProviderInfoDto `json:"grantedProviders,omitempty"`
}

// NewPermissionGrantInfoDto instantiates a new PermissionGrantInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionGrantInfoDto() *PermissionGrantInfoDto {
	this := PermissionGrantInfoDto{}
	return &this
}

// NewPermissionGrantInfoDtoWithDefaults instantiates a new PermissionGrantInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionGrantInfoDtoWithDefaults() *PermissionGrantInfoDto {
	this := PermissionGrantInfoDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PermissionGrantInfoDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGrantInfoDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PermissionGrantInfoDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PermissionGrantInfoDto) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PermissionGrantInfoDto) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGrantInfoDto) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PermissionGrantInfoDto) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PermissionGrantInfoDto) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetParentName returns the ParentName field value if set, zero value otherwise.
func (o *PermissionGrantInfoDto) GetParentName() string {
	if o == nil || IsNil(o.ParentName) {
		var ret string
		return ret
	}
	return *o.ParentName
}

// GetParentNameOk returns a tuple with the ParentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGrantInfoDto) GetParentNameOk() (*string, bool) {
	if o == nil || IsNil(o.ParentName) {
		return nil, false
	}
	return o.ParentName, true
}

// HasParentName returns a boolean if a field has been set.
func (o *PermissionGrantInfoDto) HasParentName() bool {
	if o != nil && !IsNil(o.ParentName) {
		return true
	}

	return false
}

// SetParentName gets a reference to the given string and assigns it to the ParentName field.
func (o *PermissionGrantInfoDto) SetParentName(v string) {
	o.ParentName = &v
}

// GetIsGranted returns the IsGranted field value if set, zero value otherwise.
func (o *PermissionGrantInfoDto) GetIsGranted() bool {
	if o == nil || IsNil(o.IsGranted) {
		var ret bool
		return ret
	}
	return *o.IsGranted
}

// GetIsGrantedOk returns a tuple with the IsGranted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGrantInfoDto) GetIsGrantedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGranted) {
		return nil, false
	}
	return o.IsGranted, true
}

// HasIsGranted returns a boolean if a field has been set.
func (o *PermissionGrantInfoDto) HasIsGranted() bool {
	if o != nil && !IsNil(o.IsGranted) {
		return true
	}

	return false
}

// SetIsGranted gets a reference to the given bool and assigns it to the IsGranted field.
func (o *PermissionGrantInfoDto) SetIsGranted(v bool) {
	o.IsGranted = &v
}

// GetAllowedProviders returns the AllowedProviders field value if set, zero value otherwise.
func (o *PermissionGrantInfoDto) GetAllowedProviders() []string {
	if o == nil || IsNil(o.AllowedProviders) {
		var ret []string
		return ret
	}
	return o.AllowedProviders
}

// GetAllowedProvidersOk returns a tuple with the AllowedProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGrantInfoDto) GetAllowedProvidersOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedProviders) {
		return nil, false
	}
	return o.AllowedProviders, true
}

// HasAllowedProviders returns a boolean if a field has been set.
func (o *PermissionGrantInfoDto) HasAllowedProviders() bool {
	if o != nil && !IsNil(o.AllowedProviders) {
		return true
	}

	return false
}

// SetAllowedProviders gets a reference to the given []string and assigns it to the AllowedProviders field.
func (o *PermissionGrantInfoDto) SetAllowedProviders(v []string) {
	o.AllowedProviders = v
}

// GetGrantedProviders returns the GrantedProviders field value if set, zero value otherwise.
func (o *PermissionGrantInfoDto) GetGrantedProviders() []ProviderInfoDto {
	if o == nil || IsNil(o.GrantedProviders) {
		var ret []ProviderInfoDto
		return ret
	}
	return o.GrantedProviders
}

// GetGrantedProvidersOk returns a tuple with the GrantedProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGrantInfoDto) GetGrantedProvidersOk() ([]ProviderInfoDto, bool) {
	if o == nil || IsNil(o.GrantedProviders) {
		return nil, false
	}
	return o.GrantedProviders, true
}

// HasGrantedProviders returns a boolean if a field has been set.
func (o *PermissionGrantInfoDto) HasGrantedProviders() bool {
	if o != nil && !IsNil(o.GrantedProviders) {
		return true
	}

	return false
}

// SetGrantedProviders gets a reference to the given []ProviderInfoDto and assigns it to the GrantedProviders field.
func (o *PermissionGrantInfoDto) SetGrantedProviders(v []ProviderInfoDto) {
	o.GrantedProviders = v
}

func (o PermissionGrantInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionGrantInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.ParentName) {
		toSerialize["parentName"] = o.ParentName
	}
	if !IsNil(o.IsGranted) {
		toSerialize["isGranted"] = o.IsGranted
	}
	if !IsNil(o.AllowedProviders) {
		toSerialize["allowedProviders"] = o.AllowedProviders
	}
	if !IsNil(o.GrantedProviders) {
		toSerialize["grantedProviders"] = o.GrantedProviders
	}
	return toSerialize, nil
}

type NullablePermissionGrantInfoDto struct {
	value *PermissionGrantInfoDto
	isSet bool
}

func (v NullablePermissionGrantInfoDto) Get() *PermissionGrantInfoDto {
	return v.value
}

func (v *NullablePermissionGrantInfoDto) Set(val *PermissionGrantInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionGrantInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionGrantInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionGrantInfoDto(val *PermissionGrantInfoDto) *NullablePermissionGrantInfoDto {
	return &NullablePermissionGrantInfoDto{value: val, isSet: true}
}

func (v NullablePermissionGrantInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionGrantInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



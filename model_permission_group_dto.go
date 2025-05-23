/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.17.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
)

// checks if the PermissionGroupDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionGroupDto{}

// PermissionGroupDto struct for PermissionGroupDto
type PermissionGroupDto struct {
	Name *string `json:"name,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	DisplayNameKey *string `json:"displayNameKey,omitempty"`
	DisplayNameResource *string `json:"displayNameResource,omitempty"`
	Permissions []PermissionGrantInfoDto `json:"permissions,omitempty"`
}

// NewPermissionGroupDto instantiates a new PermissionGroupDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionGroupDto() *PermissionGroupDto {
	this := PermissionGroupDto{}
	return &this
}

// NewPermissionGroupDtoWithDefaults instantiates a new PermissionGroupDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionGroupDtoWithDefaults() *PermissionGroupDto {
	this := PermissionGroupDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PermissionGroupDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PermissionGroupDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PermissionGroupDto) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PermissionGroupDto) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupDto) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PermissionGroupDto) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PermissionGroupDto) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDisplayNameKey returns the DisplayNameKey field value if set, zero value otherwise.
func (o *PermissionGroupDto) GetDisplayNameKey() string {
	if o == nil || IsNil(o.DisplayNameKey) {
		var ret string
		return ret
	}
	return *o.DisplayNameKey
}

// GetDisplayNameKeyOk returns a tuple with the DisplayNameKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupDto) GetDisplayNameKeyOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayNameKey) {
		return nil, false
	}
	return o.DisplayNameKey, true
}

// HasDisplayNameKey returns a boolean if a field has been set.
func (o *PermissionGroupDto) HasDisplayNameKey() bool {
	if o != nil && !IsNil(o.DisplayNameKey) {
		return true
	}

	return false
}

// SetDisplayNameKey gets a reference to the given string and assigns it to the DisplayNameKey field.
func (o *PermissionGroupDto) SetDisplayNameKey(v string) {
	o.DisplayNameKey = &v
}

// GetDisplayNameResource returns the DisplayNameResource field value if set, zero value otherwise.
func (o *PermissionGroupDto) GetDisplayNameResource() string {
	if o == nil || IsNil(o.DisplayNameResource) {
		var ret string
		return ret
	}
	return *o.DisplayNameResource
}

// GetDisplayNameResourceOk returns a tuple with the DisplayNameResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupDto) GetDisplayNameResourceOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayNameResource) {
		return nil, false
	}
	return o.DisplayNameResource, true
}

// HasDisplayNameResource returns a boolean if a field has been set.
func (o *PermissionGroupDto) HasDisplayNameResource() bool {
	if o != nil && !IsNil(o.DisplayNameResource) {
		return true
	}

	return false
}

// SetDisplayNameResource gets a reference to the given string and assigns it to the DisplayNameResource field.
func (o *PermissionGroupDto) SetDisplayNameResource(v string) {
	o.DisplayNameResource = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *PermissionGroupDto) GetPermissions() []PermissionGrantInfoDto {
	if o == nil || IsNil(o.Permissions) {
		var ret []PermissionGrantInfoDto
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionGroupDto) GetPermissionsOk() ([]PermissionGrantInfoDto, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *PermissionGroupDto) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []PermissionGrantInfoDto and assigns it to the Permissions field.
func (o *PermissionGroupDto) SetPermissions(v []PermissionGrantInfoDto) {
	o.Permissions = v
}

func (o PermissionGroupDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionGroupDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.DisplayNameKey) {
		toSerialize["displayNameKey"] = o.DisplayNameKey
	}
	if !IsNil(o.DisplayNameResource) {
		toSerialize["displayNameResource"] = o.DisplayNameResource
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	return toSerialize, nil
}

type NullablePermissionGroupDto struct {
	value *PermissionGroupDto
	isSet bool
}

func (v NullablePermissionGroupDto) Get() *PermissionGroupDto {
	return v.value
}

func (v *NullablePermissionGroupDto) Set(val *PermissionGroupDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionGroupDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionGroupDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionGroupDto(val *PermissionGroupDto) *NullablePermissionGroupDto {
	return &NullablePermissionGroupDto{value: val, isSet: true}
}

func (v NullablePermissionGroupDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionGroupDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the UpdatePermissionDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePermissionDto{}

// UpdatePermissionDto struct for UpdatePermissionDto
type UpdatePermissionDto struct {
	Name *string `json:"name,omitempty"`
	IsGranted *bool `json:"isGranted,omitempty"`
}

// NewUpdatePermissionDto instantiates a new UpdatePermissionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePermissionDto() *UpdatePermissionDto {
	this := UpdatePermissionDto{}
	return &this
}

// NewUpdatePermissionDtoWithDefaults instantiates a new UpdatePermissionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePermissionDtoWithDefaults() *UpdatePermissionDto {
	this := UpdatePermissionDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdatePermissionDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePermissionDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdatePermissionDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdatePermissionDto) SetName(v string) {
	o.Name = &v
}

// GetIsGranted returns the IsGranted field value if set, zero value otherwise.
func (o *UpdatePermissionDto) GetIsGranted() bool {
	if o == nil || IsNil(o.IsGranted) {
		var ret bool
		return ret
	}
	return *o.IsGranted
}

// GetIsGrantedOk returns a tuple with the IsGranted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePermissionDto) GetIsGrantedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGranted) {
		return nil, false
	}
	return o.IsGranted, true
}

// HasIsGranted returns a boolean if a field has been set.
func (o *UpdatePermissionDto) HasIsGranted() bool {
	if o != nil && !IsNil(o.IsGranted) {
		return true
	}

	return false
}

// SetIsGranted gets a reference to the given bool and assigns it to the IsGranted field.
func (o *UpdatePermissionDto) SetIsGranted(v bool) {
	o.IsGranted = &v
}

func (o UpdatePermissionDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePermissionDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.IsGranted) {
		toSerialize["isGranted"] = o.IsGranted
	}
	return toSerialize, nil
}

type NullableUpdatePermissionDto struct {
	value *UpdatePermissionDto
	isSet bool
}

func (v NullableUpdatePermissionDto) Get() *UpdatePermissionDto {
	return v.value
}

func (v *NullableUpdatePermissionDto) Set(val *UpdatePermissionDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePermissionDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePermissionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePermissionDto(val *UpdatePermissionDto) *NullableUpdatePermissionDto {
	return &NullableUpdatePermissionDto{value: val, isSet: true}
}

func (v NullableUpdatePermissionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePermissionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



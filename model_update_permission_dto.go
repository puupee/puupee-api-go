/*
Doggy API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package doggy

import (
	"encoding/json"
)

// UpdatePermissionDto struct for UpdatePermissionDto
type UpdatePermissionDto struct {
	Name NullableString `json:"name,omitempty"`
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

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdatePermissionDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatePermissionDto) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdatePermissionDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdatePermissionDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdatePermissionDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdatePermissionDto) UnsetName() {
	o.Name.Unset()
}

// GetIsGranted returns the IsGranted field value if set, zero value otherwise.
func (o *UpdatePermissionDto) GetIsGranted() bool {
	if o == nil || o.IsGranted == nil {
		var ret bool
		return ret
	}
	return *o.IsGranted
}

// GetIsGrantedOk returns a tuple with the IsGranted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePermissionDto) GetIsGrantedOk() (*bool, bool) {
	if o == nil || o.IsGranted == nil {
		return nil, false
	}
	return o.IsGranted, true
}

// HasIsGranted returns a boolean if a field has been set.
func (o *UpdatePermissionDto) HasIsGranted() bool {
	if o != nil && o.IsGranted != nil {
		return true
	}

	return false
}

// SetIsGranted gets a reference to the given bool and assigns it to the IsGranted field.
func (o *UpdatePermissionDto) SetIsGranted(v bool) {
	o.IsGranted = &v
}

func (o UpdatePermissionDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.IsGranted != nil {
		toSerialize["isGranted"] = o.IsGranted
	}
	return json.Marshal(toSerialize)
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



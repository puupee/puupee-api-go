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

// MultiTenancyInfoDto struct for MultiTenancyInfoDto
type MultiTenancyInfoDto struct {
	IsEnabled *bool `json:"isEnabled,omitempty"`
}

// NewMultiTenancyInfoDto instantiates a new MultiTenancyInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiTenancyInfoDto() *MultiTenancyInfoDto {
	this := MultiTenancyInfoDto{}
	return &this
}

// NewMultiTenancyInfoDtoWithDefaults instantiates a new MultiTenancyInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiTenancyInfoDtoWithDefaults() *MultiTenancyInfoDto {
	this := MultiTenancyInfoDto{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *MultiTenancyInfoDto) GetIsEnabled() bool {
	if o == nil || isNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiTenancyInfoDto) GetIsEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsEnabled) {
    return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *MultiTenancyInfoDto) HasIsEnabled() bool {
	if o != nil && !isNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *MultiTenancyInfoDto) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

func (o MultiTenancyInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableMultiTenancyInfoDto struct {
	value *MultiTenancyInfoDto
	isSet bool
}

func (v NullableMultiTenancyInfoDto) Get() *MultiTenancyInfoDto {
	return v.value
}

func (v *NullableMultiTenancyInfoDto) Set(val *MultiTenancyInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiTenancyInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiTenancyInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiTenancyInfoDto(val *MultiTenancyInfoDto) *NullableMultiTenancyInfoDto {
	return &NullableMultiTenancyInfoDto{value: val, isSet: true}
}

func (v NullableMultiTenancyInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiTenancyInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



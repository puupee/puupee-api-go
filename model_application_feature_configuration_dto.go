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

// ApplicationFeatureConfigurationDto struct for ApplicationFeatureConfigurationDto
type ApplicationFeatureConfigurationDto struct {
	Values map[string]string `json:"values,omitempty"`
}

// NewApplicationFeatureConfigurationDto instantiates a new ApplicationFeatureConfigurationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationFeatureConfigurationDto() *ApplicationFeatureConfigurationDto {
	this := ApplicationFeatureConfigurationDto{}
	return &this
}

// NewApplicationFeatureConfigurationDtoWithDefaults instantiates a new ApplicationFeatureConfigurationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationFeatureConfigurationDtoWithDefaults() *ApplicationFeatureConfigurationDto {
	this := ApplicationFeatureConfigurationDto{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationFeatureConfigurationDto) GetValues() map[string]string {
	if o == nil  {
		var ret map[string]string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationFeatureConfigurationDto) GetValuesOk() (*map[string]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ApplicationFeatureConfigurationDto) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given map[string]string and assigns it to the Values field.
func (o *ApplicationFeatureConfigurationDto) SetValues(v map[string]string) {
	o.Values = v
}

func (o ApplicationFeatureConfigurationDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationFeatureConfigurationDto struct {
	value *ApplicationFeatureConfigurationDto
	isSet bool
}

func (v NullableApplicationFeatureConfigurationDto) Get() *ApplicationFeatureConfigurationDto {
	return v.value
}

func (v *NullableApplicationFeatureConfigurationDto) Set(val *ApplicationFeatureConfigurationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationFeatureConfigurationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationFeatureConfigurationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationFeatureConfigurationDto(val *ApplicationFeatureConfigurationDto) *NullableApplicationFeatureConfigurationDto {
	return &NullableApplicationFeatureConfigurationDto{value: val, isSet: true}
}

func (v NullableApplicationFeatureConfigurationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationFeatureConfigurationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



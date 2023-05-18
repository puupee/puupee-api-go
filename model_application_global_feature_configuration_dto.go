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

// checks if the ApplicationGlobalFeatureConfigurationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationGlobalFeatureConfigurationDto{}

// ApplicationGlobalFeatureConfigurationDto struct for ApplicationGlobalFeatureConfigurationDto
type ApplicationGlobalFeatureConfigurationDto struct {
	EnabledFeatures []string `json:"enabledFeatures,omitempty"`
}

// NewApplicationGlobalFeatureConfigurationDto instantiates a new ApplicationGlobalFeatureConfigurationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationGlobalFeatureConfigurationDto() *ApplicationGlobalFeatureConfigurationDto {
	this := ApplicationGlobalFeatureConfigurationDto{}
	return &this
}

// NewApplicationGlobalFeatureConfigurationDtoWithDefaults instantiates a new ApplicationGlobalFeatureConfigurationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationGlobalFeatureConfigurationDtoWithDefaults() *ApplicationGlobalFeatureConfigurationDto {
	this := ApplicationGlobalFeatureConfigurationDto{}
	return &this
}

// GetEnabledFeatures returns the EnabledFeatures field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationGlobalFeatureConfigurationDto) GetEnabledFeatures() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.EnabledFeatures
}

// GetEnabledFeaturesOk returns a tuple with the EnabledFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationGlobalFeatureConfigurationDto) GetEnabledFeaturesOk() ([]string, bool) {
	if o == nil || IsNil(o.EnabledFeatures) {
		return nil, false
	}
	return o.EnabledFeatures, true
}

// HasEnabledFeatures returns a boolean if a field has been set.
func (o *ApplicationGlobalFeatureConfigurationDto) HasEnabledFeatures() bool {
	if o != nil && IsNil(o.EnabledFeatures) {
		return true
	}

	return false
}

// SetEnabledFeatures gets a reference to the given []string and assigns it to the EnabledFeatures field.
func (o *ApplicationGlobalFeatureConfigurationDto) SetEnabledFeatures(v []string) {
	o.EnabledFeatures = v
}

func (o ApplicationGlobalFeatureConfigurationDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationGlobalFeatureConfigurationDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EnabledFeatures != nil {
		toSerialize["enabledFeatures"] = o.EnabledFeatures
	}
	return toSerialize, nil
}

type NullableApplicationGlobalFeatureConfigurationDto struct {
	value *ApplicationGlobalFeatureConfigurationDto
	isSet bool
}

func (v NullableApplicationGlobalFeatureConfigurationDto) Get() *ApplicationGlobalFeatureConfigurationDto {
	return v.value
}

func (v *NullableApplicationGlobalFeatureConfigurationDto) Set(val *ApplicationGlobalFeatureConfigurationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationGlobalFeatureConfigurationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationGlobalFeatureConfigurationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationGlobalFeatureConfigurationDto(val *ApplicationGlobalFeatureConfigurationDto) *NullableApplicationGlobalFeatureConfigurationDto {
	return &NullableApplicationGlobalFeatureConfigurationDto{value: val, isSet: true}
}

func (v NullableApplicationGlobalFeatureConfigurationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationGlobalFeatureConfigurationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



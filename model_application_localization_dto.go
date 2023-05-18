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

// checks if the ApplicationLocalizationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationLocalizationDto{}

// ApplicationLocalizationDto struct for ApplicationLocalizationDto
type ApplicationLocalizationDto struct {
	Resources map[string]ApplicationLocalizationResourceDto `json:"resources,omitempty"`
}

// NewApplicationLocalizationDto instantiates a new ApplicationLocalizationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationLocalizationDto() *ApplicationLocalizationDto {
	this := ApplicationLocalizationDto{}
	return &this
}

// NewApplicationLocalizationDtoWithDefaults instantiates a new ApplicationLocalizationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationLocalizationDtoWithDefaults() *ApplicationLocalizationDto {
	this := ApplicationLocalizationDto{}
	return &this
}

// GetResources returns the Resources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationLocalizationDto) GetResources() map[string]ApplicationLocalizationResourceDto {
	if o == nil {
		var ret map[string]ApplicationLocalizationResourceDto
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationLocalizationDto) GetResourcesOk() (*map[string]ApplicationLocalizationResourceDto, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return &o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *ApplicationLocalizationDto) HasResources() bool {
	if o != nil && IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given map[string]ApplicationLocalizationResourceDto and assigns it to the Resources field.
func (o *ApplicationLocalizationDto) SetResources(v map[string]ApplicationLocalizationResourceDto) {
	o.Resources = v
}

func (o ApplicationLocalizationDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationLocalizationDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	return toSerialize, nil
}

type NullableApplicationLocalizationDto struct {
	value *ApplicationLocalizationDto
	isSet bool
}

func (v NullableApplicationLocalizationDto) Get() *ApplicationLocalizationDto {
	return v.value
}

func (v *NullableApplicationLocalizationDto) Set(val *ApplicationLocalizationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationLocalizationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationLocalizationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationLocalizationDto(val *ApplicationLocalizationDto) *NullableApplicationLocalizationDto {
	return &NullableApplicationLocalizationDto{value: val, isSet: true}
}

func (v NullableApplicationLocalizationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationLocalizationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



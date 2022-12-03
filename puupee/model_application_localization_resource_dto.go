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

// ApplicationLocalizationResourceDto struct for ApplicationLocalizationResourceDto
type ApplicationLocalizationResourceDto struct {
	Texts *map[string]string `json:"texts,omitempty"`
	BaseResources []string `json:"baseResources,omitempty"`
}

// NewApplicationLocalizationResourceDto instantiates a new ApplicationLocalizationResourceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationLocalizationResourceDto() *ApplicationLocalizationResourceDto {
	this := ApplicationLocalizationResourceDto{}
	return &this
}

// NewApplicationLocalizationResourceDtoWithDefaults instantiates a new ApplicationLocalizationResourceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationLocalizationResourceDtoWithDefaults() *ApplicationLocalizationResourceDto {
	this := ApplicationLocalizationResourceDto{}
	return &this
}

// GetTexts returns the Texts field value if set, zero value otherwise.
func (o *ApplicationLocalizationResourceDto) GetTexts() map[string]string {
	if o == nil || isNil(o.Texts) {
		var ret map[string]string
		return ret
	}
	return *o.Texts
}

// GetTextsOk returns a tuple with the Texts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLocalizationResourceDto) GetTextsOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Texts) {
    return nil, false
	}
	return o.Texts, true
}

// HasTexts returns a boolean if a field has been set.
func (o *ApplicationLocalizationResourceDto) HasTexts() bool {
	if o != nil && !isNil(o.Texts) {
		return true
	}

	return false
}

// SetTexts gets a reference to the given map[string]string and assigns it to the Texts field.
func (o *ApplicationLocalizationResourceDto) SetTexts(v map[string]string) {
	o.Texts = &v
}

// GetBaseResources returns the BaseResources field value if set, zero value otherwise.
func (o *ApplicationLocalizationResourceDto) GetBaseResources() []string {
	if o == nil || isNil(o.BaseResources) {
		var ret []string
		return ret
	}
	return o.BaseResources
}

// GetBaseResourcesOk returns a tuple with the BaseResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLocalizationResourceDto) GetBaseResourcesOk() ([]string, bool) {
	if o == nil || isNil(o.BaseResources) {
    return nil, false
	}
	return o.BaseResources, true
}

// HasBaseResources returns a boolean if a field has been set.
func (o *ApplicationLocalizationResourceDto) HasBaseResources() bool {
	if o != nil && !isNil(o.BaseResources) {
		return true
	}

	return false
}

// SetBaseResources gets a reference to the given []string and assigns it to the BaseResources field.
func (o *ApplicationLocalizationResourceDto) SetBaseResources(v []string) {
	o.BaseResources = v
}

func (o ApplicationLocalizationResourceDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Texts) {
		toSerialize["texts"] = o.Texts
	}
	if !isNil(o.BaseResources) {
		toSerialize["baseResources"] = o.BaseResources
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationLocalizationResourceDto struct {
	value *ApplicationLocalizationResourceDto
	isSet bool
}

func (v NullableApplicationLocalizationResourceDto) Get() *ApplicationLocalizationResourceDto {
	return v.value
}

func (v *NullableApplicationLocalizationResourceDto) Set(val *ApplicationLocalizationResourceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationLocalizationResourceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationLocalizationResourceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationLocalizationResourceDto(val *ApplicationLocalizationResourceDto) *NullableApplicationLocalizationResourceDto {
	return &NullableApplicationLocalizationResourceDto{value: val, isSet: true}
}

func (v NullableApplicationLocalizationResourceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationLocalizationResourceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



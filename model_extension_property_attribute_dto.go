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

// ExtensionPropertyAttributeDto struct for ExtensionPropertyAttributeDto
type ExtensionPropertyAttributeDto struct {
	TypeSimple *string `json:"typeSimple,omitempty"`
	Config map[string]map[string]interface{} `json:"config,omitempty"`
}

// NewExtensionPropertyAttributeDto instantiates a new ExtensionPropertyAttributeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionPropertyAttributeDto() *ExtensionPropertyAttributeDto {
	this := ExtensionPropertyAttributeDto{}
	return &this
}

// NewExtensionPropertyAttributeDtoWithDefaults instantiates a new ExtensionPropertyAttributeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionPropertyAttributeDtoWithDefaults() *ExtensionPropertyAttributeDto {
	this := ExtensionPropertyAttributeDto{}
	return &this
}

// GetTypeSimple returns the TypeSimple field value if set, zero value otherwise.
func (o *ExtensionPropertyAttributeDto) GetTypeSimple() string {
	if o == nil || isNil(o.TypeSimple) {
		var ret string
		return ret
	}
	return *o.TypeSimple
}

// GetTypeSimpleOk returns a tuple with the TypeSimple field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionPropertyAttributeDto) GetTypeSimpleOk() (*string, bool) {
	if o == nil || isNil(o.TypeSimple) {
    return nil, false
	}
	return o.TypeSimple, true
}

// HasTypeSimple returns a boolean if a field has been set.
func (o *ExtensionPropertyAttributeDto) HasTypeSimple() bool {
	if o != nil && !isNil(o.TypeSimple) {
		return true
	}

	return false
}

// SetTypeSimple gets a reference to the given string and assigns it to the TypeSimple field.
func (o *ExtensionPropertyAttributeDto) SetTypeSimple(v string) {
	o.TypeSimple = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *ExtensionPropertyAttributeDto) GetConfig() map[string]map[string]interface{} {
	if o == nil || isNil(o.Config) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionPropertyAttributeDto) GetConfigOk() (map[string]map[string]interface{}, bool) {
	if o == nil || isNil(o.Config) {
    return map[string]map[string]interface{}{}, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *ExtensionPropertyAttributeDto) HasConfig() bool {
	if o != nil && !isNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]map[string]interface{} and assigns it to the Config field.
func (o *ExtensionPropertyAttributeDto) SetConfig(v map[string]map[string]interface{}) {
	o.Config = v
}

func (o ExtensionPropertyAttributeDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TypeSimple) {
		toSerialize["typeSimple"] = o.TypeSimple
	}
	if !isNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableExtensionPropertyAttributeDto struct {
	value *ExtensionPropertyAttributeDto
	isSet bool
}

func (v NullableExtensionPropertyAttributeDto) Get() *ExtensionPropertyAttributeDto {
	return v.value
}

func (v *NullableExtensionPropertyAttributeDto) Set(val *ExtensionPropertyAttributeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionPropertyAttributeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionPropertyAttributeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionPropertyAttributeDto(val *ExtensionPropertyAttributeDto) *NullableExtensionPropertyAttributeDto {
	return &NullableExtensionPropertyAttributeDto{value: val, isSet: true}
}

func (v NullableExtensionPropertyAttributeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionPropertyAttributeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


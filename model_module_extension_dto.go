/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
)

// ModuleExtensionDto struct for ModuleExtensionDto
type ModuleExtensionDto struct {
	Entities *map[string]EntityExtensionDto `json:"entities,omitempty"`
	Configuration map[string]map[string]interface{} `json:"configuration,omitempty"`
}

// NewModuleExtensionDto instantiates a new ModuleExtensionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModuleExtensionDto() *ModuleExtensionDto {
	this := ModuleExtensionDto{}
	return &this
}

// NewModuleExtensionDtoWithDefaults instantiates a new ModuleExtensionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleExtensionDtoWithDefaults() *ModuleExtensionDto {
	this := ModuleExtensionDto{}
	return &this
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *ModuleExtensionDto) GetEntities() map[string]EntityExtensionDto {
	if o == nil || isNil(o.Entities) {
		var ret map[string]EntityExtensionDto
		return ret
	}
	return *o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleExtensionDto) GetEntitiesOk() (*map[string]EntityExtensionDto, bool) {
	if o == nil || isNil(o.Entities) {
    return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *ModuleExtensionDto) HasEntities() bool {
	if o != nil && !isNil(o.Entities) {
		return true
	}

	return false
}

// SetEntities gets a reference to the given map[string]EntityExtensionDto and assigns it to the Entities field.
func (o *ModuleExtensionDto) SetEntities(v map[string]EntityExtensionDto) {
	o.Entities = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *ModuleExtensionDto) GetConfiguration() map[string]map[string]interface{} {
	if o == nil || isNil(o.Configuration) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleExtensionDto) GetConfigurationOk() (map[string]map[string]interface{}, bool) {
	if o == nil || isNil(o.Configuration) {
    return map[string]map[string]interface{}{}, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *ModuleExtensionDto) HasConfiguration() bool {
	if o != nil && !isNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]map[string]interface{} and assigns it to the Configuration field.
func (o *ModuleExtensionDto) SetConfiguration(v map[string]map[string]interface{}) {
	o.Configuration = v
}

func (o ModuleExtensionDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Entities) {
		toSerialize["entities"] = o.Entities
	}
	if !isNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	return json.Marshal(toSerialize)
}

type NullableModuleExtensionDto struct {
	value *ModuleExtensionDto
	isSet bool
}

func (v NullableModuleExtensionDto) Get() *ModuleExtensionDto {
	return v.value
}

func (v *NullableModuleExtensionDto) Set(val *ModuleExtensionDto) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleExtensionDto) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleExtensionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleExtensionDto(val *ModuleExtensionDto) *NullableModuleExtensionDto {
	return &NullableModuleExtensionDto{value: val, isSet: true}
}

func (v NullableModuleExtensionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleExtensionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



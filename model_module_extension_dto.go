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

// checks if the ModuleExtensionDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModuleExtensionDto{}

// ModuleExtensionDto struct for ModuleExtensionDto
type ModuleExtensionDto struct {
	Entities map[string]EntityExtensionDto `json:"entities,omitempty"`
	Configuration map[string]interface{} `json:"configuration,omitempty"`
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

// GetEntities returns the Entities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModuleExtensionDto) GetEntities() map[string]EntityExtensionDto {
	if o == nil {
		var ret map[string]EntityExtensionDto
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModuleExtensionDto) GetEntitiesOk() (*map[string]EntityExtensionDto, bool) {
	if o == nil || IsNil(o.Entities) {
		return nil, false
	}
	return &o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *ModuleExtensionDto) HasEntities() bool {
	if o != nil && IsNil(o.Entities) {
		return true
	}

	return false
}

// SetEntities gets a reference to the given map[string]EntityExtensionDto and assigns it to the Entities field.
func (o *ModuleExtensionDto) SetEntities(v map[string]EntityExtensionDto) {
	o.Entities = v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModuleExtensionDto) GetConfiguration() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModuleExtensionDto) GetConfigurationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Configuration) {
		return map[string]interface{}{}, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *ModuleExtensionDto) HasConfiguration() bool {
	if o != nil && IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]interface{} and assigns it to the Configuration field.
func (o *ModuleExtensionDto) SetConfiguration(v map[string]interface{}) {
	o.Configuration = v
}

func (o ModuleExtensionDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModuleExtensionDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Entities != nil {
		toSerialize["entities"] = o.Entities
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	return toSerialize, nil
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



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

// ExtensionPropertyApiCreateDto struct for ExtensionPropertyApiCreateDto
type ExtensionPropertyApiCreateDto struct {
	IsAvailable *bool `json:"isAvailable,omitempty"`
}

// NewExtensionPropertyApiCreateDto instantiates a new ExtensionPropertyApiCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionPropertyApiCreateDto() *ExtensionPropertyApiCreateDto {
	this := ExtensionPropertyApiCreateDto{}
	return &this
}

// NewExtensionPropertyApiCreateDtoWithDefaults instantiates a new ExtensionPropertyApiCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionPropertyApiCreateDtoWithDefaults() *ExtensionPropertyApiCreateDto {
	this := ExtensionPropertyApiCreateDto{}
	return &this
}

// GetIsAvailable returns the IsAvailable field value if set, zero value otherwise.
func (o *ExtensionPropertyApiCreateDto) GetIsAvailable() bool {
	if o == nil || isNil(o.IsAvailable) {
		var ret bool
		return ret
	}
	return *o.IsAvailable
}

// GetIsAvailableOk returns a tuple with the IsAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionPropertyApiCreateDto) GetIsAvailableOk() (*bool, bool) {
	if o == nil || isNil(o.IsAvailable) {
    return nil, false
	}
	return o.IsAvailable, true
}

// HasIsAvailable returns a boolean if a field has been set.
func (o *ExtensionPropertyApiCreateDto) HasIsAvailable() bool {
	if o != nil && !isNil(o.IsAvailable) {
		return true
	}

	return false
}

// SetIsAvailable gets a reference to the given bool and assigns it to the IsAvailable field.
func (o *ExtensionPropertyApiCreateDto) SetIsAvailable(v bool) {
	o.IsAvailable = &v
}

func (o ExtensionPropertyApiCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IsAvailable) {
		toSerialize["isAvailable"] = o.IsAvailable
	}
	return json.Marshal(toSerialize)
}

type NullableExtensionPropertyApiCreateDto struct {
	value *ExtensionPropertyApiCreateDto
	isSet bool
}

func (v NullableExtensionPropertyApiCreateDto) Get() *ExtensionPropertyApiCreateDto {
	return v.value
}

func (v *NullableExtensionPropertyApiCreateDto) Set(val *ExtensionPropertyApiCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionPropertyApiCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionPropertyApiCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionPropertyApiCreateDto(val *ExtensionPropertyApiCreateDto) *NullableExtensionPropertyApiCreateDto {
	return &NullableExtensionPropertyApiCreateDto{value: val, isSet: true}
}

func (v NullableExtensionPropertyApiCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionPropertyApiCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



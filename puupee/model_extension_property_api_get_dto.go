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

// ExtensionPropertyApiGetDto struct for ExtensionPropertyApiGetDto
type ExtensionPropertyApiGetDto struct {
	IsAvailable *bool `json:"isAvailable,omitempty"`
}

// NewExtensionPropertyApiGetDto instantiates a new ExtensionPropertyApiGetDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionPropertyApiGetDto() *ExtensionPropertyApiGetDto {
	this := ExtensionPropertyApiGetDto{}
	return &this
}

// NewExtensionPropertyApiGetDtoWithDefaults instantiates a new ExtensionPropertyApiGetDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionPropertyApiGetDtoWithDefaults() *ExtensionPropertyApiGetDto {
	this := ExtensionPropertyApiGetDto{}
	return &this
}

// GetIsAvailable returns the IsAvailable field value if set, zero value otherwise.
func (o *ExtensionPropertyApiGetDto) GetIsAvailable() bool {
	if o == nil || isNil(o.IsAvailable) {
		var ret bool
		return ret
	}
	return *o.IsAvailable
}

// GetIsAvailableOk returns a tuple with the IsAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionPropertyApiGetDto) GetIsAvailableOk() (*bool, bool) {
	if o == nil || isNil(o.IsAvailable) {
    return nil, false
	}
	return o.IsAvailable, true
}

// HasIsAvailable returns a boolean if a field has been set.
func (o *ExtensionPropertyApiGetDto) HasIsAvailable() bool {
	if o != nil && !isNil(o.IsAvailable) {
		return true
	}

	return false
}

// SetIsAvailable gets a reference to the given bool and assigns it to the IsAvailable field.
func (o *ExtensionPropertyApiGetDto) SetIsAvailable(v bool) {
	o.IsAvailable = &v
}

func (o ExtensionPropertyApiGetDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IsAvailable) {
		toSerialize["isAvailable"] = o.IsAvailable
	}
	return json.Marshal(toSerialize)
}

type NullableExtensionPropertyApiGetDto struct {
	value *ExtensionPropertyApiGetDto
	isSet bool
}

func (v NullableExtensionPropertyApiGetDto) Get() *ExtensionPropertyApiGetDto {
	return v.value
}

func (v *NullableExtensionPropertyApiGetDto) Set(val *ExtensionPropertyApiGetDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionPropertyApiGetDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionPropertyApiGetDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionPropertyApiGetDto(val *ExtensionPropertyApiGetDto) *NullableExtensionPropertyApiGetDto {
	return &NullableExtensionPropertyApiGetDto{value: val, isSet: true}
}

func (v NullableExtensionPropertyApiGetDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionPropertyApiGetDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



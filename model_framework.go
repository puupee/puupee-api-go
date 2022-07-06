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

// Framework struct for Framework
type Framework struct {
}

// NewFramework instantiates a new Framework object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFramework() *Framework {
	this := Framework{}
	return &this
}

// NewFrameworkWithDefaults instantiates a new Framework object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameworkWithDefaults() *Framework {
	this := Framework{}
	return &this
}

func (o Framework) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableFramework struct {
	value *Framework
	isSet bool
}

func (v NullableFramework) Get() *Framework {
	return v.value
}

func (v *NullableFramework) Set(val *Framework) {
	v.value = val
	v.isSet = true
}

func (v NullableFramework) IsSet() bool {
	return v.isSet
}

func (v *NullableFramework) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFramework(val *Framework) *NullableFramework {
	return &NullableFramework{value: val, isSet: true}
}

func (v NullableFramework) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFramework) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



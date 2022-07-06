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

// DeviceStatus struct for DeviceStatus
type DeviceStatus struct {
}

// NewDeviceStatus instantiates a new DeviceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceStatus() *DeviceStatus {
	this := DeviceStatus{}
	return &this
}

// NewDeviceStatusWithDefaults instantiates a new DeviceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceStatusWithDefaults() *DeviceStatus {
	this := DeviceStatus{}
	return &this
}

func (o DeviceStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableDeviceStatus struct {
	value *DeviceStatus
	isSet bool
}

func (v NullableDeviceStatus) Get() *DeviceStatus {
	return v.value
}

func (v *NullableDeviceStatus) Set(val *DeviceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceStatus(val *DeviceStatus) *NullableDeviceStatus {
	return &NullableDeviceStatus{value: val, isSet: true}
}

func (v NullableDeviceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


